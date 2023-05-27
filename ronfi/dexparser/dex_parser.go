package dexparser

import (
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"math/big"
	"strings"
)

// Caution: all these methods MUST IN LINE WITH DexMethodsTypical
const (
	SwapExactTokensForTokens = 0x38ed1739
	SwapExactETHForTokens    = 0x7ff36ab5
	SwapExactBNBForTokens    = 0x9cf68911
	SwapExactTokensForETH    = 0x18cbafe5
	SwapExactTokensForBNB    = 0x5d616c5b

	SwapExactTokensForTokensSupportingFeeOnTransferTokens = 0x5c11d795
	SwapExactETHForTokensSupportingFeeOnTransferTokens    = 0xb6f9de95
	//SwapExactBNBForTokensSupportingFeeOnTransferTokens    = ?
	SwapExactTokensForETHSupportingFeeOnTransferTokens = 0x791ac947
	SwapExactTokensForBNBSupportingFeeOnTransferTokens = 0xd46d2f83

	SwapTokensForExactTokens = 0x8803dbee
	SwapTokensForExactETH    = 0x4a25d94a
	SwapETHForExactTokens    = 0xfb3bdb41
	SwapBNBForExactTokens    = 0x8332a963
	SwapTokensForExactBNB    = 0xd67b571e

	//swapTokensForExactTokensSupportingFeeOnTransferTokens = ?
	//swapTokensForExactETHSupportingFeeOnTransferTokens	= ?
	//swapETHForExactTokensSupportingFeeOnTransferTokens	= ?
	//swapBNBForExactTokensSupportingFeeOnTransferTokens	= ?
	//swapTokensForExactBNBSupportingFeeOnTransferTokens	= ?

	PanCakeSmartRouterSwap      = 0xa6cbf417
	PanCakeSmartRouterSwapMulti = 0xca24fde8
)

// ParsableDexMethods Caution: MUST IN LINE WITH above consts
var (
	ParsableDexMethods = map[uint32]string{
		SwapExactTokensForTokens: "swapExactTokensForTokens", // amountIn fixed
		SwapExactETHForTokens:    "swapExactETHForTokens",    // amountIn fixed,	ETH -> TOK
		SwapExactBNBForTokens:    "swapExactBNBForTokens",    // amountIn fixed
		SwapExactTokensForETH:    "swapExactTokensForETH",    // amountIn fixed,	TOK -> ETH
		SwapExactTokensForBNB:    "swapExactTokensForBNB",    // amountIn fixed,	TOK -> BNB
		//
		SwapExactTokensForTokensSupportingFeeOnTransferTokens: "swapExactTokensForTokensSupportingFeeOnTransferTokens",
		SwapExactETHForTokensSupportingFeeOnTransferTokens:    "swapExactETHForTokensSupportingFeeOnTransferTokens",
		//?       : "swapExactBNBForTokensSupportingFeeOnTransferTokens",
		SwapExactTokensForETHSupportingFeeOnTransferTokens: "swapExactTokensForETHSupportingFeeOnTransferTokens",
		SwapExactTokensForBNBSupportingFeeOnTransferTokens: "swapExactTokensForBNBSupportingFeeOnTransferTokens",
		//
		SwapTokensForExactTokens: "swapTokensForExactTokens", // amountOut fixed
		SwapTokensForExactETH:    "swapTokensForExactETH",    // amountOut fixed,	TOK -> ETH
		SwapETHForExactTokens:    "swapETHForExactTokens",    // amountOut fixed,	ETH -> TOK
		SwapBNBForExactTokens:    "swapBNBForExactTokens",    // amountOut fixed
		SwapTokensForExactBNB:    "swapTokensForExactBNB",    // amountOut fixed, TOK -> BNB

		PanCakeSmartRouterSwap:      "swap",      //PanCakeSmartRouterSwap
		PanCakeSmartRouterSwapMulti: "swapMulti", //PanCakeSmartRouterSwapMulti
	}
)

var (
	PancakeSwapSmartRouter = common.HexToAddress("0xC6665d98Efd81f47B03801187eB46cbC63F328B0")
)

type SwapParameterInfo struct {
	Router common.Address
	Path   []common.Address
	rcommon.AmountIO
}

