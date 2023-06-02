package loops

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/uniswap"
	"sort"
)

type SwapLoop struct {
	TargetToken    common.Address
	Path           []common.Address
	PoolFee        []uint64
	TokenFee       []uint64
	Fees           []uint64  // used on RonSwapX contract
	RDelta         []float64 // used on profit searching
	RDeltaProd     float64   // used to speed up profit calculation
	GasNeeded      uint64    // used to speed up estimate net profit in validateLoops()
	Dirs           []uint64
	Index          []uint64
	Count          uint64
	HasNonBriToken bool

	Cancel bool // the arb tx with this loop has been cancelled because of applyTransaction fail.

	key     string
	LoopId  common.Hash
	FlashOk uint64
}

func (p *SwapLoop) toLoopId() common.Hash {
	return rcommon.ToLoopId(p.Path, p.PoolFee, p.TokenFee)
}

func (p *SwapLoop) toKey() string {
	return rcommon.ToKey(p.Path)
}

func (p *SwapLoop) Key() string {
	return p.key
}

func (p *SwapLoop) Copy() *SwapLoop {
	path := make([]common.Address, len(p.Path))
	for _, addr := range p.Path {
		path = append(path, addr)
	}
	poolFee := make([]uint64, len(p.PoolFee))
	copy(poolFee, p.PoolFee)
	tokenFee := make([]uint64, len(p.TokenFee))
	copy(tokenFee, p.TokenFee)
	fees := make([]uint64, len(p.Fees))
	copy(fees, p.Fees)
	rDelta := make([]float64, len(p.RDelta))
	copy(rDelta, p.RDelta)
	dirs := make([]uint64, len(p.Dirs))
	copy(dirs, p.Dirs)
	index := make([]uint64, len(p.Index))
	copy(index, p.Index)

	loop := &SwapLoop{
		TargetToken:    p.TargetToken,
		Path:           path,
		PoolFee:        poolFee,
		TokenFee:       tokenFee,
		Fees:           fees,
		RDelta:         rDelta,
		RDeltaProd:     p.RDeltaProd,
		GasNeeded:      p.GasNeeded,
		Dirs:           dirs,
		Index:          index,
		Count:          p.Count,
		HasNonBriToken: p.HasNonBriToken,
		LoopId:         p.LoopId,
		Cancel:         p.Cancel,
	}

	return loop
}

func (p *SwapLoop) ToV3ArbPath(pairGasMap map[string]uint64) uniswap.V3ArbPath {
	v3ArbPath := make([]*uniswap.UniswapPool, 0, len(p.Path))
	for i := 0; i < len(p.Path); i++ {
		addr := p.Path[i]
		if i%2 != 0 {
			pf := p.PoolFee[i/2]
			tf := p.TokenFee[i/2]
			dir := p.Dirs[i/2]
			pt := uniswap.UniswapV2
			if dir>>0x2 == 1 {
				pt = uniswap.UniswapV3
				pf = 0
			}
			dir = dir & 0x1

			gasNeeded, ok := pairGasMap[fmt.Sprintf("%s-%s", p.Path[i-1].Hex(), addr.Hex())]
			if !ok {
				gasNeeded = 150000
			}
			pool := uniswap.UniswapPool{
				PoolAddr: addr,
				PoolType: pt,
				PoolFee:  int(pf),
				TokenFee: int(tf),
				TickLens: rcommon.ZeroAddress,
				TokenIn:  p.Path[i-1],
				Dir:      int(dir),
				GasNeed:  gasNeeded,
			}
			v3ArbPath = append(v3ArbPath, &pool)
		}
	}

	return v3ArbPath
}

func (p *SwapLoop) HasV3Pool() bool {
	for i, dir := range p.Dirs {
		if dir>>0x2 == 1 && (i == 0 || i == len(p.Dirs)-1) {
			return true
		}
	}

	return false
}

type SwapLoops []*SwapLoop

func (s SwapLoops) Len() int {
	return len(s)
}

func (s SwapLoops) Less(i, j int) bool {
	if s[i] != nil && s[j] != nil {
		return s[i].Count > s[j].Count
	}

	return true
}

func (s SwapLoops) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type LDB map[string][]*SwapLoop
type LMap struct {
	LoopsDB      LDB
	Pairs        map[common.Address]uint64
	PairsArray   []common.Address
	AllPairsMap  defi.PairInfoMap
	UniPairsMap  defi.PairInfoMap
	AllTokensMap map[common.Address]*defi.TokenInfo
}

