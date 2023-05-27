package defi

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	v3pool "github.com/ethereum/go-ethereum/ronfi/contracts/contract_v3pool"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type SyncPairInfo struct {
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

func (di *Info) ExtractSwapPairInfo(allPairsMap PairInfoMap, newPairsMap PairInfoMap, tx *types.Transaction, router common.Address, vLogs []*types.Log, eType RonFiExtractType) []*SwapPairInfo {
	var syncPairInfo *SyncPairInfo // there must have a 'sync' event before any 'swap'/'mint'/'burn' event.

	swapPairsInfo := make([]*SwapPairInfo, 0, len(vLogs)/2)

	pairsDb := make([]PairInfoMap, 0, 2)
	if allPairsMap != nil {
		pairsDb = append(pairsDb, allPairsMap)
	}
	if newPairsMap != nil {
		pairsDb = append(pairsDb, newPairsMap)
	}

	for _, vlog := range vLogs {
		var (
			amountIn, amountOut, reserve0, reserve1 *big.Int
			tokenIn, tokenOut, sender, to           common.Address
			info                                    *PairInfo
			poolInfo                                *PoolInfo
			key                                     string
			dir                                     uint64
			hasSwapPairInfo, knownPair, ok          bool
		)

		if len(vlog.Topics) > 0 {
			topic0 := vlog.Topics[0]
			data := vlog.Data
			address := vlog.Address
			switch topic0 {
			case state.V2PairCreatedEvent:
				if len(data) >= 32 && len(vlog.Topics) == 3 {
					//token0 := common.BytesToAddress(vlog.Topics[1].Bytes())
					//token1 := common.BytesToAddress(vlog.Topics[2].Bytes())
					//info = GetPairInfo(&token0, &token1)
					//switch eType {
					//case RonFiExtractTypeStats:
					//	//newCreatPairsMap[address] = info
					//case RonFiExtractTypeHunting:
					//	//if w != nil {
					//	//	w.newPairInfo[address] = info
					//	//}
					//default:
					//}
					log.Info("RonFi extractSwapPairInfo pair created", "dexTx", tx.Hash().String(), "pair", address)
				}
			case state.V2SyncEvent:
				if len(data) == 64 && eType == RonFiExtractTypeHunting {
					syncPairInfo = &SyncPairInfo{
						address,
						new(big.Int).SetBytes(data[18:32]), // 112 bits = 14 bytes, 32-14=18
						new(big.Int).SetBytes(data[50:64]), // 64-14=50
					}
				}
			case state.V2MintEvent, state.V2BurnEvent:
				if len(data) == 64 {
					info, ok = checkIfPairInLoops(pairsDb, address)
					if !ok { // not a known pair (i.e. none loops contain this pair), nothing we can do.
						if eType == RonFiExtractTypeStats {
							info = di.GetPairInfo(address)
						}
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
							KnownPair:    true,
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
							KnownPair:    true,
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
				if len(data) == 224 && len(vlog.Topics) == 3 {
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
						KnownPair:    false,
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
					info, ok = checkIfPairInLoops(pairsDb, address)
					if ok {
						knownPair = true
					} else {
						if eType != RonFiExtractTypeHunting { // if not hunting, here is 2nd chance to get pair token info, via rpc api.
							//log.Info("RonFi pair without loops", "dexTx", tx.Hash().String(), "pair", string(address.HexNoChecksum()), "event", state.EventName(topic0))
							info = di.GetPairInfo(address)
						}
					}
					if info == nil { // not a known pair (i.e. none loops contain this pair, and rpc query fails too), nothing we can do.
						continue
					}
					if eType == RonFiExtractTypeHunting && (syncPairInfo == nil || syncPairInfo.address != address) {
						log.Warn("RonFi extractSwapPairInfo surprise", "dexTx", tx.Hash().String(), "pair", address, "event", state.EventName(topic0))
						continue
					}
					sender = common.BytesToAddress(vlog.Topics[1].Bytes())
					if eType == RonFiExtractTypeStats && (sender != router) && (sender != address) && (sender != rcommon.RON_SWAP_F_ADDRESS) {
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
					knownPair,
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

func (di *Info) CheckIfObsTx(allPairsMap PairInfoMap, tx *types.Transaction, vLogs []*types.Log, router common.Address) (isObs bool) {
	to := tx.To()
	if to == nil {
		return
	}

	data := tx.Data()
	if len(data) < 4 {
		return
	}

	methodID := uint64(binary.BigEndian.Uint32(data[:4]))
	if _, ok := rcommon.DexMethodsTypical[methodID]; ok {
		return
	}

	swapPairsInfo := di.ExtractSwapPairInfo(allPairsMap, nil, tx, router, vLogs, RonFiExtractTypeStats)
	if len(swapPairsInfo) > 1 {
		head := swapPairsInfo[0]
		tail := swapPairsInfo[len(swapPairsInfo)-1]
		if head.TokenIn != rcommon.ZeroAddress && tail.TokenOut != rcommon.ZeroAddress &&
			head.TokenIn == tail.TokenOut && tail.AmountOut != nil &&
			head.AmountIn != nil && tail.AmountOut.Cmp(head.AmountIn) >= 0 {
			if head.Address != tail.Address || len(swapPairsInfo) > 2 {
				isObs = true
			}
		}
		if !isObs {
			return
		}

		// loops linkage check
		for i := 1; i < len(swapPairsInfo); i++ {
			tail = swapPairsInfo[i]
			if head.TokenOut == rcommon.ZeroAddress ||
				tail.TokenIn == rcommon.ZeroAddress ||
				head.TokenOut != tail.TokenIn ||
				head.To == rcommon.ZeroAddress ||
				(head.To != tail.Address && head.To != *to) {
				isObs = false
				break
			}
			head = tail
		}
	}

	return
}

func (di *Info) GetArbTxProfit(tx *types.Transaction, vLogs []*types.Log, router common.Address) (float64, bool) {
	v3Loop := false
	swapPairsInfo := di.ExtractSwapPairInfo(nil, nil, tx, router, vLogs, RonFiExtractTypeStats)
	for _, pairInfo := range swapPairsInfo {
		if pairInfo.V3 {
			v3Loop = true
		}
	}
	length := len(swapPairsInfo)
	if length > 1 {
		loops := 0
		totalProfit := 0.0
		head := 0
		for i := head; i < length; i++ {
			if profit, ok := di.checkIfLoop(swapPairsInfo[head : i+1]); ok {
				totalProfit += profit
				head = i + 1
				loops++
			}
		}

		return totalProfit, v3Loop
	}

	return 0.0, v3Loop
}

func (di *Info) checkIfLoop(swapPairsInfo []*SwapPairInfo) (profit float64, ok bool) {
	if len(swapPairsInfo) < 2 {
		return 0.0, false
	}

	head := swapPairsInfo[0]
	tail := swapPairsInfo[len(swapPairsInfo)-1]
	token := head.TokenIn

	isTradableToken := false
	isEqual := false
	amount := 0.0
	defer func() {
		// finally, we got the real profit
		if isTradableToken {
			price := GetTradingTokenPrice(token)
			amount /= price
			profit, ok = amount*GetTradingTokenPrice(rcommon.USDC), true
		} else {
			profit, ok = 0.0, false
		}
	}()

	if head.TokenIn == tail.TokenOut {
		// 1st try with tail.tokenOut
		_, isTradableToken = rcommon.TradableTokens[token]
		if tail.AmountOut.Cmp(head.AmountIn) == 0 {
			isEqual = true
		}
		if isTradableToken && !isEqual { //todo: what if this token has token fee?
			if isTradableToken {
				decimals := uint64(18)
				if tInfo := di.GetTokenInfo(token); tInfo != nil {
					decimals = tInfo.Decimals
				}
				amount = rcommon.TokenToFloat(new(big.Int).Sub(tail.AmountOut, head.AmountIn), decimals) // profit on tail
			}
			return
		}

		// 2nd try and more tries with (this.tokenOut - next.tokenIn)
		for i := len(swapPairsInfo) - 2; i >= 0; i-- {
			this := swapPairsInfo[i]
			next := swapPairsInfo[i+1]
			if this.TokenOut != next.TokenIn { // not a correct loop
				return
			}

			token = this.TokenOut
			_, isTradableToken = rcommon.TradableTokens[token]
			if this.AmountOut.Cmp(next.AmountIn) != 0 {
				isEqual = false
			}
			if isTradableToken && !isEqual {
				if isTradableToken {
					decimals := uint64(18)
					if tInfo := di.GetTokenInfo(token); tInfo != nil {
						decimals = tInfo.Decimals
					}
					amount = rcommon.TokenToFloat(new(big.Int).Sub(this.AmountOut, next.AmountIn), decimals) // profit on this
				}
				return
			}
		}
	}
	return
}
