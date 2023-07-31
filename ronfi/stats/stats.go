package stats

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/go-redis/redis"
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
	rdb           *redis.Client
	mysql         *db.Mysql
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

	obsContractStats map[common.Address]uint64
	obsMethodStats   map[uint32]uint64

	obsStats     ObsAllStatsMap
	obsPairStats ObsAllPairStatsMap
	loopsCol     *LoopsCollector
	obsCol       *ObsCollector

	pairMaxGasUsed  map[string]uint64
	dexPairs        map[common.Address]uint64
	topDexPairsInfo map[common.Address]DexPairInfo
	prevTotalVol    float64
	dexTokensVol    map[common.Address]float64
	dexPairsVol     map[common.Address]float64
	prevResetTime   time.Time
	obsMethods      map[uint64]string
	obsRouters      map[string]struct{}
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
	redis *redis.Client,
	mysql *db.Mysql,
	pairGasMap map[string]uint64,
	dexPairsMap map[common.Address]uint64,
	obsRouters map[string]struct{},
	obsMethods map[uint64]string,
) *Stats {
	s := &Stats{
		eth:    eth,
		chain:  eth.BlockChain(),
		client: nil,
		stopCh: make(chan struct{}),
	}
	s.initialBalance.Store(defi.RonFiBalance{})
	s.prevBalance.Store(defi.RonFiBalance{})
	s.v3InitialBalance.Store(defi.RonFiBalance{})
	s.v3PrevBalance.Store(defi.RonFiBalance{})

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

	s.obsContractStats = make(map[common.Address]uint64)
	s.obsMethodStats = make(map[uint32]uint64)

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
	s.rdb = redis
	s.mysql = mysql

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

	PrevBlockTxs = 0

	s.obsPairStats = make(ObsAllPairStatsMap)
	s.obsPairStats.init()
	s.obsStats = make(ObsAllStatsMap)
	s.obsStats.init()

	s.loopsCol = NewLoopsCollector(redis, mysql)
	s.loopsCol.start()

	s.obsCol = NewObsCollector(redis, mysql)
	s.obsCol.start()

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

				if blockNumber%50 == 0 && header.Number.Cmp(s.currentHeader.Number) > 0 {
					// report profit in every 10 minutes
					s.dexVolumeReport()
					for _, id := range ProfitObsIds {
						s.obsProfitReport(id)
					}

					s.UpdatePairTokensInfo()

					go func() {
						s.mysql.UpdatePairGas(s.pairMaxGasUsed)
						s.mysql.UpdateDexPairs(s.dexPairs)
					}()

					// report obs arbitrage contract/method statistics
					s.obsContractReport()

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

// collect the GasUsed of all the dex pairs
func (s *Stats) dexPairGasUsed(txs types.Transactions, receipts types.Receipts, bh common.Hash) {
	for _, tx := range txs {
		to := tx.To()
		if to == nil {
			continue
		}
		txLookup := s.chain.GetTransactionLookup(tx.Hash())
		if txLookup == nil || txLookup.BlockHash != bh {
			// maybe chain reorg
			return
		}
		if txLookup.Index >= uint64(len(receipts)) {
			log.Error("RonFi report", "dexTx", tx.Hash().String(), "BlockIndex", txLookup.BlockIndex, "Index", txLookup.Index, "receipts", len(receipts))
			return
		}

		receipt := receipts[txLookup.Index]
		if receipt.Status != 1 {
			// only collect the success dex tx gasUsed
			continue
		}

		data := tx.Data()
		if len(data) < 4 {
			continue
		}

		// collect dex pairs
		methodID := uint64(binary.BigEndian.Uint32(data[:4]))
		swapPairsInfo := s.di.ExtractSwapPairInfo(tx, *tx.To(), receipt.Logs, defi.RonFiExtractTypePairs)
		for _, swapPairInfo := range swapPairsInfo {
			// collect all dex pairs
			if frequency, exist := s.dexPairs[swapPairInfo.Address]; !exist {
				s.dexPairs[swapPairInfo.Address] = 1
			} else {
				s.dexPairs[swapPairInfo.Address] = frequency + 1
			}
		}

		// collect pair gas info
		if len(swapPairsInfo) == 1 {
			// calculate pair gas for one-hop swaps
			key := swapPairsInfo[0].Key
			if averageGasUsed, ok := s.pairMaxGasUsed[key]; ok {
				if receipt.GasUsed > averageGasUsed*2 { // Gas >200% suddenly
					log.Warn("RonFi pair gas collector, gas rise", "dexTx", tx.Hash().String(), "oldGas", averageGasUsed, "newGas", receipt.GasUsed)
				} else if receipt.GasUsed*2 < averageGasUsed { // Gas < 50% suddenly
					log.Warn("RonFi pair gas collector, gas drop", "dexTx", tx.Hash().String(), "oldGas", averageGasUsed, "newGas", receipt.GasUsed)
				}
				s.pairMaxGasUsed[key] = averageGasUsed - averageGasUsed/32 + receipt.GasUsed/32 // 31/32 * old + 1/32 * new, to filter any exceptional sharp peak
			} else {
				s.pairMaxGasUsed[key] = receipt.GasUsed
			}
		}

		// collect obs routers info
		// note: these collected obs method MUST NOT be used directly! which is highly possible to be reused by some other contracts but not obs!
		//		 best practice is to manually check these obs methods and commit into github one by one! carefully!
		_, IsObsTx := s.di.CheckIfObsTx(tx, receipt.Logs, *to)
		if IsObsTx {
			if txpool.ObsMethods != nil {
				if _, exist := txpool.ObsMethods[methodID]; !exist {
					// only if the methodId is not in the core.ObsMethods list, collect the obs routers info
					if s.obsRouters != nil && to != nil {
						routerMethod := fmt.Sprintf("%s-0x%08x", *to, methodID)
						if _, exist := s.obsRouters[routerMethod]; !exist {
							s.obsRouters[routerMethod] = struct{}{}
							s.obsCol.notifyObs(&rcommon.NewObs{
								RouterMethod: routerMethod,
							})
							log.Info("RonFi new obs found", "tx", tx.Hash().String(), "obs", tx.To(), "methodId", fmt.Sprintf("0x%08x", methodID))
						}
					}
				}
			}
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

func (s *Stats) UpdatePairTokensInfo() {
	// update pairs/tokens into mysql
	newPairsInfo := s.di.GetNewPairsInfo()
	newPoolsInfo := s.di.GetNewPoolsInfo()
	newTokensInfo := s.di.GetNewTokensInfo()
	if newPairsInfo != nil && len(newPairsInfo) > 0 {
		pRecords := make([]*db.PairInfoRecord, 0, len(newPairsInfo))
		for pair, info := range newPairsInfo {
			pRecord := &db.PairInfoRecord{
				Pair:         pair.String(),
				Name:         info.Name,
				Index:        info.Index,
				BothBriToken: info.BothBriToken,
				KeyToken:     info.KeyToken.String(),
				Token0:       info.Token0.String(),
				Token1:       info.Token1.String(),
				Factory:      info.Factory.String(),
			}
			pRecords = append(pRecords, pRecord)
		}
		s.mysql.InsertPairsInfo(pRecords)
		log.Info("RonFi Stats new pairs", "count", len(pRecords))
	} else {
		log.Info("RonFi Stats No new pairs found!")
	}

	if newPoolsInfo != nil && len(newPoolsInfo) > 0 {
		pRecords := make([]*db.PoolInfoRecord, 0, len(newPoolsInfo))
		for pool, info := range newPoolsInfo {
			pRecord := &db.PoolInfoRecord{
				Pool:        pool.String(),
				Name:        info.Name,
				Token0:      info.Token0.String(),
				Token1:      info.Token1.String(),
				Fee:         int(info.Fee.Int64()),
				TickSpacing: info.TickSpacing,
				Factory:     info.Factory.String(),
			}
			pRecords = append(pRecords, pRecord)
		}
		s.mysql.InsertPoolsInfo(pRecords)
		log.Info("RonFi Stats new pools", "count", len(newPoolsInfo))
	} else {
		log.Info("RonFi Stats No new pools found!")
	}

	if newTokensInfo != nil && len(newTokensInfo) > 0 {
		tRecords := make([]*db.TokenInfoRecord, 0, len(newTokensInfo))
		for token, info := range newTokensInfo {
			tRecord := &db.TokenInfoRecord{
				Token:    token.String(),
				Symbol:   info.Symbol,
				Decimals: info.Decimals,
			}
			tRecords = append(tRecords, tRecord)
			s.mysql.InsertTokensInfo(tRecords)
		}
	} else {
		log.Info("RonFi Stats No new tokens found!")
	}

	// Merge new pairs into all pairs map
	s.di.MergePairTokensInfo()
}
