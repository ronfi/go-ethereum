package uniswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	v2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	v3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
)

type PoolType int

const (
	V2 PoolType = iota
	V3
)

type Pool struct {
	// The Uniswap pool that the cycle is in.
	PoolAddr common.Address
	PoolType PoolType
	PoolFee  int
	TokenFee int // tokenOut fee
	TickLens common.Address
	TokenIn  common.Address
	Dir      int
	GasNeed  uint64
}

type V3ArbPath []*Pool

func (path V3ArbPath) String() string {
	var str string
	for i, pool := range path {
		if i != len(path)-1 {
			if pool.PoolType == V2 {
				str += fmt.Sprintf("%s (v2, pf: %v) -> ", pool.PoolAddr.Hex(), pool.PoolFee)
			} else {
				str += fmt.Sprintf("%s (v3) -> ", pool.PoolAddr.Hex())
			}
		} else {
			if pool.PoolType == V2 {
				str += fmt.Sprintf("%s (v2, pf: %v)", pool.PoolAddr.Hex(), pool.PoolFee)
			} else {
				str += fmt.Sprintf("%s (v3)", pool.PoolAddr.Hex())
			}
		}
	}
	return str
}

func (path V3ArbPath) SumGasNeed() uint64 {
	sum := uint64(0)
	hops := len(path)
	if hops >= len(rcommon.GasRough) {
		sum = uint64(100000 + 70000*(hops-1)) // a rough estimation
	} else {
		sum = rcommon.GasRough[hops]
	}

	for i := 1; i < len(path); i += 2 {
		gasNeed := path[i].GasNeed
		sum += gasNeed
	}

	// add a 300K gas buffer
	sum += 300000
	return sum
}

func FromAddress(
	di *defi.Info,
	tx *types.Transaction,
	statedb *state.StateDB,
	v2AmountIOs map[common.Address]map[string]*big.Int,
	v3AmountIOs map[common.Address]map[string]*v3.DetailOut,
	v2Pools map[common.Address]*v2.Pool,
	v3Pools map[common.Address]*v3.Pool,
	targetToken common.Address,
	targetPool *defi.SwapPairInfo,
	swapPoolAddrs V3ArbPath) *LPCycle {
	swapPools := make([]interface{}, 0, len(swapPoolAddrs))
	tokenFees := make([]int, 0, len(swapPoolAddrs))
	for _, pool := range swapPoolAddrs {
		var ok bool
		if pool.PoolType == V2 {
			var v2Pool *v2.Pool
			v2Pool, ok = v2Pools[pool.PoolAddr]
			if !ok {
				var stateCopy *state.StateDB
				if statedb != nil {
					stateCopy = statedb.Copy()
				}
				v2Pool = v2.NewV2Pool(di, pool.PoolAddr, uint64(pool.PoolFee), stateCopy)
				v2Pools[pool.PoolAddr] = v2Pool
			}

			if v2Pool == nil {
				return nil
			}

			swapPools = append(swapPools, v2Pool)
		} else if pool.PoolType == V3 {
			var v3Pool *v3.Pool
			v3Pool, ok = v3Pools[pool.PoolAddr]
			if !ok {
				var stateCopy *state.StateDB
				if statedb != nil {
					stateCopy = statedb.Copy()
				}
				v3Pool = v3.NewV3Pool(di, tx, pool.PoolAddr, pool.TickLens, stateCopy)
				v3Pools[pool.PoolAddr] = v3Pool
			}

			if v3Pool == nil {
				return nil
			}

			swapPools = append(swapPools, v3Pool)
		}
		tokenFees = append(tokenFees, pool.TokenFee)
	}

	tokenInfo := di.GetTokenInfo(targetToken)
	inputTokenDecimals := tokenInfo.Decimals
	inputTokenSymbol := tokenInfo.Symbol

	return NewUniswapLPCycle(di, tx, targetToken, targetPool, v2AmountIOs, v3AmountIOs, swapPools, tokenFees, inputTokenDecimals, inputTokenSymbol, swapPoolAddrs.SumGasNeed())
}

