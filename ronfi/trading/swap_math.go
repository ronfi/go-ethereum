package trading

import (
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"math/big"
)

var (
	ThousandOSK = rcommon.ToWei(1000.0, 18)
)

type MaxProfitableAmountIn struct {
	amountIn *big.Int
	profit   *big.Int
}

type FeeAndReserve struct {
	Delta    []float64
	Dataproc float64
	Reserves []defi.Reserve
}

// Calculate amountOut of a single swap
// Note: rDelta = 1 / (1.0 - feeRate), for example 0.3% pool fee => rDelta = 1 / 99.7% = 1.003009027081244
func swapOnce(rDelta float64, amountIn *big.Int, reserve defi.Reserve) (*big.Int, *big.Int) {
	xDelta := big.NewFloat(1.0 / rDelta)
	xDelta.Mul(xDelta, bf(amountIn))
	t1 := bf(reserve.Reserve1)
	t0 := bf(reserve.Reserve0)
	t1.Mul(t1, xDelta)
	t0.Add(t0, xDelta)
	amountOut, _ := t1.Quo(t1, t0).Int(nil)
	xDeltaInt, _ := xDelta.Int(nil)
	newReserveIn := xDeltaInt.Add(xDeltaInt, reserve.Reserve0)
	return amountOut, newReserveIn
}

// Recursively Calculation for C
// Alert(len(rDelta) == len(reserves))
func calculateSwapABC(rDelta []float64, reserves []defi.Reserve) *big.Float {
	size := len(rDelta)
	if size == 2 {
		C := big.NewFloat(rDelta[1])
		C.Mul(C, bf(reserves[1].Reserve0))
		C.Add(C, bf(reserves[0].Reserve1))
		return C
	} else {
		prev := calculateSwapABC(rDelta[:size-1], reserves[:size-1])
		pd3 := big.NewFloat(rDelta[size-1])
		pd3.Mul(pd3, bf(reserves[size-1].Reserve0))

		ynq := bf(reserves[0].Reserve1)
		for i := 1; i < size-1; i++ {
			ynq.Mul(ynq, bf(reserves[i].Reserve1))
		}
		C := ynq.Add(ynq, fmul(pd3, prev))
		return C
	}
}

func calAmountOutAndTestOverflow(amountIn *big.Int, rDelta []float64, reserves []defi.Reserve, isOskLimitedPair []bool) bool {
	aIn := amountIn
	for i := 0; i < len(rDelta); i++ {
		aOut, newReserveIn := swapOnce(rDelta[i], aIn, reserves[i])
		aIn = aOut
		if isOskLimitedPair[i] && reserves[i].Reserve0.Cmp(ThousandOSK) < 0 && newReserveIn.Cmp(ThousandOSK) > 0 {
			// reserve < 1000 before swap, and reserve > 1000 after swap, that's what we cant to detect! OSK limitation to max holding 1000.
			return true
		}
		//if isOskLimitedPair[i] {
		//	log.Info("RonFi debug", "reserveOld", reserves[i].Reserve0, "reserveNew", newReserveIn, "ThousandOSK", ThousandOSK)
		//}
	}
	return false
}

func calMaxProfitableAmountIn(rDelta []float64, rDeltaProd float64, reserves []defi.Reserve, testOverflow bool, isOskLimitedPair []bool) (MaxProfitableAmountIn, bool, bool) {
	if len(rDelta) != len(reserves) || len(rDelta) < 2 {
		return MaxProfitableAmountIn{}, false, false
	}

	for _, reserve := range reserves {
		if len(reserve.Reserve0.Bits()) == 0 || len(reserve.Reserve1.Bits()) == 0 { // isZero
			return MaxProfitableAmountIn{}, false, false
		}
	}

	A := bf(reserves[0].Reserve1)
	B := bf(reserves[0].Reserve0)
	for i := 1; i < len(reserves); i++ {
		A = fmul(A, bf(reserves[i].Reserve1))
		B = fmul(B, bf(reserves[i].Reserve0))
	}
	B = fmul(B, big.NewFloat(rDeltaProd))

	if A.Cmp(B) > 0 {
		C := calculateSwapABC(rDelta, reserves)
		sqrtAB := fsqrt(fmul(A, B))
		B.Sub(sqrtAB, B)
		maxAmountIn := B.Quo(B, C)
		A.Mul(A, maxAmountIn)
		A.Quo(A, sqrtAB)
		profit, _ := A.Sub(A, maxAmountIn).Int(nil)
		amountIn, _ := maxAmountIn.Int(nil)
		overflow := false
		if testOverflow {
			// special care for OSK token, to test whether this amountIn will cause 1000 limitation overflow
			overflow = calAmountOutAndTestOverflow(amountIn, rDelta, reserves, isOskLimitedPair)
		}
		return MaxProfitableAmountIn{amountIn, profit}, true, overflow
	} else {
		return MaxProfitableAmountIn{}, false, false
	}
}

// a group inline functions to make my life easier

func bf(i *big.Int) *big.Float {
	return new(big.Float).SetInt(i)
}
func fadd(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Add(a, b)
}
func fsub(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Sub(a, b)
}
func fmul(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}
func fdiv(a *big.Float, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}
func fsqrt(a *big.Float) *big.Float {
	return new(big.Float).Sqrt(a)
}

//func delta(fee uint64) *big.Float {
//	return big.NewFloat(1.0 - float64(fee)/10000)
//}
