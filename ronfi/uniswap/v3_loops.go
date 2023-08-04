package uniswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"math/big"
)

type TaggedEdge struct {
	Pair         common.Address
	Dir          int
	PoolType     PoolType
	PoolFee      uint64
	GasNeed      uint64
	BothBriToken bool
}

func (t *TaggedEdge) ID() string {
	if t.PoolType == V2 {
		return fmt.Sprintf("%s-%d-v2", t.Pair, t.Dir)
	} else if t.PoolType == V3 {
		return fmt.Sprintf("%s-%d-v3", t.Pair, t.Dir)
	} else {
		return fmt.Sprintf("%s-%d-unknown", t.Pair, t.Dir)
	}
}

type Edge struct {
	Source common.Address
	Target common.Address
	Tag    *TaggedEdge
}

type Cycle []*Edge

func (c *Cycle) String() string {
	var str string
	for _, edge := range *c {
		str += fmt.Sprintf("%s -> ", edge.Tag.ID())
	}
	//str += c[0].Source.Hex()
	return str
}

func (c *Cycle) Id() common.Hash {
	size := len(*c) * 20
	serialized := make([]byte, 0, size)
	for _, edge := range *c {
		serialized = append(serialized, edge.Tag.Pair.Bytes()...)
	}

	hash := rcommon.RawHash(serialized)
	return common.BytesToHash(hash)
}

func (c *Cycle) HasDupPair() bool {
	visited := make(map[common.Address]bool)
	for _, edge := range *c {
		if visited[edge.Tag.Pair] {
			return true
		}
		visited[edge.Tag.Pair] = true
	}
	return false
}

func (c *Cycle) CheckIfTradableInputToken() bool {
	if len(*c) == 0 {
		return false
	} else {
		edge0 := (*c)[0]
		_, ok := rcommon.TradableTokens[edge0.Source]
		if !ok {
			return false
		}
	}

	return true
}

func (c *Cycle) CheckIfAllBridgeTokens() bool {
	if len(*c) == 0 {
		return false
	} else {
		for _, edge := range *c {
			_, ok1 := rcommon.BridgeTokens[edge.Source]
			_, ok2 := rcommon.BridgeTokens[edge.Target]
			if !ok1 || !ok2 {
				return false
			}
		}
	}

	return true
}

func (c *Cycle) CheckIfV3OnHeadTail() bool {
	if len(*c) == 0 {
		return false
	}

	edgeH := (*c)[0]
	edgeT := (*c)[len(*c)-1]
	if edgeH.Tag.PoolType == V3 && edgeT.Tag.PoolType == V3 {
		return true
	}

	return false
}

type Graph struct {
	Nodes map[common.Address][]*Edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[common.Address][]*Edge),
	}
}

func (g *Graph) AddEdge(edge *Edge) {
	g.Nodes[edge.Source] = append(g.Nodes[edge.Source], edge)
}

func (g *Graph) DFS(node *Edge, targetEdge *Edge, currentPath []*Edge, result *[]Cycle, maxLevel int) {
	if node == nil || maxLevel <= 0 {
		return
	}

	if node.Tag.Pair == common.HexToAddress("0x4763784c16adbccb8fe42a8a7bae43f3bd848eee") {
		fmt.Printf("matched!\n")
	}

	// Append current node to current path
	currentPath = append(currentPath, node)

	// Check if the current path contains the sub-slice
	for _, edge := range currentPath {
		if edge.Tag.ID() == targetEdge.Tag.ID() &&
			currentPath[0].Source == currentPath[len(currentPath)-1].Target &&
			currentPath[0].Source == rcommon.WETH {
			// Add a copy of the current path to the result
			pathCopy := make(Cycle, len(currentPath))
			copy(pathCopy, currentPath)
			*result = append(*result, pathCopy)
		}
	}

	if node.Target == rcommon.WETH {
		return
	}

	// Continue DFS for neighboring nodes
	for _, neighbor := range g.Nodes[node.Target] {
		if neighbor.Tag.Pair == node.Tag.Pair {
			continue
		}
		g.DFS(neighbor, targetEdge, currentPath, result, maxLevel-1)
	}

	// Backtrack: remove current node from current path
	currentPath = currentPath[:len(currentPath)-1]
}

func (g *Graph) FindCyclesWithDedicatedEdges(dedicatedEdges []*Edge, hops int) []Cycle {
	var cycles []Cycle

	foundCycles := make([]Cycle, 0, 24)
	for _, edge := range dedicatedEdges {
		for _, src := range g.Nodes[rcommon.WETH] {
			path := make([]*Edge, 0, 6)
			g.DFS(src, edge, path, &foundCycles, hops-1)
		}

		if edge.Tag.BothBriToken {
			for _, cycle := range foundCycles {
				if cycle.CheckIfAllBridgeTokens() {
					cycles = append(cycles, cycle)
				}
			}
		} else {
			cycles = append(cycles, foundCycles...)
		}
	}
	// filter same Pair cycles and non-tradable token cycles
	filteredCycles := make([]Cycle, 0, len(cycles))
	for _, cycle := range cycles {
		if !cycle.HasDupPair() && cycle.CheckIfTradableInputToken() {
			filteredCycles = append(filteredCycles, cycle)
		}
	}

	return filteredCycles
}

