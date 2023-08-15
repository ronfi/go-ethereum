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
		}
	}

	if len(pools) > 0 {
		ronSandWich.pools = pools
		return ronSandWich
	}

	return nil
}

func (s *RonSandwich) Build() []*types.Transaction {
	txs := make([]*types.Transaction, 0, 4)

	//iterate pools
	for _, pool := range s.pools {
		amountIn := s.optimalWethIn(pool)
		if amountIn != nil && amountIn.Cmp(big.NewInt(0)) > 0 {
			log.Info("RonFi Sandwich optimalWethIn", "tx", s.targetTx.Hash(), "pair", pool.Address, "amountIn", amountIn)
		} else {
			log.Warn("RonFi Sandwich optimalWethIn failed", "tx", s.targetTx.Hash(), "pair", pool.Address)
		}
	}

	return txs
}

func (s *RonSandwich) optimalWethIn(pool *defi.SwapPairInfo) *big.Int {
	calculateF := func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) bool {
		frontRunTx := s.buildFrontRunTx(pool, amountIn)
		if frontRunTx == nil {
			log.Warn("RonFi Sandwich optimalWethIn buildFrontRunTx failed", "pair", pool.Address)
			return false
		}

		// apply transaction
		var (
			applySuccess, reverted bool
			//err                    string
		)
		frontRunTxHash := ronFiTxHash(frontRunTx.Hash())
		applySuccess, reverted, _ = s.worker.applyTransaction(frontRunTx, frontRunTxHash, appState)
		if !applySuccess || reverted {
			return false
		}

		targetTxHash := ronFiTxHash(s.targetTx.Hash())
		applySuccess, reverted, _ = s.worker.applyTransaction(s.targetTx, targetTxHash, appState)
		if !applySuccess || reverted {
			return false
		}

		return true
	}

	left := big.NewInt(0)                                              // 0 weth
	right := new(big.Int).Exp(big.NewInt(10), big.NewInt(21), nil)     // 1000 weth
	tolerance := new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil) //0.001 weth
	return s.binarySearch(pool, calculateF, left, right, tolerance)
}

func (s *RonSandwich) binarySearch(pool *defi.SwapPairInfo, calculateF func(appState *state.StateDB, pool *defi.SwapPairInfo, amountIn *big.Int) bool, left, right *big.Int, tolerance *big.Int) *big.Int {
	gap := new(big.Int).Sub(right, left)
	if gap.Cmp(tolerance) > 0 {
		mid := new(big.Int).Add(left, new(big.Int).Div(gap, big.NewInt(2)))
		if calculateF(s.appState.Copy(), pool, mid) {
			// Number go up
			return s.binarySearch(pool, calculateF, mid, right, tolerance)
		} else {
			// Number go down
			return s.binarySearch(pool, calculateF, left, mid, tolerance)
		}
	}

	ret := new(big.Int).Div(new(big.Int).Add(right, left), big.NewInt(2))
	if ret.Cmp(big.NewInt(0)) < 0 {
		return big.NewInt(0)
	}

	return ret
}

func (s *RonSandwich) buildFrontRunTx(pool *defi.SwapPairInfo, amountIn *big.Int) (frontRunTx *types.Transaction) {
	options, err := bind.NewKeyedTransactorWithChainID(s.privKey, big.NewInt(rcommon.CHAIN_ID))
	if err != nil {
		log.Warn("RonFi buildFrontRunTx get options failed!", "err", err)
		return nil
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
			log.Warn("RonFi Sandwich buildFrontRunTx NewV2pair failed", "pair", pool.Address, "err", err)
			return
		} else {
			v2PoolInst := uniswapV2.NewV2Pool(s.worker.di, pool.Address, 30, s.appState.Copy())
			if v2PoolInst == nil {
				log.Warn("RonFi Sandwich buildFrontRunTx NewV2Pool failed", "pair", pool.Address)
				return
			}

			amountOut := v2PoolInst.CalculateTokensOutFromTokensIn(pool.TokenIn, amountIn)
			if amountOut == nil {
				log.Warn("RonFi Sandwich buildFrontRunTx CalculateTokensOutFromTokensIn failed", "pair", pool.Address)
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
				log.Warn("RonFi Sandwich buildFrontRunTx Swap failed", "pair", pool.Address, "err", err)
				return
			} else {
				frontRunTx = tx
			}
		}
	} else {
		if inst, err := v3.NewV3pool(pool.Address, s.worker.client); err != nil {
			log.Warn("RonFi Sandwich buildFrontRunTx NewV3Pool failed", "pool", pool.Address, "err", err)
			return
		} else {
			poolInfo := s.worker.di.GetPoolInfo(pool.Address)
			if poolInfo == nil {
				log.Warn("RonFi Sandwich buildFrontRunTx GetPoolInfo failed", "pool", pool.Address)
				return
			} else {
				tickLens := rcommon.GetV3TickLens(poolInfo.Factory)
				if tickLens == rcommon.ZeroAddress {
					log.Warn("RonFi Sandwich buildFrontRunTx GetV3TickLens failed", "pool", pool.Address)
					return
				}

				v3PoolInst := uniswapV3.NewV3Pool(s.worker.di, s.targetTx, pool.Address, tickLens, s.appState.Copy())
				if v3PoolInst == nil {
					log.Warn("RonFi Sandwich buildFrontRunTx NewV3Pool failed", "pool", pool.Address)
					return
				}
				if tx, err := inst.Swap(options, txpool.RonFiSwapV3Address, pool.Dir == 0, amountIn, v3PoolInst.State.SqrtPriceX96, []byte{}); err != nil {
					log.Warn("RonFi Sandwich buildFrontRunTx Swap failed", "pool", pool.Address, "err", err)
					return
				} else {
					frontRunTx = tx
				}
			}
		}
	}

	return
}
