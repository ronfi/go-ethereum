package v3

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"math/big"
)

type TickData struct {
	liquidityGross *big.Int
	liquidityNet   *big.Int
}

type V3Pool struct {
	di         *defi.Info
	tx         *types.Transaction
	statedb    *state.StateDB
	tickLens   common.Address
	tickData   map[int]*TickData
	tickBitMap TickBitMap
	Address    common.Address
	PoolInfo   *defi.PoolInfo
	Name       string
	State      *PoolState
}

type StepState struct {
	sqrtPriceStartX96 *big.Int
	sqrtPriceNextX96  *big.Int
	tickNext          int
	initialized       bool
	amountIn          *big.Int
	amountOut         *big.Int
	feeAmount         *big.Int
}

type ComputeState struct {
	amountSpecifiedRemaining *big.Int
	amountCalculated         *big.Int
	sqrtPriceX96             *big.Int
	liquidity                *big.Int
	tick                     int
}

type SwapState struct {
	amount0      *big.Int
	amount1      *big.Int
	sqrtPriceX96 *big.Int
	liquidity    *big.Int
	tick         int
}

type PoolState struct {
	Tick         int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
}

func NewV3Pool(di *defi.Info, tx *types.Transaction, address common.Address, tickLens common.Address, statedb *state.StateDB) *V3Pool {
	if di == nil {
		log.Warn("RonFi NewV3Pool di is nil")
		return nil
	}

	poolInfo := di.GetPoolInfo(address)
	if poolInfo == nil {
		log.Warn("RonFi NewV3Pool PoolInfo is nil")
		return nil
	}

	token0Symbol := di.GetTokenInfo(poolInfo.Token0).Symbol
	token1Symbol := di.GetTokenInfo(poolInfo.Token1).Symbol
	name := fmt.Sprintf("%s (%s-%s V3, %d)", address, token0Symbol, token1Symbol, poolInfo.Fee.Int64())

	tickBitMap := make(TickBitMap)
	tickData := make(map[int]*TickData)
	v3Pool := &V3Pool{
		di:         di,
		tx:         tx,
		statedb:    statedb,
		Address:    address,
		PoolInfo:   poolInfo,
		Name:       name,
		tickLens:   tickLens,
		tickBitMap: tickBitMap,
		tickData:   tickData,
	}

	return v3Pool
}

func toInt24(n uint32) int32 {
	// Check if the 24th bit (sign bit) is set.
	if n&(1<<23) != 0 {
		// If the sign bit is set, extend the sign bit to the left by ORing with 0xFF000000.
		return int32(n | 0xFF000000)
	}
	// If the sign bit is not set, simply convert the value to int32.
	return int32(n)
}

func (p *V3Pool) Dump() {
	log.Info("RonFi V3Pool Dump:", "tx", p.tx.Hash().String(), "pool", p.Address)
	for k, v := range p.tickBitMap {
		log.Info("RonFi V3Pool",
			"tx", p.tx.Hash().String(), "pool", p.Address,
			"wordPos", k,
			"tickBitmap", v)
	}

	for tick, info := range p.tickData {
		log.Info("RonFi V3Pool",
			"tx", p.tx.Hash().String(), "pool", p.Address,
			"tick", tick,
			"liquidityNet", info.liquidityNet,
			"liquidityGross", info.liquidityGross)
	}
}

