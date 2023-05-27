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
	BnbOracleAddr  = common.HexToAddress("0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE")
	EthOracleAddr  = common.HexToAddress("0x9ef1B8c0E4F7dc8bF5719Ea496883DC6401d5b2e")
	BtcOracleAddr  = common.HexToAddress("0x264990fbd0A4796A3E3d8E37C4d5F87a3aCa5Ebf")
	CakeOracleAddr = common.HexToAddress("0xB6064eD41d4f67e353768aA239cA86f4F73665a1")

	bnbPriceFeed, ethPriceFeed, btcPriceFeed, cakePriceFeed *oracle.ChainlinkOracle
	bnbPrice, ethPrice, btcPrice, cakePrice                 atomic.Value
)

func InitRonFiOracle(client *ethclient.Client) bool {
	var err error
	retVal := false

	defer func() {
		if !retVal {
			bnbPrice.Store(400.0)
			ethPrice.Store(3000.0)
			btcPrice.Store(40000.0)
			cakePrice.Store(7.6)

			log.Error("RonFi Init Price Feed Oracle Failed")
		} else {
			UpdateTokenPrices()
		}
	}()

	if bnbPriceFeed, err = oracle.NewChainlinkOracle(BnbOracleAddr, client); err != nil {
		log.Error("RonFi InitRonFiOracle fail", "err", err)
		return retVal
	}

	if ethPriceFeed, err = oracle.NewChainlinkOracle(EthOracleAddr, client); err != nil {
		log.Error("RonFi InitRonFiOracle fail", "err", err)
		return retVal
	}

	if btcPriceFeed, err = oracle.NewChainlinkOracle(BtcOracleAddr, client); err != nil {
		log.Error("RonFi InitRonFiOracle fail", "err", err)
		return retVal
	}

	if cakePriceFeed, err = oracle.NewChainlinkOracle(CakeOracleAddr, client); err != nil {
		log.Error("RonFi InitRonFiOracle fail", "err", err)
		return retVal
	}

	retVal = true
	return retVal
}

func UpdateTokenPrices() {
	if bnbPriceFeed != nil {
		if price, err := bnbPriceFeed.LatestAnswer(nil); err == nil {
			priceInFloat := price2Float(price)
			if priceInFloat != 0.0 {
				bnbPrice.Store(priceInFloat)
			}
		} else {
			log.Warn("RonFi bnbPriceFeed fail", "err", err)
		}
	}

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

	if cakePriceFeed != nil {
		if price, err := cakePriceFeed.LatestAnswer(nil); err == nil {
			priceInFloat := price2Float(price)
			if priceInFloat != 0.0 {
				cakePrice.Store(priceInFloat)
			}
		} else {
			log.Warn("RonFi cakePriceFeed fail", "err", err)
		}
	}

	log.Info("RonFi PriceOracle update",
		"bnb", GetTokenPrice(rcommon.WBNB),
		"eth", GetTokenPrice(rcommon.WETH),
		"btc", GetTokenPrice(rcommon.BTCB),
		"cake", GetTokenPrice(rcommon.CAKE),
	)
}

// GetTradingTokenPrice return the oracle BNB price in token. For example, 500/4000 for ETH.
func GetTradingTokenPrice(token common.Address) float64 {
	price := 0.0
	switch token {
	case rcommon.WBNB:
		price = 1.0
	case rcommon.WETH:
		if eth := ethPrice.Load(); eth != nil {
			deno := eth.(float64)
			if bnb := bnbPrice.Load(); bnb != nil {
				nume := bnb.(float64)
				price = nume / deno
			}
		}
	case rcommon.BTCB:
		if btc := btcPrice.Load(); btc != nil {
			deno := btc.(float64)
			if bnb := bnbPrice.Load(); bnb != nil {
				nume := bnb.(float64)
				price = nume / deno
			}
		}
	case rcommon.CAKE:
		if cake := cakePrice.Load(); cake != nil {
			deno := cake.(float64)
			if bnb := bnbPrice.Load(); bnb != nil {
				nume := bnb.(float64)
				price = nume / deno
			}
		}
	case rcommon.DAI:
		fallthrough
	case rcommon.USDC:
		fallthrough
	case rcommon.USDT:
		fallthrough
	case rcommon.BUSD:
		if bnb := bnbPrice.Load(); bnb != nil {
			price = bnb.(float64)
		}
	}

	return price
}

func GetTokenPrice(token common.Address) float64 {
	price := 0.0
	switch token {
	case rcommon.WBNB:
		price = fLoad(bnbPrice)
	case rcommon.WETH:
		price = fLoad(ethPrice)
	case rcommon.BTCB:
		price = fLoad(btcPrice)
	case rcommon.CAKE:
		price = fLoad(cakePrice)
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
