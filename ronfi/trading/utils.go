package trading

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"math"
	"math/big"
	"strconv"
)

func getTokenAmount(di *defi.Info, swapPairsInfo []*defi.SwapPairInfo) (string, float64, common.Address) {
	var tag string
	var token common.Address
	var amountF float64
	var amount *big.Int

	for _, swapPairInfo := range swapPairsInfo {
		tag = "amountIn"
		token = swapPairInfo.TokenIn
		amount = swapPairInfo.AmountIn
		if tInfo := di.GetTokenInfo(swapPairInfo.TokenIn); tInfo != nil {
			amountF = rcommon.ToFloat(amount, tInfo.Decimals)
			return tag, amountF, token
		} else if tInfo := di.GetTokenInfo(swapPairInfo.TokenOut); tInfo != nil {
			tag = "amountOut"
			token = swapPairInfo.TokenOut
			amount = swapPairInfo.AmountOut

			amountF = rcommon.ToFloat(amount, tInfo.Decimals)
			return tag, amountF, token
		}
	}

	return "amountIn", 0.0, rcommon.ZeroAddress
}

func ronFiTxHash(hash common.Hash) common.Hash {
	ronfiBytes := hash.Bytes()
	copy(ronfiBytes, "ronfiarb")

	return common.BytesToHash(ronfiBytes)
}

func prefix6HexUint(number uint64) string {
	return strconv.FormatUint(number>>40, 16)
}

func toPrivateKey(privateKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(privateKey)
}

func privateKeyToAddress(privateKey string) (common.Address, error) {
	var err error
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return common.Address{}, err
	}
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, err
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func calculateProfitLevel(grossProfitInUsd float64) (level int) {
	level = 0
	if grossProfitInUsd > 1000.0 { // 30 copy for $1000 profit, 40 copy for $10000, 50 copy for $100,000, 60 copy for $1M.
		level = int(math.Ceil(10 * math.Log10(grossProfitInUsd)))
	} else if grossProfitInUsd > 250.0 {
		level = 16 // 16 copy for $500 profit
	} else if grossProfitInUsd > 100.0 {
		level = 8 // 8 copy for $100 profit
	} else if grossProfitInUsd > 25.0 {
		level = 4 // 4 copy for $25 profit
	} else if grossProfitInUsd > 10.0 {
		level = 2 // 2 copy for $10 profit
	}
	return
}

func applyTransaction(chain *core.BlockChain, chainConfig *params.ChainConfig, block *types.Block, tx *types.Transaction, txHash common.Hash, state *state.StateDB) (bool, bool, uint64, string) {
	if state == nil {
		return false, false, 0, "state == nil"
	}

	if chain == nil {
		return false, false, 0, "chain == nil"
	}

	gasPool := new(core.GasPool).AddGas(85_000_000)

	author := rcommon.ZeroAddress

	state.SetTxContext(txHash, 0)
	// 'applySuccess' is the status of ApplyTransaction, 'failed' is the status of whether an applied transaction is reverted (i.e. failed but packed into blockchain).
	if applySuccess, failed, gasUsed, err := core.ApplyRonfiTransaction(chainConfig, chain, author, gasPool, state, block.Header(), tx, *chain.GetVMConfig()); applySuccess {
		return applySuccess, failed, gasUsed, err
	} else {
		return applySuccess, failed, gasUsed, err
	}
}

func calculateSandwichProfit(pool *defi.SwapPairInfo, aPairsInfo, bPairsInfo []*defi.SwapPairInfo) *big.Int {
	var (
		aLegAmountIn, bLegAmountOut *big.Int
		profit                      *big.Int
		arbIn, arbOut               *big.Int
	)

	if len(aPairsInfo) == 0 || len(bPairsInfo) == 0 {
		for _, pairInfo := range aPairsInfo {
			if pairInfo.Address == pool.Address && pool.Dir == pairInfo.Dir {
				aLegAmountIn = pairInfo.AmountIn
				break
			}
		}

		if len(bPairsInfo) > 1 {
			// arb loop exists
			// first reversed pair is the b leg
			bLegPair := bPairsInfo[0]
			bLegAmountOut = bLegPair.AmountOut
			if aLegAmountIn != nil && bLegAmountOut != nil {
				profit = big.NewInt(0).Sub(bLegAmountOut, aLegAmountIn)
			}

			// find in and out amount for arb loop
			for _, pairInfo := range bPairsInfo[1:] {
				if pairInfo.Address == pool.Address {
					if pool.Dir == pairInfo.Dir {
						arbIn = pairInfo.AmountIn
					} else {
						arbOut = pairInfo.AmountOut
					}
				}
			}

			if arbIn != nil && arbOut != nil {
				// calculate profit
				arbProfit := big.NewInt(0).Sub(profit, big.NewInt(0).Sub(arbIn, arbOut))
				if profit != nil {
					profit.Add(profit, arbProfit)
				}
			}
		} else {
			for _, pairInfo := range aPairsInfo {
				if pairInfo.Address == pool.Address && pool.Dir != pairInfo.Dir {
					bLegAmountOut = pairInfo.AmountOut
					break
				}
			}

			if aLegAmountIn != nil && bLegAmountOut != nil {
				// calculate profit
				profit = big.NewInt(0).Sub(bLegAmountOut, aLegAmountIn)
			}
		}
	}

	return profit
}