type LPCycle struct {
	// The Uniswap pool that the cycle is in.
	di                 *defi.Info
	tx                 *types.Transaction
	InputToken         common.Address
	InputTokenDecimals uint64
	InputTokenSymbol   string
	TargetPool         *defi.SwapPairInfo
	Name               string
	LoopId             common.Hash
	SumGasNeed         uint64
	PoolAddresses      []common.Address
	PoolTokens         []*TokensPair
	SwapVectors        []*SwapVector
	pools              []interface{}
	v2AmountIOs        map[common.Address]map[string]*big.Int
	v3AmountIOs        map[common.Address]map[string]*v3.DetailOut
	best               *BestState
}

func (lpCycle *LPCycle) toLoopId() common.Hash {
	size := (len(lpCycle.PoolAddresses) + 1) * 34
	serialized := make([]byte, 0, size)
	serialized = append(serialized, lpCycle.InputToken.Bytes()...)
	for _, addr := range lpCycle.PoolAddresses {
		serialized = append(serialized, addr.Bytes()...)
	}

	hash := rcommon.RawHash(serialized)
	return common.BytesToHash(hash)
}

type TokensPair struct {
	token0 common.Address
	token1 common.Address
}

type SwapVector struct {
	TokenIn    common.Address
	TokenOut   common.Address
	ZeroForOne bool
	PoolType   uint8 //0: v2, 1: v3
	PoolFee    uint64
	TokenFee   uint64
}

type BestState struct {
	init         bool
	strategy     string
	swapAmount   *big.Int
	inputToken   common.Address
	profitAmount *big.Int
	profitToken  common.Address
	pools        []interface{}
	poolAddress  []common.Address
	//poolAmounts  []interface{}
	poolTokens []*TokensPair
}

type ProfitAmount struct {
	Iters      int
	Profitable bool
	SwapAmount *big.Int
	BestProfit *big.Int
	AmountIns  []*big.Int
	AmountOuts []*big.Int
	Boundary   []*big.Int
}

type CycleWithProfit struct {
	Cycle  *LPCycle
	Profit *ProfitAmount
}

type CycleWithProfits []*CycleWithProfit

