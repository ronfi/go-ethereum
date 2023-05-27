package defi

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
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
func (b *RonBalance) ProfitSince(previous *RonBalance, shortFee float64) (float64, float64) {
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

func (di *Info) ronTokensBalance(token common.Address, v3 bool) *big.Int {
	if di == nil {
		return nil
	}

	return nil
}

func (di *Info) GetAllBalance(executors []common.Address, v3 bool) RonBalance {
	ronBalance := RonBalance{}

	return ronBalance
}
