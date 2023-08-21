package trading

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	erc20token "github.com/ethereum/go-ethereum/ronfi/contracts/contract_erc20"
	ronswapv3fe "github.com/ethereum/go-ethereum/ronfi/contracts/contract_ronswapv3fe"
	v2 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v2pair"
	v3 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	uniswapV2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	uniswapV3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
	"strings"
)

type RonSandwich struct {
	worker      *Worker
	privKey     *ecdsa.PrivateKey
	execAddress common.Address
	targetTx    *types.Transaction
	pools       []*defi.SwapPairInfo
	appState    *state.StateDB
	lowerBound  *big.Int
	upperBound  *big.Int
	tolerance   *big.Int
}

func NewRonSandwich(worker *Worker, executorPrivKey *ecdsa.PrivateKey, executorAddress common.Address, targetTx *types.Transaction, pairsInfo []*defi.SwapPairInfo, appState *state.StateDB) *RonSandwich {
	ronSandWich := &RonSandwich{
		worker:      worker,
		privKey:     executorPrivKey,
		execAddress: executorAddress,
		targetTx:    targetTx,
		appState:    appState,
	}

	ronSandWich.lowerBound = big.NewInt(0)
	balance := worker.di.GetTokenBalance(txpool.RonFiSwapV3Address, rcommon.WETH)
	if balance == nil {
		log.Warn("RonFi Sandwich NewRonSandwich GetTokenBalance failed", "tx", targetTx.Hash(), "RonFiSwapV3Address", txpool.RonFiSwapV3Address)
		return nil
	} else {
		ronSandWich.upperBound = balance
	}
	ronSandWich.tolerance = new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil) //0.001 weth

	pools := make([]*defi.SwapPairInfo, 0, len(pairsInfo))
	for _, info := range pairsInfo {
		// only keep pair/pool whose tokenIn is WETH
		if info.TokenIn == rcommon.WETH {
			pools = append(pools, info)
		}
	}

	if len(pools) > 0 {
		ronSandWich.pools = pools
		return ronSandWich
	}

	return nil
}

