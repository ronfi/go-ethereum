package defi

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	erc20token "github.com/ethereum/go-ethereum/ronfi/contracts/contract_erc20"
	v2pair "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v2pair"
	v3pool "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
	v3TickLens "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3ticklens"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"math/big"
	"sync"
)

type Protocol int

const (
	UniSwap Protocol = iota
)

func NewInfo(client *ethclient.Client, mysql *db.Mysql) *Info {
	poolsInfo := make(map[common.Address]*PoolInfo)
	pairsInfo := make(map[common.Address]*PairInfo)
	tokensInfo := make(map[common.Address]*TokenInfo)
	newPairsInfo := make(map[common.Address]*PairInfo)
	newPoolsInfo := make(map[common.Address]*PoolInfo)
	newTokensInfo := make(map[common.Address]*TokenInfo)

	// preload pairs info from database
	pRecords := mysql.LoadPairsInfo()

	for _, record := range pRecords {
		pairsInfo[common.HexToAddress(record.Pair)] = &PairInfo{
			record.Name,
			record.Index,
			0,
			record.BothBriToken,
			common.HexToAddress(record.KeyToken),
			common.HexToAddress(record.Token0),
			common.HexToAddress(record.Token1),
			common.HexToAddress(record.Factory),
		}
	}
	log.Info("RonFi Defi preload pairs info", "size", len(pairsInfo))

	poolRecords := mysql.LoadPoolsInfo()
	for _, record := range poolRecords {
		poolsInfo[common.HexToAddress(record.Pool)] = &PoolInfo{
			record.Name,
			common.HexToAddress(record.Token0),
			common.HexToAddress(record.Token1),
			new(big.Int).SetInt64(int64(record.Fee)),
			record.TickSpacing,
			common.HexToAddress(record.Factory),
		}
	}
	log.Info("RonFi Defi preload pools info", "size", len(poolsInfo))

	// preload tokens info from database
	tRecords := mysql.LoadTokensInfo()
	for _, record := range tRecords {
		tokensInfo[common.HexToAddress(record.Token)] = &TokenInfo{
			Symbol:   record.Symbol,
			Decimals: uint64(record.Decimals),
		}
	}
	log.Info("RonFi Defi preload tokens info", "size", len(tokensInfo))

	return &Info{
		client,
		mysql,

		poolsInfo,
		pairsInfo,
		tokensInfo,
		newPairsInfo,
		newPoolsInfo,
		newTokensInfo,
		sync.RWMutex{},

		NewProxy(client),
	}
}

func (di *Info) GetTokenInfo(token common.Address) (tokenInfo *TokenInfo) {
	var ok bool

	di.lock.RLock()
	tokenInfo, ok = di.tokensInfo[token]
	di.lock.RUnlock()
	if ok {
		return
	}

	if inst, err := erc20token.NewToken(token, di.client); err == nil {
		if symbol, err := inst.Symbol(nil); err == nil {
			if decimals, err := inst.Decimals(nil); err == nil {
				tokenInfo = &TokenInfo{
					Symbol:   symbol,
					Decimals: uint64(decimals),
				}

				di.lock.Lock()
				di.tokensInfo[token] = tokenInfo
				di.newTokensInfo[token] = tokenInfo
				di.lock.Unlock()
			}
		}
	}

	return
}

func (di *Info) GetPoolInfo(address common.Address) (info *PoolInfo) {
	var ok bool
	di.lock.RLock()
	info, ok = di.poolsInfo[address]
	di.lock.RUnlock()
	if ok {
		return
	}

	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		factory, err := inst.Factory(nil)
		if err != nil {
			return
		}

		if token0, err := inst.Token0(nil); err == nil {
			if token1, err := inst.Token1(nil); err == nil {
				if tickSpacing, err := inst.TickSpacing(nil); err == nil {
					token0Info := di.GetTokenInfo(token0)
					token1Info := di.GetTokenInfo(token1)
					if token0Info == nil || token1Info == nil {
						return
					}

					var fee *big.Int
					fee, _ = inst.Fee(nil)
					poolName := ""
					if fee == nil {
						fee = big.NewInt(0)
						poolName = fmt.Sprintf("%s-%s Non Uniswap V3, %d", token0Info.Symbol, token1Info.Symbol, fee.Int64())
					} else {
						poolName = fmt.Sprintf("%s-%s Uniswap V3, %d", token0Info.Symbol, token1Info.Symbol, fee.Int64())
					}
					if token0Info != nil && token1Info != nil {
						info = &PoolInfo{
							Name:        poolName,
							Token0:      token0,
							Token1:      token1,
							Fee:         fee,
							TickSpacing: int(tickSpacing.Int64()),
							Factory:     factory,
						}

						di.lock.Lock()
						di.poolsInfo[address] = info
						di.newPoolsInfo[address] = info
						di.lock.Unlock()
					}
				}
			}
		}
	}

	return
}