func (lpsMap *LMap) Copy() *LMap {
	// copy loopsDB
	loopsDB := make(LDB, len(lpsMap.LoopsDB))
	for k, v := range lpsMap.LoopsDB {
		swapLoops := make([]*SwapLoop, 0, len(v))
		copy(swapLoops, v)
		loopsDB[k] = v
	}

	// copy allPairs
	allPairs := make(map[common.Address]uint64, len(lpsMap.Pairs))
	for addr, val := range lpsMap.Pairs {
		allPairs[addr] = val
	}

	// copy pairsArray
	pairsArray := make([]common.Address, len(lpsMap.PairsArray))
	copy(pairsArray, lpsMap.PairsArray)

	// copy allPairsMap
	allPairsMap := make(defi.PairInfoMap, len(lpsMap.AllPairsMap))
	for addr, val := range lpsMap.AllPairsMap {
		allPairsMap[addr] = val
	}

	// copy uniPairsMap
	uniPairsMap := make(defi.PairInfoMap, len(lpsMap.UniPairsMap))
	for addr, val := range lpsMap.UniPairsMap {
		uniPairsMap[addr] = val
	}

	// copy allTokensMap
	allTokensMap := make(map[common.Address]*defi.TokenInfo, len(lpsMap.AllTokensMap))
	for addr, val := range lpsMap.AllTokensMap {
		allTokensMap[addr] = val
	}

	return &LMap{loopsDB, allPairs, pairsArray, allPairsMap, uniPairsMap, allTokensMap}
}

type LIdMap map[common.Hash]*SwapLoop

func (lpm LIdMap) Copy() LIdMap {
	newLoopsIdMap := make(LIdMap)
	for k, v := range lpm {
		newLoopsIdMap[k] = v
	}

	return newLoopsIdMap
}

func NewDefaultLoopsMap() *LMap {
	loopsDB := make(LDB)
	allPairs := make(map[common.Address]uint64)
	pairsArray := make([]common.Address, 0, 100)
	allPairsMap := make(defi.PairInfoMap)
	uniPairsMap := make(defi.PairInfoMap)
	allTokensMap := make(map[common.Address]*defi.TokenInfo)

	return &LMap{loopsDB, allPairs, pairsArray, allPairsMap, uniPairsMap, allTokensMap}
}

func sumGasNeeded(hops int, loopPath []common.Address, dirs []uint64, pairGasMap map[string]uint64) (uint64, bool) {
	sum := uint64(0)
	missed := false
	if hops >= len(rcommon.GasRough) {
		sum = uint64(100000 + 70000*(hops-1)) // a rough estimation
	} else {
		sum = rcommon.GasRough[hops]
	}
	for i := 1; i < len(loopPath); i += 2 {
		pair := loopPath[i]
		dir := dirs[(i-1)/2] & 1 // 'direction' field bit0: dir, bit1: swap type, etc.
		key := fmt.Sprintf("%s-%d", pair, dir)
		if gas, ok := pairGasMap[key]; ok {
			if gas > 150000 {
				// must be an exception pair, add the extra gas needed
				sum += gas - 150000
			}
		} else {
			missed = true
		}
	}
	return sum, missed
}

const (
	Undefined        = 1 << 31
	Ok               = 1 << 1
	Missed           = 1 << 2
	Cancel           = 1 << 3
	Patched          = 1 << 4
	FlashOkT         = 1 << 5
	FlashOkH         = 1 << 6
	WbnbLoops        = 1 << 7
	NotTradableToken = 1 << 8
	TokenFeeTooHigh  = 1 << 9
	GetFactoryFail   = 1 << 10
	PairLockError    = 1 << 11
)