func (s *RonSandwich) Build() ([]*types.Transaction, bool) {
	txs := make([]*types.Transaction, 0, 4)

	//iterate pools
	for _, pool := range s.pools {
		amountIn := big.NewInt(0)
		ok := s.optimize(pool, amountIn)
		if ok && amountIn.Cmp(big.NewInt(0)) > 0 {
			//log.Info("RonFi Sandwich Build optimize succeed!", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn)
			var aLegAmountOut *big.Int
			appState := s.appState.Copy()
			// aLeg tx
			var (
				feeRate                  *big.Int
				aLegGasUsed, bLegGasUsed uint64
			)
			if aLegTx, fee, tmp, gasUsed, ok := s.sandWichTx(pool, amountIn, appState, true); !ok {
				log.Warn("RonFi Sandwich Build aLegTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
				return txs, false
			} else {
				aLegAmountOut = tmp
				aLegGasUsed = gasUsed
				txs = append(txs, aLegTx)
				feeRate = fee
			}
			log.Info("RonFi Sandwich Build aLegTx succeed!", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "amountOut", aLegAmountOut, "feeRate", feeRate)

			// target tx
			if applySuccess, reverted, _, err := s.worker.applyTransaction(s.targetTx, ronFiTxHash(s.targetTx.Hash()), appState); !applySuccess || reverted {
				log.Warn("RonFi Sandwich Build applyTransaction targetTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
				return txs, false
			} else {
				txs = append(txs, s.targetTx)
			}

			// bLeg tx
			rPool := pool.Reverse()
			bLegAmountIn := aLegAmountOut
			if feeRate != nil {
				bLegAmountIn = new(big.Int).Sub(aLegAmountOut, new(big.Int).Div(new(big.Int).Mul(aLegAmountOut, feeRate), big.NewInt(10000)))
			}
			if _, _, bLegAmountOut, gasUsed, ok := s.sandWichTx(rPool, bLegAmountIn, appState, false); !ok {
				log.Warn("RonFi Sandwich Build bLegTx failed",
					"tx", s.targetTx.Hash(),
					"pair", rPool.Address,
					"tokenIn", rPool.TokenIn,
					"amountIn", aLegAmountOut)
				return txs, false
			} else {
				bLegGasUsed = gasUsed
				log.Info("RonFi Sandwich Build bLegTx succeed!", "tx", s.targetTx.Hash(), "pair", rPool.Address, "amountIn", bLegAmountIn, "amountOut", bLegAmountOut)
				if bLegAmountOut.Cmp(amountIn) <= 0 {
					log.Warn("RonFi Sandwich Build not profitable", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "amountOut", bLegAmountOut)
					return txs, false
				} else {
					profit := new(big.Int).Sub(bLegAmountOut, amountIn)
					totalGasUsed := aLegGasUsed + bLegGasUsed
					txFee := new(big.Int).Mul(big.NewInt(int64(totalGasUsed)), s.worker.gasPrice)
					netProfit := new(big.Int).Sub(profit, txFee)
					bribeFee := new(big.Int).Div(new(big.Int).Mul(netProfit, big.NewInt(60)), big.NewInt(100)) // 80% net profit to miner
					bLegTxFee := new(big.Int).Add(new(big.Int).Mul(big.NewInt(int64(bLegGasUsed)), s.worker.gasPrice), bribeFee)
					bLegGasFee := new(big.Int).Div(bLegTxFee, big.NewInt(int64(bLegGasUsed)))
					payloads := s.generatePayloads(rPool, bLegAmountIn, appState)
					if len(payloads) == 0 {
						return txs, false
					}
					if tx := s.buildExecuteTx(pool, payloads, true, false, s.appState.GetNonce(s.execAddress)+1, bLegGasFee); tx != nil {
						txs = append(txs, tx)
						if profit.Cmp(txFee) > 0 {
							log.Info("RonFi Sandwich Build profitable",
								"tx", s.targetTx.Hash(),
								"pair", pool.Address,
								"amountIn", amountIn,
								"amountOut", bLegAmountOut,
								"profit", profit,
								"bribeFee", bribeFee,
								"ronProfit", new(big.Int).Sub(netProfit, bribeFee))
							return txs, true
						}
					}
				}
			}
		}
	}

	return txs, true
}

func (s *RonSandwich) optimize(pool *defi.SwapPairInfo, amountIn *big.Int) bool {
	calculateF := func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool, bool) {
		var (
			applySuccess, reverted bool
			//err                    string
		)

		payloads := s.generatePayloads(pool, amountIn, appState)
		if len(payloads) == 0 {
			return nil, false, false
		}

		// create tx and apply
		nonce := appState.GetNonce(s.execAddress)
		if frontRunTx := s.buildExecuteTx(pool, payloads, false, false, nonce, s.worker.gasPrice); frontRunTx != nil {
			applySuccess, reverted, _, _ = s.worker.applyTransaction(frontRunTx, ronFiTxHash(frontRunTx.Hash()), appState)
			if !applySuccess || reverted {
				return nil, false, false
			}

			targetTxHash := ronFiTxHash(s.targetTx.Hash())
			if applySuccess, reverted, _, _ = s.worker.applyTransaction(s.targetTx, targetTxHash, appState); applySuccess && !reverted {
				log.Info("RonFi Sandwich calculateF applyTransaction targetTx succeed!", "targetTx", s.targetTx.Hash().String(), "pair", pool.Address, "amountIn", amountIn)
				return frontRunTx, false, true
			} else {
				return nil, true, false
			}
		}

		return nil, false, false
	}

	return s.binarySearch(pool, calculateF, s.lowerBound, s.upperBound, s.tolerance, amountIn)
}

func (s *RonSandwich) binarySearch(pool *defi.SwapPairInfo, calculateF func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool, bool), left, right *big.Int, tolerance *big.Int, amountIn *big.Int) bool {
	gap := new(big.Int).Sub(right, left)
	if gap.Cmp(tolerance) > 0 {
		mid := new(big.Int).Add(left, new(big.Int).Div(gap, big.NewInt(2)))
		_, terminate, ok := calculateF(s.appState.Copy(), pool, mid)
		if terminate {
			// apply target tx failed, then exit the iteration
			return false
		}
		if ok {
			// Number go up
			amountIn.Set(mid)
			return s.binarySearch(pool, calculateF, mid, right, tolerance, amountIn)
		} else {
			// Number go down
			return s.binarySearch(pool, calculateF, left, mid, tolerance, amountIn)
		}
	}

	if amountIn.Cmp(big.NewInt(0)) > 0 {
		return true
	} else {
		return false
	}
}

func (s *RonSandwich) genTransTxPayload(token common.Address, to common.Address, amount *big.Int) *ronswapv3fe.RonSwapV3FEPayload {
	var payload *ronswapv3fe.RonSwapV3FEPayload

	if amount == nil || amount.Cmp(big.NewInt(0)) <= 0 {
		return nil
	}

	// Define ABI
	tokenAbi, err := abi.JSON(strings.NewReader(erc20token.TokenABI))
	if err != nil {
		log.Warn("RonFi optimize creating ABI failed!", "err", err)
		return nil
	}
	// Encode arguments
	encodedParams, err := tokenAbi.Pack("transfer", to, amount)
	if err != nil {
		log.Warn("Ronfi optimize encoding arguments failed!", "err", err)
		return nil
	}
	value := big.NewInt(0)
	payload = &ronswapv3fe.RonSwapV3FEPayload{
		Target: token,
		Data:   encodedParams,
		Value:  value,
	}

	return payload
}