func (p CycleWithProfits) Len() int {
	return len(p)
}
func (p CycleWithProfits) Less(i, j int) bool {
	return p[i].Profit.BestProfit.Cmp(p[j].Profit.BestProfit) > 0
}
func (p CycleWithProfits) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func NewUniswapLPCycle(
	di *defi.Info,
	tx *types.Transaction,
	inputToken common.Address,
	targetPool *defi.SwapPairInfo,
	v2AmountIOs map[common.Address]map[string]*big.Int,
	v3AmountIOs map[common.Address]map[string]*v3.DetailOut,
	pools []interface{},
	tokenFees []int,
	inputTokenDecimals uint64,
	inputTokenSymbol string,
	sumGasNeed uint64) *LPCycle {
	lpCycle := &LPCycle{
		di:                 di,
		tx:                 tx,
		InputToken:         inputToken,
		InputTokenDecimals: inputTokenDecimals,
		InputTokenSymbol:   inputTokenSymbol,
		TargetPool:         targetPool,
		SumGasNeed:         sumGasNeed,
		pools:              pools,
		v2AmountIOs:        v2AmountIOs,
		v3AmountIOs:        v3AmountIOs,
	}

	poolAddresses := make([]common.Address, 0, len(pools))
	poolTokens := make([]*TokensPair, 0, len(pools))
	for _, pool := range pools {
		switch pool.(type) {
		case *v2.Pool:
			v2Pool := pool.(*v2.Pool)
			poolAddresses = append(poolAddresses, v2Pool.Address)
			poolTokens = append(poolTokens, &TokensPair{
				token0: v2Pool.PairInfo.Token0,
				token1: v2Pool.PairInfo.Token1,
			})
		case *v3.Pool:
			v3Pool := pool.(*v3.Pool)
			poolAddresses = append(poolAddresses, v3Pool.Address)
			poolTokens = append(poolTokens, &TokensPair{
				token0: v3Pool.PoolInfo.Token0,
				token1: v3Pool.PoolInfo.Token1,
			})
		}
	}

	lpCycle.PoolAddresses = poolAddresses
	lpCycle.PoolTokens = poolTokens

	swapVectors := make([]*SwapVector, 0, len(pools))
	var tokenIn, tokenOut common.Address
	for i, pool := range pools {
		zeroForOne := false
		poolType := uint8(0)
		poolFee := uint64(0)
		switch pool.(type) {
		case *v2.Pool:
			v2Pool := pool.(*v2.Pool)
			poolFee = v2Pool.Fee
			if i == 0 {
				if inputToken == v2Pool.PairInfo.Token0 {
					zeroForOne = true
					tokenIn = v2Pool.PairInfo.Token0
					tokenOut = v2Pool.PairInfo.Token1
				} else if inputToken == v2Pool.PairInfo.Token1 {
					zeroForOne = false
					tokenIn = v2Pool.PairInfo.Token1
					tokenOut = v2Pool.PairInfo.Token0
				} else {
					log.Warn("RonFi Token could not be identified!")
					return nil
				}
			} else {
				if tokenOut == v2Pool.PairInfo.Token0 {
					zeroForOne = true
					tokenIn = v2Pool.PairInfo.Token0
					tokenOut = v2Pool.PairInfo.Token1
				} else if tokenOut == v2Pool.PairInfo.Token1 {
					zeroForOne = false
					tokenIn = v2Pool.PairInfo.Token1
					tokenOut = v2Pool.PairInfo.Token0
				} else {
					log.Warn("RonFi Token could not be identified!")
					return nil
				}
			}
		case *v3.Pool:
			poolType = 1
			v3Pool := pool.(*v3.Pool)
			if i == 0 {
				if inputToken == v3Pool.PoolInfo.Token0 {
					zeroForOne = true
					tokenIn = v3Pool.PoolInfo.Token0
					tokenOut = v3Pool.PoolInfo.Token1
				} else if inputToken == v3Pool.PoolInfo.Token1 {
					zeroForOne = false
					tokenIn = v3Pool.PoolInfo.Token1
					tokenOut = v3Pool.PoolInfo.Token0
				} else {
					log.Warn("RonFi Token could not be identified!")
					return nil
				}
			} else {
				if tokenOut == v3Pool.PoolInfo.Token0 {
					zeroForOne = true
					tokenIn = v3Pool.PoolInfo.Token0
					tokenOut = v3Pool.PoolInfo.Token1
				} else if tokenOut == v3Pool.PoolInfo.Token1 {
					zeroForOne = false
					tokenIn = v3Pool.PoolInfo.Token1
					tokenOut = v3Pool.PoolInfo.Token0
				}
			}
		}

		swapVectors = append(swapVectors, &SwapVector{
			TokenIn:    tokenIn,
			TokenOut:   tokenOut,
			ZeroForOne: zeroForOne,
			PoolType:   poolType,
			PoolFee:    poolFee,
			TokenFee:   uint64(tokenFees[i]),
		})
	}

	lpCycle.SwapVectors = swapVectors

	name := ""
	for i, pool := range pools {
		switch pool.(type) {
		case *v2.Pool:
			v2Pool := pool.(*v2.Pool)
			if i == 0 {
				name = v2Pool.Name
			} else {
				name += " -> " + v2Pool.Name
			}
		case *v3.Pool:
			v3Pool := pool.(*v3.Pool)
			if i == 0 {
				name = v3Pool.Name
			} else {
				name += " -> " + v3Pool.Name
			}
		}
	}
	lpCycle.Name = name

	lpCycle.best = &BestState{
		init:         true,
		strategy:     "cycle",
		swapAmount:   new(big.Int),
		inputToken:   inputToken,
		profitAmount: new(big.Int),
		profitToken:  inputToken,
		pools:        pools,
		poolAddress:  poolAddresses,
	}

	lpCycle.LoopId = lpCycle.toLoopId()

	return lpCycle
}

func (lpCycle *LPCycle) Dump() {
	for _, pool := range lpCycle.pools {
		switch pool.(type) {
		case *v3.Pool:
			v3Pool := pool.(*v3.Pool)
			v3Pool.Dump()
		}
	}
}

