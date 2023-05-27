package defi

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/contracts/contract_basev1"
	deMax "github.com/ethereum/go-ethereum/ronfi/contracts/contract_demax"
	dodoPool "github.com/ethereum/go-ethereum/ronfi/contracts/contract_dodo"
	erc20token "github.com/ethereum/go-ethereum/ronfi/contracts/contract_erc20"
	pancakePair "github.com/ethereum/go-ethereum/ronfi/contracts/contract_pancakepair"
	v3pool "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
	v3TickLens "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3ticklens"
	"math/big"
	"sync"
)

type Protocol int

const (
	PanCake Protocol = iota
	DodoSwap
	BaseV1
	DeMax
)

var (
	DoDoSwapId = crypto.Keccak256([]byte("_BASE_TOKEN_()"))[:4]
	DeMaxId    = crypto.Keccak256([]byte("FACTORY()"))[:4]
	BaseV1Id   = crypto.Keccak256([]byte("metadata()"))[:4]
	Token0Id   = crypto.Keccak256([]byte("token0()"))[:4]
)

func NewInfo(client *ethclient.Client) *Info {
	poolsInfo := make(map[common.Address]*PoolInfo)
	pairsInfo := make(map[common.Address]*PairInfo)
	tokensInfo := make(map[common.Address]*TokenInfo)
	newPairsInfo := make(map[common.Address]*PairInfo)
	newPoolsInfo := make(map[common.Address]*PoolInfo)
	newTokensInfo := make(map[common.Address]*TokenInfo)

	return &Info{
		client,

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
						poolName = fmt.Sprintf("%s-%s Non Pancake V3, %d", token0Info.Symbol, token1Info.Symbol, fee.Int64())
					} else {
						poolName = fmt.Sprintf("%s-%s Pancake V3, %d", token0Info.Symbol, token1Info.Symbol, fee.Int64())
					}
					if token0Info != nil && token1Info != nil {
						info = &PoolInfo{
							Name:        poolName,
							Token0:      token0,
							Token1:      token1,
							Fee:         fee,
							TickSpacing: int(tickSpacing.Int64()),
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

func (di *Info) GetV3Factory(address common.Address) *common.Address {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if factory, err := inst.Factory(nil); err == nil {
			return &factory
		}
	}

	return nil
}

func (di *Info) GetV3SqrtPriceX96(address common.Address) (*big.Int, error) {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if slot0, err := inst.Slot0(nil); err == nil {
			return slot0.SqrtPriceX96, nil
		}
	}

	return nil, errors.New("get v3 sqrt price failed")
}

func (di *Info) GetV3CurrentTick(address common.Address) (int, error) {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if slot0, err := inst.Slot0(nil); err == nil {
			return int(slot0.Tick.Int64()), nil
		}
	}

	return 0, errors.New("get v3 current tick failed")
}

func (di *Info) GetV3CurrentSqrtPriceX96(address common.Address) (*big.Int, error) {
	if inst, err := v3pool.NewV3pool(address, di.client); err == nil {
		if liquidity, err := inst.Liquidity(nil); err == nil {
			return liquidity, nil
		}
	}

	return nil, errors.New("get pool liquidity failed")
}

func (di *Info) GetV3CurrentLiquidity(address common.Address) (*big.Int, error) {
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

	codeAddr := pair
	if target, ok := di.proxy.detectProxyTarget(pair); ok {
		codeAddr = target
	}

	bytecode, err := di.client.CodeAt(context.Background(), codeAddr, nil)
	if err != nil || len(bytecode) <= 1 {
		return
	}

	defiProtocol := PanCake
	if bytes.Contains(bytecode, DoDoSwapId) {
		defiProtocol = DodoSwap
	} else if bytes.Contains(bytecode, BaseV1Id) {
		defiProtocol = BaseV1
	} else if bytes.Contains(bytecode, DeMaxId) {
		defiProtocol = DeMax
	} else {
		if !bytes.Contains(bytecode, Token0Id) {
			return
		}
	}

	switch defiProtocol {
	case DodoSwap:
		pairInfo = di.getDodoPairInfo(pair)
	case BaseV1:
		pairInfo = di.getBaseV1PairInfo(pair)
	default:
		pairInfo = di.getPancakePairInfo(defiProtocol, pair)
	}

	di.lock.Lock()
	if pairInfo != nil {
		di.pairsInfo[pair] = pairInfo
		di.newPairsInfo[pair] = pairInfo
	}
	di.lock.Unlock()

	return
}