func (p *SwapParameterInfo) ToKey() string {
	var str strings.Builder
	str.WriteString(p.Router.String())
	for _, token := range p.Path {
		str.WriteString("-")
		str.WriteString(token.String())
	}
	return str.String()
}

func (p *SwapParameterInfo) SymPath(di *defi.Info) string {
	var str strings.Builder
	for i, token := range p.Path {
		sym := ""
		if tInfo := di.GetTokenInfo(token); tInfo != nil {
			sym = tInfo.Symbol
		}
		str.WriteString(sym)
		if i != len(p.Path)-1 {
			str.WriteString("->")
		}
	}
	return str.String()
}

// ParseDexTxParam parseDexTxParam return parsed parameter data
func ParseDexTxParam(tx *types.Transaction) (SwapParameterInfo, bool) {
	to := tx.To()
	if to == nil {
		return SwapParameterInfo{}, false
	}

	switch *to {
	//case OneInchV3, OneInchV4:
	//	return parseOneInchMethod(w, pairsDb, tx, false)
	//case MetaMaskRouter:
	//	return parseMetaMaskMethod(w, pairsDb, tx)
	//case Agg1Router:
	//	return parseAgg1Method(w, pairsDb, tx)
	//case ZeroExRouter:
	//	return parseZeroXMethod(w, pairsDb, tx, false)
	case PancakeSwapSmartRouter:
		return parsePancakeSmartRouterMethod(tx)
	default:
		return parseUniswapMethod(tx)
	}
}

