package ronfi

import (
	"context"
	"encoding/json"
	"fmt"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	"github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/ethereum/go-ethereum/ronfi/defi"
	"github.com/ethereum/go-ethereum/ronfi/stats"
	"github.com/ethereum/go-ethereum/ronfi/uniswap"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-redis/redis"
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
	v3LoopsDb         *uniswap.V3Loops
	pairGasMap        map[string]uint64         // the gas required for a pair swap (key: pair+dir)
	dexPairsMap       map[common.Address]uint64 // collection of all dex pairs
	obsRouters        map[common.Address]uint64
	obsMethods        map[uint64]string

	client *ethclient.Client

	di    *defi.Info
	rdb   *redis.Client
	mysql *db.Mysql
}

// New Only called once when geth startup
func New(eth rcommon.Backend, chainConfig *params.ChainConfig) *RonArbiter {
	dbConf := rcommon.LoadDBConfig()
	if dbConf == nil {
		log.Error("RonFi please make sure you have a correct db_config.json")
		return nil
	}

	r := &RonArbiter{
		eth:         eth,
		chainConfig: chainConfig,
		startCh:     make(chan string),
		stopCh:      make(chan struct{}),
		running:     false,
		dryRun:      false,
	}

	r.mysql = db.NewMysql(dbConf.MysqlConf)
	r.rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", dbConf.RedisConf.RedisHost, dbConf.RedisConf.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

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

func (r *RonArbiter) NewObsRouter(router string, methodId uint32) {
	log.Info("RonFi NewObsRouter", "router", router, "methodID", methodId)
	if r.mysql != nil {
		jsonObs := &rcommon.JsonNewObs{
			Router:   router,
			MethodID: methodId,
		}
		res := r.mysql.InsertObsRouter(jsonObs)
		log.Info("RonFi NewObsRouter insert mysql done!", "router", router, "methodID", methodId, "res", res)
		if res > 0 && r.rdb != nil {
			jsonData, _ := json.Marshal(jsonObs)
			r.rdb.Publish(rcommon.RedisMsgNewObsRouter, jsonData)
			log.Info("RonFi NewObsRouter publish redis done!", "router", router, "methodID", methodId)
		}
	}
}

func (r *RonArbiter) ReloadLoops() {
	if rpc.StartTrading {
		log.Warn("RonFi arb ReloadLoops Reject on Trading")
	} else {
		r.pairGasMap = r.mysql.LoadPairGas()
		r.dexPairsMap = r.mysql.LoadDexPairs()
		r.obsRouters = r.mysql.LoadObsRouters()
		r.obsMethods = r.mysql.LoadObsMethods()
		r.eth.TxPool().SetObs(r.obsRouters, r.obsMethods)

		pairsInfo := make(defi.PairInfoMap)
		for address, info := range r.di.GetAllPairInfo() {
			if pf, ok := rcommon.ValidV2FactsMap[info.Factory]; ok {
				info.Fee = pf
				pairsInfo[address] = info
			}
		}

		poolsInfo := make(map[common.Address]*defi.PoolInfo)
		for address, info := range r.di.GetAllPoolInfo() {
			if _, ok := rcommon.ValidV3FactsMap[info.Factory]; ok {
				poolsInfo[address] = info
			}
		}

		tokensInfo := r.di.GetAllTokenInfo()
		r.v3LoopsDb = uniswap.NewV3Loops(r.di, pairsInfo, poolsInfo, tokensInfo, r.pairGasMap)

		log.Info("RonFi arb Reload Loops and White Pairs success")
	}
}

func (r *RonArbiter) Start(dryRun bool, minHuntingProfit int, totalArb int, thisIndex int, maxMatchedLoops int, maxCopyHighProfit int, doubleCopy bool, checkArbTx bool, skipElse bool, debug bool, highProfitMin int, secondChance bool, freeze bool, gamma int, isP2pHunting bool, p2pHuntingTopN int, v3Hunting bool) {
	log.Info("RonFi arb start",
		"dryRun", dryRun,
		"minHuntingProfit", minHuntingProfit,
	)

	r.dryRun = dryRun
	r.minHuntingProfit = float64(rcommon.MaxInt(minHuntingProfit, 10)) / 100 // Unit: cent. For example: 100 means $1.00, 1000 means $10.00. Minimum allowed value here is 10 means $0.10

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
	r.stats = stats.NewStats(r.eth, r.client, r.di, r.rdb, r.mysql, r.GetPairGas(), r.GetDexPairs(), r.obsRouters, r.obsMethods)
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
				r.di = defi.NewInfo(r.client, r.mysql)
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
					r.ReloadLoops()
				}
			}

			if r.oracleInitialized {
				oracle.Reset(OracleInterval)
			} else {
				oracle.Reset(OracleInitialInterval)
			}

		case <-r.startCh:
			rpc.StartTrading = true
			break

		case <-r.stopCh:
			rpc.StartTrading = false
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
