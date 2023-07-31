package defi

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/txpool"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	v3pool "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type V2SyncInfo struct {
	address  common.Address
	reserve0 *big.Int
	reserve1 *big.Int
}

type RonFiExtractType uint32

const (
	RonFiExtractTypeHunting RonFiExtractType = iota
	RonFiExtractTypeStats
	RonFiExtractTypePairs
)

func (di *Info) ExtractSwapPairInfo(tx *types.Transaction, router common.Address, vLogs []*types.Log, eType RonFiExtractType) []*SwapPairInfo {
	var syncPairInfo *V2SyncInfo // there must have a 'sync' event before any 'swap'/'mint'/'burn' event.

	swapPairsInfo := make([]*SwapPairInfo, 0, len(vLogs)/2)

	for _, vlog := range vLogs {
		var (
			amountIn, amountOut, reserve0, reserve1 *big.Int
			tokenIn, tokenOut, sender, to           common.Address
			info                                    *PairInfo
			poolInfo                                *PoolInfo
			key                                     string
			dir                                     uint64
			hasSwapPairInfo                         bool
		)

		if len(vlog.Topics) > 0 {
			topic0 := vlog.Topics[0]
			data := vlog.Data
			address := vlog.Address
			switch topic0 {
			case state.V2PairCreatedEvent:
				if len(data) >= 32 && len(vlog.Topics) == 3 {
					log.Info("RonFi extractSwapPairInfo pair created", "dexTx", tx.Hash().String(), "pair", address)
				}
			case state.V2SyncEvent:
				if len(data) == 64 && eType == RonFiExtractTypeHunting {
					syncPairInfo = &V2SyncInfo{
						address,
						new(big.Int).SetBytes(data[18:32]), // 112 bits = 14 bytes, 32-14=18
						new(big.Int).SetBytes(data[50:64]), // 64-14=50
					}
				}
			case state.V2MintEvent, state.V2BurnEvent:
				if len(data) == 64 {
					info = di.GetPairInfo(address)
					if eType != RonFiExtractTypeHunting || info == nil {
						continue
					}
					if syncPairInfo == nil || syncPairInfo.address != address {
						log.Warn("RonFi extractSwapPairInfo surprise", "dexTx", tx.Hash().String(), "pair", address, "event", state.EventName(topic0))
						continue
					}

					token0 := info.Token0
					token1 := info.Token1
					// create bi-direction swaps.
					for direction := 0; direction < 2; direction++ {
						if direction == 0 {
							amountIn = new(big.Int).SetBytes(data[18:32]) // only need uint112 (i.e. 14 bytes)
							amountOut = new(big.Int)                      // new values can be declared in the usual ways and denote 0 without further initialization
							tokenIn = token0
							tokenOut = token1
							if tokenIn == rcommon.ZeroAddress || tokenOut == rcommon.ZeroAddress {
								// Impossible reaching
								log.Warn("RonFi extractSwapPairInfo tokenIn/tokenOut nil", "pair", address, "tokenIn", tokenIn, "tokenOut", tokenOut)
								continue
							}
						} else {
							amountIn = new(big.Int).SetBytes(data[50:64]) // only need uint112 (i.e. 14 bytes)
							amountOut = new(big.Int)
							tokenIn = token1
							tokenOut = token0
						}

						reserve0 = syncPairInfo.reserve0
						reserve1 = syncPairInfo.reserve1

						key = fmt.Sprintf("%s-%d", address, dir^1)
						swapPairInfo := SwapPairInfo{
							Address:      address,
							Key:          key,
							V3:           false,
							BothBriToken: info.BothBriToken,
							TokenIn:      tokenIn,
							TokenOut:     tokenOut,
							KeyToken:     info.KeyToken,
							AmountIn:     amountIn,
							AmountOut:    amountOut,
							Reserve0:     reserve0,
							Reserve1:     reserve1,
							Dir:          uint64(dir),
						}
						swapPairsInfo = append(swapPairsInfo, &swapPairInfo)
					}
				}
			case state.V3MintEvent:
				if len(data) == 128 && len(vlog.Topics) == 4 {
					poolInfo = di.GetPoolInfo(address)
					if poolInfo == nil {
						continue
					}

					if eType != RonFiExtractTypeHunting || info == nil {
						continue
					}
					if syncPairInfo == nil || syncPairInfo.address != address {
						log.Warn("RonFi extractSwapPairInfo surprise", "dexTx", tx.Hash().String(), "pair", address, "event", state.EventName(topic0))
						continue
					}

					token0 := info.Token0
					token1 := info.Token1
					// create bi-direction swaps.
					for direction := 0; direction < 2; direction++ {
						if direction == 0 {
							amountIn = new(big.Int).SetBytes(data[64:96])   // only need uint112 (i.e. 14 bytes)
							amountOut = new(big.Int).SetBytes(data[96:128]) // new values can be declared in the usual ways and denote 0 without further initialization
							tokenIn = token0
							tokenOut = token1
							if tokenIn == rcommon.ZeroAddress || tokenOut == rcommon.ZeroAddress {
								// Impossible reaching
								log.Warn("RonFi extractSwapPairInfo tokenIn/tokenOut nil", "pair", address, "tokenIn", tokenIn, "tokenOut", tokenOut)
								continue
							}
						} else {
							amountIn = new(big.Int).SetBytes(data[96:128]) // only need uint112 (i.e. 14 bytes)
							amountOut = new(big.Int).SetBytes(data[64:96])
							tokenIn = token1
							tokenOut = token0
						}

						reserve0 = syncPairInfo.reserve0
						reserve1 = syncPairInfo.reserve1

						key = fmt.Sprintf("%s-%d", address, dir^1)
						swapPairInfo := SwapPairInfo{
							Address:      address,
							Key:          key,
							V3:           false,
							BothBriToken: info.BothBriToken,
							TokenIn:      tokenIn,
							TokenOut:     tokenOut,
							KeyToken:     info.KeyToken,
							AmountIn:     amountIn,
							AmountOut:    amountOut,
							Reserve0:     reserve0,
							Reserve1:     reserve1,
							Dir:          uint64(dir),
						}
						swapPairsInfo = append(swapPairsInfo, &swapPairInfo)
					}
				}
			case state.V3SwapEvent:
				if len(vlog.Topics) == 3 && len(data) >= 160 {
					poolInfo = di.GetPoolInfo(address)
					if poolInfo == nil {
						continue
					}

					sender = common.BytesToAddress(vlog.Topics[1].Bytes())
					if eType == RonFiExtractTypeStats && (sender != router) && (sender != address) {
						continue // when calculate profit, ignore irrelevant swap events. relevant only if sender is router address
					}
					to = common.BytesToAddress(vlog.Topics[2].Bytes())
					token0 := poolInfo.Token0
					token1 := poolInfo.Token1

					tokenAbi, err := abi.JSON(strings.NewReader(v3pool.V3poolABI))
					if err != nil {
						continue
					}
					unpack, err := tokenAbi.Unpack("Swap", data)
					if err != nil {
						continue
					}
					amount0 := unpack[0].(*big.Int)
					amount1 := unpack[1].(*big.Int)
					sqrtPriceX96 := unpack[2].(*big.Int)
					liquidity := unpack[3].(*big.Int)
					tick := int(unpack[4].(*big.Int).Int64())
					dir = 0
					if (amount0.Cmp(big.NewInt(0)) < 0) && (amount1.Cmp(big.NewInt(0)) > 0) {
						dir = 1
					}
					if dir == 0 {
						tokenIn = token0
						tokenOut = token1
						amountIn = new(big.Int).Abs(amount0)
						amountOut = new(big.Int).Abs(amount1)
					} else {
						tokenIn = token1
						tokenOut = token0
						amountIn = new(big.Int).Abs(amount1)
						amountOut = new(big.Int).Abs(amount0)
					}
					key = fmt.Sprintf("%s-%d", address, dir^1)
					swapPairInfo := SwapPairInfo{
						Address:      address,
						Key:          key,
						V3:           true,
						BothBriToken: false,
						Sender:       sender,
						To:           to,
						TokenIn:      tokenIn,
						TokenOut:     tokenOut,
						KeyToken:     tokenIn,
						AmountIn:     amountIn,
						AmountOut:    amountOut,
						Reserve0:     nil,
						Reserve1:     nil,
						Tick:         tick,
						SqrtPriceX96: sqrtPriceX96,
						Liquidity:    liquidity,
						Dir:          uint64(dir),
					}
					swapPairsInfo = append(swapPairsInfo, &swapPairInfo)
				}
			case state.V2SwapEvent:
				if len(data) == 128 && len(vlog.Topics) == 3 {
					info = di.GetPairInfo(address)
					if info == nil { // not a known pair (i.e. none loops contain this pair, and rpc query fails too), nothing we can do.
						continue
					}
					if eType == RonFiExtractTypeHunting && (syncPairInfo == nil || syncPairInfo.address != address) {
						log.Warn("RonFi extractSwapPairInfo surprise", "dexTx", tx.Hash().String(), "pair", address, "event", state.EventName(topic0))
						continue
					}
					sender = common.BytesToAddress(vlog.Topics[1].Bytes())
					if eType == RonFiExtractTypeStats && (sender != router) && (sender != address) {
						continue // when calculate profit, ignore irrelevant swap events. relevant only if sender is router address
					}
					to = common.BytesToAddress(vlog.Topics[2].Bytes())

					token0 := info.Token0
					token1 := info.Token1

					amount0In := new(big.Int).SetBytes(data[18:32])    // only need uint112 (i.e. 14 bytes)
					amount1In := new(big.Int).SetBytes(data[50:64])    //32+18
					amount0Out := new(big.Int).SetBytes(data[82:96])   //64+18
					amount1Out := new(big.Int).SetBytes(data[114:128]) //96+18

					if amount0Out.BitLen() == 0 || amount1Out.BitLen() == 0 {
						if amount0Out.BitLen() == 0 {
							dir = 0
							amountIn = amount0In
							amountOut = amount1Out
						} else {
							dir = 1
							amountIn = amount1In
							amountOut = amount0Out
						}
					} else { // if both amount0/1Out are not zero, fallback to use amount0/1In to check direction
						if amount1In.BitLen() == 0 {
							dir = 0
							amountIn = amount0In
							amountOut = amount1Out
						} else {
							dir = 1
							amountIn = amount1In
							amountOut = amount0Out
						}
					}
					if dir == 0 {
						tokenIn = token0
						tokenOut = token1
					} else {
						tokenIn = token1
						tokenOut = token0
					}
					key = fmt.Sprintf("%s-%d", address, dir^1)
					if eType == RonFiExtractTypeHunting {
						reserve0 = syncPairInfo.reserve0
						reserve1 = syncPairInfo.reserve1
					}

					hasSwapPairInfo = true
				}
			}

			if hasSwapPairInfo {
				swapPairInfo := SwapPairInfo{
					address,
					key,
					false,
					info.BothBriToken,
					sender,
					to,
					tokenIn,
					tokenOut,
					info.KeyToken,
					amountIn,
					amountOut,
					reserve0,
					reserve1,
					0,
					nil,
					nil,
					dir,
				}
				swapPairsInfo = append(swapPairsInfo, &swapPairInfo)
			}
		}
	}

	return swapPairsInfo
}

