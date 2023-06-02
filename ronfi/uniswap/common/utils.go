package common

import (
	"math/big"
)

func StrToBigInt(s string) *big.Int {
	i, _ := big.NewInt(0).SetString(s, 10)
	return i
}

func HexToBigInt(s string) *big.Int {
	i, _ := big.NewInt(0).SetString(s, 16)
	return i
}
