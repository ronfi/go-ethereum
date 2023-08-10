package v3

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
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

type DetailOut struct {
	AmountOut         *big.Int
	AmountInRemaining *big.Int
	MaxIn             *big.Int
}

type MaxSwapStepState struct {
	maxAmountIn  *big.Int
	stepStateMap *SwapStepState
}

type Pool struct {
	di           *defi.Info
	tx           *types.Transaction
	statedb      *state.StateDB
	tickLens     common.Address
	tickData     map[int]*TickData
	tickBitMap   TickBitMap
	Address      common.Address
	PoolInfo     *defi.PoolInfo
	Name         string
	State        *PoolState
	stepStateMap map[string]*MaxSwapStepState
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

func NewV3Pool(di *defi.Info, tx *types.Transaction, address common.Address, tickLens common.Address, statedb *state.StateDB) *Pool {
	if di == nil {
		log.Warn("RonFi NewV3Pool di is nil")
		return nil
	}

	poolInfo := di.GetPoolInfo(address)
	if poolInfo == nil {
		log.Warn("RonFi NewV3Pool PoolInfo is nil", "address", address)
		return nil
	}

	token0Symbol := di.GetTokenInfo(poolInfo.Token0).Symbol
	token1Symbol := di.GetTokenInfo(poolInfo.Token1).Symbol
	name := fmt.Sprintf("%s (%s-%s V3, %d)", address, token0Symbol, token1Symbol, poolInfo.Fee.Int64())

	tickBitMap := make(TickBitMap)
	tickData := make(map[int]*TickData)
	stepStateMap := make(map[string]*MaxSwapStepState)
	v3Pool := &Pool{
		di:           di,
		tx:           tx,
		statedb:      statedb,
		Address:      address,
		PoolInfo:     poolInfo,
		Name:         name,
		tickLens:     tickLens,
		tickBitMap:   tickBitMap,
		tickData:     tickData,
		stepStateMap: stepStateMap,
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

func (p *Pool) Dump() {
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

func (p *Pool) updateTickBitmapAtWord(wordPos int16) {
	if _, ok := p.tickBitMap[wordPos]; ok {
		//log.Info("updateTickBitmapAtWord matched ok", "wordPos", wordPos, "bitmap", fmt.Sprintf("%x", bitmap))
		return
	}

	if p.statedb != nil {
		tickBitmapSt := p.statedb.GetState(p.Address, common.BytesToHash(getStorageSlotIndex(int64(wordPos), big.NewInt(7))))
		singleTickBitmap := new(big.Int).SetBytes(tickBitmapSt.Bytes())
		p.tickBitMap[wordPos] = singleTickBitmap
		//log.Info("updateTickBitmapAtWord matched nok", "wordPos", wordPos, "bitmap", fmt.Sprintf("%x", singleTickBitmap))
	} else {
		// this is only used for unit test
		log.Warn("RonFi v3 updateTickBitmapAtWord, statedb is nil, use di instead (unit test only)")
		if singleTickBitmap, err := p.di.GetV3TickBitMap(p.Address, wordPos); err != nil {
			log.Warn("RonFi updateTickBitmapAtWord", "pool", p.Address, "err", err)
			return
		} else {
			if singleTickBitmap == nil {
				return
			}
			p.tickBitMap[wordPos] = singleTickBitmap
		}
	}
}

func (p *Pool) getTickData(populatedTick int) *TickData {
	if tickData, ok := p.tickData[populatedTick]; ok {
		//log.Info("getTickData matched ok", "tick", populatedTick)
		return tickData
	}

	if p.statedb != nil {
		singleTickData := p.statedb.GetState(p.Address, common.BytesToHash(getStorageSlotIndex(int64(populatedTick), big.NewInt(6))))
		liquidityNet := new(big.Int).SetBytes(singleTickData.Bytes()[:16])
		if liquidityNet.BitLen() == 128 && liquidityNet.Bit(127) == 1 {
			// Create a 128-bit mask
			mask := new(big.Int).Lsh(big.NewInt(1), 128)

			// Perform two's complement to get the negative value
			liquidityNet.Sub(mask, liquidityNet)
			liquidityNet.Neg(liquidityNet)
		}
		liquidityGross := new(big.Int).SetBytes(singleTickData.Bytes()[16:])
		tickData := &TickData{
			liquidityGross,
			liquidityNet,
		}
		p.tickData[populatedTick] = tickData
		//log.Info("getTickData matched nok", "tick", populatedTick)
		return tickData
	} else {
		// this is only used for unit test
		log.Warn("RonFi v3 getTickData, statedb is nil, use di instead (unit test only)")

		wordPos, _ := position(populatedTick / int(p.PoolInfo.TickSpacing))
		if singleTickData, err := p.di.GetV3PopulatedTicksInWord(p.tickLens, p.Address, wordPos); err != nil {
			log.Warn("RonFi updateTickDataAtWord", "pool", p.Address, "err", err)
			return nil
		} else {
			for _, data := range singleTickData {
				p.tickData[int(data.Tick.Int64())] = &TickData{
					liquidityGross: data.LiquidityGross,
					liquidityNet:   data.LiquidityNet,
				}
			}
		}
		if tickData, ok := p.tickData[populatedTick]; ok {
			return tickData
		} else {
			return nil
		}
	}
}

func (p *Pool) UpdatePoolState(v3State *PoolState) bool {
	if v3State != nil {
		// compare new 'liquidity' with 'old' one, reset maps if not same
		if p.State == nil {
			p.tickData = make(map[int]*TickData)
			p.tickBitMap = make(map[int16]*big.Int)
		}

		p.State = v3State
	} else {
		p.PoolInfo = p.di.GetPoolInfo(p.Address)
		if p.statedb != nil {
			// liquidity slot is 4 for uniswap, 5 for pancake
			slotIndex := uint64(4)
			if p.PoolInfo.Factory == common.HexToAddress("0x0BFbCF9fa4f9C56B0F40a671Ad40E0805A091865") {
				slotIndex = 5
			}
			liquidityBytes := p.statedb.GetState(p.Address, common.BigToHash(new(big.Int).SetUint64(slotIndex))).Bytes()
			liquidity := new(big.Int).SetBytes(liquidityBytes)
			//log.Info("updatePoolState: liquidity", "pool", p.Address, "liquidity", liquidity)
			// we shouldn't return false here, because liquidity == 0 only means no liquidity in current tick
			//if liquidity.BitLen() == 0 {
			//	log.Warn("updatePoolState: liquidity is 0",
			//		"pool", p.Address,
			//		"factory", p.PoolInfo.Factory,
			//		"liquidityBytes", fmt.Sprintf("%x", liquidityBytes),
			//		"liquidity", liquidity,
			//	)
			//	return false
			//}

			slot0Hash := p.statedb.GetState(p.Address, common.BigToHash(new(big.Int).SetUint64(0)))
			if slot0Hash == (common.Hash{}) {
				log.Warn("updatePoolState: slot0Hash is 0", "pool", p.Address)
				return false
			}
			slot0 := slot0Hash.Bytes()
			sqrtPriceX96 := new(big.Int).SetBytes(slot0[12:32])
			tick := int32(binary.BigEndian.Uint32(slot0[9:13])) >> 8

			if p.State == nil {
				p.tickData = make(map[int]*TickData)
				p.tickBitMap = make(map[int16]*big.Int)
			}

			p.State = &PoolState{
				Tick:         int(tick),
				SqrtPriceX96: sqrtPriceX96,
				Liquidity:    liquidity,
			}
		} else {
			// for unit test only
			log.Warn("RonFi v3 updatePoolState, statedb is nil, use di instead (unit test only)")
			var (
				err          error
				liquidity    *big.Int
				sqrtPriceX96 *big.Int
				tick         int
			)
			liquidity, err = p.di.GetV3Liquidity(p.Address)
			if err != nil {
				log.Warn("RonFi updatePoolState", "address", p.Address, "err", err)
				return false
			}

			sqrtPriceX96, err = p.di.GetV3SqrtPriceX96(p.Address)
			if err != nil {
				log.Warn("RonFi updatePoolState", "address", p.Address, "err", err)
				return false
			}

			tick, err = p.di.GetV3Tick(p.Address)
			if err != nil {
				log.Warn("RonFi updatePoolState", "address", p.Address, "err", err)
				return false
			}

			p.State = &PoolState{
				Tick:         tick,
				SqrtPriceX96: sqrtPriceX96,
				Liquidity:    liquidity,
			}
		}
	}

	return true
}

func (p *Pool) uniswapV3PoolSwap(
	zeroForOne bool,
	amountSpecified *big.Int,
	sqrtPriceLimitX96 *big.Int) (swapState *SwapState, amountSpecifiedRemaining *big.Int) {
	if amountSpecified.BitLen() == 0 {
		//log.Warn("RonFi uniswapV3PoolSwap", "err", "AS!")
		return nil, nil
	}

	liquidity := new(big.Int).Set(p.State.Liquidity)
	sqrtPriceX96 := new(big.Int).Set(p.State.SqrtPriceX96)
	tick := p.State.Tick

	if zeroForOne {
		if !(sqrtPriceLimitX96.Cmp(sqrtPriceX96) < 0 && sqrtPriceLimitX96.Cmp(MinSqrtRatio) > 0) {
			log.Warn("RonFi uniswapV3PoolSwap", "err", "SPL")
			return nil, nil
		}
	} else {
		if !(sqrtPriceLimitX96.Cmp(sqrtPriceX96) > 0 && sqrtPriceLimitX96.Cmp(MaxSqrtRatio) < 0) {
			log.Warn("RonFi uniswapV3PoolSwap", "err", "SPL")
			return nil, nil
		}
	}

	exactInput := false
	if amountSpecified.Cmp(ZERO) > 0 {
		exactInput = true
	}

	computeState := &ComputeState{
		amountSpecifiedRemaining: amountSpecified,
		amountCalculated:         big.NewInt(0),
		sqrtPriceX96:             sqrtPriceX96,
		tick:                     tick,
		liquidity:                liquidity,
	}

	startTime := mclock.Now()
	MaxSteps := 500                    // If a Swap need 500 computeSwapStep(), i.e. exhaust 500 Ticks Liquidity, the amountIn must be a huge number!
	MaxContinuousEmptyTickBitmap := 10 // At max, we allow 10 continuous empty TickBitmap words, that's 2560 ticks!
	MaxTickCount := 10
	steps := 0
	tickCount := 0
	continuousEmptyTickBitmap := 0
	for computeState.amountSpecifiedRemaining.BitLen() > 0 && computeState.sqrtPriceX96.Cmp(sqrtPriceLimitX96) != 0 && steps < MaxSteps {
		steps++
		step := &StepState{
			sqrtPriceStartX96: big.NewInt(0),
			sqrtPriceNextX96:  big.NewInt(0),
			tickNext:          0,
			initialized:       false,
			amountIn:          big.NewInt(0),
			amountOut:         big.NewInt(0),
			feeAmount:         big.NewInt(0),
		}
		step.sqrtPriceStartX96 = computeState.sqrtPriceX96

		tickCount = 0
		for tickCount < MaxTickCount {
			tickCount++
			// computeState.tick include TickSpacing
			wordPos, nextTick, initializedStatus, ok := nextInitializedTickWithinOneWord(p.tickBitMap, computeState.tick, p.PoolInfo.TickSpacing, zeroForOne)
			if ok {
				step.tickNext = nextTick // nextTick also include TickSpacing
				step.initialized = initializedStatus
				if initializedStatus {
					continuousEmptyTickBitmap = 0
				} else {
					continuousEmptyTickBitmap++
					//log.Info("uniswapV3PoolSwap", "steps", steps, "wordPos", wordPos, "continuousEmptyTickBitmap", continuousEmptyTickBitmap)
				}
				break
			} else {
				p.updateTickBitmapAtWord(wordPos)
			}
		}
		if tickCount >= MaxTickCount {
			log.Warn("RonFi uniswapV3PoolSwap TickBitmap exception", "tx", p.tx.Hash().String(), "pool", p.Address, "elapsed", mclock.Since(startTime).String())
			return nil, nil
		}

		if step.tickNext < MinTick {
			//fmt.Println("RonFi uniswapV3PoolSwap, MinTick")
			step.tickNext = MinTick
		}
		if step.tickNext > MaxTick {
			//fmt.Println("RonFi uniswapV3PoolSwap, MaxTick")
			step.tickNext = MaxTick
		}
		//log.Info("uniswapV3PoolSwap", "pool", p.Address.HexNoChecksum(), "tickNext", step.tickNext, "initialized", step.initialized, "steps", steps, "tickCount", tickCount, "elapsed", mclock.Since(startTime).String())

		step.sqrtPriceNextX96 = getSqrtRatioAtTick(step.tickNext)
		sqrtRatioTargetX96 := step.sqrtPriceNextX96
		if (zeroForOne && step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) < 0) ||
			!zeroForOne && step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) > 0 {
			sqrtRatioTargetX96 = sqrtPriceLimitX96
		}

		if computeState.sqrtPriceX96 == nil ||
			sqrtRatioTargetX96 == nil ||
			computeState.liquidity == nil ||
			computeState.amountSpecifiedRemaining == nil {
			break
		}

		var (
			compRes          *SwapStepState
			maxSwapStepState *MaxSwapStepState
			orgSqrtPriceX96  *big.Int
			ok               bool
			key              string
		)

		if exactInput {
			key = fmt.Sprintf("%d-%s-%s-%s", computeState.tick, computeState.sqrtPriceX96, sqrtRatioTargetX96, computeState.liquidity)
			reused := false
			if maxSwapStepState, ok = p.stepStateMap[key]; ok {
				if maxSwapStepState.maxAmountIn.Cmp(computeState.amountSpecifiedRemaining) <= 0 {
					reused = true
					compRes = maxSwapStepState.stepStateMap
				} else {
					key = fmt.Sprintf("%d-%s-%s-%s-%s", computeState.tick, computeState.sqrtPriceX96, sqrtRatioTargetX96, computeState.liquidity, computeState.amountSpecifiedRemaining)
					if maxSwapStepState, ok = p.stepStateMap[key]; ok {
						reused = true
						compRes = maxSwapStepState.stepStateMap
					}
				}
			}

			if !reused {
				compRes = computeSwapStep(
					computeState.sqrtPriceX96,
					sqrtRatioTargetX96,
					computeState.liquidity,
					computeState.amountSpecifiedRemaining,
					p.PoolInfo.Fee,
				)
			}
		} else {
			compRes = computeSwapStep(
				computeState.sqrtPriceX96,
				sqrtRatioTargetX96,
				computeState.liquidity,
				computeState.amountSpecifiedRemaining,
				p.PoolInfo.Fee,
			)
		}

		orgSqrtPriceX96 = new(big.Int).Set(computeState.sqrtPriceX96)
		computeState.sqrtPriceX96 = compRes.sqrtRatioNextX96
		step.amountIn = compRes.amountIn
		step.amountOut = compRes.amountOut
		step.feeAmount = compRes.feeAmount
		//log.Info("computeSwapStep", "pool", p.Address.HexNoChecksum(), "steps", steps, "step.amountIn", step.amountIn, "step.amountOut", step.amountOut, "elapsed", mclock.Since(startTime).String())

		//fmt.Printf("tick=%v -> tickNext= %v, sqrtPriceX96: %v -> sqrtPriceNextX96: %v, liquidity: %v, amountSpecifiedRemaining: %v, amountIn: %v, amountOut: %v, fee: %v\n",
		//	computeState.tick, step.tickNext, computeState.sqrtPriceX96, step.sqrtPriceNextX96, computeState.liquidity, computeState.amountSpecifiedRemaining, step.amountIn, step.amountOut, p.PoolInfo.Fee)

		if exactInput {
			computeState.amountSpecifiedRemaining = new(big.Int).Sub(
				computeState.amountSpecifiedRemaining,
				new(big.Int).Add(step.amountIn, step.feeAmount),
			)
			computeState.amountCalculated = new(big.Int).Sub(computeState.amountCalculated, step.amountOut)

			if computeState.amountSpecifiedRemaining.Cmp(ZERO) > 0 {
				key = fmt.Sprintf("%d-%s-%s-%s", computeState.tick, orgSqrtPriceX96, sqrtRatioTargetX96, computeState.liquidity)
				p.stepStateMap[key] = &MaxSwapStepState{
					maxAmountIn:  new(big.Int).Add(step.amountIn, step.feeAmount),
					stepStateMap: compRes,
				}
			} else {
				key = fmt.Sprintf("%d-%s-%s-%s-%s", computeState.tick, orgSqrtPriceX96, sqrtRatioTargetX96, computeState.liquidity, computeState.amountSpecifiedRemaining)
				p.stepStateMap[key] = &MaxSwapStepState{
					maxAmountIn:  computeState.amountSpecifiedRemaining,
					stepStateMap: compRes,
				}
			}
		} else {
			computeState.amountSpecifiedRemaining = new(big.Int).Add(computeState.amountSpecifiedRemaining, step.amountOut)
			computeState.amountCalculated = new(big.Int).Add(
				new(big.Int).Add(computeState.amountCalculated, step.amountIn),
				step.feeAmount,
			)
		}

		if computeState.sqrtPriceX96.Cmp(step.sqrtPriceNextX96) == 0 {
			if step.initialized {
				liquidityNet := big.NewInt(0)
				if data := p.getTickData(step.tickNext); data != nil {
					liquidityNet = new(big.Int).Set(data.liquidityNet)
				} else {
					log.Warn("RonFi uniswapV3PoolSwap getTickData fail", "pool", p.Address, "tick", step.tickNext)
					return nil, nil
				}
				if zeroForOne {
					liquidityNet = new(big.Int).Neg(liquidityNet)
				}

				computeState.liquidity = addDelta(computeState.liquidity, liquidityNet)
			}

			if zeroForOne {
				computeState.tick = step.tickNext - 1
			} else {
				computeState.tick = step.tickNext
			}
		} else if computeState.sqrtPriceX96.Cmp(step.sqrtPriceStartX96) != 0 {
			if tick, ok := getTickAtSqrtRatio(computeState.sqrtPriceX96); ok {
				computeState.tick = tick
			} else {
				log.Warn("RonFi uniswapV3PoolSwap", "err", "GTS")
				return nil, nil
			}
		}

		if continuousEmptyTickBitmap >= MaxContinuousEmptyTickBitmap {
			break
		}
	}

	amountSpecifiedRemaining = new(big.Int).Set(computeState.amountSpecifiedRemaining)
	amountIn := new(big.Int).Sub(amountSpecified, computeState.amountSpecifiedRemaining)
	//log.Info("uniswapV3PoolSwap", "pool", p.Address.HexNoChecksum(), "amountSpecified", amountSpecified, "amountIn", amountIn, "amountOut", computeState.amountCalculated, "steps", steps, "elapsed", mclock.Since(startTime).String())
	if steps >= MaxSteps {
		log.Warn("RonFi uniswapV3PoolSwap max steps reached",
			//"tx", p.tx.Hash().String(),
			"pool", p.Address,
			"steps", steps,
			"amountSpecified", amountSpecified,
			"amountIn", amountIn,
			"amountOut", computeState.amountCalculated,
			"elapsed", mclock.Since(startTime).String())
	}

	amount0 := big.NewInt(0)
	amount1 := big.NewInt(0)
	if zeroForOne == exactInput {
		amount0 = new(big.Int).Sub(amountSpecified, computeState.amountSpecifiedRemaining)
		amount1 = computeState.amountCalculated
	} else {
		amount0 = computeState.amountCalculated
		amount1 = new(big.Int).Sub(amountSpecified, computeState.amountSpecifiedRemaining)
	}

	swapState = &SwapState{
		amount0:      amount0,
		amount1:      amount1,
		sqrtPriceX96: computeState.sqrtPriceX96,
		liquidity:    computeState.liquidity,
		tick:         computeState.tick,
	}
	return
}

func (p *Pool) CalculateTokensOutFromTokensIn(tokenIn common.Address, amountIn *big.Int) (amountOut *big.Int, amountInRemaining *big.Int) {
	if tokenIn != p.PoolInfo.Token0 && tokenIn != p.PoolInfo.Token1 {
		log.Warn("RonFi CalculateTokensOutFromTokensIn", "pool", p.Address, "tokenIn", tokenIn)
		return
	}

	if amountIn == nil || amountIn.Cmp(ZERO) <= 0 {
		log.Warn("RonFi CalculateTokensOutFromTokensIn", "pool", p.Address, "tokenIn", tokenIn)
		return
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

	res, remaining := p.uniswapV3PoolSwap(zeroForOne, amountIn, sqrtPriceLimitX96)
	if res == nil {
		return big.NewInt(0), big.NewInt(0)
	}
	if zeroForOne {
		amountOut = new(big.Int).Neg(res.amount1)
	} else {
		amountOut = new(big.Int).Neg(res.amount0)
	}
	amountInRemaining = remaining
	return
}

func (p *Pool) CalculateTokensInFromTokensOut(tokenOut common.Address, amountOut *big.Int) (amountIn *big.Int, amountOutRemaining *big.Int) {
	if tokenOut != p.PoolInfo.Token0 && tokenOut != p.PoolInfo.Token1 {
		log.Warn("RonFi CalculateTokensInFromTokensOut", "tokenOut", tokenOut)
		return
	}

	if amountOut == nil || amountOut.Cmp(ZERO) <= 0 {
		log.Warn("RonFi CalculateTokensInFromTokensOut", "pool", p.Address, "tokenOut", tokenOut)
		return
	}

	zeroForOne := false
	if tokenOut == p.PoolInfo.Token1 {
		zeroForOne = true
	}

	var sqrtPriceLimitX96 *big.Int
	if zeroForOne {
		sqrtPriceLimitX96 = big.NewInt(0).Add(MinSqrtRatio, big.NewInt(1))
	} else {
		sqrtPriceLimitX96 = big.NewInt(0).Sub(MaxSqrtRatio, big.NewInt(1))
	}
	swapState, remaining := p.uniswapV3PoolSwap(zeroForOne, new(big.Int).Neg(amountOut), sqrtPriceLimitX96)

	if zeroForOne {
		amountIn = swapState.amount0
	} else {
		amountIn = swapState.amount1
	}
	amountOutRemaining = remaining

	return
}

func (p *Pool) MaxTokensInFromTokensOut(tokenOut common.Address) (amountIn *big.Int) {
	if tokenOut != p.PoolInfo.Token0 && tokenOut != p.PoolInfo.Token1 {
		log.Warn("RonFi CalculateTokensInFromTokensOut", "tokenOut", tokenOut)
		return
	}

	zeroForOne := false
	if tokenOut == p.PoolInfo.Token1 {
		zeroForOne = true
	}

	var sqrtPriceLimitX96 *big.Int
	if zeroForOne {
		sqrtPriceLimitX96 = big.NewInt(0).Add(MinSqrtRatio, big.NewInt(1))
	} else {
		sqrtPriceLimitX96 = big.NewInt(0).Sub(MaxSqrtRatio, big.NewInt(1))
	}
	swapState, _ := p.uniswapV3PoolSwap(zeroForOne, new(big.Int).Neg(MaxUint256), sqrtPriceLimitX96)

	if zeroForOne {
		amountIn = swapState.amount0
	} else {
		amountIn = swapState.amount1
	}

	return
}
