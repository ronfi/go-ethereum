package trading

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	erc20token "github.com/ethereum/go-ethereum/ronfi/contracts/contract_erc20"
	v2 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v2pair"
	v3 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
	wethToken "github.com/ethereum/go-ethereum/ronfi/contracts/contract_weth"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	uniswapV2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	uniswapV3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
)

type RonSandwich struct {
	worker      *Worker
	privKey     *ecdsa.PrivateKey
	execAddress common.Address
	targetTx    *types.Transaction
	pools       []*defi.SwapPairInfo
	appState    *state.StateDB
}

func NewRonSandwich(worker *Worker, executorPrivKey *ecdsa.PrivateKey, executorAddress common.Address, targetTx *types.Transaction, pairsInfo []*defi.SwapPairInfo, appState *state.StateDB) *RonSandwich {
	ronSandWich := &RonSandwich{
		worker:      worker,
		privKey:     executorPrivKey,
		execAddress: executorAddress,
		targetTx:    targetTx,
		appState:    appState,
	}

	pools := make([]*defi.SwapPairInfo, 0, len(pairsInfo))
	for _, info := range pairsInfo {
		// only keep pair/pool whose tokenIn is WETH
		if info.TokenIn != rcommon.WETH {
			//log.Warn("RonFi Sandwich NewRonSandwich invalid pair", "tx", targetTx.Hash(), "pair", info.Address)
			continue
		} else {
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
		if amountIn, ok := s.optimalWethIn(pool); ok && (amountIn != nil && amountIn.Cmp(big.NewInt(0)) > 0) {
			log.Info("RonFi Sandwich optimalWethIn succeed!", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn)
			// apply transactions to check if it's profitable
			// apply transaction
			var (
				applySuccess, reverted bool
				err                    string
			)
			appState := s.appState.Copy()

			nonce := s.appState.GetNonce(s.execAddress)
			// convert eth to weth
			{
				conTx := s.convertEthToWETH(amountIn, nonce)
				if conTx == nil {
					return nil, false
				}

				applySuccess, reverted, err = s.worker.applyTransaction(conTx, ronFiTxHash(conTx.Hash()), appState)
				if !applySuccess || reverted {
					log.Warn("RonFi Sandwich Build applyTransaction convert weth failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
					return nil, false
				}
			}

			// transfer weth to v2 pair
			{
				nonce++
				transTx := s.transferToken(rcommon.WETH, pool.Address, amountIn, nonce)
				if transTx == nil {
					return nil, false
				}

				applySuccess, reverted, err = s.worker.applyTransaction(transTx, ronFiTxHash(transTx.Hash()), appState)
				if !applySuccess || reverted {
					log.Warn("RonFi Sandwich Build applyTransaction transfer token failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "err", err)
					return nil, false
				}
			}

			// front run tx
			amountBack := big.NewInt(0)
			{
				nonce++
				frontRunTx, amountOut, ok := s.buildSandWichTx(pool, amountIn, nonce)
				if !ok || frontRunTx == nil || amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
					log.Warn("RonFi Sandwich Build create front run tx failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
					return nil, false
				}

				applySuccess, reverted, err = s.worker.applyTransaction(frontRunTx, ronFiTxHash(frontRunTx.Hash()), appState)
				if !applySuccess || reverted {
					log.Warn("RonFi Sandwich Build applyTransaction frontRunTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
					return txs, false
				} else {
					vlogs := appState.GetLogs(ronFiTxHash(frontRunTx.Hash()), s.worker.currentBlockNum, common.Hash{})
					swapPairsInfo := s.worker.di.ExtractSwapPairInfo(frontRunTx, *frontRunTx.To(), vlogs, defi.RonFiExtractTypeHunting)
					if len(swapPairsInfo) > 0 {
						swapPairInfo := swapPairsInfo[0]
						bAmountOut := swapPairInfo.AmountOut
						if bAmountOut == nil || bAmountOut.Cmp(big.NewInt(0)) <= 0 {
							log.Warn("RonFi Sandwich Build back run tx ExtractSwapPairInfo amount failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
							return txs, false
						} else {
							log.Info("RonFi Sandwich Build frontRun", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "amountOut", bAmountOut)
							amountBack = bAmountOut
						}
					} else {
						log.Warn("RonFi Sandwich Build front run tx ExtractSwapPairInfo failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
						return txs, false
					}
				}
			}

			// target tx
			{
				applySuccess, reverted, err = s.worker.applyTransaction(s.targetTx, ronFiTxHash(s.targetTx.Hash()), appState)
				if !applySuccess || reverted {
					log.Warn("RonFi Sandwich Build applyTransaction targetTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
					return txs, false
				}
			}

			rPool := pool.Reverse()
			// transfer token back
			{
				if !rPool.V3 {
					// for v2, we should transfer token first
					nonce++
					transBTx := s.transferToken(rPool.TokenIn, rPool.Address, amountBack, nonce)
					if transBTx == nil {
						return nil, false
					}

					applySuccess, reverted, err = s.worker.applyTransaction(transBTx, ronFiTxHash(transBTx.Hash()), appState)
					if !applySuccess || reverted {
						log.Warn("RonFi Sandwich Build applyTransaction transBTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
						return txs, false
					}
				}
			}

			// back run tx
			{
				nonce++
				backRunTx, _, ok := s.buildSandWichTx(rPool, amountBack, nonce)
				if !ok || backRunTx == nil {
					log.Warn("RonFi Sandwich Build create back run tx failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
					return nil, false
				}

				applySuccess, reverted, err = s.worker.applyTransaction(backRunTx, ronFiTxHash(backRunTx.Hash()), appState)
				if !applySuccess || reverted {
					log.Warn("RonFi Sandwich Build applyTransaction backRunTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
					return txs, false
				}

				// check if profitable
				vlogs := appState.GetLogs(ronFiTxHash(backRunTx.Hash()), s.worker.currentBlockNum, common.Hash{})
				swapPairsInfo := s.worker.di.ExtractSwapPairInfo(backRunTx, *backRunTx.To(), vlogs, defi.RonFiExtractTypeHunting)
				if len(swapPairsInfo) > 0 {
					swapPairInfo := swapPairsInfo[0]
					bAmountOut := swapPairInfo.AmountOut
					if bAmountOut == nil || bAmountOut.Cmp(big.NewInt(0)) <= 0 {
						log.Warn("RonFi Sandwich Build back run tx ExtractSwapPairInfo amount failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
						return txs, false
					} else {
						if bAmountOut.Cmp(amountIn) > 0 {
							log.Info("RonFi Sandwich Build profitable", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "profit", new(big.Int).Sub(bAmountOut, amountIn))
							txs = append(txs, backRunTx)
							return txs, true
						} else {
							log.Warn("RonFi Sandwich Build not profitable", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "bAmountOut", bAmountOut)
							return txs, false
						}
					}
				} else {
					log.Warn("RonFi Sandwich Build back run tx ExtractSwapPairInfo failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
					return txs, false
				}
			}
		} else {
			log.Warn("RonFi Sandwich Build optimalWethIn failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
		}
	}

	return txs, true
}

func (s *RonSandwich) optimalWethIn(pool *defi.SwapPairInfo) (*big.Int, bool) {
	calculateF := func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool) {
		// convert eth to weth
		nonce := s.appState.GetNonce(s.execAddress)
		conTx := s.convertEthToWETH(big.NewInt(800000000000000000), nonce)
		if conTx == nil {
			//log.Warn("RonFi Sandwich calculateF convertEthToWETH failed", "pair", pool.Address)
			return nil, false
		}

		// transfer weth to v2 pair
		nonce++
		transTx := s.transferToken(rcommon.WETH, pool.Address, amountIn, nonce)
		if transTx == nil {
			//log.Warn("RonFi Sandwich calculateF transferToken failed", "pair", pool.Address)
			return nil, false
		}

		nonce++
		frontRunTx, amountOut, ok := s.buildSandWichTx(pool, amountIn, nonce)
		if !ok || frontRunTx == nil || amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
			//log.Warn("RonFi Sandwich calculateF buildSandWichTx failed", "pair", pool.Address)
			return nil, false
		}

		// apply transaction
		var (
			applySuccess, reverted bool
			//err                    string
		)

		applySuccess, reverted, _ = s.worker.applyTransaction(conTx, ronFiTxHash(conTx.Hash()), appState)
		if !applySuccess || reverted {
			//log.Warn("RonFi Sandwich calculateF applyTransaction convert weth failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
			return nil, false
		}

		applySuccess, reverted, _ = s.worker.applyTransaction(transTx, ronFiTxHash(transTx.Hash()), appState)
		if !applySuccess || reverted {
			//log.Warn("RonFi Sandwich calculateF applyTransaction transfer token failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "err", err)
			return nil, false
		}

		frontRunTxHash := ronFiTxHash(frontRunTx.Hash())
		applySuccess, reverted, _ = s.worker.applyTransaction(frontRunTx, frontRunTxHash, appState)
		if !applySuccess || reverted {
			//log.Warn("RonFi Sandwich calculateF applyTransaction frontRunTx failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn, "err", err)
			return nil, false
		}

		targetTxHash := ronFiTxHash(s.targetTx.Hash())
		applySuccess, reverted, _ = s.worker.applyTransaction(s.targetTx, targetTxHash, appState)
		if !applySuccess || reverted {
			//log.Warn("RonFi Sandwich calculateF applyTransaction targetTx failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
			return nil, false
		}

		return frontRunTx, true
	}

	left := big.NewInt(0)                                              // 0 weth
	right := new(big.Int).Exp(big.NewInt(10), big.NewInt(21), nil)     // 1000 weth
	tolerance := new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil) //0.001 weth
	return s.binarySearch(pool, nil, nil, calculateF, left, right, tolerance)
}

func (s *RonSandwich) binarySearch(pool *defi.SwapPairInfo, frontRunTx *types.Transaction, amountIn *big.Int, calculateF func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool), left, right *big.Int, tolerance *big.Int) (*big.Int, bool) {
	lastTx := frontRunTx
	lastAmount := amountIn

	gap := new(big.Int).Sub(right, left)
	if gap.Cmp(tolerance) > 0 {
		mid := new(big.Int).Add(left, new(big.Int).Div(gap, big.NewInt(2)))
		if frontRunTx, ok := calculateF(s.appState.Copy(), pool, mid); ok {
			// Number go up
			//log.Warn("RonFi Sandwich binarySearch calculateF succeed!", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", mid)
			return s.binarySearch(pool, frontRunTx, mid, calculateF, mid, right, tolerance)
		} else {
			// Number go down
			return s.binarySearch(pool, lastTx, lastAmount, calculateF, left, mid, tolerance)
		}
	}

	if lastTx != nil && (lastAmount != nil && lastAmount.Cmp(big.NewInt(0)) > 0) {
		return lastAmount, true
	} else {
		return nil, false
	}
}

func (s *RonSandwich) buildSandWichTx(pool *defi.SwapPairInfo, amountIn *big.Int, nonce uint64) (frontRunTx *types.Transaction, amountOut *big.Int, ok bool) {
	options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID))
	if err != nil {
		log.Warn("RonFi buildSandWichTx get options failed!", "err", err)
		return
	}

	options.Nonce = new(big.Int).SetUint64(nonce)
	options.Value = big.NewInt(0)
	options.GasPrice = s.worker.gasPrice

	if gasLimit, ok := s.worker.pairGasMap[fmt.Sprintf("%s-%d", pool.Address, pool.Dir)]; ok {
		options.GasLimit = gasLimit
	} else {
		options.GasLimit = GasLimitOneHop
	}
	options.NoSend = true //only return signedTx

	if !pool.V3 {
		if inst, err := v2.NewV2pair(pool.Address, s.worker.client); err != nil {
			log.Warn("RonFi Sandwich buildSandWichTx NewV2pair failed", "pair", pool.Address, "err", err)
			return
		} else {
			v2PoolInst := uniswapV2.NewV2Pool(s.worker.di, pool.Address, 30, s.appState.Copy())
			if v2PoolInst == nil {
				log.Warn("RonFi Sandwich buildSandWichTx NewV2Pool failed", "pair", pool.Address)
				return
			}

			v2PoolInst.UpdateReserves(nil)
			amountOut = v2PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
			if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
				//log.Warn("RonFi Sandwich buildSandWichTx CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
				return
			}

			amount0Out := big.NewInt(0)
			amount1Out := big.NewInt(0)
			if pool.Dir == 0 {
				amount1Out = amountOut
			} else {
				amount0Out = amountOut
			}

			if tx, err := inst.Swap(options, amount0Out, amount1Out, s.execAddress, []byte{}); err != nil {
				log.Warn("RonFi Sandwich buildSandWichTx Swap failed", "pair", pool.Address, "err", err)
				return
			} else {
				frontRunTx = tx
			}
		}
	} else {
		if inst, err := v3.NewV3pool(pool.Address, s.worker.client); err != nil {
			log.Warn("RonFi Sandwich buildSandWichTx NewV3Pool failed", "pool", pool.Address, "err", err)
			return
		} else {
			poolInfo := s.worker.di.GetPoolInfo(pool.Address)
			if poolInfo == nil {
				log.Warn("RonFi Sandwich buildSandWichTx GetPoolInfo failed", "pool", pool.Address)
				return
			} else {
				tickLens := rcommon.GetV3TickLens(poolInfo.Factory)
				if tickLens == rcommon.ZeroAddress {
					log.Warn("RonFi Sandwich buildSandWichTx GetV3TickLens failed", "pool", pool.Address)
					return
				}

				v3PoolInst := uniswapV3.NewV3Pool(s.worker.di, s.targetTx, pool.Address, tickLens, s.appState.Copy())
				if v3PoolInst == nil {
					log.Warn("RonFi Sandwich buildSandWichTx NewV3Pool failed", "pool", pool.Address)
					return
				}

				v3PoolInst.UpdatePoolState(nil)
				amountOut, _ = v3PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
				if amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
					//log.Warn("RonFi Sandwich buildSandWichTx CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
					return
				}

				if tx, err := inst.Swap(options, s.execAddress, pool.Dir == 0, amountIn, v3PoolInst.State.SqrtPriceX96, []byte{}); err != nil {
					log.Warn("RonFi Sandwich buildSandWichTx Swap failed", "pool", pool.Address, "err", err)
					return
				} else {
					frontRunTx = tx
				}
			}
		}
	}

	ok = true
	return
}

func (s *RonSandwich) convertEthToWETH(amount *big.Int, nonce uint64) *types.Transaction {
	if inst, err := wethToken.NewWeth(rcommon.WETH, s.worker.client); err != nil {
		log.Warn("RonFi Sandwich convertEthToWETH NewWeth failed", "err", err)
		return nil
	} else {
		options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID))
		if err != nil {
			log.Warn("RonFi convertEthToWETH get options failed!", "err", err)
			return nil
		}

		options.Nonce = new(big.Int).SetUint64(nonce)
		options.Value = amount
		options.GasPrice = s.worker.gasPrice
		options.GasLimit = 50000
		options.NoSend = true

		if tx, err := inst.Deposit(options); err != nil {
			log.Warn("RonFi Sandwich convertEthToWETH Deposit failed", "err", err)
			return nil
		} else {
			return tx
		}
	}
}

func (s *RonSandwich) transferToken(token common.Address, to common.Address, amount *big.Int, nonce uint64) *types.Transaction {
	if inst, err := erc20token.NewToken(token, s.worker.client); err != nil {
		log.Warn("RonFi Sandwich transferToken NewToken failed", "err", err)
		return nil
	} else {
		options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID))
		if err != nil {
			log.Warn("RonFi convertEthToWETH get options failed!", "err", err)
			return nil
		}

		options.Nonce = new(big.Int).SetUint64(nonce)
		options.GasPrice = s.worker.gasPrice
		options.GasLimit = 50000
		options.NoSend = true

		if tx, err := inst.Transfer(options, to, amount); err != nil {
			log.Warn("RonFi Sandwich transferToken Deposit failed", "err", err)
			return nil
		} else {
			return tx
		}
	}
}
