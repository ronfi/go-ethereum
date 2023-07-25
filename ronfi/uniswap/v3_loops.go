package uniswap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"sync"
)

type TaggedEdge struct {
	Pair     common.Address
	Dir      int
	PoolType PoolType
	PoolFee  uint64
	GasNeed  uint64
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
	nodes   map[common.Address][]*Edge
	visited map[common.Address]bool
	lock    sync.RWMutex
}

func NewGraph() *Graph {
	return &Graph{
		nodes:   make(map[common.Address][]*Edge),
		visited: make(map[common.Address]bool),
		lock:    sync.RWMutex{},
	}
}

func (g *Graph) AddEdge(edge *Edge) {
	g.lock.Lock()
	defer func() {
		g.lock.Unlock()
	}()
	g.nodes[edge.Source] = append(g.nodes[edge.Source], edge)
}

func (g *Graph) DFS(src, target common.Address, visitedEdges map[string]bool, path []*Edge, hops int) ([]Cycle, bool) {
	if len(path) >= hops {
		return nil, false
	}

	g.lock.Lock()
	g.visited[src] = true
	g.lock.Unlock()
	found := false
	var cycles []Cycle

	for _, edge := range g.nodes[src] {
		g.lock.Lock()
		visited := g.visited[edge.Target]
		g.lock.Unlock()
		if edge.Target == target {
			cycles = append(cycles, append(path, edge))
			found = true
		} else if !visited && !visitedEdges[edge.Tag.ID()] {
			visitedEdges[edge.Tag.ID()] = true
			newPath := append(path, edge)
			foundCycles, cycleFound := g.DFS(edge.Target, target, visitedEdges, newPath, hops)
			if cycleFound {
				found = true
				cycles = append(cycles, foundCycles...)
			}
			visitedEdges[edge.Tag.ID()] = false
		}
	}

	g.lock.Lock()
	g.visited[src] = false
	g.lock.Unlock()
	return cycles, found
}

func (g *Graph) FindCyclesWithDedicatedEdges(dedicatedEdges []*Edge, hops int) []Cycle {
	var cycles []Cycle
	visitedEdges := make(map[string]bool)

	for _, edge := range dedicatedEdges {
		visitedEdges[edge.Tag.ID()] = true
	}

	for _, edge := range dedicatedEdges {
		path := []*Edge{edge}
		foundCycles, _ := g.DFS(edge.Target, edge.Source, visitedEdges, path, hops)
		cycles = append(cycles, foundCycles...)
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

	g := NewGraph()
	for addr, info := range pairsInfo {
		if isStaledBridgePair(di, addr, info) {
			log.Info("RonFi V3Loops", "skip staled bridge pair", addr)
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
		g.AddEdge(&Edge{Source: info.Token0, Target: info.Token1, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      0,
			PoolType: V2,
			PoolFee:  info.Fee,
			GasNeed:  gasNeed,
		}})

		key = fmt.Sprintf("%s-%d", addr, 1)
		gasNeed, ok = pairGasMap[key]
		if !ok {
			gasNeed = 150000
		}
		if gasNeed > 150000 {
			gasNeed = 150000
		}
		g.AddEdge(&Edge{Source: info.Token1, Target: info.Token0, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      1,
			PoolType: V2,
			PoolFee:  info.Fee,
			GasNeed:  gasNeed,
		}})
	}

	for addr, info := range poolsInfo {
		key := fmt.Sprintf("%s-%d", addr, 0)
		gasNeed, ok := pairGasMap[key]
		if !ok {
			gasNeed = 150000
		}
		if gasNeed > 150000 {
			gasNeed -= 150000
		}
		g.AddEdge(&Edge{Source: info.Token0, Target: info.Token1, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      0,
			PoolType: V3,
			GasNeed:  gasNeed,
		}})
		key = fmt.Sprintf("%s-%d", addr, 1)
		gasNeed, ok = pairGasMap[key]
		if !ok {
			gasNeed = 150000
		}
		if gasNeed > 150000 {
			gasNeed -= 150000
		}
		g.AddEdge(&Edge{Source: info.Token1, Target: info.Token0, Tag: &TaggedEdge{
			Pair:     addr,
			Dir:      1,
			PoolType: V3,
			GasNeed:  gasNeed,
		}})
	}

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
	arbs := make([]V3ArbPath, 0)
	// find cycles
	dedicatedEdges := []*Edge{edge}

	cycles := v.g.FindCyclesWithDedicatedEdges(dedicatedEdges, 2)
	if len(cycles) < 0 {
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
	}

	for _, cycle := range cycles {
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

func isStaledBridgePair(di *defi.Info, addr common.Address, info *defi.PairInfo) bool {
	_, ok0 := rcommon.BridgeTokens[info.Token0]
	_, ok1 := rcommon.BridgeTokens[info.Token1]
	if !ok0 || !ok1 {
		return false
	}

	res := di.GetPairReserves(addr)
	if res == nil {
		return false
	}
	token0Info := di.GetTokenInfo(info.Token0)
	if token0Info == nil {
		return false
	}

	value0 := rcommon.ToFloat(res.Reserve0, token0Info.Decimals) *
		defi.GetTokenPrice(info.Token0)
	if value0 < 5000.0 {
		return true
	}

	return false
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
		Pair:     pairInfo.Address,
		Dir:      int(pairInfo.Dir),
		PoolType: poolType,
		PoolFee:  0,
		GasNeed:  0,
	}

	return &Edge{
		Source: pairInfo.TokenIn,
		Target: pairInfo.TokenOut,
		Tag:    taggedEdge,
	}
}