func (di *Info) CheckIfObsTx(tx *types.Transaction, vLogs []*types.Log, router common.Address) (isDex bool, isObs bool) {
	to := tx.To()
	if to == nil {
		return
	}

	data := tx.Data()
	if len(data) < 4 {
		return
	}
	methodID := uint64(binary.BigEndian.Uint32(data[:4]))

	{
		if txpool.ObsRouters != nil {
			routerMethod := fmt.Sprintf("%s-0x%08x", *to, methodID)
			if _, isObs = txpool.ObsRouters[routerMethod]; isObs {
				return
			}
		}
	}

	if _, isDex = txpool.DexRouters[*to]; isDex {
		return
	} else if _, isDex = txpool.DexMethodsTypical[methodID]; isDex {
		return
	} else if _, isObs = txpool.ObsMethods[methodID]; isObs {
		return
	}

	swapPairsInfo := di.ExtractSwapPairInfo(tx, router, vLogs, RonFiExtractTypeStats)
	if len(swapPairsInfo) > 0 {
		isDex = true
		isObs = false

		for i := 0; i < len(swapPairsInfo); i++ {
			for j := i + 1; j < len(swapPairsInfo); j++ {
				pairs := swapPairsInfo[i : j+1]
				if len(pairs) > 1 {
					var k int
					for k = 0; k < len(pairs)-1; k++ {
						head := pairs[0]
						prev := pairs[k]
						next := pairs[k+1]
						tail := next
						if prev.To != next.Address && prev.To != *to ||
							prev.TokenOut != next.TokenIn ||
							head.TokenIn != tail.TokenOut {
							break
						}
					}
					if k == len(pairs)-1 {
						isDex = false
						isObs = true
						return
					}
				}
			}
		}
	}

	return
}

