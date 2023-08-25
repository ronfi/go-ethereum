package trading

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"fmt"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	ronswapv3fe "github.com/ethereum/go-ethereum/ronfi/contracts/contract_ronswapv3fe"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/uniswap"
	v2 "github.com/ethereum/go-ethereum/ronfi/uniswap/v2"
	v3 "github.com/ethereum/go-ethereum/ronfi/uniswap/v3"
	"github.com/metachris/flashbotsrpc"
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
	gasPrice         *big.Int
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

	ronSwapInst *ronswapv3fe.Ronswapv3fe
	flashRpc    *flashbotsrpc.BuilderBroadcastRPC

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
	var baseFee *big.Int
	if feeHis, err := w.client.FeeHistory(context.Background(), 1, new(big.Int).SetUint64(w.currentBlockNum), nil); err != nil {
		log.Warn("RonFi swap transaction, FeeHistory failed, err=%s", err)
	} else {
		baseFee = feeHis.BaseFee[0]
	}
	w.gasPrice = new(big.Int).Mul(baseFee, big.NewInt(2))
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

	balance := w.di.GetAllBalance(w.executorAddress)
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
		log.Warn("RonFi stateDbsUpdate push StateDbsCopyEvent", "error", err)
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
			log.Warn("RonFi stateDbsConsumeOneCopy push StateDbsCopyEvent", "error", err)
		}
		//log.Warn("RonFi idleStateDbs run out, fallback to local copy")
		return
	}
	stateDb = w.idleStateDbs[0]
	w.idleStateDbs = w.idleStateDbs[1:] // consume t and shrink the array
	remains := len(w.idleStateDbs)
	w.idleStateDbsLock.Unlock()
	if remains <= maxStateDbsCopies/2 { // half of stateDbs have been consumed
		if err := w.PushStateDbsCopyEvent(); err != nil {
			log.Warn("RonFi stateDbsConsumeOneCopy push StateDbsCopyEvent", "error", err)
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
	//start := mclock.Now()
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
	//elapsed := mclock.Since(start)
	//log.Info("RonFi hunting 2nd filter", "block", w.currentBlockNum, "sigCached", sigCached, "filtered", filtered, "remain", len(filteredTxs), "elapsed", elapsed)
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

					var baseFee *big.Int
					if feeHis, err := w.client.FeeHistory(context.Background(), 1, new(big.Int).SetUint64(w.currentBlockNum), nil); err != nil {
						log.Warn("RonFi swap transaction, FeeHistory failed, err=%s", err)
					} else {
						baseFee = feeHis.BaseFee[0]
					}
					w.gasPrice = new(big.Int).Add(baseFee, new(big.Int).Div(new(big.Int).Mul(baseFee, big.NewInt(2)), big.NewInt(10)))
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
		applySuccess, reverted, _, err = applyTransaction(w.chain, w.chainConfig, w.currentBlock, tx, ronfiTxHash, appState)
		if !applySuccess || reverted {
			w.ReportSkipReason(tx, SkipReasonApplyTransactionFail, err)
			return
		}
		vlogs := appState.GetLogs(ronfiTxHash, w.currentBlockNum, common.Hash{})
		swapPairsInfo := w.di.ExtractSwapPairInfo(tx, *tx.To(), vlogs, defi.RonFiExtractTypeStats)
		if len(swapPairsInfo) > 0 {
			//w.huntingTxEvent(appState, tx, 0, swapPairsInfo, handlerStartTime)
			_, newStatedb := w.stateDbsConsumeOneCopy()
			if appState == nil {
				log.Warn("RonFi handlePromotedTx newStatedb is nil")
			} else {
				w.huntingTxEvent(newStatedb, tx, 0, swapPairsInfo, handlerStartTime)
			}
		}
	}

	return
}

