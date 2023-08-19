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
		ok := s.optimalWethIn(pool, amountIn)
		if ok && amountIn.Cmp(big.NewInt(0)) > 0 {
			//log.Info("RonFi Sandwich Build optimalWethIn succeed!", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn)
			var aLegAmountOut *big.Int
			appState := s.appState.Copy()
			// aLeg tx
			if aLegTx, tmp, ok := s.sandWichTx(pool, amountIn, appState, true); !ok {
				log.Warn("RonFi Sandwich Build aLegTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
				return txs, false
			} else {
				aLegAmountOut = tmp
				txs = append(txs, aLegTx)
			}

			// target tx
			if applySuccess, reverted, err := s.worker.applyTransaction(s.targetTx, ronFiTxHash(s.targetTx.Hash()), appState); !applySuccess || reverted {
				log.Warn("RonFi Sandwich Build applyTransaction targetTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
				return txs, false
			}

			// bLeg tx
			rPool := pool.Reverse()
			if bLegTx, bLegAmountOut, ok := s.sandWichTx(rPool, aLegAmountOut, appState, false); !ok {
				log.Warn("RonFi Sandwich Build bLegTx failed",
					"tx", s.targetTx.Hash(),
					"pair", rPool.Address,
					"tokenIn", rPool.TokenIn,
					"amountIn", aLegAmountOut)
				return txs, false
			} else {
				if bLegAmountOut.Cmp(amountIn) <= 0 {
					log.Warn("RonFi Sandwich Build not profitable", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "amountOut", bLegAmountOut)
					return txs, false
				} else {
					txs = append(txs, bLegTx)
					profit := new(big.Int).Sub(bLegAmountOut, amountIn)
					log.Info("RonFi Sandwich Build profitable", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "amountOut", bLegAmountOut, "profit", profit)
					return txs, true
				}
			}
		}
	}

	return txs, true
}

