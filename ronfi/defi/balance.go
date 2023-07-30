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
	EthLowBalanceWarningAmount = big.NewInt(0).Mul(big.NewInt(50_000_000), big.NewInt(params.GWei)) // 0.05 ETH
	EthLowBalanceUrgentAmount  = big.NewInt(0).Mul(big.NewInt(10_000_000), big.NewInt(params.GWei)) // 0.01 ETH
	ChiLowBalanceWarningAmount = big.NewInt(20000)
)

// ProfitSince calculate (gross, net) profit since last balance
func (b *RonFiBalance) ProfitSince(previous *RonFiBalance, shortFee float64) (float64, float64) {
	cost := (previous.Eth - b.Eth + shortFee) * GetTradingTokenPrice(rcommon.USDC) // cost in USD

	gross := b.ContractEth - previous.ContractEth
	gross += (b.ContractEth - previous.ContractEth) / GetTradingTokenPrice(rcommon.WETH)
	gross += (b.ContractBtc - previous.ContractBtc) / GetTradingTokenPrice(rcommon.BTCB)
	gross *= GetTradingTokenPrice(rcommon.USDC)
	gross += b.ContractUsdx - previous.ContractUsdx // total gross in USD

	return gross, gross - cost
}

func GetDexVolume(tokenVols map[common.Address]float64) float64 {
	ethVol := 0.0
	usdx := 0.0
	for token := range rcommon.BridgeTokens {
		if v, e := tokenVols[token]; e {
			switch token {
			case rcommon.WETH:
				ethVol += v
			case rcommon.BTCB:
				ethVol += v / GetTradingTokenPrice(rcommon.BTCB)
			case rcommon.USDC:
				fallthrough
			case rcommon.USDT:
				fallthrough
			case rcommon.DAI:
				usdx += v
			}
		}
	}

	totalVol := ethVol * GetTradingTokenPrice(rcommon.USDC)
	totalVol += usdx
	return totalVol
}

func GetAmountInUSD(v float64, token common.Address) float64 {
	amount := 0.0
	if _, exist := rcommon.BridgeTokens[token]; !exist {
		return amount
	}

	switch token {
	case rcommon.WETH:
		amount = v * GetTradingTokenPrice(rcommon.USDC)
	case rcommon.BTCB:
		amount = v * (GetTradingTokenPrice(rcommon.USDC) / GetTradingTokenPrice(rcommon.BTCB))
	case rcommon.USDC:
		fallthrough
	case rcommon.USDT:
		fallthrough
	case rcommon.DAI:
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

func (di *Info) ronTokensBalance(token common.Address) *big.Int {
	if di == nil {
		return nil
	} else {
		if balance := di.tokenBalance(token, rcommon.RON_V3_SWAP_ADDRESS); balance == nil {
			log.Error("RonFi GetRonFiTokensBalance on RonSwapV3 failed!")
		} else {
			return balance
		}
	}

	return nil
}

func (di *Info) GetAllBalance(executors []common.Address, v3 bool) RonFiBalance {
	ronBalance := RonFiBalance{}

	totalValue := 0.0
	contractValue := 0.0

	// Eth in All Executors Wallet
	{
		lowBalanceExecutors := 0
		allLowBalanceExecutorsEth := big.NewInt(0)
		allExecutorsEth := big.NewInt(0)
		for _, executor := range executors {
			if balance := di.balance(executor); balance == nil {
				log.Warn("RonFi Get Executor Balance Failed", "Executor", rcommon.AbbrHexString(executor.String()))
			} else {
				if balance.Cmp(EthLowBalanceWarningAmount) < 0 {
					lowBalanceExecutors++
					if balance.Cmp(EthLowBalanceUrgentAmount) < 0 {
						log.Warn("RonFi arb low balance", "executor", rcommon.AbbrHexString(executor.String()), "ETH Balance", toEth(balance))
					}
					allLowBalanceExecutorsEth = new(big.Int).Add(allLowBalanceExecutorsEth, balance)
				}
				allExecutorsEth = new(big.Int).Add(allExecutorsEth, balance)
			}
		}
		ronBalance.Eth = rcommon.ToFloat(allExecutorsEth, 18)
		if lowBalanceExecutors > 0 {
			log.Warn("RonFi arb low balance", "executors", lowBalanceExecutors, "Average ETH Balance", toEth(new(big.Int).Div(allLowBalanceExecutorsEth, big.NewInt(int64(lowBalanceExecutors)))))
		}
	}
	totalValue += ronBalance.Eth

	for token, symbol := range rcommon.TradableTokens {
		if balance := di.ronTokensBalance(token); balance == nil {
			log.Warn("RonFi GetAllBalance GetRonTokensBalance failed", "token", symbol)
		} else {
			balanceInFloat := rcommon.ToFloat(balance, 18)
			switch symbol {
			case "WETH":
				ronBalance.ContractEth = balanceInFloat
			case "BTCB":
				ronBalance.ContractBtc = balanceInFloat
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
		if balance := di.ronTokensBalance(rcommon.CHI); balance == nil {
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