func (w *Worker) sandwichTx(tx *types.Transaction, pairInfo *defi.SwapPairInfo, appState *state.StateDB, handlerStartTime mclock.AbsTime) {
	var (
		applySuccess, reverted bool
		err                    string
		gasUsed, realBLegGas   uint64
		tokenPairsAndFee       []*big.Int
	)

	randomExecutorId := tx.Hash().TailUint64()
	ringId := randomExecutorId % w.totalExecutors
	executorPrivKey := w.executorPrivKey[ringId]
	executorAddress := w.executorAddress[ringId]

	ronSandwich := NewRonSandwich(w.client, w.di, w.ronSwapInst, executorPrivKey, executorAddress, w.chain, w.chainConfig, w.currentBlock, tx, pairInfo, appState.Copy())
	if ronSandwich == nil {
		return
	}

	amountIn := big.NewInt(0)
	if !ronSandwich.optimize(pairInfo, amountIn) {
		//log.Warn("RonFi sandwichTx optimize fail", "tx", tx.Hash().String(), "pair", pairInfo.Address)
		return
	}

	if res := ronSandwich.prepare(pairInfo, amountIn); res != nil {
		bLegAmount := res.bLegAmountIn

		log.Info("RonFi sandwichTx prepare succeed",
			"tx", tx.Hash().String(),
			"pair", pairInfo.Address,
			"amountIn", rcommon.EthBigInt2Float64(amountIn),
			"bLegAmount", rcommon.EthBigInt2Float64(bLegAmount))

		txs := make([]*types.Transaction, 0, 3)

		// build bundle
		aLegGas := res.aLegGasUsed + 500000 // add 500K gas for aLeg
		bLegGas := res.bLegGasUsed + 500000 // add 500K gas for bLeg

		statedbCopy := appState.Copy()
		aLegNonce := statedbCopy.GetNonce(executorAddress)
		bLegNonce := aLegNonce + 1
		// aLegTx
		aLegPayloads, _ := ronSandwich.generatePayloads(pairInfo, amountIn, nil, statedbCopy)
		aLegTxFee := new(big.Int).Mul(w.gasPrice, new(big.Int).SetUint64(res.aLegGasUsed))
		aLegTx := ronSandwich.buildExecuteTx(aLegPayloads, true, []*big.Int{}, big.NewInt(0), big.NewInt(0), aLegNonce, w.gasPrice, aLegGas)
		// apply aLegTx
		if applySuccess, reverted, gasUsed, err = applyTransaction(w.chain, w.chainConfig, w.currentBlock, aLegTx, ronFiTxHash(aLegTx.Hash()), statedbCopy); !applySuccess || reverted {
			log.Warn("RonFi sandwichTx apply aLegTx fail", "tx", tx.Hash().String(), "pair", pairInfo.Address, "before gas", res.aLegGasUsed, "gas", gasUsed, "err", err)
			return
		}
		aLogs := statedbCopy.GetLogs(ronFiTxHash(aLegTx.Hash()), w.currentBlockNum, common.Hash{})
		txs = append(txs, aLegTx)

		// targetTx
		if applySuccess, reverted, _, err = applyTransaction(w.chain, w.chainConfig, w.currentBlock, tx, ronFiTxHash(tx.Hash()), statedbCopy); !applySuccess || reverted {
			log.Warn("RonFi sandwichTx apply target tx fail", "tx", tx.Hash().String(), "pair", pairInfo.Address, "err", err)
			return
		}
		txs = append(txs, tx)

		//bLegTx
		cycle, swapAmountIn, hasArb := w.sandwichBackRun(res.appState, tx, pairInfo, ronSandwich, handlerStartTime)
		rPairInfo := pairInfo.Reverse()
		bLegPayloads, _ := ronSandwich.generatePayloads(rPairInfo, bLegAmount, res.tokenFee, statedbCopy)
		var bLegTx *types.Transaction
		if hasArb {
			path := make([]common.Address, 0, len(cycle.PoolAddresses)*2)
			dirs := make([]uint8, 0, len(cycle.PoolAddresses))
			poolFee := make([]uint64, 0, len(cycle.PoolAddresses))
			tokenFee := make([]uint64, 0, len(cycle.PoolAddresses))
			poolType := make([]uint8, 0, len(cycle.PoolAddresses))
			if len(cycle.PoolAddresses) != len(cycle.SwapVectors) {
				log.Error("RonFi sandwichTx, swapVectors and poolAddress length mismatch")
				return
			}

			for i := 0; i < len(cycle.PoolAddresses); i++ {
				pool := cycle.PoolAddresses[i]
				swapVector := cycle.SwapVectors[i]
				token := swapVector.TokenOut
				dir := uint8(0)
				if !swapVector.ZeroForOne {
					dir = 1
				}
				dirs = append(dirs, dir)
				path = append(path, pool)
				path = append(path, token)
				poolFee = append(poolFee, swapVector.PoolFee)
				tokenFee = append(tokenFee, swapVector.TokenFee)
				poolType = append(poolType, swapVector.PoolType)
			}

			tokenPairsAndFee = make([]*big.Int, 2*len(cycle.PoolAddresses))
			for i := 0; i < len(cycle.PoolAddresses); i++ {
				tmp := new(big.Int).SetBytes(path[2*i].Bytes())
				tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(poolFee[i]), 160))
				tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(dirs[i])), 176))
				tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(poolType[i])), 192))
				tokenPairsAndFee[2*i] = tmp

				tmp = new(big.Int).SetBytes(path[2*i+1].Bytes())
				tmp = tmp.Add(tmp, new(big.Int).Lsh(new(big.Int).SetUint64(tokenFee[i]), 160))
				tokenPairsAndFee[2*i+1] = tmp
			}

			bLegTx = ronSandwich.buildExecuteTx(bLegPayloads, false, tokenPairsAndFee, swapAmountIn, big.NewInt(0), bLegNonce, w.gasPrice, bLegGas)
		} else {
			swapAmountIn = big.NewInt(0)
			bLegTx = ronSandwich.buildExecuteTx(bLegPayloads, false, tokenPairsAndFee, swapAmountIn, big.NewInt(0), bLegNonce, w.gasPrice, bLegGas)
		}
		if applySuccess, reverted, realBLegGas, err = applyTransaction(w.chain, w.chainConfig, w.currentBlock, bLegTx, ronFiTxHash(bLegTx.Hash()), statedbCopy); !applySuccess || reverted {
			log.Warn("RonFi sandwichTx apply bLegTx+arbTx fail", "tx", tx.Hash().String(), "hasArb", hasArb, "err", err)
			return
		}
		bLogs := statedbCopy.GetLogs(ronFiTxHash(bLegTx.Hash()), w.currentBlockNum, common.Hash{})

		aPairsInfo := w.di.ExtractSwapPairInfo(aLegTx, *aLegTx.To(), aLogs, defi.RonFiExtractTypeStats)
		bPairsInfo := w.di.ExtractSwapPairInfo(bLegTx, *bLegTx.To(), bLogs, defi.RonFiExtractTypeStats)
		grossProfit := calculateSandwichProfit(pairInfo, aPairsInfo, bPairsInfo)
		if grossProfit == nil || grossProfit.Cmp(big.NewInt(0)) <= 0 {
			log.Warn("RonFi sandwichTx calculateSandwichProfit fail", "tx", tx.Hash().String(), "pair", pairInfo.Address)
			return
		}

		bLegTxFee := new(big.Int).Mul(w.gasPrice, new(big.Int).SetUint64(res.bLegGasUsed))
		baseBundleTxFee := new(big.Int).Add(aLegTxFee, bLegTxFee)
		if grossProfit.Cmp(baseBundleTxFee) > 0 {
			bLegTxFee = new(big.Int).Div(new(big.Int).Mul(new(big.Int).Sub(grossProfit, aLegTxFee), big.NewInt(60)), big.NewInt(100))
			realBLegGas += 500000 // add 500k gas for bLeg
			bLegTxGasPrice := new(big.Int).Div(bLegTxFee, big.NewInt(int64(realBLegGas)))
			if bLegTxGasPrice.Cmp(w.gasPrice) < 0 {
				log.Warn("RonFi sandwichTx minimal profit!", "baseFee", w.gasPrice, "bLegGasPrice", bLegTxGasPrice, "bLegTxFee", rcommon.EthBigInt2Float64(bLegTxFee), "bLegGas", realBLegGas)
				return
			}
			bLegTx = ronSandwich.buildExecuteTx(bLegPayloads, false, tokenPairsAndFee, swapAmountIn, bLegTxFee, bLegNonce, bLegTxGasPrice, realBLegGas)
			txs = append(txs, bLegTx)

			// then simulate the txs, and send them to chain
			FlashbotSandWich(w.flashRpc, w.currentBlock, w.currentBlockNum, txs)
			totalArbTxs := atomic.AddUint64(&w.totalArbTxs, 1)
			log.Info(
				"RonFi sandwichTx bundle sent!",
				"target tx", tx.Hash().String(),
				"grossProfit", rcommon.EthBigInt2Float64(grossProfit),
				"bribe fee", rcommon.EthBigInt2Float64(bLegTxFee),
				"bLegTxGasPrice", bLegTxGasPrice,
				"#", totalArbTxs)
		}
	}
}