func (di *Info) GetV3SqrtPriceX96(address common.Address) (*big.Int, error) {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if slot0, err := inst.Slot0(nil); err == nil {
			return slot0.SqrtPriceX96, nil
		}
	}

	return nil, errors.New("get v3 sqrt price failed")
}

func (di *Info) GetV3Tick(address common.Address) (int, error) {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if slot0, err := inst.Slot0(nil); err == nil {
			return int(slot0.Tick.Int64()), nil
		}
	}

	return 0, errors.New("get v3 current tick failed")
}

func (di *Info) GetV3Liquidity(address common.Address) (*big.Int, error) {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if liquidity, err := inst.Liquidity(nil); err == nil {
			return liquidity, nil
		}
	}

	return nil, errors.New("get pool liquidity failed")
}

func (di *Info) GetV3TickBitMap(address common.Address, wordPos int16) (*big.Int, error) {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if tickBitMap, err := inst.TickBitmap(nil, wordPos); err == nil {
			return tickBitMap, nil
		}
	}

	return nil, errors.New("get tick bitmap failed")
}

func (di *Info) GetV3PopulatedTicksInWord(tickLensAddr common.Address, poolAddr common.Address, wordPos int16) ([]v3TickLens.ITickLensPopulatedTick, error) {
	if inst, err := v3TickLens.NewV3TickLens(tickLensAddr, di.client); err == nil {
		if tick, err := inst.GetPopulatedTicksInWord(nil, poolAddr, wordPos); err == nil {
			return tick, nil
		} else {
			return nil, err
		}
	}

	return nil, errors.New("get populated ticks in word failed")
}

func (di *Info) GetPairInfo(pair common.Address) (pairInfo *PairInfo) {
	var ok bool

	di.lock.RLock()
	pairInfo, ok = di.pairsInfo[pair]
	di.lock.RUnlock()
	if ok {
		return
	}

	pairInfo = di.getUniSwapPairInfo(pair)
	if pairInfo == nil {
		log.Warn("RonFi Defi get pair info failed", "pair", pair)
	}

	di.lock.Lock()
	if pairInfo != nil {
		di.pairsInfo[pair] = pairInfo
		di.newPairsInfo[pair] = pairInfo
	}
	di.lock.Unlock()

	return
}

func (di *Info) getUniSwapPairInfo(pair common.Address) *PairInfo {
	{
		if inst, err := v2pair.NewV2pair(pair, di.client); err == nil {
			if token0, err := inst.Token0(nil); err == nil {
				if token1, err := inst.Token1(nil); err == nil {
					if reserves, err := inst.GetReserves(nil); err == nil {
						factory, err := inst.Factory(nil)
						if err != nil {
							factory = rcommon.ZeroAddress
						}

						name, err := inst.Name(nil)
						if err != nil {
							name = "pancake lp"
						}
						return di.buildPairInfo(pair, factory, token0, token1, name, reserves.Reserve0, reserves.Reserve1)
					}
				}
			}
		}
	}

	return nil
}

func (di *Info) buildPairInfo(pair, factory, token0, token1 common.Address, name string, reserve0, reserve1 *big.Int) *PairInfo {
	keyToken := rcommon.ZeroAddress
	bothBriToken := false
	_, ok0 := rcommon.BridgeTokens[token0]
	_, ok1 := rcommon.BridgeTokens[token1]

	switch {
	case ok0 && ok1:
		bothBriToken = true
	case ok0 && !ok1:
		keyToken = token1
	case ok1 && !ok0:
		keyToken = token0
	default:
	}

	index := di.findPairResvIndex(pair, reserve0, reserve1)
	return &PairInfo{
		name,
		index,
		0,
		bothBriToken,
		keyToken,
		token0,
		token1,
		factory,
	}
}