//// parseMetaMaskMethod return parsed data and noMatchedPair
//func parseMetaMaskMethod(w *Worker, pairsDb []PairInfoMap, tx *types.Transaction) []*SwapPairInfo {
//	if tx == nil {
//		return []*SwapPairInfo{}
//	}
//
//	data := tx.Data()
//	methodID := uint64(binary.BigEndian.Uint32(data[:4]))
//	if methodID != uint64(MetaMaskSwap) {
//		return []*SwapPairInfo{}
//	} else {
//		if len(data) < 512 {
//			return []*SwapPairInfo{}
//		}
//		aggrIdEndPos := bytes.Index(data[164:196], MetaMaskAggIdHint)
//		aggrId := data[164 : 164+aggrIdEndPos]
//		if string(aggrId) == MetaMask0x {
//			return parseZeroXMethod(w, pairsDb, tx, true)
//		} else if string(aggrId) == MetaMask1InchV4 {
//			return parseOneInchMethod(w, pairsDb, tx, true)
//		}
//
//		return []*SwapPairInfo{}
//	}
//}
//
//// parseZeroXMethod
//func parseZeroXMethod(w *Worker, pairsDb []PairInfoMap, tx *types.Transaction, fromMetaMask bool) []*SwapPairInfo {
//	var (
//		data []byte
//
//		amountIn, amountOut *big.Int
//	)
//
//	if tx == nil {
//		return []*SwapPairInfo{}
//	}
//
//	if !fromMetaMask {
//		data = tx.Data()
//	} else {
//		data = tx.Data()[516:]
//	}
//
//	methodID := uint64(binary.BigEndian.Uint32(data[:4]))
//	offset := 4
//
//	if methodID == uint64(ZeroExMethod1) {
//		data = data[offset:]
//		amountIn = new(big.Int).SetBytes(data[32:64])
//		router := ZeroXDefaultRouter
//		rounterIndex := binary.BigEndian.Uint64(data[120:128])
//		if upRouter, ok := ZeroXRouterMap[rounterIndex]; ok {
//			router = upRouter
//		}
//		tokenNum := binary.BigEndian.Uint64(data[152:160])
//		if tokenNum < 2 {
//			log.Info("RonFi parseZeroXMethod: number of tokens less than 2", "dexTx", tx.Hash().String(), "tokenNum", tokenNum)
//		} else {
//			var tokenIn, tokenOut *common.Address
//			pos := 192
//			swapPairsInfo := make([]*SwapPairInfo, 0, tokenNum-1)
//
//			tokenIn = BytesToAddress(data[pos-32 : pos])
//			if *tokenIn == OneInchNullAddress {
//				tokenIn = &WBNB
//			}
//
//			for i := uint64(0); i < tokenNum-1; i++ {
//				tokenOut = BytesToAddress(data[pos : pos+32])
//				if *tokenOut == OneInchNullAddress {
//					tokenOut = &WBNB
//				}
//
//				aIn, aOut := swapPairAmounts(int(i), int(tokenNum-1), amountIn, amountOut)
//				swapPairInfo := buildSwapPairInfo(pairsDb, &router, tx, nil, tokenIn, tokenOut, aIn, aOut)
//				if swapPairInfo != nil {
//					swapPairsInfo = append(swapPairsInfo, swapPairInfo)
//				} else {
//					return []*SwapPairInfo{} // one pair fail, all fail
//				}
//
//				tokenIn = tokenOut
//				pos += 32
//			}
//
//			return swapPairsInfo
//		}
//	} else if methodID == uint64(ZeroExMethod2) {
//		data = data[offset:]
//		tokenIn := BytesToAddress(data[:32])
//		if *tokenIn == OneInchNullAddress {
//			tokenIn = &WBNB
//		}
//		tokenOut := BytesToAddress(data[32:64])
//		if *tokenOut == OneInchNullAddress {
//			tokenOut = &WBNB
//		}
//		amountIn = new(big.Int).SetBytes(data[64:96])
//		pairsMap := make(map[common.Address]int, 10)
//		for i := 0; i < int(binary.BigEndian.Uint64(data[184:192]))-1; i++ {
//			pos := int(binary.BigEndian.Uint64(data[224+i*32+24 : 224+i*32+32]))
//			addr := BytesToAddress(data[pos-32 : pos])
//			if _, ok := RouterFactoryMap[*addr]; !ok {
//				addr = BytesToAddress(data[pos-96 : pos-64])
//				if _, ok := RouterFactoryMap[*addr]; !ok {
//					continue
//				} else {
//					pairsMap[*addr] = pos - 32
//				}
//			} else {
//				pairsMap[*addr] = pos + 32
//			}
//		}
//
//		if len(pairsMap) > 0 {
//			swapPairsInfo := make([]*SwapPairInfo, 0, len(pairsMap))
//			for router, pos := range pairsMap {
//				tokenNum := binary.BigEndian.Uint64(data[pos+24 : pos+32])
//				if tokenNum >= 2 {
//					var tokenIn, tokenOut *common.Address
//					offset := pos + 64
//					tokenIn = BytesToAddress(data[offset-32 : offset])
//					if *tokenIn == OneInchNullAddress {
//						tokenIn = &WBNB
//					}
//
//					for i := uint64(0); i < tokenNum-1; i++ {
//						tokenOut = BytesToAddress(data[offset : offset+32])
//						if *tokenOut == OneInchNullAddress {
//							tokenOut = &WBNB
//						}
//
//						aIn, aOut := swapPairAmounts(int(i), int(tokenNum-1), amountIn, amountOut)
//						swapPairInfo := buildSwapPairInfo(pairsDb, &router, tx, nil, tokenIn, tokenOut, aIn, aOut)
//						if swapPairInfo != nil {
//							swapPairsInfo = append(swapPairsInfo, swapPairInfo)
//						} else {
//							return []*SwapPairInfo{} // one pair fail, all fail
//						}
//
//						tokenIn = tokenOut
//						offset += 32
//					}
//				}
//			}
//
//			return swapPairsInfo
//		} else {
//			log.Info("RonFi parseZeroXMethod ZeroExMethod2: parse router failed!!!", "tx", tx.Hash().String(), "metamask", fromMetaMask)
//		}
//	}
//
//	return []*SwapPairInfo{}
//}
//
//// parseOneInchMethod return parsed data and noMatchedPair
//func parseOneInchMethod(w *Worker, pairsDb []PairInfoMap, tx *types.Transaction, fromMetaMask bool) []*SwapPairInfo {
//	var (
//		data []byte
//		amountIn, amountOut *big.Int
//	)
//	if tx == nil {
//		return []*SwapPairInfo{}
//	}
//
//	if !fromMetaMask {
//		data = tx.Data()
//	} else {
//		data = tx.Data()[516:]
//	}
//
//	methodID := uint64(binary.BigEndian.Uint32(data[:4]))
//	switch methodID {
//	case OneInchSwap1:
//		if len(data) < 516 {
//			return []*SwapPairInfo{}
//		}
//		pairsNum := bytes.Count(data, OneInchSwapKW)
//		if pairsNum == 0 {
//			return []*SwapPairInfo{}
//		}
//
//		//tokenIn := BytesToAddress(data[100:132])
//		//if *tokenIn == OneInchNullAddress {
//		//	tokenIn = &WBNB
//		//}
//		//tokenOut := BytesToAddress(data[132:164])
//		//if *tokenOut == OneInchNullAddress {
//		//	tokenOut = &WBNB
//		//}
//
//		// srcReceiver data[164:196]
//		// dstReceiver data[196:228]
//		amountIn = new(big.Int).SetBytes(data[228:260])
//		// minReturnAmount data[260:292]
//
//		var poses []int
//		curPos := 0
//		for i := 0; i < pairsNum; i++ {
//			index := bytes.Index(data[curPos:], OneInchSwapKW)
//			curPos += index + 4
//			poses = append(poses, curPos)
//		}
//
//		swapPairsInfo := make([]*SwapPairInfo, 0, len(poses))
//		for i, pos := range poses {
//			address := BytesToAddress(data[pos : pos+32])
//			tokenIn := BytesToAddress(data[pos+32 : pos+64])
//			tokenOut := BytesToAddress(data[pos+64 : pos+96])
//
//			aIn, aOut := swapPairAmounts(i, len(poses), amountIn, amountOut)
//			swapPairInfo := buildSwapPairInfo(pairsDb, tx.To(), tx, address, tokenIn, tokenOut, aIn, aOut)
//			if swapPairInfo != nil {
//				swapPairsInfo = append(swapPairsInfo, swapPairInfo)
//			} else {
//				return []*SwapPairInfo{} // one pair fail, all fail
//			}
//		}
//
//		return swapPairsInfo
//	case OneInchSwap2:
//		if len(data) < 228 {
//			return nil
//		}
//		swapPairsInfo := make([]*SwapPairInfo, 0, 1)
//		address := BytesToAddress(data[100:132])		// order.maker
//		tokenIn := BytesToAddress(data[68:100])			// order.takerAsset
//		tokenOut := BytesToAddress(data[36:68])			// order.makerAsset
//		// data[162:196] // order.makingAmount
//		amountIn = new(big.Int).SetBytes(data[196:228])	// order.takingAmount
//
//		swapPairInfo := buildSwapPairInfo(pairsDb, tx.To(), tx, address, tokenIn, tokenOut, amountIn, amountOut)
//		if swapPairInfo != nil {
//			swapPairsInfo = append(swapPairsInfo, swapPairInfo)
//		} else {
//			return []*SwapPairInfo{} // one pair fail, all fail
//		}
//		return swapPairsInfo
//	}
//
//	return []*SwapPairInfo{}
//}
//
//// parseAgg1Method return parsed data and noMatchedPair
//func parseAgg1Method(w *Worker, pairsDb []PairInfoMap, tx *types.Transaction) []*SwapPairInfo {
//	var (
//		data []byte
//		amountIn, amountOut *big.Int
//	)
//	if tx == nil {
//		return []*SwapPairInfo{}
//	}
//
//	data = tx.Data()
//	methodID := uint64(binary.BigEndian.Uint32(data[:4]))
//	router := BytesToAddress(data[36:68])
//	amountIn = new(big.Int).SetBytes(data[100:132])
//
//	if _, exist := RouterFactoryMap[*router]; !exist {
//		log.Info("RonFi parseAgg1Method on router unknown", "router", router)
//		return []*SwapPairInfo{} // early return here if a router can not be parsed. i.e. 'PairFor(router, token0, token1)' does not work.
//	}
//
//	if methodID == uint64(Agg1Method) {
//		if len(data) < 260 {
//			return []*SwapPairInfo{}
//		}
//		pairsNum := uint64(data[259]) - 1
//		if pairsNum == 0 {
//			return []*SwapPairInfo{}
//		}
//
//		var tokenIn, tokenOut *common.Address
//		pos := uint64(32)
//		data = data[260:]
//		tokenIn = BytesToAddress(data[pos-32 : pos])
//		if *tokenIn == ZeroAddress {
//			tokenIn = &WBNB
//		}
//		swapPairsInfo := make([]*SwapPairInfo, 0, pairsNum)
//
//		for i := uint64(0); i < pairsNum; i++ {
//			tokenOut = BytesToAddress(data[pos : pos+32])
//			if *tokenOut == ZeroAddress {
//				tokenOut = &WBNB
//			}
//
//			aIn, aOut := swapPairAmounts(int(i), int(pairsNum), amountIn, amountOut)
//			swapPairInfo := buildSwapPairInfo(pairsDb, tx.To(), tx, nil, tokenIn, tokenOut, aIn, aOut)
//			if swapPairInfo != nil {
//				swapPairsInfo = append(swapPairsInfo, swapPairInfo)
//			} else {
//				return []*SwapPairInfo{} // one pair fail, all fail
//			}
//
//			tokenIn = tokenOut
//			pos += 32
//		}
//
//		return swapPairsInfo
//	}
//
//	return []*SwapPairInfo{}
//}