func (w *Worker) huntingTxEvent(appState *state.StateDB, tx *types.Transaction, pairId int, pairsInfo []*defi.SwapPairInfo, handlerStartTime mclock.AbsTime) {
	// first check sandwich
	for _, info := range pairsInfo {
		if info.TokenIn == rcommon.WETH {
			w.sandwichTx(tx, info, appState, handlerStartTime)
		}

		//else {
		//	w.backRun(appState, tx, pairId, pairsInfo, handlerStartTime)
		//}
	}

	return
}

func (w *Worker) sandwichBackRun(appState *state.StateDB, tx *types.Transaction, pairInfo *defi.SwapPairInfo, ronSandwich *RonSandwich, handlerStartTime mclock.AbsTime) (*uniswap.LPCycle, *big.Int, bool) {
	v3States := make(map[common.Address]*v3.PoolState)
	v2States := make(map[common.Address]*v2.PoolState)

	edge := uniswap.ToV3Edge(pairInfo)
	if edge == nil {
		return nil, nil, false
	}

	arbs := w.v3LoopsDb.FindLoops(edge)
	if len(arbs) == 0 {
		log.Info("RonFi backRun no matched loops!", "tx", tx.Hash().String(), "pair", pairInfo.Address)
		return nil, nil, false
	}

	profits := make(ProfitDetails, 0, len(arbs))
	for _, arb := range arbs {
		lpCycle := uniswap.FromAddress(w.di, tx, appState, arb[0].TokenIn, pairInfo, arb)
		if lpCycle == nil {
			log.Warn("RonFi backRun: uniswap.FromAddress fail", "loopId", arb.String(), "tx", tx.Hash().String(), "pair", pairInfo.Address)
			continue
		}

		if !lpCycle.AutoUpdate(v2States, v3States) {
			log.Info("RonFi backRun, lpCycle.AutoUpdate fail", "loopId", arb.String(), "tx", tx.Hash().String(), "pair", pairInfo.Address)
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

			txFeeInEth := rcommon.EthBigInt2Float64(new(big.Int).Mul(w.gasPrice, new(big.Int).SetUint64(profit.Cycle.SumGasNeed)))
			txFeeInUsd := txFeeInEth * defi.GetTradingTokenPrice(rcommon.USDC)
			txFeeInToken := price * txFeeInEth
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

		}
	}

	if len(profits) > 0 {
		sort.Sort(profits)
		highestProfit := profits[0]

		return highestProfit.uniProfit.Cycle, highestProfit.uniProfit.Profit.SwapAmount, true
	}

	return nil, nil, false
}

