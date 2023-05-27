package common

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"sort"
	"strings"
)

// Backend wraps all methods required for arbing.
type Backend interface {
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	GetTradingDexTxCh() chan types.Transactions
	RonBroadcastTransaction(tx *types.Transaction)
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

type KnowledgeAmount struct {
	Kid    int
	Profit float32
	AmountIO
}

func (in *KnowledgeAmount) Less(other *KnowledgeAmount) bool {
	if other.AmountIn != nil {
		if in.AmountIn != nil {
			if other.AmountIn.Cmp(in.AmountIn) < 0 {
				return false
			}
		} else {
			return true
		}
	} else if other.AmountOut != nil {
		if in.AmountOut != nil {
			if other.AmountOut.Cmp(in.AmountOut) < 0 {
				return false
			}
		} else {
			return true
		}
	}

	return true
}

type KnowledgeAmounts []*KnowledgeAmount

func (klA KnowledgeAmounts) Sort() {
	sort.Slice(klA, func(i, j int) bool {
		return klA[i].Less(klA[j])
	})
}

type AmountProfit struct {
	KAmounts  KnowledgeAmounts
	SumProfit float32
}

func NewAmountProfit(info IAmountIO, profit float32, kid int) *AmountProfit {
	i := info.GetAmount()
	kAmounts := make(KnowledgeAmounts, 0, 10)
	af := &AmountProfit{
		KAmounts:  kAmounts,
		SumProfit: 0.0,
	}
	af.SumProfit = profit

	kAmount := &KnowledgeAmount{
		Kid:    kid,
		Profit: profit,
	}

	// must not take info.amount directly! otherwise, the whole info is hard to recycle.
	if i.AmountIn != nil {
		kAmount.AmountIn = new(big.Int).Set(i.AmountIn)
	} else if i.AmountOut != nil {
		kAmount.AmountOut = new(big.Int).Set(i.AmountOut)
	}

	af.KAmounts = append(af.KAmounts, kAmount)
	return af
}

// EvaluateAmount Evaluate whether this dex tx has an amount bigger than ever hunted
func (in *AmountProfit) EvaluateAmount(param IAmountIO) (greenLight bool) {
	greenLight = false
	if len(in.KAmounts) < 1 {
		return
	}
	minAm := in.KAmounts[0]
	p := param.GetAmount()

	if p.AmountIn != nil && minAm.AmountIn != nil {
		if p.AmountIn.Cmp(minAm.AmountIn) >= 0 {
			greenLight = true
		}
	} else if p.AmountOut != nil && minAm.AmountOut != nil {
		if p.AmountOut.Cmp(minAm.AmountOut) >= 0 {
			greenLight = true
		}
	}
	return
}

func (in *AmountProfit) AddNewKnowledge(info IAmountIO, profit float32, kid int) {
	kAmount := &KnowledgeAmount{
		Kid:    kid,
		Profit: profit,
	}

	i := info.GetAmount()
	if i.AmountIn != nil {
		kAmount.AmountIn = new(big.Int).Set(i.AmountIn)
	} else if i.AmountOut != nil {
		kAmount.AmountOut = new(big.Int).Set(i.AmountOut)
	}

	in.KAmounts = append(in.KAmounts, kAmount)
	in.KAmounts.Sort()
}

type ProfitSorted struct {
	Profit float32
	LoopId common.Hash
}

func (e *ProfitSorted) Less(other *ProfitSorted) bool {
	return e.Profit > other.Profit
}

type LoopIdMapSorted struct {
	LoopIdMap map[common.Hash]*AmountProfit
	Profits   []*ProfitSorted
}

type RonKnowledgeMap map[string]*LoopIdMapSorted

func (ronLidMap RonKnowledgeMap) Copy() RonKnowledgeMap {
	newRonKnowledgeMap := make(RonKnowledgeMap)
	for k, v := range ronLidMap {
		newRecord := v.Copy()
		newRonKnowledgeMap[k] = &newRecord
	}

	return newRonKnowledgeMap
}

func (ronLidMap RonKnowledgeMap) CancelKnowledge(kid int) bool {
	marked := false

	for _, v := range ronLidMap {
		for loopId, af := range v.LoopIdMap {
			found := false
			for _, kAmount := range af.KAmounts {
				if kAmount.Kid == kid {
					found = true
					marked = true
					break
				}
			}

			if found {
				// if more than 1, delete first one; otherwise delete all
				if len(af.KAmounts) > 1 {
					currentMin := af.KAmounts[0]
					af.KAmounts = af.KAmounts[1:]
					af.SumProfit -= currentMin.Profit
				} else {
					//delete loopId here
					delete(v.LoopIdMap, loopId)
				}
			}
		}
	}

	return marked
}

func (li *LoopIdMapSorted) Copy() LoopIdMapSorted {
	newLoopIdMap := make(map[common.Hash]*AmountProfit)
	for k, v := range li.LoopIdMap {
		newLoopIdMap[k] = v
	}

	newProfits := make([]*ProfitSorted, len(li.Profits))
	copy(newProfits[:], li.Profits[:])

	return LoopIdMapSorted{
		LoopIdMap: newLoopIdMap,
		Profits:   newProfits,
	}
}

func (li *LoopIdMapSorted) SortTopN() {
	profits := make([]*ProfitSorted, 0, len(li.LoopIdMap))
	for loopId, ap := range li.LoopIdMap {
		profits = append(profits, &ProfitSorted{
			Profit: ap.SumProfit,
			LoopId: loopId,
		})
	}

	if len(profits) > 1 {
		sort.Slice(profits, func(i, j int) bool {
			return profits[i].Less(profits[j])
		})

		if len(profits) >= LoopIdTopN {
			profits = profits[:LoopIdTopN]
		}
	}
	li.Profits = profits
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
	str.WriteString(p.Router.HexNoChecksum())
	for _, token := range p.Token {
		str.WriteString("-")
		str.WriteString(token.HexNoChecksum())
	}
	return str.String()
}

type JsonNewObs struct {
	Router   string `json:"router"`
	MethodID uint32 `json:"methodId"`
}

func (record *JsonNewObs) ToNewObs() *NewObs {
	return &NewObs{
		Router:   common.HexToAddress(record.Router),
		MethodID: record.MethodID,
	}
}

type NewObs struct {
	Router   common.Address
	MethodID uint32
}

func (newObs *NewObs) ToJsonNewObs() *JsonNewObs {
	return &JsonNewObs{
		Router:   newObs.Router.String(),
		MethodID: newObs.MethodID,
	}
}