func (lpCycle *LPCycle) AutoUpdate(v3States map[common.Address]*v3.PoolState) bool {
	//var addr common.Address
	for _, pool := range lpCycle.pools {
		updated := false

		switch pool.(type) {
		case *v2.Pool:
			v2Pool := pool.(*v2.Pool)
			//addr = v2Pool.Address
			updated = v2Pool.UpdateReserves()
			//if updated {
			//	lpCycle.shareAmounts.V2PoolsIOLock.Lock()
			//	v2AmountIO, ok := lpCycle.shareAmounts.V2PoolsAmountIOmap[addr]
			//	if ok && v2AmountIO != nil {
			//		if v2AmountIO.Reserve0.Cmp(v2Pool.State.Reserve0) != 0 {
			//			v2AmountIO.Reserve0 = v2Pool.State.Reserve0
			//			v2AmountIO.AmountIOmap = make(map[string]*big.Int)
			//		}
			//	} else {
			//		amountIOmap := make(map[string]*big.Int)
			//		lpCycle.shareAmounts.V2PoolsAmountIOmap[addr] = &V2AmountsIO{
			//			Reserve0:    v2Pool.State.Reserve0,
			//			AmountIOmap: amountIOmap,
			//		}
			//	}
			//	lpCycle.shareAmounts.V2PoolsIOLock.Unlock()
			//}
		case *v3.Pool:
			v3Pool := pool.(*v3.Pool)
			//addr = v3Pool.Address
			v3State, _ := v3States[v3Pool.Address]
			updated = v3Pool.UpdatePoolState(v3State)
			//if updated {
			//	lpCycle.shareAmounts.V3PoolsIOLock.Lock()
			//	v3AmountIO, ok := lpCycle.shareAmounts.V3PoolsAmountIOmap[addr]
			//	if ok && v3AmountIO != nil {
			//		if v3AmountIO.Liquidity.Cmp(v3Pool.State.Liquidity) != 0 {
			//			v3AmountIO.Liquidity = v3Pool.State.Liquidity
			//			v3AmountIO.AmountIOmap = make(map[string]*v3.DetailOut)
			//		}
			//	} else {
			//		amountIOmap := make(map[string]*v3.DetailOut)
			//		lpCycle.shareAmounts.V3PoolsAmountIOmap[addr] = &V3AmountsIO{
			//			Liquidity:   v3Pool.State.Liquidity,
			//			AmountIOmap: amountIOmap,
			//		}
			//	}
			//	lpCycle.shareAmounts.V3PoolsIOLock.Unlock()
			//}
		}

		if !updated {
			//log.Warn("LPCycle::AutoUpdate Update fail", "pool", addr.HexNoChecksum())
			return false
		}
	}

	return true
}

type V2AmountOut struct {
	poolType  PoolType
	amountIn  *big.Int
	amountOut *big.Int
}

type V3AmountOut struct {
	poolType          PoolType
	zeroForOne        bool
	amountSpecified   *big.Int
	sqrtPriceLimitX96 *big.Int
}