func LoadSwapLoops(
	mysql *db.Mysql,
	di *defi.Info,
	loopsIdMap LIdMap,
	feePatchMap map[common.Address]uint64,
	pairGasMap map[string]uint64,
	flashNokPairs map[common.Address]uint64,
) *LMap {
	var (
		totalMissedPairGas, totalLoopsLoaded, totalLoopsOK, totalCancelLoops, totalPatchedLoops, totalFakeWbnbLoops, totalPairLockLoops, totalNotTradableLoops, totalFlashOkT, totalFlashOkH int
	)

	startTime := mclock.Now()
	loopsDB := make(LDB)
	allPairs := make(map[common.Address]uint64)
	allPairsMap := make(defi.PairInfoMap)
	uniPairsMap := make(defi.PairInfoMap)
	allTokensMap := make(map[common.Address]*defi.TokenInfo)

	allLoops := mysql.LoadLoops()
	totalLoopsLoaded = len(allLoops)

	for i, loop := range allLoops {
		if i > 5000 {
			break
		}
		if loop.Cancel {
			totalCancelLoops++
			continue
		}

		swapLoop, status := HandleSwapLoop(loop, loopsDB, allPairs, allPairsMap, uniPairsMap, allTokensMap, feePatchMap, pairGasMap, flashNokPairs, di)
		if (status & Ok) != 0 {
			totalLoopsOK++
			if swapLoop != nil {
				// todo only keep v3 loop here.
				if swapLoop.HasV3Pool() {
					loopId := swapLoop.LoopId
					loopsIdMap[loopId] = swapLoop // build a map from loopId to the true SwapLoop, ONLY in initialization stage!
					//Note: the latter one with same loopId will overwrite the former one.
				}
			}
		}
		if (status & Missed) != 0 {
			totalMissedPairGas++
		}
		if (status & Patched) != 0 {
			totalPatchedLoops++
		}
		if (status & WbnbLoops) != 0 {
			totalFakeWbnbLoops++
		}
		if status&FlashOkT != 0 {
			totalFlashOkT++
		}
		if status&FlashOkH != 0 {
			totalFlashOkH++
		}
		if status&PairLockError != 0 {
			totalPairLockLoops++
		}
		if status&NotTradableToken != 0 {
			totalNotTradableLoops++
		}
	}

	totalSizeOfLoopsDB := 0
	for _, loops := range loopsDB {
		sort.Sort(SwapLoops(loops))
		totalSizeOfLoopsDB += len(loops)
	}
	log.Info("RonFi LoadSwapLoops", "loopsTotal", totalLoopsOK, "loopsDB", totalSizeOfLoopsDB)

	i := 0
	pairsArray := make([]common.Address, len(allPairs))
	for pair := range allPairs {
		pairsArray[i] = pair
		i++
	}
	log.Info("RonFi LoadSwapLoops success",
		"allPairs", len(allPairs),
		"loops", totalLoopsLoaded,
		"loopsOk", totalLoopsOK,
		"loopsCancel", totalCancelLoops,
		"loopsMissedPairGas", totalMissedPairGas,
		"loopsPatched", totalPatchedLoops,
		"loopsWbnb'", totalFakeWbnbLoops,
		"loopsPairLockErr", totalPairLockLoops,
		"loopsNotTradable", totalNotTradableLoops,
		"loopsFlashOkT", totalFlashOkT,
		"loopsFlashOkH", totalFlashOkH,
		"elapsed", mclock.Since(startTime))

	// show the top 10 flash swap NOK factory
	if flashNokPairs != nil {
		type kv struct {
			Key   common.Address
			Value uint64
		}
		ss := make([]kv, 0, len(flashNokPairs))
		for k, v := range flashNokPairs {
			ss = append(ss, kv{k, v})
		}
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value > ss[j].Value
		})
		for j := 0; j < 10; j++ {
			if j >= len(ss)-1 {
				break
			}
			log.Info("RonFi flashNok factory", "no.", j+1, "factory", ss[j].Key, "number", ss[j].Value)
		}
	}

	return &LMap{loopsDB, allPairs, pairsArray, allPairsMap, uniPairsMap, allTokensMap}
}