func (di *Info) GetArbTxProfit(tx *types.Transaction, vLogs []*types.Log, router common.Address) (float64, bool) {
	v3Loop := false
	swapPairsInfo := di.ExtractSwapPairInfo(tx, router, vLogs, RonFiExtractTypeStats)
	for _, pairInfo := range swapPairsInfo {
		if pairInfo.V3 {
			v3Loop = true
		}
	}

	totalProfit := 0.0
	if len(swapPairsInfo) > 1 {
		for i := 0; i < len(swapPairsInfo); i++ {
			for j := i + 1; j < len(swapPairsInfo); j++ {
				pairs := swapPairsInfo[i : j+1]
				if len(pairs) > 1 {
					var k int
					for k = 0; k < len(pairs)-1; k++ {
						head := pairs[0]
						prev := pairs[k]
						next := pairs[k+1]
						tail := next
						if prev.To != next.Address && prev.To != *tx.To() ||
							prev.TokenOut != next.TokenIn ||
							head.TokenIn != tail.TokenOut {
							break
						}
					}
					if k == len(pairs)-1 {
						totalProfit += di.loopProfit(pairs)
					}
				}
			}
		}
	}

	return totalProfit, v3Loop
}

func (di *Info) loopProfit(swapPairsInfo []*SwapPairInfo) (profit float64) {
	if len(swapPairsInfo) < 2 {
		return
	}

	prev := swapPairsInfo[0]
	next := swapPairsInfo[1]
	// check if head/tail right?
	if prev.AmountOut.Cmp(next.AmountIn) != 0 &&
		next.AmountOut.Cmp(prev.AmountIn) == 0 {
		// reverse
		for i := len(swapPairsInfo)/2 - 1; i >= 0; i-- {
			opp := len(swapPairsInfo) - 1 - i
			swapPairsInfo[i], swapPairsInfo[opp] = swapPairsInfo[opp], swapPairsInfo[i]
		}
	}

	head := swapPairsInfo[0]
	tail := swapPairsInfo[len(swapPairsInfo)-1]
	token := head.TokenIn

	isTradableToken := false
	amount := 0.0
	defer func() {
		// finally, we got the real profit
		if isTradableToken {
			price := GetTradingTokenPrice(token)
			amount /= price
			profit = amount * GetTradingTokenPrice(rcommon.USDC)
		}
	}()

	if head.TokenIn == tail.TokenOut {
		if _, isTradableToken = rcommon.TradableTokens[token]; isTradableToken {
			decimals := uint64(18)
			if tInfo := di.GetTokenInfo(token); tInfo != nil {
				decimals = tInfo.Decimals
			}
			amount = rcommon.TokenToFloat(new(big.Int).Sub(tail.AmountOut, head.AmountIn), decimals) // profit on tail
		}
	}
	return
}
