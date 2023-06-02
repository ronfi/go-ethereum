package v3

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ronfi/uniswap/common"
	"math/big"
)

func mostSignificantBit(v *big.Int) int {
	if v.Cmp(big.NewInt(0)) <= 0 {
		log.Warn("RonFi mostSignificantBit: v <= 0")
		return -1
	}

	x := new(big.Int).Set(v)
	r := big.NewInt(0)

	if x.Cmp(common.HexToBigInt("100000000000000000000000000000000")) >= 0 {
		x.Rsh(x, 128)
		r.Add(r, big.NewInt(128))
	}

	if x.Cmp(common.HexToBigInt("10000000000000000")) >= 0 {
		x.Rsh(x, 64)
		r.Add(r, big.NewInt(64))
	}

	if x.Cmp(common.HexToBigInt("100000000")) >= 0 {
		x.Rsh(x, 32)
		r.Add(r, big.NewInt(32))
	}

	if x.Cmp(common.HexToBigInt("10000")) >= 0 {
		x.Rsh(x, 16)
		r.Add(r, big.NewInt(16))
	}

	if x.Cmp(common.HexToBigInt("100")) >= 0 {
		x.Rsh(x, 8)
		r.Add(r, big.NewInt(8))
	}

	if x.Cmp(common.HexToBigInt("10")) >= 0 {
		x.Rsh(x, 4)
		r.Add(r, big.NewInt(4))
	}

	if x.Cmp(common.HexToBigInt("4")) >= 0 {
		x.Rsh(x, 2)
		r.Add(r, big.NewInt(2))
	}

	if x.Cmp(common.HexToBigInt("2")) >= 0 {
		r.Add(r, big.NewInt(1))
	}

	return int(r.Int64())
}

func leastSignificantBit(v *big.Int) int {
	if v.Cmp(big.NewInt(0)) <= 0 {
		log.Warn("RonFi: leastSignificantBit: v <= 0")
		return -1
	}

	x := new(big.Int).Set(v)
	r := big.NewInt(255)

	if new(big.Int).And(x, MaxUint128).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(128))
	} else {
		x.Rsh(x, 128)
	}

	if new(big.Int).And(x, MaxUint64).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(64))
	} else {
		x.Rsh(x, 64)
	}

	if new(big.Int).And(x, MaxUint32).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(32))
	} else {
		x.Rsh(x, 32)
	}

	if new(big.Int).And(x, MaxUint16).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(16))
	} else {
		x.Rsh(x, 16)
	}

	if new(big.Int).And(x, MaxUint8).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(8))
	} else {
		x.Rsh(x, 8)
	}

	if new(big.Int).And(x, big.NewInt(15)).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(4))
	} else {
		x.Rsh(x, 4)
	}

	if new(big.Int).And(x, big.NewInt(3)).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(2))
	} else {
		x.Rsh(x, 2)
	}

	if new(big.Int).And(x, big.NewInt(1)).Cmp(big.NewInt(0)) > 0 {
		r.Sub(r, big.NewInt(1))
	}

	return int(r.Int64())
}
