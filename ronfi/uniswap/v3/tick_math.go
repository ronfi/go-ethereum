package v3

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/ronfi/uniswap/common"
	"math/big"
)

func getSqrtRatioAtTick(tick int) *big.Int {
	absTick := tick
	if absTick < 0 {
		absTick = -tick
	}

	if absTick > MaxTick || absTick < 0 {
		log.Warn("RonFi getSqrtRatioAtTick T")
		return nil
	}

	ratio := big.NewInt(0)
	if (absTick & 0x1) != 0 {
		ratio = common.HexToBigInt("FFFCB933BD6FAD37AA2D162D1A594001")
	} else {
		ratio = common.HexToBigInt("100000000000000000000000000000000")
	}

	if (absTick & 0x2) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FFF97272373D413259A46990580E213A"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x4) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FFF2E50F5F656932EF12357CF3C7FDCC"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x8) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FFE5CACA7E10E4E61C3624EAA0941CD0"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x10) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FFCB9843D60F6159C9DB58835C926644"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x20) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FF973B41FA98C081472E6896DFB254C0"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x40) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FF2EA16466C96A3843EC78B326B52861"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x80) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FE5DEE046A99A2A811C461F1969C3053"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x100) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("FCBE86C7900A88AEDCFFC83B479AA3A4"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x200) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("F987A7253AC413176F2B074CF7815E54"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x400) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("F3392B0822B70005940C7A398E4B70F3"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x800) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("E7159475A2C29B7443B29C7FA6E889D9"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x1000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("D097F3BDFD2022B8845AD8F792AA5825"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x2000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("A9F746462D870FDF8A65DC1F90E061E5"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x4000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("70D869A156D2A1B890BB3DF62BAF32F7"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x8000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("31BE135F97D08FD981231505542FCFA6"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x10000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("9AA508B5B7A84E1C677DE54F3E99BC9"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x20000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("5D6AF8DEDB81196699C329225EE604"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x40000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("2216E584F5FA1EA926041BEDFE98"))
		ratio.Rsh(ratio, 128)
	}
	if (absTick & 0x80000) != 0 {
		ratio.Mul(ratio, common.HexToBigInt("48A170391F7DC42444E8FA2"))
		ratio.Rsh(ratio, 128)
	}

	if tick > 0 {
		ratio = new(big.Int).Div(MaxUint256, ratio)
	}

	shifted := new(big.Int).Rsh(ratio, 32)
	remainder := new(big.Int).And(ratio, big.NewInt(0xffffffff))
	if remainder.Cmp(big.NewInt(0)) != 0 {
		shifted.Add(shifted, big.NewInt(1))
	}

	return shifted
}

func getTickAtSqrtRatio(sqrtPriceX96 *big.Int) (int, bool) {
	if sqrtPriceX96.Cmp(big.NewInt(0)) < 0 || sqrtPriceX96.Cmp(MaxUint160) > 0 {
		log.Warn("RonFi - getTickAtSqrtRatio - not a valid uint160")
		return 0, false
	}

	if sqrtPriceX96.Cmp(MinSqrtRatio) < 0 || sqrtPriceX96.Cmp(MaxSqrtRatio) > 0 {
		return 0, false
	}

	ratio := new(big.Int).Lsh(sqrtPriceX96, 32)
	r := new(big.Int).Set(ratio)
	msb := big.NewInt(0)

	var f *big.Int

	if r.Cmp(common.HexToBigInt("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")) > 0 {
		f = new(big.Int).Lsh(big.NewInt(1), 7)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")) > 0 {
		f = new(big.Int).Lsh(big.NewInt(1), 7)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("FFFFFFFFFFFFFFFF")) > 0 {
		f = new(big.Int).Lsh(big.NewInt(1), 6)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("FFFFFFFF")) == 1 {
		f = new(big.Int).Lsh(big.NewInt(1), 5)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("FFFF")) == 1 {
		f = new(big.Int).Lsh(big.NewInt(1), 4)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("FF")) == 1 {
		f = new(big.Int).Lsh(big.NewInt(1), 3)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("F")) == 1 {
		f = new(big.Int).Lsh(big.NewInt(1), 2)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("3")) == 1 {
		f = new(big.Int).Lsh(big.NewInt(1), 1)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	if r.Cmp(common.HexToBigInt("1")) == 1 {
		f = new(big.Int).Lsh(big.NewInt(1), 1)
	} else {
		f = big.NewInt(0)
	}
	msb.Or(msb, f)

	if msb.Cmp(big.NewInt(128)) >= 0 {
		r = new(big.Int).Rsh(ratio, uint(msb.Uint64())-127)
	} else {
		r = new(big.Int).Lsh(ratio, 127-uint(msb.Uint64()))
	}

	log2 := new(big.Int).Lsh(new(big.Int).Sub(msb, big.NewInt(128)), 64)

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 63))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 62))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 61))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 60))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 59))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 58))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 57))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 56))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 55))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 54))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 53))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 52))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 51))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	r = new(big.Int).Rsh(new(big.Int).Exp(r, big.NewInt(2), nil), 127)
	f = new(big.Int).Rsh(r, 128)
	log2 = new(big.Int).Or(log2, new(big.Int).Lsh(f, 50))
	r = new(big.Int).Rsh(r, uint(f.Uint64()))

	logSqrt1001 := new(big.Int).Mul(log2, common.StrToBigInt("255738958999603826347141"))

	tickLo := new(big.Int).Rsh(new(big.Int).Sub(logSqrt1001, common.StrToBigInt("3402992956809132418596140100660247210")), 128)
	tickHi := new(big.Int).Rsh(new(big.Int).Add(logSqrt1001, common.StrToBigInt("291339464771989622907027621153398088495")), 128)

	var tick *big.Int
	if tickLo.Cmp(tickHi) == 0 {
		tick = tickLo
	} else {
		if getSqrtRatioAtTick(int(tickHi.Int64())).Cmp(sqrtPriceX96) <= 0 {
			tick = tickHi
		} else {
			tick = tickLo
		}
	}

	return int(tick.Int64()), true
}