type V3Loops struct {
	di         *defi.Info
	pairsInfo  defi.PairInfoMap
	poolsInfo  map[common.Address]*defi.PoolInfo
	tokensInfo map[common.Address]*defi.TokenInfo
	pairGasMap map[string]uint64

	g *Graph
}

type V3PoolInfo struct {
	address common.Address
	dir     int
}

func NewV3Loops(
	di *defi.Info,
	pairsInfo defi.PairInfoMap,
	poolsInfo map[common.Address]*defi.PoolInfo,
	tokensInfo map[common.Address]*defi.TokenInfo,
	pairGasMap map[string]uint64,
) *V3Loops {

	totalStaledV2Pairs := 0
	totalStaledV3Pools := 0
	g := NewGraph()
	for addr, info := range pairsInfo {
		if isStaledPools(di, addr, info.Token0, info.Token1) {
			totalStaledV2Pairs++
			//log.Info("RonFi V3Loops", "skip staled pair", addr)
			continue
		}

		key := fmt.Sprintf("%s-%d", addr, 0)
		gasNeed, ok := pairGasMap[key]
		if !ok {
			gasNeed = 150000
		}
		if gasNeed > 150000 {
			gasNeed = 150000
		}

		source := info.Token0
		target := info.Token1
		if source.Big().Cmp(target.Big()) > 0 {
			source, target = target, source
		}

		edge := &Edge{Source: source, Target: target, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      0,
			PoolType: V2,
			PoolFee:  info.Fee,
			GasNeed:  gasNeed,
		}}

		key = fmt.Sprintf("%s-%d", addr, 1)
		gasNeed, ok = pairGasMap[key]
		if !ok {
			gasNeed = 150000
		}
		if gasNeed > 150000 {
			gasNeed = 150000
		}

		source, target = target, source
		reversedEdge := &Edge{Source: source, Target: target, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      1,
			PoolType: V2,
			PoolFee:  info.Fee,
			GasNeed:  gasNeed,
		}}

		g.AddEdge(edge)
		g.AddEdge(reversedEdge)
	}

	for addr, info := range poolsInfo {
		if isStaledPools(di, addr, info.Token0, info.Token1) {
			totalStaledV3Pools++
			//log.Info("RonFi V3Loops", "skip staled pool", addr)
			continue
		}

		key := fmt.Sprintf("%s-%d", addr, 0)
		gasNeed, ok := pairGasMap[key]
		if !ok {
			gasNeed = 150000
		}
		if gasNeed > 150000 {
			gasNeed -= 150000
		}

		source := info.Token0
		target := info.Token1
		if source.Big().Cmp(target.Big()) > 0 {
			source, target = target, source
		}

		edge := &Edge{Source: source, Target: target, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      0,
			PoolType: V3,
			GasNeed:  gasNeed,
		}}
		key = fmt.Sprintf("%s-%d", addr, 1)
		gasNeed, ok = pairGasMap[key]
		if !ok {
			gasNeed = 150000
		}
		if gasNeed > 150000 {
			gasNeed -= 150000
		}

		source, target = target, source
		reversedEdge := &Edge{Source: source, Target: target, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      1,
			PoolType: V3,
			GasNeed:  gasNeed,
		}}

		g.AddEdge(edge)
		g.AddEdge(reversedEdge)
	}
	log.Info("RonFi V3Loops", "total staled v2 pairs", totalStaledV2Pairs, "total staled v3 pools", totalStaledV3Pools)

	return &V3Loops{
		di:         di,
		pairsInfo:  pairsInfo,
		poolsInfo:  poolsInfo,
		tokensInfo: tokensInfo,
		pairGasMap: pairGasMap,
		g:          g,
	}
}