func (w *Worker) backRun(appState *state.StateDB, tx *types.Transaction, pairId int, pairsInfo []*defi.SwapPairInfo, handlerStartTime mclock.AbsTime) {
	v3States := make(map[common.Address]*v3.PoolState)
	v2States := make(map[common.Address]*v2.PoolState)

	// update v2/v3 states from target dexTx
	for _, info := range pairsInfo {
		if !info.V3 {
			v2States[info.Address] = &v2.PoolState{
				Reserve0: info.Reserve0,
				Reserve1: info.Reserve1,
			}
		} else {
			v3States[info.Address] = &v3.PoolState{
				Tick:         info.Tick,
				SqrtPriceX96: info.SqrtPriceX96,
				Liquidity:    info.Liquidity,
			}
		}
	}

	for i, info := range pairsInfo {
		edge := uniswap.ToV3Edge(info)
		if edge == nil {
			continue
		}

		arbs := w.v3LoopsDb.FindLoops(edge)
		if len(arbs) == 0 {
			log.Info("RonFi huntingTxEvent no matched loops!", "idx", i, "tx", tx.Hash().String(), "pair", info.Address)
			continue
		} else {
			log.Info("RonFi huntingTxEvent matched loops", "idx", i, "tx", tx.Hash().String(), "pair", info.Address, "loops", len(arbs))
		}
		profits := make(ProfitDetails, 0, len(arbs))
		for _, arb := range arbs {
			lpCycle := uniswap.FromAddress(w.di, tx, appState, arb[0].TokenIn, info, arb)
			if lpCycle == nil {
				log.Warn("RonFi huntingTxEvent: uniswap.FromAddress fail", "idx", i, "loopId", arb.String(), "tx", tx.Hash().String(), "pair", info.Address)
				continue
			}

			if !lpCycle.AutoUpdate(v2States, v3States) {
				log.Info("RonFi huntingV3TxEvent, lpCycle.AutoUpdate fail", "idx", i, "loopId", arb.String(), "tx", tx.Hash().String(), "pair", info.Address)
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

				txFeeInEth := rcommon.EthBigInt2Float64(new(big.Int).Mul(w.gasPrice, new(big.Int).SetUint64(profit.Cycle.SumGasNeed)))
				txFeeInUsd := txFeeInEth * defi.GetTradingTokenPrice(rcommon.USDC)
				txFeeInToken := price * txFeeInEth
				netProfitInUsd := grossProfitInUsd/2.0 - txFeeInUsd
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
			}
		}

		if len(profits) > 0 {
			sort.Sort(profits)
			highestProfit := profits[0]

			if highestProfit.netProfitInUsd > w.minHuntingProfit {
				log.Info("RonFi huntingTxEvent hunting",
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
	//profitMin := rcommon.ToWei(info.txFeeInToken, 18)
	randomExecutorId := bestProfit.Cycle.LoopId.Uint64() + tx.Hash().TailUint64()
	// run the ring to draw an executor
	ringId := randomExecutorId % w.totalExecutors
	var arbTx *types.Transaction

	if ok, arb := w.DexSwapHunting(w.executorPrivKey[ringId], &w.executorAddress[ringId], tx, swapPairInfo, gasMaxLimit, bestProfit, w.dryRun, handlerStartTime); !ok {
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