func (lpCycle *LPCycle) CalculateArbitrage() *ProfitAmount {
	for i, pool := range lpCycle.pools {
		switch pool.(type) {
		case *v2.Pool:
			v2Pool := pool.(*v2.Pool)
			if v2Pool.State.Reserve0.Cmp(big.NewInt(1)) < 0 || v2Pool.State.Reserve1.Cmp(big.NewInt(1)) < 0 {
				log.Warn("RonFi V2 pool has no liquidity", "pool", v2Pool.Address)
				return nil
			}
		case *v3.Pool:
			v3Pool := pool.(*v3.Pool)
			if v3Pool.State.Liquidity.BitLen() == 0 {
				if lpCycle.SwapVectors[i].ZeroForOne && v3Pool.State.SqrtPriceX96.Cmp(new(big.Int).Add(v3.MinSqrtRatio, big.NewInt(1))) == 0 {
					log.Warn("RonFi V3 pool has no liquidity", "pool", v3Pool.Address)
					return nil
				}
				if !lpCycle.SwapVectors[i].ZeroForOne && v3Pool.State.SqrtPriceX96.Cmp(new(big.Int).Sub(v3.MaxSqrtRatio, big.NewInt(1))) == 0 {
					log.Warn("RonFi V3 pool has no liquidity", "pool", v3Pool.Address)
					return nil
				}
			}
		}
	}

	arbitrage := func(c interface{}, x *big.Int) (*big.Int, *big.Int, []*big.Int, []*big.Int) {
		maxIn := big.NewInt(0)
		cycle := c.(*LPCycle)

		amountIns := make([]*big.Int, 0, len(cycle.pools))
		amountOuts := make([]*big.Int, 0, len(cycle.pools))
		if x == nil || x.BitLen() == 0 {
			return nil, nil, amountIns, amountOuts
		}

		amountIn := new(big.Int).Set(x)
		var (
			amountOut         *big.Int
			amountInRemaining *big.Int
			profit            *big.Int
		)
		for i, pool := range cycle.pools {
			poolAddress := cycle.PoolAddresses[i]
			tokenIn := cycle.SwapVectors[i].TokenIn
			tokenOut := cycle.SwapVectors[i].TokenOut
			isV3Pool := false
			switch pool.(type) {
			case *v2.Pool:
				v2Pool := pool.(*v2.Pool)
				// check v2PoolsAmountIOmap firstly, directly take the value if having
				reuseOK := false
				isNewMap := true
				dir := 0
				if tokenIn == v2Pool.PairInfo.Token1 {
					dir = 1
				}
				amountIO, ok := cycle.v2AmountIOs[poolAddress]
				if ok {
					isNewMap = false
					key := fmt.Sprintf("%d-%s", dir, amountIn.String())
					exist := false
					if amountOut, exist = amountIO[key]; exist {
						reuseOK = true
					}
				}

				if !reuseOK {
					amountOut = v2Pool.CalculateTokensOutFromTokensIn(tokenIn, amountIn)
					if isNewMap {
						amountIO := make(map[string]*big.Int)
						cycle.v2AmountIOs[poolAddress] = amountIO
					}

					// save this calculation into map
					key := fmt.Sprintf("%d-%s", dir, amountIn.String())
					cycle.v2AmountIOs[poolAddress][key] = amountOut
				}

			case *v3.Pool:
				isV3Pool = true
				v3Pool := pool.(*v3.Pool)
				// check v3PoolsAmountIOmap firstly, directly take the value if having
				reuseOK := false
				isNewMap := true
				dir := 0
				if tokenIn == v3Pool.PoolInfo.Token1 {
					dir = 1
				}
				amountIO, ok := cycle.v3AmountIOs[poolAddress]
				if ok {
					isNewMap = false
					key := fmt.Sprintf("%d-%s", dir, amountIn.String())
					if out, exist := amountIO[key]; exist {
						reuseOK = true
						amountOut = out.AmountOut
						amountInRemaining = out.AmountInRemaining
					}
				}

				if !reuseOK {
					amountOut, amountInRemaining = v3Pool.CalculateTokensOutFromTokensIn(tokenIn, amountIn)
					detailOut := v3.DetailOut{AmountOut: new(big.Int).Set(amountOut), AmountInRemaining: new(big.Int).Set(amountInRemaining), MaxIn: nil}
					if isNewMap {
						amountIOmap := make(map[string]*v3.DetailOut)
						cycle.v3AmountIOs[poolAddress] = amountIOmap
					}
					if amountInRemaining != nil && amountInRemaining.BitLen() != 0 {
						if i == 0 {
							maxIn = new(big.Int).Sub(x, amountInRemaining)
						} else {
							maxOut := amountOut
							if maxOut == nil {
								maxOut = cycle.di.GetTokenBalance(poolAddress, tokenOut)
							}
							maxIn = cycle.CalculateMaxInAmount(maxOut, i)
						}
					}
					// save this calculation into map
					key := fmt.Sprintf("%d-%s", dir, amountIn.String())
					cycle.v3AmountIOs[poolAddress][key] = &detailOut
				}
			}

			if amountOut == nil || amountOut.BitLen() == 0 {
				if !isV3Pool {
					return nil, nil, amountIns, amountOuts
				} else {
					return maxIn, nil, amountIns, amountOuts
				}
			}

			amountIns = append(amountIns, amountIn)
			amountOuts = append(amountOuts, amountOut)
			amountIn = amountOut
		}

		profit = new(big.Int).Neg(new(big.Int).Sub(amountOut, x))
		return maxIn, profit, amountIns, amountOuts
	}

	l := -1
	h := 3
	var minIn, maxIn, lower, upper *big.Int
	epsExp := int64(1e14) // i.e. 10^-4
	switch lpCycle.InputToken {
	case rcommon.BTCB:
		l = -7
		h = 4
		epsExp = 1e13 // i.e. $0.3
	case rcommon.DAI:
	case rcommon.USDT:
	case rcommon.USDC:
		l = -3
		h = 7
		epsExp = 1e17 // i.e. $0.3
	case rcommon.WETH:
		l = -6
		h = 5
		epsExp = 1e15 // i.e. (10^-3)*1800 = $0.18
	}

	eps := big.NewInt(epsExp)
	minIn = new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(int(lpCycle.InputTokenDecimals)+l)), nil)
	maxIn = new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(int(lpCycle.InputTokenDecimals)+h)), nil)

	lower = minIn
	upper = maxIn
	bound := []*big.Int{lower, upper}
	options := &MinOptions{
		MaxIters: 100,
		EPS:      eps,
	}

	swapAmount := big.NewInt(0)
	bestProfit := big.NewInt(0)
	profitable := false

	startTime := mclock.Now()
	if opt, iters, err := minimizeScalar(lpCycle, arbitrage, bound, options); err == nil {
		x := opt.x
		profit := new(big.Int).Neg(opt.fun)
		amountIns := opt.amountIns
		amountOuts := opt.amountOuts

		if profit.Cmp(bestProfit) > 0 {
			swapAmount = x
			bestProfit = profit
		}

		if bestProfit.Cmp(v3.ZERO) > 0 {
			profitable = true
		}

		lpCycle.best.swapAmount = new(big.Int).Set(swapAmount)
		lpCycle.best.profitAmount = new(big.Int).Set(bestProfit)
		log.Warn("RonFi CalculateArbitrage succeed", "hops", len(lpCycle.pools), "tx", lpCycle.tx.Hash(), "loopId", lpCycle.LoopId, "iters", iters, "lower", lower, "upper", upper, "elapsed", mclock.Since(startTime).String())

		return &ProfitAmount{
			Iters:      iters,
			Profitable: profitable,
			SwapAmount: swapAmount,
			BestProfit: bestProfit,
			AmountIns:  amountIns,
			AmountOuts: amountOuts,
			Boundary:   bound,
		}
	} else {
		elapsed := mclock.Since(startTime)
		log.Warn("RonFi CalculateArbitrage fail", "hops", len(lpCycle.pools), "tx", lpCycle.tx.Hash(), "loopId", lpCycle.LoopId, "iters", iters, "lower", lower, "upper", upper, "eps", eps, "elapsed", elapsed.String(), "err", err)
	}

	return nil
}