func parseUniswapMethod(tx *types.Transaction) (SwapParameterInfo, bool) {
	var (
		data                   []byte
		pathLength, pathOffset uint64
		amountIn, amountOut    *big.Int
	)

	if tx == nil {
		return SwapParameterInfo{}, false
	}

	data = tx.Data()
	if len(data) < 228 {
		return SwapParameterInfo{}, false
	}

	methodID := binary.BigEndian.Uint32(data[:4])
	if _, parsable := ParsableDexMethods[methodID]; !parsable {
		return SwapParameterInfo{}, false // early return if unable to parse
	}

	switch methodID {
	case SwapExactETHForTokens, SwapExactBNBForTokens, SwapExactETHForTokensSupportingFeeOnTransferTokens:
		//0:4      methodId
		//4:36     amountOutMin
		//36:68    path offset
		//68:100   to
		//100:132  deadline
		//132:164  path array length
		//164:     path array
		pathLength = uint64(binary.BigEndian.Uint32(data[160:164])) //data[132:164]
		pathOffset = 164
		amountIn = tx.Value()
		//amountOut = big.NewInt(0)
		//amountInMax = big.NewInt(0)
	case SwapExactTokensForETH,
		SwapExactTokensForETHSupportingFeeOnTransferTokens,
		SwapExactTokensForBNB,
		SwapExactTokensForBNBSupportingFeeOnTransferTokens,
		SwapExactTokensForTokens,
		SwapExactTokensForTokensSupportingFeeOnTransferTokens:
		//0:4      methodId
		//4:36     amountIn
		//36:68    amountOutMin
		//68:100   path offset
		//100:132  to
		//132:164  deadline
		//164:196  path array length
		//196:     path array
		pathLength = uint64(binary.BigEndian.Uint32(data[192:196])) //data[164:196]
		pathOffset = 196
		amountIn = new(big.Int).SetBytes(data[4:36])
		//amountOut = big.NewInt(0)
		//amountInMax = big.NewInt(0)
	case SwapETHForExactTokens,
		SwapBNBForExactTokens:
		//0:4      methodId
		//4:36     amountOut
		//36:68    path offset
		//68:100   to
		//100:132  deadline
		//132:164  path array length
		//164:     path array
		pathLength = uint64(binary.BigEndian.Uint32(data[160:164])) //data[132:164]
		pathOffset = 164
		//amountIn = big.NewInt(0)
		//amountOutMin = big.NewInt(0)
		amountOut = new(big.Int).SetBytes(data[4:36])
		//amountInMax = big.NewInt(0)
	case SwapTokensForExactETH,
		SwapTokensForExactBNB,
		SwapTokensForExactTokens:
		//0:4      methodId
		//4:36     amountOut
		//36:68    amountInMax
		//68:100   path offset
		//100:132  to
		//132:164  deadline
		//164:196  path array length
		//196:     path array
		pathLength = uint64(binary.BigEndian.Uint32(data[192:196])) //data[164:196]
		pathOffset = 196
		//amountIn = big.NewInt(0)
		//amountOutMin = big.NewInt(0)
		amountOut = new(big.Int).SetBytes(data[4:36])
	default:
		return SwapParameterInfo{}, false
	}

	//make sure pathLength reasonable
	if pathLength > 1 && pathLength < 100 {
		// pair length = path length - 1
		swapParamsInfo := SwapParameterInfo{Router: *tx.To()}
		if amountIn != nil {
			swapParamsInfo.AmountIn = amountIn
		} else if amountOut != nil {
			swapParamsInfo.AmountOut = amountOut
		} else {
			log.Warn("RonFi parseUniswapMethod fail, amount missed", "dexTx", tx.Hash().String(), "len(data)", len(data))
			return SwapParameterInfo{}, false // early return if amount missing
		}
		swapParamsInfo.Path = make([]common.Address, pathLength)
		var tokenIn, tokenOut common.Address
		pos := pathOffset + 32
		tokenIn = common.BytesToAddress(data[pos-32 : pos])
		swapParamsInfo.Path[0] = tokenIn
		for i := uint64(1); i < pathLength; i++ {
			if pos+32 > uint64(len(data)) {
				log.Warn("RonFi parseUniswapMethod fail", "dexTx", tx.Hash().String(), "i", i, "pos", pos, "len(data)", len(data))
				return SwapParameterInfo{}, false
			}
			tokenOut = common.BytesToAddress(data[pos : pos+32])
			swapParamsInfo.Path[i] = tokenOut

			tokenIn = tokenOut
			pos += 32
		}

		return swapParamsInfo, true
	}

	return SwapParameterInfo{}, false
}