func (s *RonSandwich) genSwapTxPayload(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (payload ronswapv3fe.RonSwapV3FEPayload, amountOut *big.Int, ok bool) {
	if !pool.V3 {
		v2PoolInst := uniswapV2.NewV2Pool(s.worker.di, pool.Address, 30, appState)
		if v2PoolInst == nil {
			log.Warn("RonFi Sandwich genSwapTxPayload NewV2Pool failed", "pair", pool.Address)
			return
		}

		if !v2PoolInst.UpdateReserves(nil) {
			//log.Warn("RonFi Sandwich genSwapTxPayload UpdateReserves failed", "pair", pool.Address)
			return
		}
		amountOut = v2PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
		if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
			//log.Warn("RonFi Sandwich genSwapTxPayload CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
			return
		}

		amount0Out := big.NewInt(0)
		amount1Out := big.NewInt(0)
		if pool.Dir == 0 {
			amount1Out = amountOut
		} else {
			amount0Out = amountOut
		}

		if amount0Out.Cmp(big.NewInt(0)) <= 0 && amount1Out.Cmp(big.NewInt(0)) <= 0 {
			log.Warn("RonFi Sandwich genSwapTxPayload amount0Out and amount1Out both <= 0", "pair", pool.Address)
			return
		}

		// Define ABI
		v2PairAbi, err := abi.JSON(strings.NewReader(v2.V2pairABI))
		if err != nil {
			log.Warn("RonFi genSwapTxPayload v2 creating ABI failed!", "err", err)
			return
		}

		var encodedParams []byte
		// Encode arguments
		encodedParams, err = v2PairAbi.Pack("swap", amount0Out, amount1Out, txpool.RonFiSwapV3Address, []byte{})
		if err != nil {
			log.Warn("RonFi genSwapTxPayload v2 encoding arguments failed!", "err", err)
			return
		}

		//log.Info("RonFi genSwapTxPayload", "pair", pool.Address, "amountIn", amountIn, "amountOut", amountOut, "amount0Out", amount0Out, "amount1Out", amount1Out)
		value := big.NewInt(0)
		payload = ronswapv3fe.RonSwapV3FEPayload{
			Target: pool.Address,
			Data:   encodedParams,
			Value:  value,
		}
	} else {
		poolInfo := s.worker.di.GetPoolInfo(pool.Address)
		if poolInfo == nil {
			log.Warn("RonFi Sandwich genSwapTxPayload GetPoolInfo failed", "pool", pool.Address)
			return
		} else {
			tickLens := rcommon.GetV3TickLens(poolInfo.Factory)
			if tickLens == rcommon.ZeroAddress {
				log.Warn("RonFi Sandwich genSwapTxPayload GetV3TickLens failed", "pool", pool.Address)
				return
			}

			v3PoolInst := uniswapV3.NewV3Pool(s.worker.di, s.targetTx, pool.Address, tickLens, appState)
			if v3PoolInst == nil {
				log.Warn("RonFi Sandwich genSwapTxPayload NewV3Pool failed", "pool", pool.Address)
				return
			}

			if !v3PoolInst.UpdatePoolState(nil) {
				//log.Warn("RonFi Sandwich genSwapTxPayload UpdatePoolState failed", "pool", pool.Address)
				return
			}
			amountOut, _ = v3PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
			if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
				//log.Warn("RonFi Sandwich genSwapTxPayload CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
				return
			}

			// Define ABI
			v3PoolAbi, err := abi.JSON(strings.NewReader(v3.V3poolABI))
			if err != nil {
				log.Warn("RonFi Sandwich genSwapTxPayload v3 creating ABI failed!", "err", err)
				return
			}

			var encodedParams []byte
			encodedParams, err = v3PoolAbi.Pack("swap", txpool.RonFiSwapV3Address, pool.Dir == 0, amountIn, v3PoolInst.State.SqrtPriceX96, []byte{})
			if err != nil {
				log.Warn("RonFi Sandwich genSwapTxPayload v3  encoding arguments failed!", "err", err)
				return
			}
			value := big.NewInt(0)

			payload = ronswapv3fe.RonSwapV3FEPayload{
				Target: pool.Address,
				Data:   encodedParams,
				Value:  value,
			}
		}
	}

	ok = true
	return
}

