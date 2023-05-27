package stats

import (
	"context"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common/mclock"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
)

type Stats struct {
	eth           rcommon.Backend
	chain         *core.BlockChain
	client        *ethclient.Client
	di            *defi.Info
	currentHeader *types.Header

	stopCh chan struct{}

	startTime           mclock.AbsTime
	initialBalance      atomic.Value
	prevBalance         atomic.Value
	v3InitialBalance    atomic.Value
	v3PrevBalance       atomic.Value
	prevReportResetTime time.Time

	chiBalance   uint64
	bnbBalance   float64
	v3ChiBalance uint64
	v3BnbBalance float64

	totalArbTxs           uint64
	reportedTotalArbTxs   uint64
	txCount               uint64
	reportedTxCount       uint64
	missedTxCount         uint64
	reportedMissedTxCount uint64
	reportedTotalCancel   int

	totals                sync.Map
	initialPairGasMapSize int
	totalSuccess          int
	totalFail             int
	totalCancel           int
	totalCopied           int
	totalException        int
	totalLoss             float64
	totalFee              float64
	reportedTotalFee      float64
	dexTotalFail          int
	dexTotal              int

	obsStats     ObsAllStatsMap
	obsPairStats ObsAllPairStatsMap

	pairMaxGasUsed  map[string]uint64
	dexPairs        map[common.Address]uint64
	topDexPairsInfo map[common.Address]DexPairInfo
	prevTotalVol    float64
	dexTokensVol    map[common.Address]float64
	dexPairsVol     map[common.Address]float64
	prevResetTime   time.Time
	obsMethods      map[uint64]string
	obsRouters      map[common.Address]uint64
}

type miscStatCnt struct {
	totalFail      uint64
	totalException uint64
	totalSuccess   uint64
	totalFee       float64
}

func NewStats(
	eth rcommon.Backend,
	client *ethclient.Client,
	di *defi.Info,
	pairGasMap map[string]uint64,
	dexPairsMap map[common.Address]uint64,
	obsRouters map[common.Address]uint64,
	obsMethods map[uint64]string,
) *Stats {
	s := &Stats{
		eth:    eth,
		chain:  eth.BlockChain(),
		client: nil,
		stopCh: make(chan struct{}),
	}
	s.initialBalance.Store(defi.RonBalance{})
	s.prevBalance.Store(defi.RonBalance{})
	s.v3InitialBalance.Store(defi.RonBalance{})
	s.v3PrevBalance.Store(defi.RonBalance{})

	s.startTime = mclock.Now()
	s.prevResetTime = time.Now()
	s.prevReportResetTime = time.Now()
	s.pairMaxGasUsed = make(map[string]uint64)
	s.dexPairs = make(map[common.Address]uint64)
	s.dexTokensVol = make(map[common.Address]float64)
	s.dexPairsVol = make(map[common.Address]float64)
	s.topDexPairsInfo = make(map[common.Address]DexPairInfo)
	s.obsRouters = obsRouters
	s.obsMethods = obsMethods

	// initialize the 'pairMaxGasUsed' as the loaded 'pair_gas.json' file contents.
	s.initialPairGasMapSize = len(pairGasMap)
	for key, gas := range pairGasMap {
		s.pairMaxGasUsed[key] = gas
	}

	// deep copy dexPairs
	for pair, frequency := range dexPairsMap {
		s.dexPairs[pair] = frequency
	}

	s.client = client
	s.di = di

	balance := s.di.GetAllBalance(rcommon.AllTradingExecutors, false)
	balanceV3 := s.di.GetAllBalance(rcommon.AllV3TradingExecutors, true)
	log.Info("RonFi Stats Init",
		"Initial Executor Eth", balance.Eth,
		"Initial Contract Balance in USD", balance.ContractTotal,
		"Initial V3 Executor Eth", balanceV3.Eth,
		"Initial V3 Contract Balance in USD", balanceV3.ContractTotal)
	s.initialBalance.Store(balance.Copy())
	s.prevBalance.Store(balance)
	s.chiBalance = balance.ContractChi
	s.bnbBalance = balance.Eth
	s.v3InitialBalance.Store(balanceV3.Copy())
	s.v3PrevBalance.Store(balanceV3.Copy())
	s.v3ChiBalance = balanceV3.ContractChi
	s.v3BnbBalance = balanceV3.Eth

	PrevBlockTargetDexTxs = make(map[uint64]TargetDexInfo)
	PrevBlockTxs = 0

	s.obsPairStats = make(ObsAllPairStatsMap)
	s.obsPairStats.init()
	s.obsStats = make(ObsAllStatsMap)
	s.obsStats.init()

	return s
}

