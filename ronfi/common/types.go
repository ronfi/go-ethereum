package common

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

// Backend wraps all methods required for arbing.
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *txpool.TxPool
	GetTradingDexTxCh() chan types.Transactions
	//RonBroadcastTransaction(tx *types.Transaction)
}

type TokenDecimals struct {
	Symbol   string
	Decimals int
}

type MysqlConfig struct {
	DbHost string `json:"host"`
	DbPort string `json:"port"`
	DbUser string `json:"user"`
	DbPass string `json:"pass"`
	DbData string `json:"data"`
}

type RedisConfig struct {
	RedisHost string `json:"host"`
	RedisPort string `json:"port"`
}

type DBConfig struct {
	MysqlConf MysqlConfig `json:"mysql"`
	RedisConf RedisConfig `json:"redis"`
}

type AmountIO struct {
	AmountIn  *big.Int
	AmountOut *big.Int
}

func (a *AmountIO) GetAmount() *AmountIO {
	return a
}

func (a *AmountIO) CopyAmount(from IAmountIO) {
	f := from.GetAmount()
	if f.AmountIn != nil {
		a.AmountIn = new(big.Int).Set(f.AmountIn)
	} else {
		a.AmountIn = nil
	}
	if f.AmountOut != nil {
		a.AmountOut = new(big.Int).Set(f.AmountOut)
	} else {
		a.AmountOut = nil
	}
}

func (a *AmountIO) Amount() string {
	if a.AmountIn != nil {
		return ToFloatToken(a.AmountIn, 18)
	} else if a.AmountOut != nil {
		return ToFloatToken(a.AmountOut, 18)
	} else {
		return "nil"
	}
}

type IAmountIO interface {
	GetAmount() *AmountIO
}

type ProfitSorted struct {
	Profit float32
	LoopId common.Hash
}

func (e *ProfitSorted) Less(other *ProfitSorted) bool {
	return e.Profit > other.Profit
}

type JsonRawKnowledge struct {
	Kid       int     `json:"kid"`
	Router    string  `json:"router"`
	Token     string  `json:"token"`
	AmountIn  string  `json:"amountIn"`
	AmountOut string  `json:"amountOut"`
	Profit    float64 `json:"profit"`
	LoopId    string  `json:"loopId"`
}

type RawKnowledge struct {
	Kid    int
	Router common.Address
	Token  []common.Address
	AmountIO
	Profit float32
	LoopId common.Hash
}

func (js *JsonRawKnowledge) ToRawKnowledge() *RawKnowledge {
	var (
		kid       int
		loopId    common.Hash
		router    common.Address
		tokenStr  []string
		token     []common.Address
		aIn, aOut *big.Int
		profit    float32
		ok        bool
	)

	kid = js.Kid

	loopId = common.HexToHash(js.LoopId)

	router = common.HexToAddress(js.Router)
	if e := json.Unmarshal([]byte(js.Token), &tokenStr); e != nil {
		return nil
	}
	for _, addr := range tokenStr {
		token = append(token, common.HexToAddress(addr))
	}

	if aIn, ok = new(big.Int).SetString(js.AmountIn, 10); !ok {
		return nil
	}
	if aOut, ok = new(big.Int).SetString(js.AmountOut, 10); !ok {
		return nil
	}

	amountIO := AmountIO{
		AmountIn:  aIn,
		AmountOut: aOut,
	}

	profit = float32(js.Profit)

	return &RawKnowledge{
		Kid:      kid,
		Router:   router,
		Token:    token,
		AmountIO: amountIO,
		Profit:   profit,
		LoopId:   loopId,
	}
}

func (p *RawKnowledge) ToKey() string {
	var str strings.Builder
	str.WriteString(p.Router.String())
	for _, token := range p.Token {
		str.WriteString("-")
		str.WriteString(token.String())
	}
	return str.String()
}

type JsonNewObs struct {
	RouterMethod string `json:"router"`
}

func (record *JsonNewObs) ToNewObs() *NewObs {
	return &NewObs{
		RouterMethod: record.RouterMethod,
	}
}

type NewObs struct {
	RouterMethod string
}

func (newObs *NewObs) ToJsonNewObs() *JsonNewObs {
	return &JsonNewObs{
		RouterMethod: newObs.RouterMethod,
	}
}

type JsonNewDex struct {
	Router string `json:"router"`
}

type NewDex struct {
	Router common.Address
}

func (record *JsonNewDex) ToNewDex() *NewDex {
	return &NewDex{
		Router: common.HexToAddress(record.Router),
	}
}

func (newDex *NewDex) ToJsonNewDex() *JsonNewDex {
	return &JsonNewDex{
		Router: newDex.Router.String(),
	}
}