func HandleSwapLoop(
	loop *db.DBLoop,
	loopsDB LDB,
	allPairs map[common.Address]uint64,
	allPairsMap defi.PairInfoMap,
	uniPairsMap defi.PairInfoMap,
	allTokensMap map[common.Address]*defi.TokenInfo,
	feePatchMap map[common.Address]uint64,
	pairGasMap map[string]uint64,
	flashNokPairs map[common.Address]uint64,
	defiInfo *defi.Info,
) (loopRet *SwapLoop, status uint32) {
	status = Undefined

	if len(loop.Path)%2 != 0 {
		log.Warn("RonFi HandleSwapLoop exception loops, path length must be even", "loopId", loop.LoopId)
		return
	}
	hops := len(loop.Path) / 2
	//skip surprising hops
	if hops > rcommon.HopsMaxAllowed {
		log.Warn("RonFi HandleSwapLoop exception loops, too many hops", "loopId", loop.LoopId, "hops", hops)
		return
	}

	var swapLoop SwapLoop
	if hops != len(loop.PoolFee) || hops != len(loop.TokenFee) || hops != len(loop.Direction) || hops != len(loop.Index) {
		log.Warn("RonFi HandleSwapLoop exception loops, arrays size not same", "loopId", loop.LoopId)
		return
	}

	validityCheckOk := true
	swapLoop.Dirs = make([]uint64, len(loop.Direction))
	copy(swapLoop.Dirs, loop.Direction)
	//for _, dir := range swapLoop.Dirs {
	//	if dir > 3 {
	//		validityCheckOk = false
	//		break
	//	}
	//}

	swapLoop.Index = make([]uint64, len(loop.Index))
	copy(swapLoop.Index, loop.Index)

	swapLoop.Count = loop.Count
	//if !validityCheckOk {
	//	log.Warn("RonFi LoadSwapLoops exception loops, direction > 3", "loopId", loop.LoopId)
	//	return
	//}

	path := make([]common.Address, 0, len(loop.Path))
	for i, addr := range loop.Path {
		path = append(path, addr)
		if addr == rcommon.ZeroAddress {
			if !(i == 1 || i == len(loop.Path)-1) {
				// 0x00..01 special 'pair' in the middle of a loop, that's impossible and not allowed
				validityCheckOk = false
				break
			}
		}
	}
	swapLoop.Path = path
	if !validityCheckOk {
		log.Warn("RonFi LoadSwapLoops exception loops, 0x00..01 in middle", "loopId", loop.LoopId)
		return
	}

	// Generate loop Key
	swapLoop.key = swapLoop.toKey()

	// Deep shallow copy of the array
	swapLoop.PoolFee = make([]uint64, len(loop.PoolFee))
	copy(swapLoop.PoolFee, loop.PoolFee)
	swapLoop.TokenFee = make([]uint64, len(loop.TokenFee))
	copy(swapLoop.TokenFee, loop.TokenFee)

	// Generate loop id
	swapLoop.LoopId = swapLoop.toLoopId()
	if swapLoop.LoopId != loop.LoopId {
		validityCheckOk = false
	}
	if !validityCheckOk {
		log.Warn("RonFi LoadSwapLoops exception loops, wrong loopId", "loopId", loop.LoopId)
		return
	}

	// Apply Fee Patches
	patched := false
	for i, p := range loop.Path {
		if patch, exist := feePatchMap[p]; exist {
			if i%2 == 0 { // token
				if patch != swapLoop.TokenFee[i/2] {
					swapLoop.TokenFee[i/2] = patch
					patched = true
				}
			} else { // pair
				if patch != swapLoop.PoolFee[i/2] {
					swapLoop.PoolFee[i/2] = patch
					patched = true
				}
			}
		}
	}
	if patched {
		status |= Patched
	}

	//swapLoop.Fees = make([]uint64, len(loop.PoolFee))
	//swapLoop.RDelta = make([]float64, len(loop.PoolFee))
	//rDeltaProd := 1.0
	//for i := 0; i < len(loop.PoolFee); i++ {
	//	pf := loop.PoolFee[i]
	//	tf := loop.TokenFee[i]
	//	if path[2*i+1] != rcommon.ZeroAddress { // normal pair
	//		fee := pf + tf - pf*tf/10000
	//		if fee >= 6000 { // 60% fee is impossible, sth must be wrong!
	//			validityCheckOk = false
	//			break
	//		}
	//		swapLoop.Fees[i] = fee
	//		rDelta := 1.0 / (1.0 - float64(fee)/10000)
	//		swapLoop.RDelta[i] = rDelta
	//		rDeltaProd *= rDelta
	//	} else { // fake pair as wbnb' deposit/withdraw which does not have a fee
	//		//swapLoop.fees[i] = 0
	//		//swapLoop.rDelta[i] = 1.0
	//		status |= WbnbLoops
	//		return // now we don't want these wbnb' loops
	//	}
	//}
	//swapLoop.RDeltaProd = rDeltaProd // 1/d_1 * 1/d_2 * 1/d_3 * ...
	//if !validityCheckOk {
	//	status |= TokenFeeTooHigh
	//	log.Warn("RonFi HandleSwapLoop exception loops, fee > 60%", "loopId", loop.LoopId)
	//	return
	//}

	//first path is the target token
	swapLoop.TargetToken = path[0]

	//pre-calculation of the gas needed for this loop
	missed := false
	swapLoop.GasNeeded, missed = sumGasNeeded(hops, path, swapLoop.Dirs, pairGasMap)

	//skip path which has a target not in my TradableTokens list
	if _, ok := rcommon.TradableTokens[swapLoop.TargetToken]; !ok {
		status |= NotTradableToken
		log.Warn("RonFi HandleSwapLoop exception loops, NotTradableToken", "loopId", loop.LoopId)
		return
	}

	//check whether same pair used multiple times in one loop
	{
		for i := 0; i < hops-1; i++ {
			thisPair := path[2*i+1]
			for j := i + 1; j < hops; j++ {
				otherPair := path[2*j+1]
				if thisPair == otherPair {
					status |= PairLockError
					log.Warn("RonFi HandleSwapLoop exception loops, PairLockError", "loopId", loop.LoopId)
					return
				}
			}
		}
	}

	//check whether this loop can use flashSwap
	{
		twoFactory := make([]common.Address, 0, 2)
		for i := 0; i < 2; i++ {
			var (
				factory common.Address
				pair    common.Address
				index   = 0
			)
			if i == 0 {
				index = hops - 1 // check last hop
			} else {
				index = 0 // check 1st hop
			}
			pair = path[2*index+1]
			pairInfo := defiInfo.GetPairInfo(pair)
			if pairInfo != nil && pairInfo.Factory != rcommon.ZeroAddress {
				twoFactory = append(twoFactory, factory)
			} else {
				status |= GetFactoryFail
			}

			//use swap() function to check whether this pair support flash swap, i.e. if swap(..., bytes calldata data), then flash swap is ok.
			if swapLoop.Dirs[index]&2 == 0 {
				if i == 0 {
					status |= FlashOkT   // last pair flash swap ok
					swapLoop.FlashOk = 1 // 1: flash on last pair
				} else {
					status |= FlashOkH   // first pair flash swap ok
					swapLoop.FlashOk = 2 // 2: flash on first pair
				}
				break
			}
		}
		if flashNokPairs != nil && (status&FlashOkT == 0) && (status&FlashOkH == 0) {
			for _, factory := range twoFactory {
				if old, exist := flashNokPairs[factory]; exist {
					flashNokPairs[factory] = old + 1 // count the times of this factory
				} else {
					flashNokPairs[factory] = 1
				}
			}
		}
	}

	//build (pair+dir)s trie
	swapLoop.HasNonBriToken = false
	pairs := make([]common.Address, hops)
	for i := 0; i < hops; i++ {
		token := path[2*i]
		if _, ok := rcommon.BridgeTokens[token]; !ok {
			swapLoop.HasNonBriToken = true
		}
		if tokenInfo, ok := allTokensMap[token]; !ok {
			info := defiInfo.GetTokenInfo(token)
			if info == nil {
				continue
			}
			info.Fee = swapLoop.TokenFee[i]
			allTokensMap[token] = info
		} else {
			orgFee := tokenInfo.Fee
			newFee := swapLoop.TokenFee[i]
			if newFee > orgFee {
				tokenInfo.Fee = newFee
				allTokensMap[token] = tokenInfo
			}
		}

		pair := path[2*i+1]
		if index, ok := allPairs[pair]; !ok {
			allPairs[pair] = swapLoop.Index[i]
		} else if index != swapLoop.Index[i] {
			log.Warn("RonFi LoadNewSwapLoops conflict index in json", "loopId", loop.LoopId, "pair", pair, "old", index, "new", swapLoop.Index[i])
		}

		pairs[i] = pair
		if pairInfo, ok := allPairsMap[pair]; !ok {
			var token0, token1, tmp common.Address
			if i != hops-1 {
				token0 = path[2*i]
				token1 = path[2*(i+1)]
			} else {
				token0 = path[2*i]
				token1 = path[0]
			}

			// swap
			if loop.Direction[i]&1 == 1 {
				tmp = token0
				token0 = token1
				token1 = tmp
			}

			info := defiInfo.GetPairInfo(pair)
			if info == nil {
				continue
			}
			info.Fee = swapLoop.PoolFee[i]
			allPairsMap[pair] = info

			if loop.Direction[i] <= 1 {
				uniPairsMap[pair] = info
			}
		} else {
			orgFee := pairInfo.Fee
			newFee := swapLoop.PoolFee[i]
			if newFee > orgFee {
				pairInfo.Fee = newFee
				allPairsMap[pair] = pairInfo
			}
		}
	}

	// Assign return SwapLoop
	loopRet = &swapLoop

	for i := 0; i < len(pairs); i++ {
		key := fmt.Sprintf("%s-%d", pairs[i], loop.Direction[i]&1)

		loops := []*SwapLoop{&swapLoop}
		if tmp, ok := loopsDB[key]; ok {
			loops = deDupLoops(tmp, &swapLoop)
		}
		loopsDB[key] = loops
	}

	if missed {
		status |= Missed
	}

	status |= Ok
	return
}

