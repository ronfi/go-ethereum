package v3

import (
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

func mulDiv(a, b, denominator *big.Int) *big.Int {
	if !(a.Cmp(MinUint256) >= 0 && a.Cmp(MaxUint256) <= 0) {
		log.Warn("RonFi mulDiv: a does not fit into uint256")
		return nil
	}

	if !(b.Cmp(MinUint256) >= 0 && b.Cmp(MaxUint256) <= 0) {
		log.Warn("RonFi mulDiv: b does not fit into uint256")
		return nil
	}

	if denominator.Sign() == 0 {
		log.Warn("RonFi mulDiv: DIVISION BY ZERO")
		return nil
	}

	result := new(big.Int).Div(new(big.Int).Mul(a, b), denominator)

	if !(result.Cmp(MinUint256) >= 0 && result.Cmp(MaxUint256) <= 0) {
		log.Warn("RonFi mulDiv: result does not fit into uint256")
		return nil
	}

	return result
}

func mulDivRoundingUp(a, b, denominator *big.Int) *big.Int {
	res := mulDiv(a, b, denominator)
	if res == nil {
		log.Warn("RonFi mulDivRoundingUp: mulDiv returned nil")
		return nil
	}

	if new(big.Int).Mod(new(big.Int).Mul(a, b), denominator).Cmp(big.NewInt(0)) > 0 {
		res = new(big.Int).Add(res, big.NewInt(1))
	}

	return res
}
