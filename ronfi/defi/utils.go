package defi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"math/big"
	"strconv"
	"sync/atomic"
)

func price2Float(val *big.Int) float64 {
	if price, err := strconv.ParseFloat(rcommon.ToFloatToken(val, 8), 64); err != nil {
		return 0.0
	} else {
		return price
	}
}

// fLoad Must Not Load on Nil, Make Sure Initialize it Before Using this
func fLoad(floatAtomic atomic.Value) float64 {
	if floatAtomic.Load() != nil {
		return floatAtomic.Load().(float64)
	} else {
		return 1.0
	}
}

// toEth wei to eth
func toEth(amount *big.Int) string {
	compactAmount := big.NewInt(0)
	reminder := big.NewInt(0)
	divisor := big.NewInt(params.Ether)
	compactAmount.QuoRem(amount, divisor, reminder)
	return fmt.Sprintf("%v.%018s", compactAmount.String(), reminder.String())
}

// check if target pair has a loop in our loops.
func checkIfPairInLoops(pairsDB []PairInfoMap, address common.Address) (*PairInfo, bool) {
	for _, pairsMap := range pairsDB {
		if info, ok := pairsMap[address]; ok {
			return info, ok
		}
	}
	return nil, false
}
