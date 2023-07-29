package trading

import (
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"fmt"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/uniswap"
	v2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	v3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"math/big"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/cmap"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

type HuntingSkipReason int

func (r HuntingSkipReason) String() string {
	switch r {
	case SkipReasonInQueueLong:
		return "InQueueLong"
	case SkipReasonNodeOutOfSync:
		return "NodeOutOfSync"
	case SkipReasonGasPriceHigh:
		return "GasPriceHigh"
	case SkipReasonGasPriceLow:
		return "GasPriceLow"
	case SkipReasonSplitToOtherTradingNode:
		return "Split"
	case SkipReasonOtherTradingNode:
		return "OtherTradingNode"
	case SkipReasonNoMatchedPair:
		return "NoMatchedPair"
	case SkipReasonApplyTransactionFail:
		return "ApplyTxFail"
	case SkipReasonApplyTransactionNoSwapPairs:
		return "ApplyTxNoSwapPairs"
	case SkipReasonNoProfit:
		return "NoProfit"
	case SkipReasonTrivialProfit:
		return "TrivialProfit"
	case SkipReasonImpossible:
		return "Impossible"
	case SkipReasonNoBalanceOrTokenUnSupport:
		return "NoBalanceOrTokenUnSupport"
	case SkipReasonDexSwapHuntingFail:
		return "DexSwapHuntingFail"
	case SkipReasonMaxHops:
		return "MaxHops"
	case SkipReasonDuplicatedSenderNonce:
		return "DupSNonce"
	case SkipReasonGiveWayForVipHunting:
		return "GiveWayForVipHunting"
	case SkipReasonSlippage:
		return "Slippage"
	case SkipReasonHuntingTaskPoolBusy:
		return "Busy"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

const (
	SkipReasonInQueueLong HuntingSkipReason = iota
	SkipReasonNodeOutOfSync
	SkipReasonGasPriceHigh
	SkipReasonGasPriceLow
	SkipReasonSplitToOtherTradingNode
	SkipReasonOtherTradingNode
	SkipReasonNoMatchedPair
	SkipReasonApplyTransactionFail
	SkipReasonApplyTransactionNoSwapPairs
	SkipReasonNoProfit
	SkipReasonTrivialProfit
	SkipReasonImpossible
	SkipReasonNoBalanceOrTokenUnSupport
	SkipReasonDexSwapHuntingFail
	SkipReasonMaxHops
	SkipReasonDuplicatedSenderNonce
	SkipReasonGiveWayForVipHunting
	SkipReasonSlippage
	SkipReasonHuntingTaskPoolBusy
	HuntingSkipReasonMax
)

var (
	// ErrChannelFull is a special error for identifying overflowing channel buffers
	ErrChannelFull = errors.New("channel full")
)

type ProfitDetail struct {
	loopName         string
	targetToken      common.Address
	amountIn         *big.Int
	txFeeInToken     float64
	txFeeInUsd       float64
	netProfitInUsd   float64
	grossProfitInUsd float64
	profitInToken    float64
	uniProfit        *uniswap.CycleWithProfit
}

type ProfitDetails []*ProfitDetail

func (p ProfitDetails) Len() int {
	return len(p)
}
func (p ProfitDetails) Less(i, j int) bool {
	return p[i].netProfitInUsd > p[j].netProfitInUsd
}
func (p ProfitDetails) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Worker struct {
	eth              rcommon.Backend
	chain            *core.BlockChain
	chainConfig      *params.ChainConfig
	signer           types.Signer
	tradingTxCh      chan types.Transactions
	client           *ethclient.Client
	di               *defi.Info
	stateDbsCopyCh   chan struct{}
	idleStateDbsLock sync.Mutex
	idleStateDbs     []*state.StateDB
	idleState        *state.StateDB
	appState         *state.StateDB
	appLock          sync.Mutex
	nonceState       *state.StateDB
	nonceLock        sync.Mutex
	currentBlockHash common.Hash
	currentBlockNum  uint64
	currentBlock     *types.Block
	done             chan struct{}

	currentBlockReceivedTime mclock.AbsTime

	txCount    uint64
	pairGasMap map[string]uint64
	//newPairLock sync.RWMutex

	executorPrivKey []*ecdsa.PrivateKey
	executorAddress []common.Address

	txDuplicatedSenderNonce uint64

	initialBalance      atomic.Value
	prevBalance         atomic.Value
	prevReportResetTime time.Time

	dryRun           bool
	minHuntingProfit float64 // Trigger the hunting if profitInToken >= txFeeInToken*this
	v3LoopsDb        *uniswap.V3Loops

	startTime           mclock.AbsTime
	totalArbTxs         uint64
	totalArbTxsCanceled uint64
	reportedTotalArbTxs uint64
	totalExecutors      uint64
	chiBalance          uint64
	bnbBalance          float64

	huntedRecords   cmap.ConcurrentMapset
	huntingSkipStat [HuntingSkipReasonMax]int

	skipReasonFile *os.File

	runningApp    int64
	huntingTxs    int64
	huntingPairs  int
	logHuntingTxs bool

	numCPU int
}

func NewWorker(eth rcommon.Backend, chainConfig *params.ChainConfig, client *ethclient.Client, di *defi.Info, dryRun bool, minHuntingProfit float64) *Worker {
	return &Worker{
		eth:              eth,
		chain:            eth.BlockChain(),
		chainConfig:      chainConfig,
		signer:           types.MakeSigner(chainConfig, big.NewInt(17034870), 1681266455), // shanghai upgrade
		tradingTxCh:      eth.GetTradingDexTxCh(),
		client:           client,
		di:               di,
		dryRun:           dryRun,
		minHuntingProfit: minHuntingProfit,
		numCPU:           runtime.NumCPU(),
	}
}

func (w *Worker) Init(executorPrivateKey string, pairGasMap map[string]uint64, v3LoopsDb *uniswap.V3Loops) bool {
	var (
		err error
	)

	w.initialBalance.Store(defi.RonFiBalance{})
	w.prevBalance.Store(defi.RonFiBalance{})
	w.prevReportResetTime = time.Now()

	w.skipReasonFile, err = os.OpenFile("/home/bsc/bin/logs/skip-reason.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Error("RonFi Init: fail to open/create 'skip-reason.json' file")
		return false
	}

	for i := range w.huntingSkipStat {
		w.huntingSkipStat[i] = 0
	}

	atomic.StoreUint64(&w.txCount, 0)
	atomic.StoreUint64(&w.totalArbTxs, 0)
	w.txDuplicatedSenderNonce = 0

	w.pairGasMap = pairGasMap
	w.v3LoopsDb = v3LoopsDb

	w.stateDbsCopyCh = make(chan struct{}, 128)
	w.startTime = mclock.Now()
	w.currentBlockNum = w.chain.CurrentBlock().Number.Uint64()
	w.currentBlock = w.chain.GetBlockByNumber(w.currentBlockNum)
	w.currentBlockHash = w.currentBlock.Hash()
	w.appState, _ = w.chain.StateAt(w.currentBlock.Root())
	w.idleState = w.appState.Copy()
	w.nonceState = w.appState.Copy()
	w.stateDbsUpdate()
	log.Info("RonFi worker Init. stateDbsUpdate", "stateCopies", len(w.idleStateDbs))
	w.currentBlockReceivedTime = w.startTime

	//get account from private key
	keys := strings.Split(executorPrivateKey, "-")
	if len(keys) == 0 {
		log.Error("RonFi Init: Executor Private Key is Empty")
		return false
	}

	w.totalExecutors = uint64(len(keys))
	w.executorPrivKey = make([]*ecdsa.PrivateKey, len(keys))
	w.executorAddress = make([]common.Address, len(keys))
	tmpMap := make(map[common.Address]bool)
	for i, key := range keys {
		privateKeyECDSA, err := toPrivateKey(key)
		if err != nil {
			log.Error("RonFi Generate ECDSA private key failed", "index", i, "err", err)
			return false
		}
		w.executorPrivKey[i] = privateKeyECDSA

		executorAddress, err := privateKeyToAddress(key)
		if err != nil {
			log.Error("RonFi Generate Account by private key failed", "index", i, "err", err)
			return false
		}
		log.Debug("RonFi Init", "executorAddress", rcommon.AbbrHexString(executorAddress.String()))
		w.executorAddress[i] = executorAddress
		tmpMap[executorAddress] = true
	}
	log.Info("RonFi Init executors of this node", "total", len(keys))

	for _, addr := range rcommon.AllTradingExecutors {
		if _, ok := tmpMap[addr]; !ok {
			w.executorAddress = append(w.executorAddress, addr)
			tmpMap[addr] = true
		}
	}
	log.Info("RonFi Init all executors", "total", len(w.executorAddress))
	tmpMap = nil

	if !w.InitRonFiSwap() {
		log.Error("RonFi InitRonFiSwap() failed")
		return false
	}

	balance := w.di.GetAllBalance(w.executorAddress, false)
	log.Info("RonFi Init", "Initial Executor Eth", balance.Eth, "Initial Contract Balance in USD", balance.ContractTotal)
	//if balance.Eth == 0 || balance.ContractTotal == 0 {
	//	log.Error("RonFi wallet initial Eth and/or contract initial balance is zero")
	//	return false
	//}

	w.initialBalance.Store(balance.Copy())
	w.prevBalance.Store(balance)
	w.chiBalance = balance.ContractChi
	w.bnbBalance = balance.Eth

	w.huntedRecords = cmap.NewMapset(256, maxWalletNonceKnownWallet)

	go w.stateDbLoop()
	go w.workLoop()

	return true
}

func (w *Worker) UnInit() {
	close(w.done)
	if !w.huntedRecords.IsEmpty() {
		totalDexTxCount := atomic.LoadUint64(&w.txCount)
		log.Info("RonFi UnInit report", "huntedRecords", w.huntedRecords.Cardinality(), "duplicated(sender+nonce)(%)", float64(w.txDuplicatedSenderNonce)/float64(totalDexTxCount)*100.0)
	}
	atomic.StoreUint64(&w.txCount, 0)
	atomic.StoreUint64(&w.totalArbTxs, 0)
	if w.skipReasonFile != nil {
		_ = w.skipReasonFile.Close()
	}
}

func (w *Worker) UpdateDefiInfo(di *defi.Info) {
	w.di = di
}

// SetCurrentBlock used for debug only
func (w *Worker) SetCurrentBlock(block *types.Block) {
	w.currentBlock = block
	w.currentBlockHash = block.Hash()
	w.currentBlockNum = block.NumberU64()
}

func (w *Worker) LogSkipReason() {
	totalDexTxCount := atomic.LoadUint64(&w.txCount)
	if totalDexTxCount == 0 {
		return
	}
	log.Info("RonFi LogSkipReason")
	percentage := float64(0)
	for reason, count := range w.huntingSkipStat {
		percentage = float64(count) / float64(totalDexTxCount) * 100
		log.Info("RonFi skip reason stat", rcommon.Float2Str(percentage, 2)+"%", HuntingSkipReason(reason).String(), "skip", count)
	}

	log.Info("RonFi", "huntedRecords", w.huntedRecords.Cardinality(), "duplicated(Sender+Nonce)(%)", float64(w.txDuplicatedSenderNonce)/float64(totalDexTxCount)*100.0)
}

func (w *Worker) PushStateDbsCopyEvent() error {
	select {
	case w.stateDbsCopyCh <- struct{}{}:
		return nil
	default:
		return ErrChannelFull
	}
}

func (w *Worker) stateDbLoop() {
	for {
		select {
		case <-w.stateDbsCopyCh:
			for { // to avoid long lock, only produce one copy within each loop
				w.idleStateDbsLock.Lock()
				copiesNeed := maxStateDbsCopies - len(w.idleStateDbs) // when idle stateDbs is too little, produce more copy
				if copiesNeed > 0 {
					w.idleStateDbs = append(w.idleStateDbs, w.idleState.Copy())
				}
				w.idleStateDbsLock.Unlock()
				if copiesNeed <= 0 {
					break
				}
			}

		case <-w.done:
			return
		}
	}
}

func (w *Worker) stateDbsUpdate() {
	w.idleStateDbsLock.Lock()
	w.idleStateDbs = make([]*state.StateDB, 0, maxStateDbsCopies) // drop all current idle stateDbs and copy the new w.appState
	// note: to avoid long lock here, we produce one copy only instead of full copies, and let 'stateDbLoop' do the remaining job
	w.idleStateDbs = append(w.idleStateDbs, w.idleState.Copy())
	w.idleStateDbsLock.Unlock()
	// let 'stateDbLoop' do the remaining copy job
	if err := w.PushStateDbsCopyEvent(); err != nil {
		log.Warn("RonFi push StateDbsCopyEvent", "error", err)
	}
}

// For task which is not time critical, we can produce a copy slowly here!
func (w *Worker) stateDbProduceOneCopy() (stateDb *state.StateDB) {
	w.idleStateDbsLock.Lock()
	stateDb = w.idleState.Copy()
	w.idleStateDbsLock.Unlock()
	return
}

// For hunting task, we need get one stateDB copy asap! here is a short-cut to get a pre-copied stateDB with most quick speed.
func (w *Worker) stateDbsConsumeOneCopy() (blockHash common.Hash, stateDb *state.StateDB) {
	w.idleStateDbsLock.Lock()
	blockHash = w.currentBlockHash
	if len(w.idleStateDbs) == 0 {
		stateDb = w.idleState.Copy()
		w.idleStateDbsLock.Unlock()
		if err := w.PushStateDbsCopyEvent(); err != nil {
			log.Warn("RonFi push StateDbsCopyEvent", "error", err)
		}
		log.Warn("RonFi idleStateDbs run out, fallback to local copy")
		return
	}
	stateDb = w.idleStateDbs[0]
	w.idleStateDbs = w.idleStateDbs[1:] // consume t and shrink the array
	remains := len(w.idleStateDbs)
	w.idleStateDbsLock.Unlock()
	if remains <= maxStateDbsCopies/2 { // half of stateDbs have been consumed
		if err := w.PushStateDbsCopyEvent(); err != nil {
			log.Warn("RonFi push StateDbsCopyEvent", "error", err)
		}
	}
	return
}

// in case a stateDb consumed but without any state change (i.e. there's no applyTransaction), this consumed stateDb can be recycled
func (w *Worker) stateDbsRecycleOneCopy(blockHash common.Hash, stateDb *state.StateDB) {
	w.idleStateDbsLock.Lock()
	if blockHash == w.currentBlockHash { // avoid recycle an old state db
		w.idleStateDbs = append(w.idleStateDbs, stateDb)
	}
	w.idleStateDbsLock.Unlock()
	return
}

// / queueTxsFilter filter txs to remove same From + same To
func (w *Worker) queueTxsFilter(txs types.Transactions) (filteredTxs types.Transactions) {
	filteredTxs = make(types.Transactions, 0, len(txs))
	if len(txs) == 0 {
		return
	}
	sameFromAndTo := make(map[common.Address]struct{}, len(txs))
	filtered := 0
	sigCached := 0
	start := mclock.Now()
	for _, tx := range txs {
		if sender, cache, err := types.RonFiSender(w.signer, tx); err != nil {
			continue
		} else {
			if cache {
				sigCached++
			}
			fromAndTo := new(big.Int).SetBytes(sender.Bytes())
			fromAndTo = fromAndTo.Add(fromAndTo, new(big.Int).SetBytes(tx.To().Bytes()))
			key := common.BigToAddress(fromAndTo)
			if _, same := sameFromAndTo[key]; same {
				w.ReportSkipReason(tx, SkipReasonInQueueLong, "sameFrmTo")
				filtered++
				continue
			} else {
				sameFromAndTo[key] = struct{}{}
				filteredTxs = append(filteredTxs, tx)
			}
		}
	}
	elapsed := mclock.Since(start)
	log.Info("RonFi hunting 2nd filter", "block", w.currentBlockNum, "sigCached", sigCached, "filtered", filtered, "remain", len(filteredTxs), "elapsed", elapsed)
	return
}

func (w *Worker) workLoop() {
	// subscribe new block event
	ChanBufferSize := 128
	chainHeadCh := make(chan core.ChainHeadEvent, ChanBufferSize)
	headersDone := make(chan struct{})
	huntingDone := make(chan struct{})
	chainHeadSub := w.eth.BlockChain().SubscribeChainHeadEvent(chainHeadCh)
	w.done = make(chan struct{})
	defer func() {
		chainHeadSub.Unsubscribe()
		close(chainHeadCh)
		close(huntingDone)
		close(headersDone)
	}()

	// block header event loop
	go func() {
		reportInterval := 30 * time.Minute
		reportTimer := time.NewTimer(reportInterval)

		defer func() {
			reportTimer.Stop()
		}()

		skipBlocks := 0
		for {
			select {
			case ev := <-chainHeadCh:
				if len(chainHeadCh) > 0 {
					// if more blocks event there, just skip current and go to next block
					skipBlocks++
					continue
				}
				if skipBlocks > 0 {
					log.Warn("RonFi workLoop some chain head events skipped", "blocks", skipBlocks)
					skipBlocks = 0
				}

				if ev.Block != nil && w.chain != nil {
					block := ev.Block
					w.appLock.Lock() //lock here, because LogEvChan will be attached with it.
					if tmpState, err := w.chain.StateAt(block.Root()); err != nil {
						log.Error("RonFi workLoop StateAtWithSharedPool fail", "err", err)
						w.appLock.Unlock()
						continue
					} else {
						w.appState = tmpState
					}
					w.idleState = w.appState.Copy()
					w.nonceState = w.appState.Copy()
					w.currentBlockHash = block.Hash()
					w.currentBlockNum = block.NumberU64()
					w.currentBlock = block
					w.appLock.Unlock()
					w.stateDbsUpdate()
					w.currentBlockReceivedTime = mclock.Now()
				}

			case <-reportTimer.C:
				reportTimer.Reset(reportInterval)
				log.Info("worker report", "huntedRecords", w.huntedRecords.Cardinality())
				log.Info("worker report", "totalArbTxs", w.totalArbTxs, "totalCancel", w.totalArbTxsCanceled)
				log.Info("worker report", "duration", mclock.Since(w.startTime))

			case err := <-chainHeadSub.Err():
				log.Error("RonFi  chainHead subscribe error!", "chainHeadSub.Err", err)
				return
			case <-huntingDone:
				log.Warn("RonFi block header event loop exit!!!")
				return
			}
		}
	}()

	for {
		select {
		case txs := <-w.tradingTxCh:

			queueSize := len(w.tradingTxCh)
			if queueSize >= DexTxInQueueFilterMin || len(txs) >= DexTxInQueueFilterMin {
				queueTxs := make(types.Transactions, 0, DexTxInQueueSkipMax*8)
				queueTxs = append(queueTxs, txs...)
				inQueue := len(txs)
				for i := 0; i < queueSize; i++ {
					remains := <-w.tradingTxCh
					queueTxs = append(queueTxs, remains...)
					inQueue += len(remains)
				}

				// 2nd filter, remove same From + same To
				filteredTxs := w.queueTxsFilter(queueTxs)

				if len(filteredTxs) >= DexTxInQueueSkipMax { // skip all txs if we see a surprising long queue here!
					// in busy case, a batch of txs received, skip it to avoid big delay, and log warning
					log.Warn("RonFi hunting long queue, skipped", "block", w.currentBlockNum, "filtered", len(filteredTxs), "inQueue", inQueue, "queueSize", queueSize)
					blockStr := fmt.Sprintf("b=%d", w.currentBlockNum)
					for _, tx := range filteredTxs {
						w.ReportSkipReason(tx, SkipReasonInQueueLong, blockStr)
					}
					continue
				}
				txs = filteredTxs
			}

			for i, tx := range txs {
				txCopy := tx
				iCopy := i
				if !common.RonHuntingTaskDispatch(func() {
					if w.handlePromotedTx(txCopy) {
						if w.logHuntingTxs {
							w.logHuntingTxs = false
							log.Info("RonFi arb", "huntingTxs", atomic.LoadInt64(&w.huntingTxs), "huntingPairs", w.huntingPairs, "queue", queueSize, "txs", len(txs), "i", iCopy)
						}
					}
				}) {
					w.ReportSkipReason(tx, SkipReasonHuntingTaskPoolBusy, fmt.Sprintf("b=%d", w.currentBlockNum))
					time.Sleep(time.Millisecond)
				}
			}

		case <-w.done:
			rpc.StartTrading = false // redundant! make sure no new filling of tradingTxCh chan
			queueSize := len(w.tradingTxCh)
			for i := 0; i < queueSize; i++ {
				_ = <-w.tradingTxCh // clear the chan before quit
			}

			headersDone <- struct{}{}
			log.Warn("RonFi worker loop done!")
			return
		}
	}
}

func (w *Worker) handlePromotedTx(tx *types.Transaction) (hunting bool) {
	if current := atomic.AddInt64(&w.huntingTxs, 1); current >= common.MaxRonHuntingTasks/2 {
		w.logHuntingTxs = true
	}
	defer func() {
		atomic.AddInt64(&w.huntingTxs, -1)
	}()

	hunting = false
	handlerStartTime := mclock.Now()
	atomic.AddUint64(&w.txCount, 1)

	// in case the node is out of sync, skip the hunting. 6s > 1block
	if !w.dryRun && handlerStartTime.Sub(w.currentBlockReceivedTime).Milliseconds() > 6000 {
		w.ReportSkipReason(tx, SkipReasonNodeOutOfSync, "")
		return
	}

	gasPriceMinAllowed := GasPriceMinAllowedIdle
	if tx.GasPrice().Cmp(gasPriceMinAllowed) < 0 {
		// gas price less than 5 Gwei is illegal, and rise this '5' to '5.01' in case blockchain busy.
		w.ReportSkipReason(tx, SkipReasonGasPriceLow, "")
		return
	}

	if rpc.FilterDupTxs {
		// check if duplicated dexTx
		if sender, err := types.Sender(w.signer, tx); err != nil {
			w.ReportSkipReason(tx, SkipReasonImpossible, "sig fail")
			return
		} else {
			nonce := tx.Nonce()
			key := binary.BigEndian.Uint64(sender[12:]) // Last 8-bytes of wallet address
			if lastNonce, exist := w.huntedRecords.Get(key); exist {
				if nonce <= lastNonce {
					w.txDuplicatedSenderNonce++
					//log.Debug("RonFi handlePromotedTx duplicated dexTx", "tx", tx.Hash().String(), "from", sender)
					w.ReportSkipReason(tx, SkipReasonDuplicatedSenderNonce, fmt.Sprintf("nonce=%d, last=%d", nonce, lastNonce))
					return
				}
			}
			w.huntedRecords.Set(key, nonce)
		}
	}

	// apply transaction
	var (
		applySuccess, reverted bool
		err                    string
	)
	ronfiTxHash := ronFiTxHash(tx.Hash())
	_, appState := w.stateDbsConsumeOneCopy()
	if appState == nil {
		log.Warn("RonFi handlePromotedTx appState is nil")
	} else {
		applySuccess, reverted, err = w.applyTransaction(tx, ronfiTxHash, appState)
		if !applySuccess || reverted {
			w.ReportSkipReason(tx, SkipReasonApplyTransactionFail, err)
			return
		}
		vlogs := appState.GetLogs(ronfiTxHash, w.currentBlockNum, common.Hash{})
		swapPairsInfo := w.di.ExtractSwapPairInfo(tx, *tx.To(), vlogs, defi.RonFiExtractTypeHunting)
		if len(swapPairsInfo) > 0 {
			w.huntingTxEvent(appState, tx, 0, swapPairsInfo, handlerStartTime)
		}
	}

	return
}

func (w *Worker) huntingTxEvent(appState *state.StateDB, tx *types.Transaction, pairId int, pairsInfo []*defi.SwapPairInfo, handlerStartTime mclock.AbsTime) {
	v3States := make(map[common.Address]*v3.PoolState)
	v2Pools := make(map[common.Address]*v2.Pool)
	v3Pools := make(map[common.Address]*v3.Pool)
	v2AmountIOs := make(map[common.Address]map[string]*big.Int)
	v3AmountIOs := make(map[common.Address]map[string]*v3.DetailOut)

	for i, info := range pairsInfo {
		edge := uniswap.ToV3Edge(info)
		if edge == nil {
			continue
		}

		arbs := w.v3LoopsDb.FindLoops(edge)
		profits := make(ProfitDetails, 0, len(arbs))
		for _, arb := range arbs {
			lpCycle := uniswap.FromAddress(w.di, tx, appState, v2AmountIOs, v3AmountIOs, v2Pools, v3Pools, arb[0].TokenIn, info, arb)
			if lpCycle == nil {
				log.Warn("RonFi huntingTxEvent: uniswap.FromAddress fail", "idx", i, "loopId", arb.String(), "tx", tx.Hash().String(), "pair", info.Address)
				continue
			}

			if !lpCycle.AutoUpdate(v3States) {
				//log.Info("RonFi huntingV3TxEvent, lpCycle.AutoUpdate fail", "idx", idx, "loopId", arb.LoopId, "tx", tx.Hash().String(), "pair", pairInfo.Address)
				continue
			}

			res := lpCycle.CalculateArbitrage()
			if res != nil && res.Profitable {
				profit := &uniswap.CycleWithProfit{
					Cycle:  lpCycle,
					Profit: res,
				}
				price := defi.GetTradingTokenPrice(profit.Cycle.InputToken)
				profitInToken := rcommon.EthBigInt2Float64(profit.Profit.BestProfit)
				grossProfitInUsd := profitInToken / price * defi.GetTradingTokenPrice(rcommon.USDC)

				txFeeInBnb := rcommon.EthBigInt2Float64(new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(profit.Cycle.SumGasNeed)))
				txFeeInUsd := txFeeInBnb * defi.GetTradingTokenPrice(rcommon.USDC)
				txFeeInToken := price * txFeeInBnb
				netProfitInUsd := grossProfitInUsd - txFeeInUsd
				profitDetail := ProfitDetail{
					loopName:         arb.String(),
					targetToken:      profit.Cycle.InputToken,
					amountIn:         profit.Profit.SwapAmount,
					txFeeInToken:     txFeeInToken,
					txFeeInUsd:       txFeeInUsd,
					netProfitInUsd:   netProfitInUsd,
					grossProfitInUsd: grossProfitInUsd,
					profitInToken:    profitInToken,
					uniProfit:        profit,
				}
				profits = append(profits, &profitDetail)

				log.Warn("RonFi huntingTxEvent, lpCycle.CalculateArbitrage --- succeed",
					"tx", tx.Hash().String(),
					"idx", i,
					"input token", profit.Cycle.InputToken,
					"loop", arb.String(),
					"pair", info.Address,
					"amountIn", res.SwapAmount,
					"grossProfitInUsd", profitDetail.grossProfitInUsd,
					"netProfitInUsd", profitDetail.netProfitInUsd)
			}
		}

		if len(profits) > 0 {
			sort.Sort(profits)
			highestProfit := profits[0]

			if highestProfit.netProfitInUsd > w.minHuntingProfit {
				log.Info("RonFi huntingTxEvent early hunting",
					"i", i, "loops", len(arbs),
					"loop", highestProfit.loopName, "netProfitInUsd", highestProfit.netProfitInUsd)

				amountIns := ""
				for _, amountIn := range highestProfit.uniProfit.Profit.AmountIns {
					amountIns = fmt.Sprintf("%s %s", amountIns, amountIn.String())
				}
				amountOuts := ""
				for _, amountOut := range highestProfit.uniProfit.Profit.AmountOuts {
					amountOuts = fmt.Sprintf("%s %s", amountOuts, amountOut.String())
				}
				log.Info("RonFi huntingTxEvent found profit",
					"tx", tx.Hash().String(),
					"inputToken", highestProfit.uniProfit.Cycle.InputToken,
					"target pair", info.Address,
					"target amountIn", info.AmountIn.String(),
					"target amountOut", info.AmountOut.String(),
					"lower", highestProfit.uniProfit.Profit.Boundary[0],
					"upper", highestProfit.uniProfit.Profit.Boundary[1],
					"target dir", info.Dir,
					"amountIns", amountIns,
					"amountOuts", amountOuts,
					"loop", highestProfit.loopName,
					"cycle", highestProfit.uniProfit.Cycle.Name,
					"iters", highestProfit.uniProfit.Profit.Iters,
					"amountIn", highestProfit.uniProfit.Profit.SwapAmount,
					"profit", highestProfit.uniProfit.Profit.BestProfit,
					"grossProfitInUsd", highestProfit.grossProfitInUsd,
					"txFeeInUsd", highestProfit.txFeeInUsd,
					"netProfitInUsd", highestProfit.netProfitInUsd)

				w.huntingTxPair(tx, i, handlerStartTime, w.currentBlockNum, info, highestProfit.uniProfit, highestProfit)
			}
		}
	}
}

func (w *Worker) huntingTxPair(tx *types.Transaction, pairId int, handlerStartTime mclock.AbsTime, blockNumber uint64, swapPairInfo *defi.SwapPairInfo, bestProfit *uniswap.CycleWithProfit, info *ProfitDetail) {
	w.huntingPairs++
	defer func() {
		w.huntingPairs--
	}()

	gasMaxLimit := bestProfit.Cycle.SumGasNeed
	profitMin := rcommon.ToWei(info.txFeeInToken, 18)
	randomExecutorId := bestProfit.Cycle.LoopId.Uint64() + tx.Hash().TailUint64()
	// run the ring to draw an executor
	ringId := randomExecutorId % w.totalExecutors
	var arbTx *types.Transaction

	if ok, arb := w.DexSwapHunting(w.executorPrivKey[ringId], &w.executorAddress[ringId], tx, profitMin, gasMaxLimit, bestProfit, w.dryRun, handlerStartTime); !ok {
		w.huntingSkipStat[SkipReasonDexSwapHuntingFail]++
	} else {
		arbTx = arb
	}

	totalArbTxs := atomic.AddUint64(&w.totalArbTxs, 1)
	if arbTx != nil {
		log.Info(
			"RonFi handler",
			"dexTx", tx.Hash(),
			"arbTx", arbTx.Hash(),
			"pair", swapPairInfo.Address,
			"id", pairId,
			"f", rcommon.Float2Str(info.txFeeInToken, 6),
			"p", rcommon.Float2Str(info.profitInToken, 6),
			"t", w.tokenSymbol(info.targetToken),
			"#", totalArbTxs)
		{
			// log and compare the original dex tx amount and the arb tx amount
			tag, amount, token := getTokenAmount(w.di, []*defi.SwapPairInfo{
				swapPairInfo,
			})
			if amount != 0 {
				log.Info("RonFi handler",
					"dexTx", tx.Hash(),
					"peer", prefix6HexUint(tx.PeerId),
					"rank", tx.PeerRank,
					"t", w.tokenSymbol(token),
					tag, rcommon.Float2Str(amount, 6),
					"block", blockNumber)
			} else {
				log.Info("RonFi handler",
					"dexTx", tx.Hash(),
					"peer", prefix6HexUint(tx.PeerId),
					"rank", tx.PeerRank,
					"block", blockNumber)
			}
		}
	}
}

func (w *Worker) applyTransaction(tx *types.Transaction, txHash common.Hash, state *state.StateDB) (bool, bool, string) {
	if state == nil {
		return false, false, "state == nil"
	}

	if w.chain == nil {
		return false, false, "w.chain == nil"
	}

	gasPool := new(core.GasPool).AddGas(85_000_000)

	author := rcommon.ZeroAddress

	state.SetTxContext(txHash, 0)
	// 'applySuccess' is the status of ApplyTransaction, 'failed' is the status of whether an applied transaction is reverted (i.e. failed but packed into blockchain).
	if applySuccess, failed, _, err := core.ApplyRonfiTransaction(w.chainConfig, w.chain, author, gasPool, state, w.currentBlock.Header(), tx, *w.chain.GetVMConfig()); applySuccess {
		return applySuccess, failed, err
	} else {
		return applySuccess, failed, err
	}
}

func (w *Worker) ReportSkipReason(tx *types.Transaction, skipReason HuntingSkipReason, extraInfo string) {
	w.huntingSkipStat[skipReason]++
	if rpc.LogSkipReason {
		reason := ""
		if len(extraInfo) != 0 {
			reason = fmt.Sprintf("0x%010x, %s, %s\n", tx.Hash().Uint64()>>24, skipReason.String(), extraInfo)
		} else {
			reason = fmt.Sprintf("0x%010x, %s\n", tx.Hash().Uint64()>>24, skipReason.String())
		}
		if w.skipReasonFile != nil {
			_, _ = w.skipReasonFile.WriteString(reason)
		}
	}
}

func (w *Worker) tokenSymbol(token common.Address) string {
	if info := w.di.GetTokenInfo(token); info != nil {
		return info.Symbol
	} else {
		return "undefined"
	}
}