func (p *V3Pool) updateTickDataAtWord(wordPos int16) {
	if _, ok := p.tickBitMap[wordPos]; ok {
		return
	}

	if p.statedb != nil {
		tickBitmapSt := p.statedb.GetState(p.Address, common.BytesToHash(getStorageSlotIndex(int64(wordPos), big.NewInt(7))))
		singleTickBitmap := new(big.Int).SetBytes(tickBitmapSt.Bytes())
		if singleTickBitmap.Cmp(ZERO) != 0 {
			for i := int64(0); i < 256; i++ {
				if new(big.Int).And(singleTickBitmap, new(big.Int).Lsh(big.NewInt(1), uint(i))).Cmp(ZERO) > 0 {
					populatedTick := toInt24(uint32((toInt24(uint32(int64(wordPos)<<8 + i))) * toInt24(uint32(p.PoolInfo.TickSpacing))))
					singleTickData := p.statedb.GetState(p.Address, common.BytesToHash(getStorageSlotIndex(int64(populatedTick), big.NewInt(6))))
					liquidityNet := new(big.Int).SetBytes(singleTickData.Bytes()[:16])
					if liquidityNet.BitLen() == 128 && liquidityNet.Bit(127) == 1 {
						// Create a 128-bit mask
						mask := new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil)

						// Perform two's complement to get the negative value
						liquidityNet.Sub(mask, liquidityNet)
						liquidityNet.Neg(liquidityNet)
					}
					liquidityGross := new(big.Int).SetBytes(singleTickData.Bytes()[16:])
					p.tickData[int(populatedTick)] = &TickData{
						liquidityGross,
						liquidityNet,
					}
				}
			}

			//singleTickBitmapCntr, err := p.di.GetV3TickBitMap(p.Address, wordPos)
			//if err == nil && singleTickBitmap.Cmp(singleTickBitmapCntr) != 0 {
			//	log.Info("RonFi updateTickDataAtWord, different singleTickBitmap",
			//		"pool", p.Address,
			//		"wordPos", wordPos,
			//		"singleTickBitmap", singleTickBitmap,
			//		"singleTickBitmapCntr", singleTickBitmapCntr)
			//}
			//if singleTickData, err := p.di.GetV3PopulatedTicksInWord(p.tickLens, p.Address, wordPos); err != nil {
			//	log.Warn("RonFi updateTickDataAtWord", "pool", p.Address, "err", err)
			//	return
			//} else {
			//	for _, data := range singleTickData {
			//		org, ok := p.tickData[int(data.Tick.Int64())]
			//		if !ok || org.liquidityNet.Cmp(data.LiquidityNet) != 0 ||
			//			org.liquidityGross.Cmp(data.LiquidityGross) != 0 {
			//			log.Warn("RonFi updateTickDataAtWord, different tick data",
			//				"singleTickBitmap", singleTickBitmap,
			//				"singleTickBitmapCntr", singleTickBitmapCntr,
			//				"tick", data.Tick.Int64(),
			//				"org.liquidityNet", org.liquidityNet,
			//				"data.LiquidityNet", data.LiquidityNet,
			//				"org.liquidityGross", org.liquidityGross,
			//				"data.LiquidityGross", data.LiquidityGross)
			//		}
			//
			//	}
			//}

			p.tickBitMap[wordPos] = singleTickBitmap
		}
	} else {
		if singleTickBitmap, err := p.di.GetV3TickBitMap(p.Address, wordPos); err != nil {
			log.Warn("RonFi updateTickDataAtWord", "pool", p.Address, "err", err)
			return
		} else {
			if singleTickBitmap == nil {
				return
			}

			if singleTickBitmap.Cmp(ZERO) != 0 {
				if singleTickData, err := p.di.GetV3PopulatedTicksInWord(p.tickLens, p.Address, wordPos); err != nil {
					log.Warn("RonFi updateTickDataAtWord", "pool", p.Address, "err", err)
					return
				} else {
					for _, data := range singleTickData {
						p.tickData[int(data.Tick.Int64())] = &TickData{
							liquidityGross: data.LiquidityGross,
							liquidityNet:   data.LiquidityNet,
						}
					}
				}
			}

			p.tickBitMap[wordPos] = singleTickBitmap
		}
	}
}

type V3PoolState struct {
	Tick         int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
}

func (p *V3Pool) UpdatePoolState(v3States *V3PoolState) bool {
	//defer func() {
	//	fmt.Printf("Name: %v\n", p.Name)
	//	fmt.Printf("Token0: %v\n", p.PoolInfo.Token0)
	//	fmt.Printf("Token1: %v\n", p.PoolInfo.Token1)
	//	fmt.Printf("Liquidity: %v\n", p.State.Liquidity)
	//	fmt.Printf("SqrtPriceX96: %v\n", p.State.SqrtPriceX96)
	//	fmt.Printf("Tick: %v\n", p.State.Tick)
	//}()

	if v3States != nil {
		p.State = &PoolState{
			Tick:         v3States.Tick,
			SqrtPriceX96: v3States.SqrtPriceX96,
			Liquidity:    v3States.Liquidity,
		}
		return true
	} else {
		var err error
		p.PoolInfo = p.di.GetPoolInfo(p.Address)
		liquidity, err := p.di.GetV3CurrentLiquidity(p.Address)
		if err != nil {
			log.Warn("RonFi updatePoolState", "address", p.Address, "err", err)
			return false
		}

		sqrtPriceX96, err := p.di.GetV3SqrtPriceX96(p.Address)
		if err != nil {
			log.Warn("RonFi updatePoolState", "address", p.Address, "err", err)
			return false
		}
		tick, err := p.di.GetV3CurrentTick(p.Address)
		if err != nil {
			log.Warn("RonFi updatePoolState", "address", p.Address, "err", err)
			return false
		}

		p.State = &PoolState{
			Tick:         tick,
			SqrtPriceX96: sqrtPriceX96,
			Liquidity:    liquidity,
		}

		return true
	}
}

