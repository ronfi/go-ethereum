package stats

type ObsId string

type ObsStats struct {
	grossProfit         float64
	prevGrossProfit     float64
	grossV3Profit       float64
	prevGrossV3Profit   float64
	prevSlush           float64
	totalFee            float64
	reportedTotalFee    float64
	totalArbTxs         uint64
	reportedTotalArbTxs uint64
	totalSuccess        uint64
	totalFail           uint64
	totalCopied         uint64
}

type ObsAllStatsMap map[ObsId]*ObsStats

var (
	ProfitObsIds = []ObsId{Obs1, Obsx}
)

func (oa ObsAllStatsMap) init() {
	for _, id := range ProfitObsIds {
		oa[id] = &ObsStats{}
	}
}

func (oa ObsAllStatsMap) updateGrossProfit(id ObsId, grossProfit float64) {
	oa[id].grossProfit += grossProfit
}

func (oa ObsAllStatsMap) getGrossProfit(id ObsId) float64 {
	return oa[id].grossProfit
}

func (oa ObsAllStatsMap) updatePrevGrossProfit(id ObsId, prevGrossProfit float64) {
	oa[id].prevGrossProfit = prevGrossProfit
}

func (oa ObsAllStatsMap) getPrevGrossProfit(id ObsId) float64 {
	return oa[id].prevGrossProfit
}

func (oa ObsAllStatsMap) updatePrevSlush(id ObsId, prevSlush float64) {
	oa[id].prevSlush = prevSlush
}

func (oa ObsAllStatsMap) getPrevSlush(id ObsId) float64 {
	return oa[id].prevSlush
}

func (oa ObsAllStatsMap) updateGrossV3Profit(id ObsId, grossProfit float64) {
	oa[id].grossV3Profit += grossProfit
}

func (oa ObsAllStatsMap) getGrossV3Profit(id ObsId) float64 {
	return oa[id].grossV3Profit
}

func (oa ObsAllStatsMap) updatePrevGrossV3Profit(id ObsId, prevGrossProfit float64) {
	oa[id].prevGrossV3Profit = prevGrossProfit
}

func (oa ObsAllStatsMap) getPrevGrossV3Profit(id ObsId) float64 {
	return oa[id].prevGrossV3Profit
}

func (oa ObsAllStatsMap) updateTotalFee(id ObsId, totalFee float64) {
	oa[id].totalFee += totalFee
}

func (oa ObsAllStatsMap) getTotalFee(id ObsId) float64 {
	return oa[id].totalFee
}

func (oa ObsAllStatsMap) updateReportedTotalFee(id ObsId, reportedTotalFee float64) {
	oa[id].reportedTotalFee = reportedTotalFee
}

func (oa ObsAllStatsMap) getReportedTotalFee(id ObsId) float64 {
	return oa[id].reportedTotalFee
}

func (oa ObsAllStatsMap) updateReportedTotalArbTxs(id ObsId, reportedTotalArbTxs uint64) {
	oa[id].reportedTotalArbTxs = reportedTotalArbTxs
}

func (oa ObsAllStatsMap) getReportedTotalArbTxs(id ObsId) uint64 {
	return oa[id].reportedTotalArbTxs
}

func (oa ObsAllStatsMap) updateTotalArbTxs(id ObsId) {
	oa[id].totalArbTxs++
}

func (oa ObsAllStatsMap) getTotalArbTxs(id ObsId) uint64 {
	return oa[id].totalArbTxs
}

func (oa ObsAllStatsMap) updateTotalSuccess(id ObsId) {
	oa[id].totalSuccess++
}

func (oa ObsAllStatsMap) getTotalSuccess(id ObsId) uint64 {
	return oa[id].totalSuccess
}

func (oa ObsAllStatsMap) updateTotalFail(id ObsId) {
	oa[id].totalFail++
}

func (oa ObsAllStatsMap) getTotalFail(id ObsId) uint64 {
	return oa[id].totalFail
}

func (oa ObsAllStatsMap) updateTotalCopied(id ObsId) {
	oa[id].totalCopied++
}

func (oa ObsAllStatsMap) getTotalCopied(id ObsId) uint64 {
	return oa[id].totalCopied
}

func (oa ObsAllStatsMap) ratio(id ObsId) float64 {
	totalSuccess := oa.getTotalSuccess(id)
	totalFail := oa.getTotalFail(id)
	total := totalSuccess + totalFail

	if total == 0 {
		return 0.0
	} else {
		return float64(totalSuccess) / float64(total) * 100
	}
}

func (oa ObsAllStatsMap) shortFee(id ObsId) float64 {
	return oa.getTotalFee(id) - oa.getReportedTotalFee(id)
}

func (oa ObsAllStatsMap) copyRate(id ObsId) float64 {
	totalCopied := oa.getTotalCopied(id)
	totalSuccess := oa.getTotalSuccess(id)
	totalFail := oa.getTotalFail(id)
	total := totalSuccess + totalFail

	if total == 0 {
		return 0.0
	} else {
		return float64(totalCopied) / float64(total) * 100
	}
}
