package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	erc20token "github.com/ethereum/go-ethereum/ronfi/contracts/contract_erc20"
	"math/big"
)

var (
	BnbLowBalanceWarningAmount = big.NewInt(0).Mul(big.NewInt(50_000_000), big.NewInt(params.GWei)) // 0.05 BNB
	BnbLowBalanceUrgentAmount  = big.NewInt(0).Mul(big.NewInt(10_000_000), big.NewInt(params.GWei)) // 0.01 BNB
	ChiLowBalanceWarningAmount = big.NewInt(20000)
)

// ProfitSince calculate (gross, net) profit since last balance
func (b *RonBalance) ProfitSince(previous *RonBalance, shortFee float64) (float64, float64) {
	cost := (previous.Bnb - b.Bnb + shortFee) * GetTradingTokenPrice(rcommon.USDC) // cost in USD

	gross := b.ContractBnb - previous.ContractBnb
	gross += (b.ContractEth - previous.ContractEth) / GetTradingTokenPrice(rcommon.WETH)
	gross += (b.ContractCake - previous.ContractCake) / GetTradingTokenPrice(rcommon.CAKE)
	gross += (b.ContractBtc - previous.ContractBtc) / GetTradingTokenPrice(rcommon.BTCB)
	gross *= GetTradingTokenPrice(rcommon.USDC)
	gross += b.ContractUsdx - previous.ContractUsdx // total gross in USD

	return gross, gross - cost
}

func GetDexVolume(tokenVols map[common.Address]float64) float64 {
	bnbVol := 0.0
	usdx := 0.0
	for token := range rcommon.BridgeTokens {
		if v, e := tokenVols[token]; e {
			switch token {
			case rcommon.WBNB:
				bnbVol += v
			case rcommon.WETH:
				bnbVol += v / GetTradingTokenPrice(rcommon.WETH)
			case rcommon.CAKE:
				bnbVol += v / GetTradingTokenPrice(rcommon.CAKE)
			case rcommon.BTCB:
				bnbVol += v / GetTradingTokenPrice(rcommon.BTCB)
			case rcommon.USDC:
				fallthrough
			case rcommon.USDT:
				fallthrough
			case rcommon.DAI:
				fallthrough
			case rcommon.BUSD:
				usdx += v
			}
		}
	}

	totalVol := bnbVol * GetTradingTokenPrice(rcommon.USDC)
	totalVol += usdx
	return totalVol
}

func GetAmountInUSD(v float64, token common.Address) float64 {
	amount := 0.0
	if _, exist := rcommon.BridgeTokens[token]; !exist {
		return amount
	}

	switch token {
	case rcommon.WBNB:
		amount = v * GetTradingTokenPrice(rcommon.USDC)
	case rcommon.WETH:
		amount = v * (GetTradingTokenPrice(rcommon.USDC) / GetTradingTokenPrice(rcommon.WETH))
	case rcommon.CAKE:
		amount = v * (GetTradingTokenPrice(rcommon.USDC) / GetTradingTokenPrice(rcommon.CAKE))
	case rcommon.BTCB:
		amount = v * (GetTradingTokenPrice(rcommon.USDC) / GetTradingTokenPrice(rcommon.BTCB))
	case rcommon.USDC:
		fallthrough
	case rcommon.USDT:
		fallthrough
	case rcommon.DAI:
		fallthrough
	case rcommon.BUSD:
		amount = v
	}

	return amount
}

func (di *Info) balance(account common.Address) *big.Int {
	if balance, err := di.client.BalanceAt(context.Background(), account, nil); err == nil {
		return balance
	}

	return nil
}

func (di *Info) tokenBalance(token, account common.Address) *big.Int {
	if inst, err := erc20token.NewToken(token, di.client); err == nil {
		if balance, err := inst.BalanceOf(nil, account); err == nil {
			return balance
		}
	}

	return nil
}