func (s *Stats) StopStats() {
	s.stopCh <- struct{}{}
}

func (s *Stats) Stop() {
}

func (s *Stats) IsSyncing() bool {
	syncProgress, _ := s.client.SyncProgress(context.Background())
	if syncProgress != nil {
		return true
	} else {
		return false
	}
}

func (s *Stats) Run() {
	//syncing := true
	// subscribe new block event
	headers := make(chan *types.Header, 16)
	blockSub, err := s.client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Error("RonFi Stats SubscribeNewHead() fail", "err", err)
		close(headers)
		return
	}

	defer func() {
		e := recover()
		if e != nil {
			log.Error("RonFi Stats Run exception", "err", e)
			log.Error("RonFi Stats Run exception", "stack trace", string(debug.Stack()))
		}

		s.Stop()

		blockSub.Unsubscribe()
		close(headers)
	}()

	// block header event loop
	var blockNumber uint64
	for {
		select {
		case header := <-headers:
			if header == nil {
				continue
			}

			blockNumber = header.Number.Uint64()
			{
				s.report(header)

				if blockNumber%200 == 0 && header.Number.Cmp(s.currentHeader.Number) > 0 {
					// report profit in every 10 minutes
					s.dexVolumeReport()

					log.Info("RonFi arb") // splitter of profit report

					// reset report automatically at every day 8:00am
					if blockTime := time.Unix(int64(s.currentHeader.Time), 0); blockTime.Day() != s.prevReportResetTime.Day() {
						s.prevReportResetTime = blockTime
						s.initialBalance.Store(s.di.GetAllBalance(rcommon.AllTradingExecutors, false))
						s.v3InitialBalance.Store(s.di.GetAllBalance(rcommon.AllV3TradingExecutors, true))
						s.ResetStats()
						s.startTime = mclock.Now()
						log.Info("RonFi arb reset report for a wonderful new day", "block", s.currentHeader.Number)

						s.obsPairStats = make(ObsAllPairStatsMap)
						s.obsPairStats.init()

						// clear dex volume
						s.prevTotalVol = 0.0
						for k := range s.dexTokensVol {
							delete(s.dexTokensVol, k)
						}

						for k := range s.dexPairsVol {
							delete(s.dexPairsVol, k)
						}
					}
				}
			}
			s.currentHeader = header
			//if !rpc.StartTrading && header.Number.Uint64()%10 == 0 { // update syncing status in half minute
			//	syncing = s.IsSyncing()
			//}
		// System shutdown.
		case err := <-blockSub.Err():
			log.Error("RonFi  Stats block subscribe error!", "blockSub.Err", err)
			return

		case <-s.stopCh:
			log.Info("RonFi stats exit")
			return
		}
	}
}

func (s *Stats) ResetStats() {
	s.totals.Range(func(key interface{}, value interface{}) bool {
		s.totals.Delete(key)
		return true
	})
	s.totalSuccess = 0
	s.totalFail = 0
	s.totalCancel = 0
	s.totalCopied = 0
	s.totalException = 0
	s.totalLoss = 0
	s.dexTotalFail = 0
	s.dexTotal = 0
	s.totalFee = 0
	s.reportedTotalFee = 0
	s.reportedTotalCancel = 0

	s.txCount = 0
	s.reportedTxCount = 0
	s.missedTxCount = 0
	s.reportedMissedTxCount = 0

	atomic.StoreUint64(&s.totalArbTxs, 0)
	s.reportedTotalArbTxs = 0

	s.obsStats = make(ObsAllStatsMap)
	s.obsStats.init()
}
