package v3

import (
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

func addDelta(x, y *big.Int) *big.Int {
	if x.Cmp(big.NewInt(0)) < 0 || x.Cmp(MaxUint128) > 0 {
		log.Warn("RonFi addDelta: x not a valid uint128")
		return nil
	}

	if y.Cmp(MinInt128) < 0 || y.Cmp(MaxInt128) > 0 {
		log.Warn("RonFi addDelta: y not a valid int128")
		return nil
	}

	var z *big.Int
	if y.Cmp(big.NewInt(0)) < 0 {
		yNeg := new(big.Int).Neg(y)
		z = new(big.Int).Sub(x, yNeg)
		if z.Cmp(big.NewInt(0)) < 0 || z.Cmp(MaxUint128) > 0 {
			//log.Warn("RonFi addDelta: LS")
			return nil
		}
	} else {
		z = new(big.Int).Add(x, y)
		if z.Cmp(big.NewInt(0)) < 0 || z.Cmp(MaxUint128) > 0 {
			//log.Warn("RonFi addDelta: LA")
			return nil
		}
	}

	return z
}