func (s *RonSandwich) optimalWethIn(pool *defi.SwapPairInfo, amountIn *big.Int) bool {
	calculateF := func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool, bool) {
		var (
			applySuccess, reverted bool
			//err                    string
		)
		payloads := make([]ronswapv3fe.RonSwapV3FEPayload, 0, 10)
		if !pool.V3 {
			if payload := s.generateTransferTx(rcommon.WETH, pool.Address, amountIn); payload != nil {
				payloads = append(payloads, *payload)
			} else {
				return nil, false, false
			}
		}
		if payload, amountOut, ok := s.buildSwapTx(appState, pool, amountIn); ok && amountOut != nil && amountOut.Cmp(big.NewInt(0)) > 0 {
			payloads = append(payloads, payload)
		} else {
			return nil, false, false
		}

		if len(payloads) == 0 {
			return nil, false, false
		}

		// create tx and apply
		nonce := appState.GetNonce(s.execAddress)
		if frontRunTx := s.buildTx(pool, payloads, false, false, nonce); frontRunTx != nil {
			applySuccess, reverted, _ = s.worker.applyTransaction(frontRunTx, ronFiTxHash(frontRunTx.Hash()), appState)
			if !applySuccess || reverted {
				return nil, false, false
			}

			targetTxHash := ronFiTxHash(s.targetTx.Hash())
			if applySuccess, reverted, _ = s.worker.applyTransaction(s.targetTx, targetTxHash, appState); applySuccess && !reverted {
				//log.Info("RonFi Sandwich calculateF applyTransaction targetTx succeed!", "targetTx", s.targetTx.Hash().String(), "pair", pool.Address, "amountIn", amountIn)
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

func (s *RonSandwich) buildSwapTx(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (payload ronswapv3fe.RonSwapV3FEPayload, amountOut *big.Int, ok bool) {
	if !pool.V3 {
		v2PoolInst := uniswapV2.NewV2Pool(s.worker.di, pool.Address, 30, appState)
		if v2PoolInst == nil {
			log.Warn("RonFi Sandwich buildSwapTx NewV2Pool failed", "pair", pool.Address)
			return
		}

		v2PoolInst.UpdateReserves(nil)
		amountOut = v2PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
		if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
			//log.Warn("RonFi Sandwich buildSwapTx CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
			return
		}

		amount0Out := big.NewInt(0)
		amount1Out := big.NewInt(0)
		if pool.Dir == 0 {
			amount1Out = amountOut
		} else {
			amount0Out = amountOut
		}

		// Define ABI
		v2PairAbi, err := abi.JSON(strings.NewReader(v2.V2pairABI))
		if err != nil {
			log.Warn("RonFi buildSwapTx v2 creating ABI failed!", "err", err)
			return
		}

		var encodedParams []byte
		// Encode arguments
		encodedParams, err = v2PairAbi.Pack("swap", amount0Out, amount1Out, txpool.RonFiSwapV3Address, []byte{})
		if err != nil {
			log.Warn("RonFi buildSwapTx v2 encoding arguments failed!", "err", err)
			return
		}

		//log.Info("RonFi buildSwapTx", "pair", pool.Address, "amountIn", amountIn, "amountOut", amountOut, "amount0Out", amount0Out, "amount1Out", amount1Out)
		value := big.NewInt(0)
		payload = ronswapv3fe.RonSwapV3FEPayload{
			Target: pool.Address,
			Data:   encodedParams,
			Value:  value,
		}
	} else {
		poolInfo := s.worker.di.GetPoolInfo(pool.Address)
		if poolInfo == nil {
			log.Warn("RonFi Sandwich buildSwapTx GetPoolInfo failed", "pool", pool.Address)
			return
		} else {
			tickLens := rcommon.GetV3TickLens(poolInfo.Factory)
			if tickLens == rcommon.ZeroAddress {
				log.Warn("RonFi Sandwich buildSwapTx GetV3TickLens failed", "pool", pool.Address)
				return
			}

			v3PoolInst := uniswapV3.NewV3Pool(s.worker.di, s.targetTx, pool.Address, tickLens, appState)
			if v3PoolInst == nil {
				log.Warn("RonFi Sandwich buildSwapTx NewV3Pool failed", "pool", pool.Address)
				return
			}

			v3PoolInst.UpdatePoolState(nil)
			amountOut, _ = v3PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
			if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
				//log.Warn("RonFi Sandwich buildSwapTx CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
				return
			}

			// Define ABI
			v3PoolAbi, err := abi.JSON(strings.NewReader(v3.V3poolABI))
			if err != nil {
				log.Warn("RonFi Sandwich buildSwapTx v3 creating ABI failed!", "err", err)
				return
			}

			var encodedParams []byte
			encodedParams, err = v3PoolAbi.Pack("swap", txpool.RonFiSwapV3Address, pool.Dir == 0, amountIn, v3PoolInst.State.SqrtPriceX96, []byte{})
			if err != nil {
				log.Warn("RonFi Sandwich buildSwapTx v3  encoding arguments failed!", "err", err)
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

func (s *RonSandwich) generateTransferTx(token common.Address, to common.Address, amount *big.Int) *ronswapv3fe.RonSwapV3FEPayload {
	var payload *ronswapv3fe.RonSwapV3FEPayload

	if amount == nil || amount.Cmp(big.NewInt(0)) <= 0 {
		return nil
	}

	// Define ABI
	tokenAbi, err := abi.JSON(strings.NewReader(erc20token.TokenABI))
	if err != nil {
		log.Warn("RonFi optimalWethIn creating ABI failed!", "err", err)
		return nil
	}
	// Encode arguments
	encodedParams, err := tokenAbi.Pack("transfer", to, amount)
	if err != nil {
		log.Warn("Ronfi optimalWethIn encoding arguments failed!", "err", err)
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

func (s *RonSandwich) buildTx(pool *defi.SwapPairInfo, payloads []ronswapv3fe.RonSwapV3FEPayload, isSandwich, isAleg bool, nonce uint64) *types.Transaction {
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
		options.GasPrice = s.worker.gasPrice
		options.NoSend = true //only return signedTx

		if tx, err := s.worker.ronSwapInst.ExecutePayloads(options, isSandwich, isAleg, payloads); err == nil {
			return tx
		}
	}

	return nil
}

func (s *RonSandwich) sandWichTx(pool *defi.SwapPairInfo, amountIn *big.Int, appState *state.StateDB, isAleg bool) (*types.Transaction, *big.Int, bool) {
	var (
		applySuccess, reverted bool
		err                    string
	)

	payloads := make([]ronswapv3fe.RonSwapV3FEPayload, 0, 10)
	if !pool.V3 {
		if payload := s.generateTransferTx(pool.TokenIn, pool.Address, amountIn); payload != nil {
			payloads = append(payloads, *payload)
		} else {
			log.Warn("RonFi Sandwich generateTransferTx failed",
				"tx", s.targetTx.Hash(),
				"pair", pool.Address,
				"token", pool.TokenIn,
				"isAleg", isAleg,
				"amountIn", amountIn)
			return nil, nil, false
		}
	}
	if payload, amountOut, ok := s.buildSwapTx(appState, pool, amountIn); ok && amountOut != nil && amountOut.Cmp(big.NewInt(0)) > 0 {
		payloads = append(payloads, payload)
	} else {
		log.Warn("RonFi Sandwich buildSwapTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "isAleg", isAleg, "amountIn", amountIn)
		return nil, nil, false
	}

	if len(payloads) == 0 {
		log.Warn("RonFi Sandwich payloads empty", "tx", s.targetTx.Hash(), "pair", pool.Address, "isAleg", isAleg)
		return nil, nil, false
	}

	if tx := s.buildTx(pool, payloads, true, isAleg, appState.GetNonce(s.execAddress)); tx != nil {
		applySuccess, reverted, err = s.worker.applyTransaction(tx, ronFiTxHash(tx.Hash()), appState)
		if !applySuccess || reverted {
			log.Warn("RonFi Sandwich applyTransaction failed",
				"tx", s.targetTx.Hash(),
				"pair", pool.Address,
				"tokenIn", pool.TokenIn,
				"amountIn", amountIn,
				"isAleg", isAleg,
				"err", err)
			return nil, nil, false
		}

		vlogs := appState.GetLogs(ronFiTxHash(tx.Hash()), s.worker.currentBlockNum, common.Hash{})
		if bAmountOut := s.worker.di.ExtractTransferAmount(vlogs, pool.TokenOut, pool.Address, txpool.RonFiSwapV3Address); bAmountOut != nil && bAmountOut.Cmp(big.NewInt(0)) > 0 {
			return tx, bAmountOut, true
		}
	}

	return nil, nil, false
}