func (lpCycle *LPCycle) CalculateMaxInAmount(maxAmountOut *big.Int, hops int) (maxIn *big.Int) {
	defer func() {
		if maxIn == nil {
			maxIn = new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(25)), nil)
		}
	}()

	amountIn := big.NewInt(0)
	balance := big.NewInt(0)
	if hops > len(lpCycle.pools)-1 {
		return
	}
	for i := hops; i >= 0; i-- {
		pool := lpCycle.pools[i]
		poolAddress := lpCycle.PoolAddresses[i]
		swapVector := lpCycle.SwapVectors[i]
		if pool == nil || swapVector == nil {
			return
		}

		balance = lpCycle.di.GetTokenBalance(poolAddress, swapVector.TokenOut)
		if balance == nil || balance.BitLen() == 0 {
			return
		}

		amountOut := amountIn
		if i == hops {
			amountOut = maxAmountOut
		}

		if amountOut == nil || amountOut.BitLen() == 0 {
			return
		}

		if amountOut.Cmp(balance) > 0 {
			amountOut = balance
		}

		switch pool.(type) {
		case *v2.Pool:
			v2Pool := pool.(*v2.Pool)
			amountIn = v2Pool.CalculateTokensInFromTokensOut(swapVector.TokenOut, amountOut)
		case *v3.Pool:
			v3Pool := pool.(*v3.Pool)
			amountIn, _ = v3Pool.CalculateTokensInFromTokensOut(swapVector.TokenOut, amountOut)
		}
	}

	if amountIn != nil && amountIn.Cmp(v3.ZERO) > 0 {
		maxIn = amountIn
	}

	return
}
