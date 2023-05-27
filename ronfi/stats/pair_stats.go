package stats

import (
	"github.com/ethereum/go-ethereum/common"
	"sort"
)

type PairFrequency struct {
	Key   common.Address
	Value uint64
}

type ObsAllPairStatsMap map[ObsId]map[common.Address]uint64

var (
	PairsObsIds = []ObsId{Obs1, Obs3, Obs4, Obs5, Obs6, Obs7, Ron}
)

func (sp ObsAllPairStatsMap) init() {
	for _, id := range PairsObsIds {
		sp[id] = make(map[common.Address]uint64)
	}
}

func (sp ObsAllPairStatsMap) update(id ObsId, pair common.Address) {
	sp[id][pair]++
}

func (sp ObsAllPairStatsMap) count(id ObsId) int {
	return len(sp[id])
}

func (sp ObsAllPairStatsMap) topN(id ObsId, num int) []PairFrequency {
	obsPairStats := sp[id]
	ss := make([]PairFrequency, 0, len(obsPairStats))
	for pair, freq := range obsPairStats {
		ss = append(ss, PairFrequency{pair, freq})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	if len(ss) <= num {
		return ss
	} else {
		return ss[:num]
	}
}

func (sp ObsAllPairStatsMap) freqInObs(id ObsId, pair common.Address) uint64 {
	if freq, ok := sp[id][pair]; ok {
		return freq
	} else {
		return 0
	}
}
