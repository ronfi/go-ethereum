package stats

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	db "github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/go-redis/redis"
)

type ObsRecord struct {
	tx    *types.Transaction
	obsId ObsId
	loops *ObsParsedResult
}

type ObsCollector struct {
	rdb   *redis.Client
	mysql *db.Mysql

	newObsCh    chan *rcommon.NewObs
	obsRecordCh chan *ObsRecord
	stopCh      chan struct{}
}

func NewObsCollector(rdb *redis.Client, sql *db.Mysql) *ObsCollector {
	newObsCh := make(chan *rcommon.NewObs, 128)
	obsRecordCh := make(chan *ObsRecord, 4096)
	stopCh := make(chan struct{})

	knCol := &ObsCollector{
		rdb,
		sql,
		newObsCh,
		obsRecordCh,
		stopCh,
	}

	return knCol
}

func (oc *ObsCollector) start() {
	log.Info("RonFi NewObsCollector start")
	go oc.run()
}

func (oc *ObsCollector) stop() {
	log.Info("RonFi NewObsCollector stop")
	close(oc.stopCh)
}

func (oc *ObsCollector) run() {
	for {
		select {
		case newObs := <-oc.newObsCh:
			{
				record := newObs.ToJsonNewObs()
				res := oc.mysql.InsertObsRouter(record)
				if res > 0 {
					jsonData, _ := json.Marshal(record)
					oc.rdb.Publish(rcommon.RedisMsgNewObsRouter, string(jsonData))
				}
			}

		case obsRecord := <-oc.obsRecordCh:
			{
				loops, _ := json.Marshal(obsRecord.loops)
				oc.mysql.InsertObsAll(&db.LoopObsRecord{
					Tx:    obsRecord.tx.Hash().String(),
					ObsId: string(obsRecord.obsId),
					Loops: string(loops),
				})
			}

		case <-oc.stopCh:
			return
		}
	}
}

func (oc *ObsCollector) notifyObs(newObs *rcommon.NewObs) {
	oc.newObsCh <- newObs
}

func (oc *ObsCollector) notifyObsRecord(obsRecord *ObsRecord) {
	oc.obsRecordCh <- obsRecord
}