func parsePancakeSmartRouterMethod(tx *types.Transaction) (SwapParameterInfo, bool) {
	var (
		data                   []byte
		pathLength, pathOffset uint64
		amountIn, amountOut    *big.Int
	)

	if tx == nil {
		return SwapParameterInfo{}, false
	}

	data = tx.Data()
	if len(data) < 100 {
		return SwapParameterInfo{}, false
	}

	methodID := binary.BigEndian.Uint32(data[:4])
	if _, parsable := ParsableDexMethods[methodID]; !parsable {
		return SwapParameterInfo{}, false // early return if unable to parse
	}

	switch methodID {
	case PanCakeSmartRouterSwap:
		{
			pos := 4
			tokenIn := common.BytesToAddress(data[pos : pos+32])
			pos += 32
			tokenOut := common.BytesToAddress(data[pos : pos+32])
			pos += 32
			amountIn = new(big.Int).SetBytes(data[pos : pos+32])
			swapParamsInfo := SwapParameterInfo{
				Router: *tx.To(),
				Path:   []common.Address{tokenIn, tokenOut},
			}
			swapParamsInfo.AmountIn = amountIn

			return swapParamsInfo, true
		}
	case PanCakeSmartRouterSwapMulti:
		//0:4      methodId
		//4:36     path offset
		//36:68    amountIn
		amountIn = new(big.Int).SetBytes(data[36:68])
		pathLength = uint64(binary.BigEndian.Uint32(data[160:164])) //data[132:164]
		pathOffset = 164
	default:
		return SwapParameterInfo{}, false
	}

	//make sure pathLength reasonable
	if pathLength > 1 && pathLength < 100 {
		// pair length = path length - 1
		swapParamsInfo := SwapParameterInfo{Router: *tx.To()}
		if amountIn != nil {
			swapParamsInfo.AmountIn = amountIn
		} else if amountOut != nil {
			swapParamsInfo.AmountOut = amountOut
		} else {
			log.Warn("RonFi parseUniswapMethod fail, amount missed", "dexTx", tx.Hash().String(), "len(data)", len(data))
			return SwapParameterInfo{}, false // early return if amount missing
		}
		swapParamsInfo.Path = make([]common.Address, pathLength)
		var tokenIn, tokenOut common.Address
		pos := pathOffset + 32
		tokenIn = common.BytesToAddress(data[pos-32 : pos])
		swapParamsInfo.Path[0] = tokenIn
		for i := uint64(1); i < pathLength; i++ {
			if pos+32 > uint64(len(data)) {
				log.Warn("RonFi parseUniswapMethod fail", "dexTx", tx.Hash().String(), "i", i, "pos", pos, "len(data)", len(data))
				return SwapParameterInfo{}, false
			}
			tokenOut = common.BytesToAddress(data[pos : pos+32])
			swapParamsInfo.Path[i] = tokenOut

			tokenIn = tokenOut
			pos += 32
		}

		return swapParamsInfo, true
	}

	return SwapParameterInfo{}, false
}