func (p *V3Pool) uniswapV3PoolSwap(
	zeroForOne bool,
	amountSpecified *big.Int,
	sqrtPriceLimitX96 *big.Int) *SwapState {
	if amountSpecified.Cmp(ZERO) == 0 {
		//log.Warn("RonFi uniswapV3PoolSwap", "err", "AS!")
		return nil
	}

	liquidity := new(big.Int).Set(p.State.Liquidity)
	sqrtPriceX96 := new(big.Int).Set(p.State.SqrtPriceX96)
	tick := p.State.Tick

	if zeroForOne {
		if !(sqrtPriceLimitX96.Cmp(sqrtPriceX96) < 0 && sqrtPriceLimitX96.Cmp(MinSqrtRatio) > 0) {
			log.Warn("RonFi uniswapV3PoolSwap", "err", "SPL")
			return nil
		}
	} else {
		if !(sqrtPriceLimitX96.Cmp(sqrtPriceX96) > 0 && sqrtPriceLimitX96.Cmp(MaxSqrtRatio) < 0) {
			log.Warn("RonFi uniswapV3PoolSwap", "err", "SPL")
			return nil
		}
	}

	exactInput := false
	if amountSpecified.Cmp(ZERO) > 0 {
		exactInput = true
	}

	swapState := &ComputeState{
		amountSpecifiedRemaining: amountSpecified,
		amountCalculated:         big.NewInt(0),
		sqrtPriceX96:             sqrtPriceX96,
		tick:                     tick,
		liquidity:                liquidity,
	}

	MaxCount := 20
	MaxTickCount := 50
	count := 0
	tickCount := 0
	for swapState.amountSpecifiedRemaining.Cmp(ZERO) != 0 && swapState.sqrtPriceX96.Cmp(sqrtPriceLimitX96) != 0 && count < MaxCount {
		count++
		step := &StepState{
			sqrtPriceStartX96: big.NewInt(0),
			sqrtPriceNextX96:  big.NewInt(0),
			tickNext:          0,
			initialized:       false,
			amountIn:          big.NewInt(0),
			amountOut:         big.NewInt(0),
			feeAmount:         big.NewInt(0),
		}
		step.sqrtPriceStartX96 = swapState.sqrtPriceX96

		for tickCount < MaxTickCount {
			tickCount++
			wordPos, nextTick, initializedStatus, ok := nextInitializedTickWithinOneWord(p.tickBitMap, swapState.tick, p.PoolInfo.TickSpacing, zeroForOne)
			if ok {
				step.tickNext = nextTick
				step.initialized = initializedStatus

				tickNextWord, _ := position(step.tickNext)
				if _, ok := p.tickBitMap[tickNextWord]; !ok {
					p.updateTickDataAtWord(tickNextWord)
				}

				break
			} else {
				p.updateTickDataAtWord(wordPos)
			}
		}

		if step.tickNext < MinTick {
			step.tickNext = MinTick
		}
		if step.tickNext > MaxTick {
			step.tickNext = MaxTick
		}

		step.sqrtPriceNextX96 = getSqrtRatioAtTick(step.tickNext)
		sqrtRatioTargetX96 := step.sqrtPriceNextX96
		if (zeroForOne && step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) < 0) ||
			!zeroForOne && step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) > 0 {
			sqrtRatioTargetX96 = sqrtPriceLimitX96
		}

		if swapState.sqrtPriceX96 == nil ||
			sqrtRatioTargetX96 == nil ||
			swapState.liquidity == nil ||
			swapState.amountSpecifiedRemaining == nil {
			break
		}

		compRes := computeSwapStep(
			swapState.sqrtPriceX96,
			sqrtRatioTargetX96,
			swapState.liquidity,
			swapState.amountSpecifiedRemaining,
			p.PoolInfo.Fee,
		)

		swapState.sqrtPriceX96 = compRes.sqrtRatioNextX96
		step.amountIn = compRes.amountIn
		step.amountOut = compRes.amountOut
		step.feeAmount = compRes.feeAmount

		//fmt.Printf("tick=%v -> tickNext= %v, sqrtPriceX96: %v -> sqrtPriceNextX96: %v, liquidity: %v, amountSpecifiedRemaining: %v, amountIn: %v, amountOut: %v, fee: %v\n",
		//	swapState.tick, step.tickNext, swapState.sqrtPriceX96, step.sqrtPriceNextX96, swapState.liquidity, swapState.amountSpecifiedRemaining, step.amountIn, step.amountOut, p.PoolInfo.Fee)

		if exactInput {
			swapState.amountSpecifiedRemaining = new(big.Int).Sub(
				swapState.amountSpecifiedRemaining,
				new(big.Int).Add(step.amountIn, step.feeAmount),
			)
			swapState.amountCalculated = new(big.Int).Sub(swapState.amountCalculated, step.amountOut)
		} else {
			swapState.amountSpecifiedRemaining = new(big.Int).Add(swapState.amountSpecifiedRemaining, step.amountOut)
			swapState.amountCalculated = new(big.Int).Add(
				new(big.Int).Add(swapState.amountCalculated, step.amountIn),
				step.feeAmount,
			)
		}

		if swapState.sqrtPriceX96.Cmp(step.sqrtPriceNextX96) == 0 {
			if step.initialized {
				liquidityNet := big.NewInt(0)
				if data, ok := p.tickData[step.tickNext]; ok {
					liquidityNet = new(big.Int).Set(data.liquidityNet)
				} else {
					//log.Warn("RonFi uniswapV3PoolSwap", "err", "TBD")
					return nil
				}
				if zeroForOne {
					liquidityNet = new(big.Int).Neg(liquidityNet)
				}

				swapState.liquidity = addDelta(swapState.liquidity, liquidityNet)
			}

			if zeroForOne {
				swapState.tick = step.tickNext - 1
			} else {
				swapState.tick = step.tickNext
			}
		} else if swapState.sqrtPriceX96.Cmp(step.sqrtPriceStartX96) != 0 {
			if tick, ok := getTickAtSqrtRatio(swapState.sqrtPriceX96); ok {
				swapState.tick = tick
			} else {
				log.Warn("RonFi uniswapV3PoolSwap", "err", "GTS")
				return nil
			}
		}
	}

	if count >= MaxCount || tickCount >= MaxTickCount {
		//log.Warn("RonFi uniswapV3PoolSwap count overflow!", "tx", p.tx.Hash().String(), "pool", p.Address, "count", count, "tickCount", tickCount)
		return nil
	}

	amount0 := big.NewInt(0)
	amount1 := big.NewInt(0)
	if zeroForOne == exactInput {
		amount0 = new(big.Int).Sub(amountSpecified, swapState.amountSpecifiedRemaining)
		amount1 = swapState.amountCalculated
	} else {
		amount0 = swapState.amountCalculated
		amount1 = new(big.Int).Sub(amountSpecified, swapState.amountSpecifiedRemaining)
	}

	return &SwapState{
		amount0:      amount0,
		amount1:      amount1,
		sqrtPriceX96: swapState.sqrtPriceX96,
		liquidity:    swapState.liquidity,
		tick:         swapState.tick,
	}
}

func (p *V3Pool) CalculateTokensOutFromTokensIn(tokenIn common.Address, amountIn *big.Int) *big.Int {
	if tokenIn != p.PoolInfo.Token0 && tokenIn != p.PoolInfo.Token1 {
		log.Warn("RonFi calculateTokensOutFromTokensIn", "tokenIn", tokenIn)
		return nil
	}

	zeroForOne := false
	if tokenIn == p.PoolInfo.Token0 {
		zeroForOne = true
	}

	var sqrtPriceLimitX96 *big.Int
	if zeroForOne {
		sqrtPriceLimitX96 = big.NewInt(0).Add(MinSqrtRatio, big.NewInt(1))
	} else {
		sqrtPriceLimitX96 = big.NewInt(0).Sub(MaxSqrtRatio, big.NewInt(1))
	}

	res := p.uniswapV3PoolSwap(zeroForOne, amountIn, sqrtPriceLimitX96)
	if res == nil {
		return new(big.Int).SetUint64(0)
	}
	if zeroForOne {
		return new(big.Int).Neg(res.amount1)
	} else {
		return new(big.Int).Neg(res.amount0)
	}
}
