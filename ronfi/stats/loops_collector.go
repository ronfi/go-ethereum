package stats

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	db "github.com/ethereum/go-ethereum/ronfi/db"
	"github.com/go-redis/redis"
)

type LoopsCollector struct {
	rdb   *redis.Client
	mysql *db.Mysql

	newDBLoopCh chan *db.DBLoop
	stopCh      chan struct{}
}

func NewLoopsCollector(rdb *redis.Client, sql *db.Mysql) *LoopsCollector {
	newDBLoopCh := make(chan *db.DBLoop, 128)
	stopCh := make(chan struct{})

	loopsDis := &LoopsCollector{
		rdb,
		sql,
		newDBLoopCh,
		stopCh,
	}

	return loopsDis
}

func (ldp *LoopsCollector) start() {
	log.Info("RonFi LoopsCollector start")
	go ldp.run()
}

func (ldp *LoopsCollector) stop() {
	log.Info("RonFi LoopsCollector stop")
	close(ldp.stopCh)
}

func (ldp *LoopsCollector) run() {
	for {
		select {
		case dbLoop := <-ldp.newDBLoopCh:
			{
				record := dbLoop.ToLoopRecord()
				res := ldp.mysql.InsertLoop(record)
				if res > 0 {
					jsonLoop, _ := json.Marshal(record)
					ldp.rdb.Publish(rcommon.RedisMsgNewLoop, string(jsonLoop))
					log.Info("RonFi LoopsCollector new loops found!", "loopId", dbLoop.LoopId)
				} else {
					log.Info("RonFi LoopsCollector new loops found, but inserted db failed!", "loopId", dbLoop.LoopId)
				}
			}
		case <-ldp.stopCh:
			return
		}
	}
}

func (ldp *LoopsCollector) notify(dbLoop *db.DBLoop) {
	ldp.newDBLoopCh <- dbLoop
}