func (di *Info) ronTokensBalance(token common.Address, v3 bool) *big.Int {
	if di == nil {
		return nil
	} else {
		if !v3 {
			if balanceX := di.tokenBalance(token, rcommon.RON_SWAP_X_ADDRESS); balanceX == nil {
				log.Error("RonFi GetRonTokensBalance on RonSwapX failed!")
			} else {
				return balanceX
			}
		} else {
			if balance := di.tokenBalance(token, rcommon.RON_V3_SWAP_ADDRESS); balance == nil {
				log.Error("RonFi GetRonTokensBalance on RonSwapV3 failed!")
			} else {
				return balance
			}
		}
	}

	return nil
}

func (di *Info) GetAllBalance(executors []common.Address, v3 bool) RonBalance {
	ronBalance := RonBalance{}

	totalValue := 0.0
	contractValue := 0.0

	// BNB in All Executors Wallet
	{
		lowBalanceExecutors := 0
		allLowBalanceExecutorsBnb := big.NewInt(0)
		allExecutorsBnb := big.NewInt(0)
		for _, executor := range executors {
			if balance := di.balance(executor); balance == nil {
				log.Warn("RonFi Get Executor Balance Failed", "Executor", rcommon.AbbrHexString(executor.String()))
			} else {
				if balance.Cmp(BnbLowBalanceWarningAmount) < 0 {
					lowBalanceExecutors++
					if balance.Cmp(BnbLowBalanceUrgentAmount) < 0 {
						log.Warn("RonFi arb low balance", "executor", rcommon.AbbrHexString(executor.String()), "BNB Balance", toEth(balance))
					}
					allLowBalanceExecutorsBnb = new(big.Int).Add(allLowBalanceExecutorsBnb, balance)
				}
				allExecutorsBnb = new(big.Int).Add(allExecutorsBnb, balance)
			}
		}
		ronBalance.Bnb = rcommon.ToFloat(allExecutorsBnb, 18)
		if lowBalanceExecutors > 0 {
			log.Warn("RonFi arb low balance", "executors", lowBalanceExecutors, "Average BNB Balance", toEth(new(big.Int).Div(allLowBalanceExecutorsBnb, big.NewInt(int64(lowBalanceExecutors)))))
		}
	}
	totalValue += ronBalance.Bnb

	for token, symbol := range rcommon.TradableTokens {
		if balance := di.ronTokensBalance(token, v3); balance == nil {
			log.Warn("RonFi GetAllBalance GetRonTokensBalance failed", "token", symbol)
		} else {
			balanceInFloat := rcommon.ToFloat(balance, 18)
			switch symbol {
			case "WBNB":
				ronBalance.ContractBnb = balanceInFloat
			case "WETH":
				ronBalance.ContractEth = balanceInFloat
			case "BTCB":
				ronBalance.ContractBtc = balanceInFloat
			case "CAKE":
				ronBalance.ContractCake = balanceInFloat
			case "DAI":
				fallthrough
			case "USDC":
				fallthrough
			case "USDT":
				fallthrough
			case "BUSD":
				ronBalance.ContractUsdx += balanceInFloat
			default:
				log.Warn("RonFi new trading token?", "token", token, "token", symbol)
			}
			log.Info("RonFi arb GetAllBalance", "Token", symbol, "Balance", balanceInFloat)

			priceInFloat := GetTradingTokenPrice(token)
			contractValue += balanceInFloat / priceInFloat
		}
	}

	totalValue += contractValue

	totalValueInUsd := totalValue * GetTradingTokenPrice(rcommon.USDC)
	contractValueInUsd := contractValue * GetTradingTokenPrice(rcommon.USDC)

	ronBalance.Total = totalValueInUsd
	ronBalance.ContractTotal = contractValueInUsd

	// Chi Gas Token
	{
		if balance := di.ronTokensBalance(rcommon.CHI, v3); balance == nil {
			log.Warn("RonFi GetAllBalance GetRonTokensBalance failed", "token", "CHI")
		} else {
			ronBalance.ContractChi = balance.Uint64()
			log.Info("RonFi GetAllBalance", "Token", "Chi", "Balance", ronBalance.ContractChi)
			if balance.Cmp(ChiLowBalanceWarningAmount) < 0 {
				log.Warn("RonFi arb low balance", "Chi Balance", ronBalance.ContractChi)
			}
		}
	}

	return ronBalance
}
