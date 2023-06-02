package v3

import (
	"github.com/ethereum/go-ethereum/ronfi/uniswap/common"
	"math/big"
)

var (
	MinTick      = -887272
	MaxTick      = -MinTick
	MinSqrtRatio = big.NewInt(4295128739)
	MaxSqrtRatio = common.StrToBigInt("1461446703485210103287273052203988822378723970342")

	MinUint8 = big.NewInt(0)
	MaxUint8 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 8), big.NewInt(1))

	MinUint16 = big.NewInt(0)
	MaxUint16 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 16), big.NewInt(1))

	MinUint32 = big.NewInt(0)
	MaxUint32 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 32), big.NewInt(1))

	MinUint64 = big.NewInt(0)
	MaxUint64 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 64), big.NewInt(1))

	MinInt128 = new(big.Int).Neg(new(big.Int).Sub(new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 127), big.NewInt(1)), big.NewInt(1)))
	MaxInt128 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 127), big.NewInt(1))

	MinUint128 = big.NewInt(0)
	MaxUint128 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(1))

	MinUint160 = big.NewInt(0)
	MaxUint160 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 160), big.NewInt(1))

	MinUint256 = big.NewInt(0)
	MaxUint256 = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))

	Q96        = common.HexToBigInt("1000000000000000000000000")
	RESOLUTION = uint(96)

	FEE_BASE = big.NewInt(1000000)
	ZERO     = big.NewInt(0)
)