func deDupLoops(loops []*SwapLoop, newLoop *SwapLoop) []*SwapLoop {
	newLoops := make([]*SwapLoop, 0, len(loops)+1)

	// duplicated loops from existing 'loops' removing
	for _, loop := range loops {
		if loop.Key() != newLoop.Key() {
			newLoops = append(newLoops, loop)
		}
	}
	// new loops (or latest loops) always taken
	newLoops = append(newLoops, newLoop)

	return newLoops
}

// MatchLoops search the best-matching loops for profit
func MatchLoops(swapPairsInfo []*defi.SwapPairInfo, loopsMap *LMap, newLoopsMap *LMap, loopsWanted int) (matchedPairsLoops map[int][]*SwapLoop) {
	deDup := make(map[string]struct{})
	matchedPairsLoops = make(map[int][]*SwapLoop)
	loopsDBs := make([]LDB, 0, 2)
	if newLoopsMap != nil {
		loopsDBs = append(loopsDBs, newLoopsMap.LoopsDB)
	}
	if loopsMap != nil {
		loopsDBs = append(loopsDBs, loopsMap.LoopsDB)
	}

	// search loops matched with keypair
	for index, swapPairInfo := range swapPairsInfo {
		if swapPairInfo.V3 {
			continue
		}
		coupPairsMap := make(map[common.Address]struct{})
		matchedLoops := make([]*SwapLoop, 0, loopsWanted)
		loopsMaxReached := false
		pairAddr := swapPairInfo.Address
		keyToken := swapPairInfo.KeyToken

		for _, loopsDB := range loopsDBs {
			if loops := loopsDB[swapPairInfo.Key]; loops != nil && !loopsMaxReached {
				for _, loop := range loops {
					var (
						keyTokenPos, keyPairPos, coupPairPos int
						validLoop                            bool
					)

					if loop.Cancel { // if a cancel happen once, this loop will be cancelled also, never use it anymore!
						continue
					}

					// pair contain one bridging token
					if keyToken != rcommon.ZeroAddress {
						for i := 0; i < len(loop.Path); i++ {
							addr := loop.Path[i]
							if addr == keyToken {
								keyTokenPos = i
							}
							if addr == pairAddr {
								keyPairPos = i
							}
						}

						// targetToken shouldn't be the loop trading token(in position 0).
						if keyTokenPos == 0 {
							continue
						}

						if keyTokenPos < keyPairPos {
							coupPairPos = keyTokenPos - 1
						} else {
							coupPairPos = keyTokenPos + 1
						}
						if coupPairPos >= len(loop.Path) {
							continue
						}
						coupPair := loop.Path[coupPairPos]
						if _, exist := coupPairsMap[coupPair]; !exist {
							coupPairsMap[coupPair] = struct{}{}
							validLoop = true
						}
					} else {
						if swapPairInfo.BothBriToken {
							// target pair: both tokens are bridge token
							if !loop.HasNonBriToken {
								// pick loops composed by bridge tokens.
								validLoop = true
							}
						} else {
							// target pair: both tokens are not bridge Token
							validLoop = true
						}
					}

					if validLoop {
						loopKey := loop.Key()
						if _, exist := deDup[loopKey]; !exist {
							matchedLoops = append(matchedLoops, loop)
							deDup[loopKey] = struct{}{}
							if len(matchedLoops) >= loopsWanted {
								loopsMaxReached = true
								break
							}
						}
					}
				}
			}
		}

		if len(matchedLoops) > 0 {
			matchedPairsLoops[index] = matchedLoops
		}
	}

	return
}
