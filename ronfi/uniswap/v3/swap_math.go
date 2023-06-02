package v3

import (
	"math/big"
)

type SwapStepState struct {
	sqrtRatioNextX96 *big.Int
	amountIn         *big.Int
	amountOut        *big.Int
	feeAmount        *big.Int
}

func computeSwapStep(
	sqrtRatioCurrentX96 *big.Int,
	sqrtRatioTargetX96 *big.Int,
	liquidity *big.Int,
	amountRemaining *big.Int,
	feePips *big.Int) *SwapStepState {
	var (
		sqrtRatioNextX96 *big.Int
		feeAmount        *big.Int
	)

	zeroForOne := false
	if sqrtRatioCurrentX96.Cmp(sqrtRatioTargetX96) >= 0 {
		zeroForOne = true
	}

	exactIn := false
	if amountRemaining.Cmp(ZERO) >= 0 {
		exactIn = true
	}

	amountIn := big.NewInt(0)
	amountOut := big.NewInt(0)
	negAmountRemaining := new(big.Int).Neg(amountRemaining)

	if exactIn {
		amountRemainingLessFee := mulDiv(amountRemaining, new(big.Int).Sub(FEE_BASE, feePips), FEE_BASE)
		if zeroForOne {
			amountIn = getAmount0Delta(sqrtRatioTargetX96, sqrtRatioCurrentX96, liquidity, true)
		} else {
			amountIn = getAmount1Delta(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, true)
		}

		if amountRemainingLessFee.Cmp(amountIn) >= 0 {
			sqrtRatioNextX96 = sqrtRatioTargetX96
		} else {
			sqrtRatioNextX96 = getNextSqrtPriceFromInput(sqrtRatioCurrentX96, liquidity, amountRemainingLessFee, zeroForOne)
		}
	} else {
		if zeroForOne {
			amountOut = getAmount1Delta(sqrtRatioTargetX96, sqrtRatioCurrentX96, liquidity, false)
		} else {
			amountOut = getAmount0Delta(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, false)
		}

		if negAmountRemaining.Cmp(amountOut) >= 0 {
			sqrtRatioNextX96 = sqrtRatioTargetX96
		} else {
			sqrtRatioNextX96 = getNextSqrtPriceFromOutput(sqrtRatioCurrentX96, liquidity, negAmountRemaining, zeroForOne)
		}
	}

	max := false
	if sqrtRatioTargetX96.Cmp(sqrtRatioNextX96) == 0 {
		max = true
	}

	if zeroForOne {
		if max && exactIn {
			amountIn = getAmount0Delta(sqrtRatioNextX96, sqrtRatioCurrentX96, liquidity, true)
		}

		if !(max && !exactIn) {
			amountOut = getAmount1Delta(sqrtRatioNextX96, sqrtRatioCurrentX96, liquidity, false)
		}
	} else {
		if !(max && exactIn) {
			amountIn = getAmount1Delta(sqrtRatioCurrentX96, sqrtRatioNextX96, liquidity, true)
		}

		if !(max && !exactIn) {
			amountOut = getAmount0Delta(sqrtRatioCurrentX96, sqrtRatioNextX96, liquidity, false)
		}
	}

	if !exactIn && amountOut.Cmp(negAmountRemaining) > 0 {
		amountOut = negAmountRemaining
	}

	if exactIn && sqrtRatioNextX96.Cmp(sqrtRatioTargetX96) != 0 {
		feeAmount = new(big.Int).Sub(amountRemaining, amountIn)
	} else {
		feeAmount = mulDivRoundingUp(amountIn, feePips, new(big.Int).Sub(FEE_BASE, feePips))
	}

	return &SwapStepState{
		sqrtRatioNextX96: sqrtRatioNextX96,
		amountIn:         amountIn,
		amountOut:        amountOut,
		feeAmount:        feeAmount,
	}
}
