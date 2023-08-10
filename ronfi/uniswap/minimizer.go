package uniswap

import (
	"fmt"
	"math/big"
)

type MinOptions struct {
	MaxIters int
	EPS      *big.Int
}

type MinPoint struct {
	x          *big.Int
	fun        *big.Int
	amountIns  []*big.Int
	amountOuts []*big.Int
}

func floatToInt(float *big.Float) *big.Int {
	integer := new(big.Int)
	float.Int(integer)
	return integer
}

func intToFloat(integer *big.Int) *big.Float {
	float := new(big.Float)
	float.SetInt(integer)
	return float
}

func minimizeScalar(
	inst interface{},
	f func(interface{}, *big.Int) (*big.Int, *big.Int, []*big.Int, []*big.Int),
	bounds []*big.Int,
	options *MinOptions,
) (*MinPoint, int, error) {
	// Unpack bounds
	lower, upper := bounds[0], bounds[1]
	a, b := intToFloat(lower), intToFloat(upper)
	var (
		fa, fb, fc, fd, maxB  *big.Int
		amountIns, amountOuts []*big.Int
	)

	_, fa, amountIns, amountOuts = f(inst, floatToInt(a))
	maxB, fb, amountIns, amountOuts = f(inst, floatToInt(b))
	if maxB != nil && maxB.Cmp(lower) > 0 && maxB.Cmp(upper) < 0 {
		b = intToFloat(maxB)
		_, fb, amountIns, amountOuts = f(inst, floatToInt(b))
	}

	if a.Cmp(b) > 0 {
		return nil, 0, fmt.Errorf("lower bound must be less than upper bound")
	}

	if fa == nil || fb == nil {
		return nil, 0, fmt.Errorf("function must be defined on bounds")
	}

	var min *MinPoint
	if fa.Cmp(fb) < 0 {
		min = &MinPoint{floatToInt(a), fa, amountIns, amountOuts}
	} else {
		min = &MinPoint{floatToInt(b), fb, amountIns, amountOuts}
	}

	maxIters := options.MaxIters
	eps := intToFloat(options.EPS)

	GR := new(big.Float).Quo(
		new(big.Float).Add(
			new(big.Float).SetInt64(1),
			new(big.Float).Sqrt(new(big.Float).SetInt64(5))),
		new(big.Float).SetInt64(2))

	gap := new(big.Float).Sub(b, a)
	c := new(big.Float).Sub(b, new(big.Float).Quo(gap, GR))
	d := new(big.Float).Add(a, new(big.Float).Quo(gap, GR))
	_, fc, amountIns, amountOuts = f(inst, floatToInt(c))
	_, fd, amountIns, amountOuts = f(inst, floatToInt(d))
	if fc == nil || fd == nil {
		return nil, 0, fmt.Errorf("function must be defined on bounds")
	}

	iter := 0
	for iter < maxIters && gap.Cmp(eps) > 0 {
		if fc.Cmp(fd) <= 0 {
			b = d
			d = c
			fd = fc
			gap = new(big.Float).Sub(b, a)
			c = new(big.Float).Sub(b, new(big.Float).Quo(gap, GR))
			_, fc, amountIns, amountOuts = f(inst, floatToInt(c))
			//log.Info("minimizeScalar (1)", "iter", iter, "c", floatToInt(c))
			if fc == nil {
				return nil, iter, fmt.Errorf("function must be defined on bounds")
			}
		} else {
			a = c
			c = d
			fc = fd
			gap = new(big.Float).Sub(b, a)
			d = new(big.Float).Add(a, new(big.Float).Quo(gap, GR))
			_, fd, amountIns, amountOuts = f(inst, floatToInt(d))
			//log.Info("minimizeScalar (2)", "iter", iter, "d", floatToInt(d))
			if fd == nil {
				return nil, iter, fmt.Errorf("function must be defined on bounds")
			}
		}

		if floatToInt(c).Cmp(lower) <= 0 {
			c = intToFloat(lower)
			_, fc, amountIns, amountOuts = f(inst, floatToInt(c))
			//log.Info("minimizeScalar (3)", "iter", iter, "c", floatToInt(c))
			if fc == nil {
				return nil, iter, fmt.Errorf("function must be defined on bounds")
			}
		}

		if floatToInt(d).Cmp(upper) >= 0 {
			d = intToFloat(upper)
			_, fd, amountIns, amountOuts = f(inst, floatToInt(d))
			//log.Info("minimizeScalar (4)", "iter", iter, "d", floatToInt(d))
			if fd == nil {
				return nil, iter, fmt.Errorf("function must be defined on bounds")
			}
		}

		iter++
	}

	if fc.Cmp(fd) <= 0 {
		if fc.Cmp(min.fun) < 0 {
			min = &MinPoint{floatToInt(c), fc, amountIns, amountOuts}
		}
	} else {
		if fd.Cmp(min.fun) < 0 {
			min = &MinPoint{floatToInt(d), fd, amountIns, amountOuts}
		}
	}

	return min, iter, nil
}

func gradientApproximation(
	inst interface{},
	f func(interface{}, *big.Rat) *big.Rat,
	x *big.Rat, h *big.Rat) *big.Rat {
	xPlusH := new(big.Rat).Add(x, h)
	xMinusH := new(big.Rat).Sub(x, h)

	fxPlusH := f(inst, xPlusH)
	fxMinusH := f(inst, xMinusH)
	if fxPlusH == nil || fxMinusH == nil {
		return big.NewRat(0, 1)
	}

	numerator := new(big.Rat).Sub(fxPlusH, fxMinusH)
	gradient := new(big.Rat).Quo(numerator, new(big.Rat).Mul(h, big.NewRat(2, 1)))
	//fmt.Printf("xPlusH: %v, xMinusH: %v, fxPlushH: %v, fxMinusH: %v, numerator: %v, gradient: %v\n", xPlusH, xMinusH, fxPlusH, fxMinusH, numerator, gradient)

	return gradient
}

func gradientDescent(
	inst interface{},
	f func(interface{}, *big.Rat) *big.Rat,
	learningRate *big.Rat, initialX *big.Rat, iterations int, h *big.Rat) (*big.Rat, *big.Rat) {
	currentX := new(big.Rat).Set(initialX)

	for i := 0; i < iterations; i++ {
		grad := gradientApproximation(inst, f, currentX, h)
		update := new(big.Rat).Mul(learningRate, grad)
		currentX.Sub(currentX, update)
		//fmt.Printf("grad: %v, update: %v, currentX: %v\n", grad, update, currentX)
	}

	return currentX, f(inst, currentX)
}