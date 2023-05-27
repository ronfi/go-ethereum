package ronfi

import (
	"context"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/stats"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

const (
	ClientInitialInterval = 5 * time.Second
	OracleInitialInterval = 10 * time.Second
	OracleInterval        = 5 * time.Minute
)

type RonArbiter struct {
	eth               rcommon.Backend
	chainConfig       *params.ChainConfig
	startCh           chan string
	stopCh            chan struct{}
	stats             *stats.Stats
	oracleInitialized bool
	running           bool
	dryRun            bool
	minHuntingProfit  float64 // Trigger the hunting if profitInToken >= txFeeInToken*this
	totalArb          int     // How many arbitrate trading node I have totally
	thisArbIndex      int     // The index of this arbitrate among all my arbitrate nodes
	maxMatchedLoops   int     // MaxMatchedLoopsAllowed, originally we define it as a const with 96, now it's a parameter when start arb.
	maxCopyHighProfit int     // Max copy of arb tx when there's high profit for example >= $10.
	doubleCopy        bool    // Whether double high profit copy (>$25, 2 copy; >$100, 3 copy; >$500, 5 copy)
	checkArbTx        bool    // Whether cancel failed arbTx
	skipElse          bool    // Whether skip hunted else for <$10 profit
	debug             bool    // Whether log debug info
	highProfitMin     float64 // The threshold of a high profit chance which will be treated as VipHunting
	secondChance      bool    // Whether hunting 2nd chance
	freeze            bool    // Whether freeze all other activities when a VipHunting (big profit)
	isP2pHunting      bool
	p2pHuntingTopN    int     // Control hunting on the Top N learning loops, only meaningful for p2p hunting
	Gamma             float64 // Beta  = 0.998498873309329 // β = Sqrt( 1 - 0.3% ) = 0.998498873309329
	// Gamma = 0.001503383459709 // γ = ( 1-β ) / β
	pairGasMap         map[string]uint64         // the gas required for a pair swap (key: pair+dir)
	feePatchMap        map[common.Address]uint64 // the patch for pool fee and/or token fee
	flashNokPairs      map[common.Address]uint64
	dexPairsMap        map[common.Address]uint64 // collection of all dex pairs
	oskLimitedPairsMap map[common.Address]struct{}
	obsRouters         map[common.Address]uint64
	obsMethods         map[uint64]string

	client *ethclient.Client

	di *defi.Info
}

// New Only called once when geth startup
func New(eth rcommon.Backend, chainConfig *params.ChainConfig) *RonArbiter {
	r := &RonArbiter{
		eth:             eth,
		chainConfig:     chainConfig,
		startCh:         make(chan string),
		stopCh:          make(chan struct{}),
		running:         false,
		dryRun:          false,
		totalArb:        1,
		maxMatchedLoops: 48,
	}

	go r.mainLoop()
	return r
}

func (r *RonArbiter) GetPairGas() map[string]uint64 {
	return r.pairGasMap
}

func (r *RonArbiter) GetDexPairs() map[common.Address]uint64 {
	return r.dexPairsMap
}

func (r *RonArbiter) ResetStats() {
	if r.stats != nil {
		r.stats.ResetStats()
	}
}

func (r *RonArbiter) Start(dryRun bool, minHuntingProfit int, totalArb int, thisIndex int, maxMatchedLoops int, maxCopyHighProfit int, doubleCopy bool, checkArbTx bool, skipElse bool, debug bool, highProfitMin int, secondChance bool, freeze bool, gamma int, isP2pHunting bool, p2pHuntingTopN int, v3Hunting bool) {
	log.Info("RonFi arb start",
		"dryRun", dryRun,
		"minHuntingProfit", minHuntingProfit,
		"totalArb", totalArb,
		"thisIndex", thisIndex,
		"maxMatchedLoops", maxMatchedLoops,
		"maxCopyHighProfit", maxCopyHighProfit,
		"doubleCopy", doubleCopy,
		"checkArbTx", checkArbTx,
		"skipElse", skipElse,
		"debug", debug,
		"highProfitMin", highProfitMin,
		"secondChance", secondChance,
		"freeze", freeze,
		"Gamma", gamma,
		"isP2pHunting", isP2pHunting,
		"p2pHuntingTopN", p2pHuntingTopN,
		"v3Hunting", v3Hunting,
	)

	r.dryRun = dryRun
	r.minHuntingProfit = float64(rcommon.MaxInt(minHuntingProfit, 10)) / 100 // Unit: cent. For example: 100 means $1.00, 1000 means $10.00. Minimum allowed value here is 10 means $0.10
	r.totalArb = rcommon.MaxInt(totalArb, 1)
	r.thisArbIndex = rcommon.MinInt(thisIndex, r.totalArb+1)
	r.maxMatchedLoops = rcommon.MaxInt(maxMatchedLoops, rcommon.MaxMatchedLoopsAllowed) // don't allow setting below 2
	r.maxCopyHighProfit = rcommon.MinInt(rcommon.MaxInt(maxCopyHighProfit, 1), 8)       // must within [1, 8]
	r.doubleCopy = doubleCopy
	r.checkArbTx = checkArbTx
	r.skipElse = skipElse
	r.debug = debug
	r.highProfitMin = float64(rcommon.MaxInt(highProfitMin, 10)) / 100 // Unit: cent. must >= $0.1
	r.secondChance = secondChance
	r.freeze = freeze
	r.Gamma = float64(rcommon.MaxInt(gamma, 15)) / 10000 // Unit: 1/10000. must >= 0.0015
	// Beta  = 0.998498873309329 // β = Sqrt( 1 - 0.3% ) = 0.998498873309329
	// Gamma = 0.001503383459709 // γ = ( 1-β ) / β
	r.isP2pHunting = isP2pHunting
	r.p2pHuntingTopN = rcommon.MinInt(rcommon.MaxInt(p2pHuntingTopN, 1), 10) // must within [1, 10]

	r.startCh <- "start"
}

func (r *RonArbiter) Stop() {
	log.Info("RonFi arb Stop")
	if r.running {
		r.stopCh <- struct{}{}
	}
}

func (r *RonArbiter) Arbing() bool {
	return r.running
}

func (r *RonArbiter) StartStats() {
	if r.stats != nil {
		log.Warn("RonFi stats service already started")
		return
	}
	// start stats service
	r.stats = stats.NewStats(r.eth, r.client, r.di, r.GetPairGas(), r.GetDexPairs(), r.obsRouters, r.obsMethods)
	if r.stats == nil {
		log.Warn("RonFi stats service started failed")
	} else {
		log.Info("RonFi stats service started")
		go r.stats.Run()
	}
}

func (r *RonArbiter) StopStats() {
	if r.stats != nil {
		log.Info("RonFi stats Stop")
		r.stats.StopStats()
		r.stats = nil
	} else {
		log.Warn("RonFi stats not started")
	}
}

func (r *RonArbiter) mainLoop() {
	oracle := time.NewTimer(OracleInitialInterval)
	clientInit := time.NewTimer(ClientInitialInterval)

	for {
		select {
		case <-clientInit.C:
			client, err := ethclient.Dial(rcommon.CLIENT_URL)
			if err != nil {
				log.Error("RonFi mainLoop: create ethClient failed!", "reason", err)
				clientInit.Reset(ClientInitialInterval) // retry in 5 seconds
			} else {
				r.client = client
				r.di = defi.NewInfo(r.client)
			}

		case <-oracle.C:
			if r.oracleInitialized {
				//if err := gopool.Submit(func() {
				//	// update token price in every 5 minutes
				//	defi.UpdateTokenPrices()
				//}); err != nil {
				//	log.Error("RonFi mainLoop: gopool.Submit() failed", "reason", err)
				//	continue
				//}
			} else {
				// price oracle initialization
				if r.client != nil && !defi.InitRonFiOracle(r.client) {
					log.Error("RonFi mainLoop: InitRonFiOracle() failed")
				} else {
					r.oracleInitialized = true
				}
			}

			if r.oracleInitialized {
				oracle.Reset(OracleInterval)
			} else {
				oracle.Reset(OracleInitialInterval)
			}

		case <-r.startCh:
			break

		case <-r.stopCh:
			break
		}
	}
}

func (r *RonArbiter) GetTransaction(hash string) *types.Transaction {
	// connect speedy client
	client, err := ethclient.Dial("https://nd-814-711-835.p2pify.com/049f42c7290c310495b6940701e2ae14")
	if err != nil {
		log.Info("RonFi GetTransaction connect rpc provider failed", "tx", hash)
		return nil
	}

	defer func() {
		if client != nil {
			client.Close()
		}
	}()

	b, err := hexutil.Decode(hash)
	if err != nil {
		log.Info("RonFi GetTransaction decode string to bytes failed!", "tx", hash)
		return nil
	}

	tx, _, err := client.TransactionByHash(context.Background(), common.BytesToHash(b))
	if err != nil {
		log.Info("RonFi GetTransaction retrieve transaction failed!", "tx", hash)
		return nil
	}

	return tx
}