func (s *RonSandwich) buildExecuteTx(pool *defi.SwapPairInfo, payloads []ronswapv3fe.RonSwapV3FEPayload, isSandwich, isAleg bool, nonce uint64, gasFee *big.Int) *types.Transaction {
	if options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID)); err != nil {
		return nil
	} else {
		options.Nonce = new(big.Int).SetUint64(nonce)
		options.Value = big.NewInt(0)
		if gasLimit, ok := s.worker.pairGasMap[fmt.Sprintf("%s-%d", pool.Address, pool.Dir)]; ok {
			options.GasLimit = gasLimit
		} else {
			options.GasLimit = GasLimitOneHop
		}
		options.GasPrice = gasFee
		options.NoSend = true //only return signedTx

		if tx, err := s.worker.ronSwapInst.ExecutePayloads(options, isSandwich, isAleg, payloads); err == nil {
			return tx
		}
	}

	return nil
}

func (s *RonSandwich) generatePayloads(pool *defi.SwapPairInfo, amountIn *big.Int, appState *state.StateDB) []ronswapv3fe.RonSwapV3FEPayload {
	payloads := make([]ronswapv3fe.RonSwapV3FEPayload, 0, 10)
	if !pool.V3 {
		if payload := s.genTransTxPayload(pool.TokenIn, pool.Address, amountIn); payload != nil {
			payloads = append(payloads, *payload)
		} else {
			return payloads
		}
	}

	if payload, amountOut, ok := s.genSwapTxPayload(appState, pool, amountIn); ok && amountOut != nil && amountOut.Cmp(big.NewInt(0)) > 0 {
		payloads = append(payloads, payload)
	}

	return payloads
}

func (s *RonSandwich) sandWichTx(pool *defi.SwapPairInfo, amountIn *big.Int, appState *state.StateDB, isAleg bool) (*types.Transaction, *big.Int, *big.Int, uint64, bool) {
	var (
		applySuccess, reverted bool
		gasUsed                uint64
		err                    string
	)

	payloads := s.generatePayloads(pool, amountIn, appState)
	if len(payloads) == 0 {
		log.Warn("RonFi Sandwich payloads empty", "tx", s.targetTx.Hash(), "pair", pool.Address, "isAleg", isAleg)
		return nil, nil, nil, 0, false
	}

	if tx := s.buildExecuteTx(pool, payloads, true, isAleg, appState.GetNonce(s.execAddress), s.worker.gasPrice); tx != nil {
		applySuccess, reverted, gasUsed, err = s.worker.applyTransaction(tx, ronFiTxHash(tx.Hash()), appState)
		if !applySuccess || reverted {
			log.Warn("RonFi Sandwich applyTransaction failed",
				"tx", s.targetTx.Hash(),
				"pair", pool.Address,
				"tokenIn", pool.TokenIn,
				"amountIn", amountIn,
				"gasUsed", gasUsed,
				"isAleg", isAleg,
				"applySuccess", applySuccess,
				"reverted", reverted,
				"err", err)
			return nil, nil, nil, 0, false
		}

		vlogs := appState.GetLogs(ronFiTxHash(tx.Hash()), s.worker.currentBlockNum, common.Hash{})
		swapPairsInfo := s.worker.di.ExtractSwapPairInfo(tx, *tx.To(), vlogs, defi.RonFiExtractTypeHunting)
		if len(swapPairsInfo) == 0 {
			return nil, nil, nil, 0, false
		} else {
			info := swapPairsInfo[0]
			swapAmountOut := info.AmountOut
			if bAmountOut := s.worker.di.ExtractTransferAmount(vlogs, pool.TokenOut, pool.Address, txpool.RonFiSwapV3Address); bAmountOut != nil && bAmountOut.Cmp(big.NewInt(0)) > 0 {
				var feeRate *big.Int
				if bAmountOut.Cmp(swapAmountOut) != 0 {
					log.Warn("RonFi Sandwich bAmountOut != swapAmountOut",
						"tx", s.targetTx.Hash(),
						"pair", pool.Address,
						"tokenIn", pool.TokenIn,
						"amountIn", amountIn,
						"isAleg", isAleg,
						"swapAmountOut", swapAmountOut,
						"bAmountOut", bAmountOut)
					feeRate = new(big.Int).Div(new(big.Int).Mul(big.NewInt(10000), new(big.Int).Sub(swapAmountOut, bAmountOut)), swapAmountOut)
					feeRate = new(big.Int).Add(feeRate, big.NewInt(1))
				}
				return tx, feeRate, bAmountOut, gasUsed, true
			}
		}

	}

	return nil, nil, nil, 0, false
}
