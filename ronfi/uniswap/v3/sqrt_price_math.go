package v3

import (
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

func getAmount0Delta(sqrtRatioAX96 *big.Int, sqrtRatioBX96 *big.Int, liquidity *big.Int, roundUp bool) *big.Int {
	if !(liquidity.Cmp(MinUint128) >= 0 && liquidity.Cmp(MaxUint128) <= 0) {
		log.Warn("RonFi getAmount0Delta: Liquidity does not fit into uint128")
		return nil
	}

	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	numerator1 := new(big.Int).Lsh(liquidity, RESOLUTION)
	numerator2 := new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96)

	if sqrtRatioAX96.Cmp(big.NewInt(0)) <= 0 {
		log.Warn("RonFi getAmount0Delta: require sqrtRatioAX96 > 0")
		return nil
	}

	res := divRoundingUp(mulDivRoundingUp(numerator1, numerator2, sqrtRatioBX96), sqrtRatioAX96)
	if !roundUp {
		res = new(big.Int).Div(mulDiv(numerator1, numerator2, sqrtRatioBX96), sqrtRatioAX96)
	}

	return res
}

func getAmount1Delta(sqrtRatioAX96 *big.Int, sqrtRatioBX96 *big.Int, liquidity *big.Int, roundUp bool) *big.Int {
	if liquidity.Cmp(MinUint128) < 0 || liquidity.Cmp(MaxUint128) > 0 {
		log.Warn("RonFi getAmount1Delta: Liquidity does not fit into uint128")
		return nil
	}

	if sqrtRatioAX96.Cmp(sqrtRatioBX96) > 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	res := mulDivRoundingUp(liquidity, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96), Q96)
	if !roundUp {
		res = mulDiv(liquidity, new(big.Int).Sub(sqrtRatioBX96, sqrtRatioAX96), Q96)
	}

	return res
}

func getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96 *big.Int, liquidity *big.Int, amount *big.Int, add bool) *big.Int {
	if amount.Sign() == 0 {
		return sqrtPX96
	}

	numerator1 := new(big.Int).Lsh(liquidity, RESOLUTION)
	product := new(big.Int).Mul(amount, sqrtPX96)
	if add {
		if new(big.Int).Div(product, amount).Cmp(sqrtPX96) == 0 {
			denominator := new(big.Int).Add(numerator1, product)
			if denominator.Cmp(numerator1) >= 0 {
				return mulDivRoundingUp(numerator1, sqrtPX96, denominator)
			}
		}

		return divRoundingUp(numerator1, new(big.Int).Add(new(big.Int).Div(numerator1, sqrtPX96), amount))
	} else {
		if new(big.Int).Div(product, amount).Cmp(sqrtPX96) != 0 && numerator1.Cmp(product) > 0 {
			log.Warn("RonFi getNextSqrtPriceFromAmount0RoundingUp: product / amount == sqrtPX96 && numerator1 > product")
			return nil
		}

		denominator := new(big.Int).Sub(numerator1, product)
		return mulDivRoundingUp(numerator1, sqrtPX96, denominator)
	}
}

func getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96 *big.Int, liquidity *big.Int, amount *big.Int, add bool) *big.Int {
	quotient := big.NewInt(0)

	if add {
		if amount.Cmp(MaxUint160) <= 0 {
			quotient = new(big.Int).Div(new(big.Int).Lsh(amount, RESOLUTION), liquidity)
		} else {
			quotient = mulDiv(amount, Q96, liquidity)
		}
		return new(big.Int).Add(sqrtPX96, quotient)
	} else {
		if amount.Cmp(MaxUint160) <= 0 {
			quotient = divRoundingUp(new(big.Int).Lsh(amount, RESOLUTION), liquidity)
		} else {
			quotient = mulDivRoundingUp(amount, Q96, liquidity)
		}

		if sqrtPX96.Cmp(quotient) <= 0 {
			log.Warn("RonFi getNextSqrtPriceFromAmount1RoundingDown: require sqrtPX96 > quotient")
			return nil
		}

		return new(big.Int).Sub(sqrtPX96, quotient)
	}
}

func getNextSqrtPriceFromInput(sqrtPX96 *big.Int, liquidity *big.Int, amountIn *big.Int, zeroForOne bool) *big.Int {
	if sqrtPX96.Cmp(MinUint160) <= 0 {
		log.Warn("RonFi getNextSqrtPriceFromInput: require sqrtPX96 > 0")
		return nil
	}

	if liquidity.Cmp(MinUint160) <= 0 {
		log.Warn("RonFi getNextSqrtPriceFromInput: require Liquidity > 0")
		return nil
	}

	if zeroForOne {
		return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountIn, true)
	} else {
		return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountIn, true)
	}
}

func getNextSqrtPriceFromOutput(sqrtPX96 *big.Int, liquidity *big.Int, amountOut *big.Int, zeroForOne bool) *big.Int {
	if sqrtPX96.Cmp(MinUint160) <= 0 {
		log.Warn("RonFi getNextSqrtPriceFromOutput: require sqrtPX96 > 0")
		return nil
	}

	if liquidity.Cmp(MinUint160) <= 0 {
		log.Warn("RonFi getNextSqrtPriceFromOutput: require Liquidity > 0")
		return nil
	}

	if zeroForOne {
		return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountOut, false)
	} else {
		return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountOut, false)
	}
}
