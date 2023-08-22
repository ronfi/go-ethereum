package defi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"math/big"
	"sync"
)

type RonFiBalance struct {
	Eth float64

	Total         float64
	ContractTotal float64

	ContractEth  float64
	ContractBtc  float64
	ContractUsdx float64
	ContractChi  uint64
}

// Copy execute a deep copy
func (b *RonFiBalance) Copy() RonFiBalance {
	return RonFiBalance{
		Eth:           b.Eth,
		Total:         b.Total,
		ContractTotal: b.ContractTotal,

		ContractEth:  b.ContractEth,
		ContractBtc:  b.ContractBtc,
		ContractUsdx: b.ContractUsdx,
		ContractChi:  b.ContractChi,
	}
}

type PoolInfo struct {
	Name        string
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing int
	Factory     common.Address
}

type PairInfo struct {
	Name         string
	Index        uint64
	Fee          uint64
	BothBriToken bool
	CanFlashLoan bool
	KeyToken     common.Address
	Token0       common.Address
	Token1       common.Address
	Factory      common.Address
}

type PairInfoMap map[common.Address]*PairInfo

func (p PairInfoMap) Copy() PairInfoMap {
	newPairInfoMap := make(PairInfoMap)
	for k, v := range p {
		newPairInfoMap[k] = v
	}

	return newPairInfoMap
}

type TokenInfo struct {
	Symbol   string
	Decimals uint64
	Fee      uint64
}

type Reserve struct {
	Reserve0  *big.Int
	Reserve1  *big.Int
	Timestamp uint32
}

func (m *Reserve) Switch() Reserve {
	return Reserve{m.Reserve1, m.Reserve0, m.Timestamp}
}

func (p PairInfo) Equals(s PairInfo) bool {
	if p.Name != s.Name ||
		p.Index != s.Index ||
		p.BothBriToken != s.BothBriToken ||
		p.CanFlashLoan != s.CanFlashLoan ||
		p.KeyToken != s.KeyToken ||
		p.Token0 != s.Token0 ||
		p.Token1 != s.Token1 ||
		p.Factory != s.Factory {
		return false
	}

	return true
}

type Info struct {
	client *ethclient.Client
	mysql  *db.Mysql
	signer types.Signer

	poolsInfo     map[common.Address]*PoolInfo
	pairsInfo     map[common.Address]*PairInfo
	tokensInfo    map[common.Address]*TokenInfo
	newPairsInfo  map[common.Address]*PairInfo
	newPoolsInfo  map[common.Address]*PoolInfo
	newTokensInfo map[common.Address]*TokenInfo
	lock          sync.RWMutex

	proxy *Proxy
}

type SwapPairInfo struct {
	Address      common.Address
	Key          string
	V3           bool
	BothBriToken bool
	Sender       common.Address
	To           common.Address
	TokenIn      common.Address
	TokenOut     common.Address
	KeyToken     common.Address
	AmountIn     *big.Int
	AmountOut    *big.Int
	Reserve0     *big.Int
	Reserve1     *big.Int
	Tick         int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Dir          uint64
}

func (p *SwapPairInfo) Reverse() *SwapPairInfo {
	dir := 1 - p.Dir
	key := fmt.Sprintf("%s-%d", p.Address, dir^1)

	return &SwapPairInfo{
		Address:      p.Address,
		Key:          key,
		V3:           p.V3,
		BothBriToken: p.BothBriToken,
		Sender:       p.Sender,
		To:           p.To,
		TokenIn:      p.TokenOut,
		TokenOut:     p.TokenIn,
		KeyToken:     p.KeyToken,
		AmountIn:     nil,
		AmountOut:    nil,
		Reserve0:     nil,
		Reserve1:     nil,
		Tick:         p.Tick,
		SqrtPriceX96: p.SqrtPriceX96,
		Liquidity:    p.Liquidity,
		Dir:          dir,
	}
}

type TxAndReceipt struct {
	Tx      *types.Transaction
	Receipt *types.Receipt
}
