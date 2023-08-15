package trading

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	v2 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v2pair"
	v3 "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
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
			log.Info("RonFi Sandwich optimalWethIn", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn)
			if frontRunTx, amountOut, ok := s.buildSandWichTx(pool, amountIn); ok && (frontRunTx != nil && amountOut != nil && amountOut.Cmp(big.NewInt(0)) > 0) {
				txs = append(txs, frontRunTx)
				// add target tx
				txs = append(txs, s.targetTx)
				// build back run tx
				rPool := pool.Reverse()
				if backRunTx, _, ok := s.buildSandWichTx(rPool, amountOut); ok && backRunTx != nil {
					// apply transactions to check if it's profitable
					// apply transaction
					var (
						applySuccess, reverted bool
						err                    string
					)
					appState := s.appState.Copy()
					applySuccess, reverted, err = s.worker.applyTransaction(frontRunTx, ronFiTxHash(frontRunTx.Hash()), appState)
					if !applySuccess || reverted {
						log.Warn("RonFi Sandwich Build applyTransaction frontRunTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
						return txs, false
					}

					applySuccess, reverted, _ = s.worker.applyTransaction(s.targetTx, ronFiTxHash(s.targetTx.Hash()), appState)
					if !applySuccess || reverted {
						log.Warn("RonFi Sandwich Build applyTransaction targetTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
						return txs, false
					}

					applySuccess, reverted, _ = s.worker.applyTransaction(backRunTx, ronFiTxHash(backRunTx.Hash()), appState)
					if !applySuccess || reverted {
						log.Warn("RonFi Sandwich Build applyTransaction backRunTx failed", "tx", s.targetTx.Hash(), "pair", pool.Address, "err", err)
						return txs, false
					}

					// check if profitable
					vlogs := appState.GetLogs(ronFiTxHash(backRunTx.Hash()), s.worker.currentBlockNum, common.Hash{})
					swapPairsInfo := s.worker.di.ExtractSwapPairInfo(s.targetTx, *s.targetTx.To(), vlogs, defi.RonFiExtractTypeHunting)
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
				log.Warn("RonFi Sandwich Build front run transaction failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
			}
		} else {
			log.Warn("RonFi Sandwich Build optimalWethIn failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
		}
	}

	return txs, true
}

func (s *RonSandwich) optimalWethIn(pool *defi.SwapPairInfo) (*big.Int, bool) {
	calculateF := func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) (*types.Transaction, bool) {
		frontRunTx, amountOut, ok := s.buildSandWichTx(pool, amountIn)
		if !ok || frontRunTx == nil || amountOut == nil || amountOut.Cmp(big.NewInt(0)) <= 0 {
			log.Warn("RonFi Sandwich optimalWethIn buildSandWichTx failed", "pair", pool.Address)
			return nil, false
		}

		// apply transaction
		var (
			applySuccess, reverted bool
			//err                    string
		)
		frontRunTxHash := ronFiTxHash(frontRunTx.Hash())
		applySuccess, reverted, _ = s.worker.applyTransaction(frontRunTx, frontRunTxHash, appState)
		if !applySuccess || reverted {
			//log.Warn("RonFi Sandwich optimalWethIn applyTransaction frontRunTx failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address)
			return nil, false
		}

		targetTxHash := ronFiTxHash(s.targetTx.Hash())
		applySuccess, reverted, _ = s.worker.applyTransaction(s.targetTx, targetTxHash, appState)
		if !applySuccess || reverted {
			//log.Warn("RonFi Sandwich optimalWethIn applyTransaction targetTx failed", "targetTx", s.targetTx.Hash(), "pair", pool.Address)
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
	gap := new(big.Int).Sub(right, left)
	if gap.Cmp(tolerance) > 0 {
		mid := new(big.Int).Add(left, new(big.Int).Div(gap, big.NewInt(2)))
		if frontRunTx, ok := calculateF(s.appState.Copy(), pool, mid); ok {
			// Number go up
			return s.binarySearch(pool, frontRunTx, mid, calculateF, mid, right, tolerance)
		} else {
			// Number go down
			return s.binarySearch(pool, nil, nil, calculateF, left, mid, tolerance)
		}
	}

	if frontRunTx != nil && (amountIn != nil && amountIn.Cmp(big.NewInt(0)) > 0) {
		return amountIn, true
	} else {
		return nil, false
	}
}

func (s *RonSandwich) buildSandWichTx(pool *defi.SwapPairInfo, amountIn *big.Int) (frontRunTx *types.Transaction, amountOut *big.Int, ok bool) {
	options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID))
	if err != nil {
		log.Warn("RonFi buildSandWichTx get options failed!", "err", err)
		return
	}

	options.Nonce = new(big.Int).SetUint64(s.appState.GetNonce(s.execAddress))
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

			if tx, err := inst.Swap(options, amount0Out, amount1Out, pool.Address, []byte{}); err != nil {
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

				if tx, err := inst.Swap(options, txpool.RonFiSwapV3Address, pool.Dir == 0, amountIn, v3PoolInst.State.SqrtPriceX96, []byte{}); err != nil {
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