func (di *Info) findPairResvIndex(pair common.Address, reserve0, reserve1 *big.Int) uint64 {
	index := uint64(0xff)

	for i := uint64(0); i < 20; i++ {
		if data, err := di.client.StorageAt(context.Background(), pair, common.BytesToHash([]byte{byte(i)}), nil); err == nil {
			if len(data) != 32 {
				continue
			}
			r0 := new(big.Int).SetBytes(data[18:32])
			r1 := new(big.Int).SetBytes(data[4:18])
			if r0.Cmp(reserve0) == 0 && r1.Cmp(reserve1) == 0 {
				index = i
				break
			}
		}
	}

	return index
}

func (di *Info) GetPairReserves(pair common.Address) *Reserve {
	var resv *Reserve
	resv = di.getPancakePairReserves(pair)
	if resv == nil {
		log.Warn("RonFi Defi get pancake pair reserves failed", "pair", pair.Hex())
	}

	return resv
}

func (di *Info) getPancakePairReserves(pair common.Address) *Reserve {
	if inst, err := v2pair.NewV2pairCaller(pair, di.client); err == nil {
		if reserves, err := inst.GetReserves(nil); err == nil {
			return &Reserve{
				reserves.Reserve0,
				reserves.Reserve1,
				reserves.BlockTimestampLast,
			}
		}
	}

	return nil
}

func (di *Info) GetNewPairsInfo() map[common.Address]*PairInfo {
	newPairsInfoCopy := make(map[common.Address]*PairInfo)
	di.lock.RLock()
	for k, v := range di.newPairsInfo {
		newPairsInfoCopy[k] = v
	}
	di.lock.RUnlock()

	return newPairsInfoCopy
}

func (di *Info) GetNewPoolsInfo() map[common.Address]*PoolInfo {
	newPoolsInfoCopy := make(map[common.Address]*PoolInfo)
	di.lock.RLock()
	for k, v := range di.newPoolsInfo {
		newPoolsInfoCopy[k] = v
	}
	di.lock.RUnlock()

	return newPoolsInfoCopy
}

func (di *Info) GetNewTokensInfo() map[common.Address]*TokenInfo {
	newTokensInfoCopy := make(map[common.Address]*TokenInfo)
	di.lock.RLock()
	for k, v := range di.newTokensInfo {
		newTokensInfoCopy[k] = v
	}
	di.lock.RUnlock()

	return newTokensInfoCopy
}

func (di *Info) GetAllPairInfo() map[common.Address]*PairInfo {
	allPairInfoCopy := make(map[common.Address]*PairInfo)
	di.lock.RLock()
	for k, v := range di.pairsInfo {
		allPairInfoCopy[k] = v
	}
	for k, v := range di.newPairsInfo {
		allPairInfoCopy[k] = v
	}
	di.lock.RUnlock()

	return allPairInfoCopy
}

func (di *Info) GetAllPoolInfo() map[common.Address]*PoolInfo {
	allPoolsInfoCopy := make(map[common.Address]*PoolInfo)
	di.lock.RLock()
	for k, v := range di.poolsInfo {
		allPoolsInfoCopy[k] = v
	}
	for k, v := range di.newPoolsInfo {
		allPoolsInfoCopy[k] = v
	}
	di.lock.RUnlock()

	return allPoolsInfoCopy
}

func (di *Info) GetAllTokenInfo() map[common.Address]*TokenInfo {
	allTokenInfoCopy := make(map[common.Address]*TokenInfo)
	di.lock.RLock()
	for k, v := range di.tokensInfo {
		allTokenInfoCopy[k] = v
	}
	for k, v := range di.newTokensInfo {
		allTokenInfoCopy[k] = v
	}
	di.lock.RUnlock()

	return allTokenInfoCopy
}

func (di *Info) GetTokenBalance(account common.Address, token common.Address) *big.Int {
	if inst, err := erc20token.NewToken(token, di.client); err == nil {
		if balance, err := inst.BalanceOf(nil, account); err == nil {
			return balance
		}
	}

	return nil
}

func (di *Info) MergePairTokensInfo() {
	di.lock.Lock()
	for k, v := range di.newPairsInfo {
		di.pairsInfo[k] = v
	}
	di.newPairsInfo = make(map[common.Address]*PairInfo)

	for k, v := range di.newTokensInfo {
		di.tokensInfo[k] = v
	}
	di.newTokensInfo = make(map[common.Address]*TokenInfo)

	for k, v := range di.newPoolsInfo {
		di.poolsInfo[k] = v
	}
	di.newPoolsInfo = make(map[common.Address]*PoolInfo)
	di.lock.Unlock()
}