func (v *V3Loops) FindLoops(edge *Edge) []V3ArbPath {
	redupMaps := make(map[common.Hash]*Cycle)
	arbs := make([]V3ArbPath, 0)
	// find cycles
	dedicatedEdges := []*Edge{v.ReversedEdge(edge)}

	cycles := v.g.FindCyclesWithDedicatedEdges(dedicatedEdges, 2)
	cycles3 := v.g.FindCyclesWithDedicatedEdges(dedicatedEdges, 3)
	if len(cycles3) < 0 {
		cycles4 := v.g.FindCyclesWithDedicatedEdges(dedicatedEdges, 4)
		cycles = append(cycles, cycles4...)
	} else {
		cycles = append(cycles, cycles3...)
	}
	if len(cycles) > 24 {
		cycles = cycles[:24]
	}

	for _, cycle := range cycles {
		if _, ok := redupMaps[cycle.Id()]; !ok {
			redupMaps[cycle.Id()] = &cycle
		} else {
			continue
		}

		validCycle := true
		path := make(V3ArbPath, 0, len(cycle))
		for _, edge := range cycle {
			tokenFee := 0

			if v.tokensInfo[edge.Target] == nil {
				//log.Info("RonFi V3Loops", "unKnown token", edge.Target.HexNoChecksum())
				validCycle = false
				break
			} else {
				tokenFee = int(v.tokensInfo[edge.Target].Fee)
			}
			if edge.Tag.PoolType == V2 {
				poolFee := 30
				if v.pairsInfo[edge.Tag.Pair] == nil {
					log.Info("RonFi V3Loops", "unKnown v2 pair", edge.Tag.Pair)
					validCycle = false
					break
				} else {
					poolFee = int(v.pairsInfo[edge.Tag.Pair].Fee)
				}
				path = append(path, &Pool{
					PoolAddr: edge.Tag.Pair,
					PoolType: V2,
					PoolFee:  poolFee,
					TokenFee: tokenFee,
					TickLens: rcommon.ZeroAddress,
					TokenIn:  edge.Source,
					Dir:      edge.Tag.Dir,
					GasNeed:  edge.Tag.GasNeed,
				})
			} else if edge.Tag.PoolType == V3 {
				path = append(path, &Pool{
					PoolAddr: edge.Tag.Pair,
					PoolType: V3,
					PoolFee:  0,
					TokenFee: tokenFee,
					TickLens: rcommon.UniswapV3TicklensAddress,
					TokenIn:  edge.Source,
					Dir:      edge.Tag.Dir,
					GasNeed:  edge.Tag.GasNeed,
				})
			} else {
				break
			}
		}
		if len(path) > 0 && validCycle {
			arbs = append(arbs, path)
		}
	}
	return arbs
}

func (v *V3Loops) ReversedEdge(edge *Edge) *Edge {
	if edge == nil {
		return nil
	}

	gasNeed := v.pairGasMap[fmt.Sprintf("%s-%d", edge.Tag.Pair, 1-edge.Tag.Dir)]

	return &Edge{
		Source: edge.Target,
		Target: edge.Source,
		Tag: &TaggedEdge{
			Pair:         edge.Tag.Pair,
			Dir:          1 - edge.Tag.Dir,
			PoolType:     edge.Tag.PoolType,
			PoolFee:      edge.Tag.PoolFee,
			GasNeed:      gasNeed,
			BothBriToken: edge.Tag.BothBriToken,
		},
	}
}

func isStaledPools(di *defi.Info, addr, token0, token1 common.Address) bool {
	return false

	isBriPool := false
	_, ok0 := rcommon.BridgeTokens[token0]
	_, ok1 := rcommon.BridgeTokens[token1]
	if ok0 && ok1 {
		isBriPool = true
	}

	token0Bal := di.GetTokenBalance(addr, token0)
	token0Info := di.GetTokenInfo(token0)
	if token0Info == nil {
		return true
	}

	// if bridge pool and total vol < 10000 usd, ignore it
	if isBriPool {
		value0 := rcommon.ToFloat(token0Bal, token0Info.Decimals) *
			defi.GetTokenPrice(token0)
		if value0 < 5000.0 {
			return true
		} else {
			return false
		}
	} else {
		threshold := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token0Info.Decimals-5)), nil)
		if token0Bal == nil || token0Bal.Cmp(threshold) < 0 {
			return true
		}

		token1Bal := di.GetTokenBalance(addr, token1)
		token1Info := di.GetTokenInfo(token1)
		if token1Info == nil {
			return true
		}
		threshold = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token1Info.Decimals-5)), nil)
		if token1Bal == nil || token1Bal.Cmp(threshold) == 0 {
			return true
		}

		return false
	}
}

func ToV3Edge(pairInfo *defi.SwapPairInfo) *Edge {
	if pairInfo == nil {
		return nil
	}

	poolType := V2
	if pairInfo.V3 {
		poolType = V3
	}

	taggedEdge := &TaggedEdge{
		Pair:         pairInfo.Address,
		Dir:          int(pairInfo.Dir),
		PoolType:     poolType,
		PoolFee:      0,
		GasNeed:      0,
		BothBriToken: pairInfo.BothBriToken,
	}

	return &Edge{
		Source: pairInfo.TokenIn,
		Target: pairInfo.TokenOut,
		Tag:    taggedEdge,
	}
}
