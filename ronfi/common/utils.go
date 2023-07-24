package common

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
	"golang.org/x/crypto/sha3"
	"math/big"
	"strconv"
)

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MinUint64(x, y uint64) uint64 {
	if x < y {
		return x
	}
	return y
}

func MaxUint64(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func AbbrHexString(longHexStr string) string {
	result := make([]byte, 0, 14)
	result = append(result, []byte(longHexStr[2:8])...)
	result = append(result, ".."...)
	result = append(result, []byte(longHexStr[len(longHexStr)-6:])...)

	return string(result)
}

// ToWei decimals to wei.
// Be Careful! Use 'ToWei(1000.0, 18)' for example, MUST NOT use 'ToWei(1000, 18)'.
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func Float2Str(amount float64, precision int) string {
	if precision < 0 {
		return strconv.FormatFloat(amount, 'f', -1, 64)
	} else {
		return strconv.FormatFloat(amount, 'f', precision, 64)
	}
}

// ToFloat panic if token is not in the list of bscTokens[]
func TokenToFloat(amount *big.Int, decimals uint64) float64 {
	return ToFloat(amount, decimals)
}

func ToFloat(amount *big.Int, decimals uint64) float64 {
	if int(decimals) >= len(decimalsToBig) {
		log.Warn("RonFi ToFloat super big token decimals", "decimals", decimals)
		return 0.0
	}
	divisor := decimalsToBig[decimals]
	quotient, _ := new(big.Float).Quo(new(big.Float).SetInt(amount), new(big.Float).SetInt(divisor)).Float64()
	return quotient
}

func ToFloatToken(amount *big.Int, decimals int) string {
	if int(decimals) >= len(decimalsToBig) {
		log.Warn("RonFi ToFloatToken super big token decimals", "decimals", decimals)
		return "0.0"
	}
	divisor := decimalsToBig[decimals]
	quotient, _ := new(big.Float).Quo(new(big.Float).SetInt(amount), new(big.Float).SetInt(divisor)).Float64()
	return strconv.FormatFloat(quotient, 'f', -1, 64)
}

func RawHash(data ...[]byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	for _, d := range data {
		hasher.Write(d)
	}
	hash := hasher.Sum(nil)
	return hash
}

// Big converts an address to a big integer.
func Address2Big(addr common.Address) *big.Int { return new(big.Int).SetBytes(addr.Bytes()) }

func SortTokens(tokenA, tokenB common.Address) (common.Address, common.Address, uint64) {
	if Address2Big(tokenA).Cmp(Address2Big(tokenB)) < 0 {
		return tokenA, tokenB, 0
	} else {
		return tokenB, tokenA, 1
	}
}

func EthBigInt2Float64(num *big.Int) float64 {
	return float64(new(big.Int).Div(num, big.NewInt(params.GWei)).Int64()) / params.GWei
}

func ToLoopId(path []common.Address, poolFee []uint64, tokenFee []uint64) common.Hash {
	size := len(path) * 34
	serialized := make([]byte, 0, size)
	for _, addr := range path {
		serialized = append(serialized, addr.Bytes()...)
	}
	for _, fee := range poolFee {
		serialized = append(serialized, byte(fee))
		serialized = append(serialized, byte(fee>>8))
	}
	for _, fee := range tokenFee {
		serialized = append(serialized, byte(fee))
		serialized = append(serialized, byte(fee>>8))
	}

	hash := RawHash(serialized)
	return common.BytesToHash(hash)
}

func ToKey(path []common.Address) string {
	// take every address' first and last 4 bytes and gaped addresses by '-'
	result := ""
	for i, addr := range path {
		addrBytes := addr.Bytes()
		tmp := make([]byte, 0, 8*len(path))
		tmp = append(tmp, addrBytes[:4]...)
		tmp = append(tmp, addrBytes[16:]...)
		if i == 0 {
			result = hex.EncodeToString(tmp)
		} else {
			result = fmt.Sprintf("%s-%s", result, hex.EncodeToString(tmp))
		}
	}

	return result
}

func IsValidFactory(factory common.Address) bool {
	ok := false

	_, ok = ValidV2FactsMap[factory]
	if !ok {
		_, ok = ValidV3FactsMap[factory]
	}

	return ok
}