func (di *Info) getDodoPairInfo(pair common.Address) *PairInfo {
	if inst, err := dodoPool.NewDppAdvanced(pair, di.client); err == nil {
		if baseToken, err := inst.BASETOKEN(nil); err == nil {
			if quoteToken, err := inst.QUOTETOKEN(nil); err == nil {
				token0, token1, dir := rcommon.SortTokens(baseToken, quoteToken)
				reserve0, _ := inst.BASERESERVE(nil)
				reserve1, _ := inst.QUOTERESERVE(nil)
				if reserve0 == nil || reserve1 == nil {
					return nil
				}
				if dir == 1 {
					tmp := reserve0
					reserve0 = reserve1
					reserve1 = tmp
				}

				return di.buildPairInfo(pair, rcommon.ZeroAddress, token0, token1, "dodo", reserve0, reserve1)
			}
		}
	}

	return nil
}

func (di *Info) getBaseV1PairInfo(pair common.Address) *PairInfo {
	if inst, err := basev1.NewBasev1(pair, di.client); err == nil {
		if meta, err := inst.Metadata(nil); err == nil {
			factory, err := inst.Factory(nil)
			if err != nil {
				factory = rcommon.ZeroAddress
			}

			name, err := inst.Name(nil)
			if err != nil {
				name = "basev1"
			}

			return di.buildPairInfo(pair, factory, meta.T0, meta.T1, name, meta.R0, meta.R1)
		}
	}

	return nil
}

func (di *Info) getPancakePairInfo(defiProtocol Protocol, pair common.Address) *PairInfo {
	if defiProtocol == DeMax {
		if inst, err := deMax.NewDemax(pair, di.client); err == nil {
			if token0, err := inst.Token0(nil); err == nil {
				if token1, err := inst.Token1(nil); err == nil {
					if reserves, err := inst.GetReserves(nil); err == nil {
						factory, err := inst.FACTORY(nil)
						if err != nil {
							factory = rcommon.ZeroAddress
						}

						name, err := inst.Name(nil)
						if err != nil {
							name = "demax"
						}
						return di.buildPairInfo(pair, factory, token0, token1, name, reserves.Reserve0, reserves.Reserve1)
					}
				}
			}
		}
	} else {
		if inst, err := pancakePair.NewPancakepair(pair, di.client); err == nil {
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
	codeAddr := pair
	if target, ok := di.proxy.detectProxyTarget(pair); ok {
		codeAddr = target
	}

	bytecode, err := di.client.CodeAt(context.Background(), codeAddr, nil)
	if err != nil || len(bytecode) <= 1 {
		return nil
	}

	defiProtocol := PanCake
	if bytes.Contains(bytecode, DoDoSwapId) {
		defiProtocol = DodoSwap
	} else if bytes.Contains(bytecode, BaseV1Id) {
		defiProtocol = BaseV1
	} else if bytes.Contains(bytecode, DeMaxId) {
		defiProtocol = DeMax
	} else {
		if !bytes.Contains(bytecode, Token0Id) {
			return nil
		}
	}

	switch defiProtocol {
	case DodoSwap:
		return di.getDoDoPairReserves(pair)
	case BaseV1:
		return di.getBaseV1PairReserves(pair)
	default:
		return di.getPancakePairReserves(pair)
	}
}

func (di *Info) getDoDoPairReserves(pair common.Address) *Reserve {
	if inst, err := dodoPool.NewDppAdvanced(pair, di.client); err == nil {
		if baseToken, err := inst.BASETOKEN(nil); err == nil {
			if quoteToken, err := inst.QUOTETOKEN(nil); err == nil {
				_, _, dir := rcommon.SortTokens(baseToken, quoteToken)
				reserve0, _ := inst.BASERESERVE(nil)
				reserve1, _ := inst.QUOTERESERVE(nil)
				if dir == 1 {
					tmp := reserve0
					reserve0 = reserve1
					reserve1 = tmp
				}
				return &Reserve{
					reserve0,
					reserve1,
					uint32(0),
				}
			}
		}
	}

	return nil
}

func (di *Info) getBaseV1PairReserves(pair common.Address) *Reserve {
	if inst, err := basev1.NewBasev1(pair, di.client); err == nil {
		if meta, err := inst.Metadata(nil); err == nil {
			return &Reserve{
				meta.R0,
				meta.R1,
				uint32(0),
			}
		}
	}

	return nil
}

func (di *Info) getPancakePairReserves(pair common.Address) *Reserve {
	if inst, err := pancakePair.NewPancakepair(pair, di.client); err == nil {
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
	allPoolInfoCopy := make(map[common.Address]*PoolInfo)
	di.lock.RLock()
	for k, v := range di.poolsInfo {
		if factory := di.GetV3Factory(k); factory != nil && *factory == rcommon.PANCAKE_V3_FACOTRY_ADDRESS {
			allPoolInfoCopy[k] = v
		}
	}
	for k, v := range di.newPoolsInfo {
		if factory := di.GetV3Factory(k); factory != nil && *factory == rcommon.PANCAKE_V3_FACOTRY_ADDRESS {
			allPoolInfoCopy[k] = v
		}
	}
	di.lock.RUnlock()

	return allPoolInfoCopy
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
