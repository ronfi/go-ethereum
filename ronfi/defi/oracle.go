package defi

import (
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	oracle "github.com/ethereum/go-ethereum/ronfi/contracts/contract_oracle"
)

var (
	EthOracleAddr = common.HexToAddress("0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419")
	BtcOracleAddr = common.HexToAddress("0xf4030086522a5beea4988f8ca5b36dbc97bee88c")

	ethPriceFeed, btcPriceFeed *oracle.ChainlinkOracle
	ethPrice, btcPrice         atomic.Value
)

func InitRonFiOracle(client *ethclient.Client) bool {
	var err error
	retVal := false

	defer func() {
		if !retVal {
			ethPrice.Store(3000.0)
			btcPrice.Store(40000.0)

			log.Error("RonFi Init Price Feed Oracle Failed")
		} else {
			UpdateTokenPrices()
		}
	}()

	if ethPriceFeed, err = oracle.NewChainlinkOracle(EthOracleAddr, client); err != nil {
		log.Error("RonFi InitRonFiOracle fail", "err", err)
		return false
	}

	if btcPriceFeed, err = oracle.NewChainlinkOracle(BtcOracleAddr, client); err != nil {
		log.Error("RonFi InitRonFiOracle fail", "err", err)
		return false
	}

	retVal = true
	return retVal
}

func UpdateTokenPrices() {
	if ethPriceFeed != nil {
		if price, err := ethPriceFeed.LatestAnswer(nil); err == nil {
			priceInFloat := price2Float(price)
			if priceInFloat != 0.0 {
				ethPrice.Store(priceInFloat)
			}
		} else {
			log.Warn("RonFi ethPriceFeed fail", "err", err)
		}
	}

	if btcPriceFeed != nil {
		if price, err := btcPriceFeed.LatestAnswer(nil); err == nil {
			priceInFloat := price2Float(price)
			if priceInFloat != 0.0 {
				btcPrice.Store(priceInFloat)
			}
		} else {
			log.Warn("RonFi btcPriceFeed fail", "err", err)
		}
	}

	log.Info("RonFi PriceOracle update",
		"eth", GetTokenPrice(rcommon.WETH),
		"btc", GetTokenPrice(rcommon.BTCB),
	)
}

// GetTradingTokenPrice return the oracle WETH price in token. For example, 500/4000 for ETH.
func GetTradingTokenPrice(token common.Address) float64 {
	price := 0.0
	switch token {
	case rcommon.WETH:
		price = 1.0

	case rcommon.BTCB:
		if btc := btcPrice.Load(); btc != nil {
			deno := btc.(float64)
			if eth := ethPrice.Load(); eth != nil {
				nume := eth.(float64)
				price = nume / deno
			}
		}

	case rcommon.DAI:
		fallthrough
	case rcommon.USDC:
		fallthrough
	case rcommon.USDT:
		if eth := ethPrice.Load(); eth != nil {
			price = eth.(float64)
		}
	}

	return price
}

func GetTokenPrice(token common.Address) float64 {
	price := 0.0
	switch token {
	case rcommon.WETH:
		price = fLoad(ethPrice)
	case rcommon.BTCB:
		price = fLoad(btcPrice)
	case rcommon.DAI:
		fallthrough
	case rcommon.USDC:
		fallthrough
	case rcommon.USDT:
		fallthrough
	case rcommon.BUSD:
		price = 1.0
	}

	return price
}
