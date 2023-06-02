package uniswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	v2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	v3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math"
	"math/big"
)

type PoolType int

const (
	UniswapV2 PoolType = iota
	UniswapV3
)

type UniswapPool struct {
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

type V3ArbPath []*UniswapPool

func (path V3ArbPath) String() string {
	var str string
	for i, pool := range path {
		if i != len(path)-1 {
			if pool.PoolType == UniswapV2 {
				str += fmt.Sprintf("%s (v2) -> ", pool.PoolAddr.Hex())
			} else {
				str += fmt.Sprintf("%s (v3) -> ", pool.PoolAddr.Hex())
			}
		} else {
			if pool.PoolType == UniswapV2 {
				str += fmt.Sprintf("%s (v2)", pool.PoolAddr.Hex())
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

func FromAddress(di *defi.Info, tx *types.Transaction, statedb *state.StateDB, targetToken common.Address, swapPoolAddrs V3ArbPath) *LPCycle {
	swapPools := make([]interface{}, 0, len(swapPoolAddrs))
	tokenFees := make([]int, 0, len(swapPoolAddrs))
	for _, pool := range swapPoolAddrs {
		if pool.PoolType == UniswapV2 {
			v2Pool := v2.NewV2Pool(di, pool.PoolAddr, uint64(pool.PoolFee), statedb)
			if v2Pool == nil {
				return nil
			}
			swapPools = append(swapPools, v2Pool)
		} else if pool.PoolType == UniswapV3 {
			v3Pool := v3.NewV3Pool(di, tx, pool.PoolAddr, pool.TickLens, statedb)
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

	// maxInput Value set to 100 * 10^inputTokenDecimals
	maxInput := new(big.Int).Mul(big.NewInt(100), big.NewInt(int64(math.Pow10(int(inputTokenDecimals)))))
	switch targetToken {
	case rcommon.USDT:
	case rcommon.USDC:
	case rcommon.DAI:
		maxInput = new(big.Int).Mul(big.NewInt(100000), big.NewInt(int64(math.Pow10(int(inputTokenDecimals)))))
	}

	return NewUniswapLPCycle(tx, targetToken, maxInput, swapPools, tokenFees, inputTokenDecimals, inputTokenSymbol, swapPoolAddrs.SumGasNeed())
}

type LPCycle struct {
	// The Uniswap pool that the cycle is in.
	tx                 *types.Transaction
	InputToken         common.Address
	InputTokenDecimals uint64
	InputTokenSymbol   string
	Name               string
	LoopId             common.Hash
	SumGasNeed         uint64
	MaxIn              *big.Int
	PoolAddresses      []common.Address
	PoolTokens         []*TokensPair
	SwapVectors        []*SwapVector
	pools              []interface{}
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
	Profitable bool
	SwapAmount *big.Int
	BestProfit *big.Int
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

func NewUniswapLPCycle(tx *types.Transaction, inputToken common.Address, maxInput *big.Int, pools []interface{}, tokenFees []int, inputTokenDecimals uint64, inputTokenSymbol string, sumGasNeed uint64) *LPCycle {
	lpCycle := &LPCycle{
		tx:                 tx,
		InputToken:         inputToken,
		InputTokenDecimals: inputTokenDecimals,
		InputTokenSymbol:   inputTokenSymbol,
		SumGasNeed:         sumGasNeed,
		pools:              pools,
	}

	lpCycle.MaxIn = maxInput
	poolAddresses := make([]common.Address, 0, len(pools))
	poolTokens := make([]*TokensPair, 0, len(pools))
	for _, pool := range pools {
		switch pool.(type) {
		case *v2.V2Pool:
			v2Pool := pool.(*v2.V2Pool)
			poolAddresses = append(poolAddresses, v2Pool.Address)
			poolTokens = append(poolTokens, &TokensPair{
				token0: v2Pool.PairInfo.Token0,
				token1: v2Pool.PairInfo.Token1,
			})
		case *v3.V3Pool:
			v3Pool := pool.(*v3.V3Pool)
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
		case *v2.V2Pool:
			v2Pool := pool.(*v2.V2Pool)
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
		case *v3.V3Pool:
			poolType = 1
			v3Pool := pool.(*v3.V3Pool)
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
		case *v2.V2Pool:
			v2Pool := pool.(*v2.V2Pool)
			if i == 0 {
				name = v2Pool.Name
			} else {
				name += " -> " + v2Pool.Name
			}
		case *v3.V3Pool:
			v3Pool := pool.(*v3.V3Pool)
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
		case *v3.V3Pool:
			v3Pool := pool.(*v3.V3Pool)
			v3Pool.Dump()
		}
	}
}

func (lpCycle *LPCycle) AutoUpdate(v3States map[common.Address]*v3.V3PoolState) {
	for _, pool := range lpCycle.pools {
		switch pool.(type) {
		case *v2.V2Pool:
			v2Pool := pool.(*v2.V2Pool)
			v2Pool.UpdateReserves()
		case *v3.V3Pool:
			v3Pool := pool.(*v3.V3Pool)
			if v3State, ok := v3States[v3Pool.Address]; ok && v3State != nil {
				v3Pool.UpdatePoolState(v3State)
			} else {
				v3Pool.UpdatePoolState(nil)
			}
		}
	}
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
		case *v2.V2Pool:
			v2Pool := pool.(*v2.V2Pool)
			if v2Pool.State.Reserve0.Cmp(big.NewInt(1)) < 0 || v2Pool.State.Reserve1.Cmp(big.NewInt(1)) < 0 {
				log.Warn("RonFi V2 pool has no liquidity")
				return nil
			}
		case *v3.V3Pool:
			v3Pool := pool.(*v3.V3Pool)
			if v3Pool.State.Liquidity.Cmp(big.NewInt(0)) == 0 {
				if lpCycle.SwapVectors[i].ZeroForOne && v3Pool.State.SqrtPriceX96.Cmp(new(big.Int).Add(v3.MinSqrtRatio, big.NewInt(1))) == 0 {
					log.Warn("RonFi V3 pool has no liquidity")
					return nil
				}
				if !lpCycle.SwapVectors[i].ZeroForOne && v3Pool.State.SqrtPriceX96.Cmp(new(big.Int).Sub(v3.MaxSqrtRatio, big.NewInt(1))) == 0 {
					log.Warn("RonFi V3 pool has no liquidity")
					return nil
				}
			}
		}
	}

	arbitrage := func(c interface{}, x *big.Int) *big.Int {
		if x == nil || x.Cmp(big.NewInt(0)) == 0 {
			return nil
		}

		cycle := c.(*LPCycle)
		amountIn := new(big.Int).Set(x)
		var amountOut *big.Int
		for i, pool := range cycle.pools {
			tokenIn := cycle.SwapVectors[i].TokenIn
			switch pool.(type) {
			case *v2.V2Pool:
				v2Pool := pool.(*v2.V2Pool)
				amountOut = v2Pool.CalculateTokensOutFromTokensIn(tokenIn, amountIn)
			case *v3.V3Pool:
				v3Pool := pool.(*v3.V3Pool)
				amountOut = v3Pool.CalculateTokensOutFromTokensIn(tokenIn, amountIn)
			}

			if amountOut == nil || amountOut.Cmp(big.NewInt(0)) == 0 {
				return nil
			}
			//fmt.Println("RonFi arbitrage", "amountIn", amountIn, "amountOut", amountOut)
			amountIn = amountOut
		}

		return new(big.Int).Mul(new(big.Int).Sub(amountOut, x), big.NewInt(-1))
	}

	l := -1
	h := 3
	switch lpCycle.InputToken {
	case rcommon.BTCB:
		l = -2
		h = 3
	case rcommon.DAI:
	case rcommon.USDT:
	case rcommon.USDC:
		l = 2
		h = 6
	}

	lower := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(int(lpCycle.InputTokenDecimals)+l)), nil)
	upper := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetInt64(int64(int(lpCycle.InputTokenDecimals)+h)), nil)
	bound := []*big.Int{lower, upper}
	options := &MinOptions{
		MaxIters: 100,
		EPS:      big.NewInt(2),
	}

	swapAmount := big.NewInt(0)
	bestProfit := big.NewInt(0)
	profitable := false

	//swapAmount, _ = new(big.Int).SetString("763955629297960282560", 10)
	//bestProfit = arbitrage(lpCycle, swapAmount)
	//bestProfit = new(big.Int).Mul(bestProfit, big.NewInt(-1))
	//if bestProfit.Cmp(big.NewInt(0)) > 0 {
	//	profitable = true
	//}
	//return &ProfitAmount{
	//	Profitable: profitable,
	//	SwapAmount: swapAmount,
	//	BestProfit: bestProfit,
	//}

	if opt, err := minimizeScalar(lpCycle, arbitrage, bound, options); err == nil {
		x := opt.x
		profit := new(big.Int).Mul(opt.fun, big.NewInt(-1))
		if profit.Cmp(bestProfit) > 0 {
			swapAmount = x
			bestProfit = profit
		}

		if bestProfit.Cmp(big.NewInt(0)) > 0 {
			profitable = true
		}

		lpCycle.best.swapAmount = new(big.Int).Set(swapAmount)
		lpCycle.best.profitAmount = new(big.Int).Set(bestProfit)

		return &ProfitAmount{
			Profitable: profitable,
			SwapAmount: swapAmount,
			BestProfit: bestProfit,
		}
	}

	return nil
}
