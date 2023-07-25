package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/mclock"
	"github.com/ethereum/go-ethereum/log"
	rcommon "github.com/ethereum/go-ethereum/ronfi/common"
	_ "github.com/go-sql-driver/mysql"
	"sort"
	"strconv"
)

type LoopRecord struct {
	LoopsId   string `json:"loopsId"`
	Path      string `json:"path"`
	PoolFee   string `json:"poolFee"`
	TokenFee  string `json:"tokenFee"`
	Direction string `json:"direction"`
	Index     string `json:"index"`
	Counts    uint64 `json:"counts"`
	Cancel    bool   `json:"cancel"`
}

func (record *LoopRecord) ToDBLoop() *DBLoop {
	var (
		err       error
		key       string
		loopId    common.Hash
		pathArr   []string
		path      []common.Address
		poolFee   []uint64
		tokenFee  []uint64
		direction []uint64
		index     []uint64
		count     uint64
	)

	loopId = common.HexToHash(record.LoopsId)
	if err = json.Unmarshal([]byte(record.Path), &pathArr); err != nil {
		return nil
	}
	for _, addr := range pathArr {
		path = append(path, common.HexToAddress(addr))
	}

	if err = json.Unmarshal([]byte(record.PoolFee), &poolFee); err != nil {
		return nil
	}

	if err = json.Unmarshal([]byte(record.TokenFee), &tokenFee); err != nil {
		return nil
	}

	if err = json.Unmarshal([]byte(record.Direction), &direction); err != nil {
		return nil
	}

	if err = json.Unmarshal([]byte(record.Index), &index); err != nil {
		return nil
	}

	count = record.Counts

	return &DBLoop{
		Key:       key,
		LoopId:    loopId,
		Path:      path,
		PoolFee:   poolFee,
		TokenFee:  tokenFee,
		Direction: direction,
		Index:     index,
		Count:     count,
	}
}

type DBLoop struct {
	Key       string
	LoopId    common.Hash
	Path      []common.Address
	PoolFee   []uint64
	TokenFee  []uint64
	Direction []uint64
	Index     []uint64
	Count     uint64
	Cancel    bool
	HasV3     bool
}

func (loop *DBLoop) ToLoopRecord() *LoopRecord {
	path := "["
	for i, addr := range loop.Path {
		if i != len(loop.Path)-1 {
			path += fmt.Sprintf("\"%s\", ", addr.String())
		} else {
			path += fmt.Sprintf("\"%s\"", addr.String())
		}
	}
	path += "]"

	poolFee := "["
	for i, pf := range loop.PoolFee {
		if i != len(loop.PoolFee)-1 {
			poolFee += fmt.Sprintf("%d, ", pf)
		} else {
			poolFee += fmt.Sprintf("%d", pf)
		}
	}
	poolFee += "]"

	tokenFee := "["
	for i, tf := range loop.TokenFee {
		if i != len(loop.TokenFee)-1 {
			tokenFee += fmt.Sprintf("%d, ", tf)
		} else {
			tokenFee += fmt.Sprintf("%d", tf)
		}
	}
	tokenFee += "]"

	direction := "["
	for i, dir := range loop.Direction {
		if i != len(loop.Direction)-1 {
			direction += fmt.Sprintf("%d, ", dir)
		} else {
			direction += fmt.Sprintf("%d", dir)
		}
	}
	direction += "]"

	index := "["
	for i, val := range loop.Index {
		if i != len(loop.Index)-1 {
			index += fmt.Sprintf("%d, ", val)
		} else {
			index += fmt.Sprintf("%d", val)
		}
	}
	index += "]"

	return &LoopRecord{
		loop.LoopId.String(),
		path,
		poolFee,
		tokenFee,
		direction,
		index,
		loop.Count,
		loop.Cancel,
	}
}

func (loop *DBLoop) Equals(dbLoop *DBLoop) bool {
	if len(loop.Path) != len(dbLoop.Path) {
		return false
	} else {
		for i := 0; i < len(loop.Path); i++ {
			if loop.Path[i] != dbLoop.Path[i] {
				return false
			}
		}
	}

	if len(loop.PoolFee) != len(dbLoop.PoolFee) {
		return false
	} else {
		for i := 0; i < len(loop.PoolFee); i++ {
			if loop.PoolFee[i] != dbLoop.PoolFee[i] {
				return false
			}
		}
	}

	if len(loop.TokenFee) != len(dbLoop.TokenFee) {
		return false
	} else {
		for i := 0; i < len(loop.TokenFee); i++ {
			if loop.TokenFee[i] != dbLoop.TokenFee[i] {
				return false
			}
		}
	}

	if len(loop.Direction) != len(dbLoop.Direction) {
		return false
	} else {
		for i := 0; i < len(loop.Direction); i++ {
			if loop.Direction[i] != dbLoop.Direction[i] {
				return false
			}
		}
	}

	if len(loop.Index) != len(dbLoop.Index) {
		return false
	} else {
		for i := 0; i < len(loop.Index); i++ {
			if loop.Index[i] != dbLoop.Index[i] {
				return false
			}
		}
	}

	return true
}

type Mysql struct {
	db *sql.DB
}

func NewMysql(conf rcommon.MysqlConfig) *Mysql {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.DbUser, conf.DbPass, conf.DbHost, conf.DbPort, conf.DbData)
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		log.Error("RonFi Mysql create db connection to(%s) failed!", conf.DbHost)
		return nil
	}

	db.SetMaxOpenConns(100)

	return &Mysql{
		db,
	}
}

func (sql *Mysql) Close() {
	if sql.db != nil {
		if err := sql.db.Close(); err != nil {
			log.Warn("RonFi Mysql close failed!", "err", err)
		}
	}
}

type PairGasRecord struct {
	pairDir string
	gas     uint64
}

func (sql *Mysql) LoadPairGas() map[string]uint64 {
	startTime := mclock.Now()

	pairDirGasMap := make(map[string]uint64)
	records := sql.fetchPairGasInParallel()
	for _, record := range records {
		pairDirGasMap[record.pairDir] = record.gas
	}

	log.Info("RonFi Mysql LoadPairGas", "total pairs", len(pairDirGasMap), "elapsed", mclock.Since(startTime))

	return pairDirGasMap
}

func (sql *Mysql) fetchPairGasInParallel() []PairGasRecord {
	ids := make([]int, 0)
	if rows, err := sql.db.Query("select id from pair_dir_gas;"); err == nil {
		for rows.Next() {
			var id int
			if err = rows.Scan(&id); err == nil {
				ids = append(ids, id)
			}
		}
	}
	sort.Ints(ids)
	count := len(ids)
	records := make([]PairGasRecord, 0, count)
	bucketSize := 1000
	resultCount := 0
	resultChannel := make(chan []PairGasRecord, 0)
	groups := count / bucketSize
	if count%bucketSize != 0 {
		groups += 1
	}
	for i := 0; i < groups; i++ {
		begin := i * bucketSize
		end := (i + 1) * bucketSize
		if end >= count {
			end = count - 1
		}
		beginID := ids[begin]
		endId := ids[end]
		go func(beginId int, endId int) {
			querySQL := fmt.Sprintf("select pairDir, gas from pair_dir_gas where id between %d and %d;", beginId, endId)
			rows, err := sql.db.Query(querySQL)
			if err != nil {
				log.Warn("RonFi Mysql LoadPairGas query data failed", "querySQL", querySQL, "err", err)
				return
			}
			defer func() {
				_ = rows.Close()
			}()

			currentRecords := make([]PairGasRecord, 0, bucketSize)
			for rows.Next() {
				var (
					pairDir string
					gas     uint64
				)

				if err = rows.Scan(&pairDir, &gas); err == nil {
					currentRecords = append(currentRecords, PairGasRecord{pairDir, gas})
				}
			}

			resultChannel <- currentRecords
		}(beginID, endId)
		resultCount += 1
	}
	for i := 0; i < resultCount; i++ {
		currentRecords := <-resultChannel
		records = append(records, currentRecords...)
	}

	return records
}

func (sql *Mysql) UpdatePairGas(pairDirGasMap map[string]uint64) int64 {
	startTime := mclock.Now()

	updateSQL := fmt.Sprintf("insert into pair_dir_gas (pairDir, gas) values ")
	index := 0
	length := len(pairDirGasMap)
	for key, value := range pairDirGasMap {
		if index == length-1 {
			updateSQL = fmt.Sprintf("%s (\"%s\", %d) ", updateSQL, key, value)
		} else {
			updateSQL = fmt.Sprintf("%s (\"%s\", %d), ", updateSQL, key, value)
		}
		index++
	}
	updateSQL = fmt.Sprintf("%s as new on duplicate key update gas=new.gas", updateSQL)

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			log.Info("RonFi Mysql UpdatePairGas", "total pairs", len(pairDirGasMap), "elapsed", mclock.Since(startTime))
			return lastId
		}
	}

	log.Warn("RonFi Mysql UpdatePairGas failed", "err", err)
	return -1
}

func (sql *Mysql) LoadDexPairs() map[common.Address]uint64 {
	querySQL := fmt.Sprintf("select pair, frequency from dex_pairs;")
	rows, err := sql.db.Query(querySQL)
	if err != nil {
		log.Warn("RonFi Mysql LoadDexPairs query data failed", "querySQL", querySQL, "err", err)
		return nil
	}
	defer func() {
		_ = rows.Close()
	}()

	dexPairsMap := make(map[common.Address]uint64)
	for rows.Next() {
		var (
			pair      string
			frequency uint64
		)

		if err = rows.Scan(&pair, &frequency); err == nil {
			dexPairsMap[common.HexToAddress(pair)] = frequency
		}
	}

	return dexPairsMap
}

func (sql *Mysql) UpdateDexPairs(dexPairs map[common.Address]uint64) int64 {
	startTime := mclock.Now()

	updateSQL := fmt.Sprintf("insert into dex_pairs (pair, frequency) values ")
	index := 0
	length := len(dexPairs)
	for pair, frequency := range dexPairs {
		if index == length-1 {
			updateSQL = fmt.Sprintf("%s (\"%s\", %d) ", updateSQL, pair.String(), frequency)
		} else {
			updateSQL = fmt.Sprintf("%s (\"%s\", %d), ", updateSQL, pair.String(), frequency)
		}
		index++
	}
	updateSQL = fmt.Sprintf("%s as new on duplicate key update frequency=new.frequency", updateSQL)

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			log.Info("RonFi Mysql UpdateDexPairs", "total pairs", len(dexPairs), "elapsed", mclock.Since(startTime))
			return lastId
		}
	}

	log.Warn("RonFi Mysql UpdateDexPairs failed", "err", err)
	return -1
}

func (sql *Mysql) LoadObsRouters() map[common.Address]uint64 {
	querySQL := fmt.Sprintf("select router, methodID from obs_routers;")
	rows, err := sql.db.Query(querySQL)
	if err != nil {
		log.Warn("RonFi Mysql LoadObsRouters query data failed", "querySQL", querySQL, "err", err)
		return nil
	}
	defer func() {
		_ = rows.Close()
	}()

	obsRoutersMap := make(map[common.Address]uint64)
	for rows.Next() {
		var (
			router   string
			methodID uint64
		)

		if err = rows.Scan(&router, &methodID); err == nil {
			obsRoutersMap[common.HexToAddress(router)] = methodID
		}
	}

	return obsRoutersMap
}

func (sql *Mysql) UpdateObsRouters(obsRouters map[common.Address]uint64) int64 {
	startTime := mclock.Now()

	updateSQL := fmt.Sprintf("insert into obs_routers (router, methodID) values ")
	index := 0
	length := len(obsRouters)
	for router, methodID := range obsRouters {
		if index == length-1 {
			updateSQL = fmt.Sprintf("%s (\"%s\", %d) ", updateSQL, router.String(), methodID)
		} else {
			updateSQL = fmt.Sprintf("%s (\"%s\", %d), ", updateSQL, router.String(), methodID)
		}
		index++
	}
	updateSQL = fmt.Sprintf("%s as new on duplicate key update methodID=new.methodID", updateSQL)

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			log.Info("RonFi Mysql UpdateObsRouters", "total routers", len(obsRouters), "elapsed", mclock.Since(startTime))
			return lastId
		}
	}

	log.Warn("RonFi Mysql UpdateObsRouters failed", "err", err)
	return -1
}

func (sql *Mysql) InsertObsRouter(record *rcommon.JsonNewObs) int64 {
	updateSQL := fmt.Sprintf("insert ignore into obs_routers (router, methodID) values (\"%s\", %d);",
		record.Router, record.MethodID)
	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			return lastId
		}
	}

	log.Warn("RonFi Mysql InsertObsRouter failed", "err", err)
	return -1
}

func (sql *Mysql) LoadObsMethods() map[uint64]string {
	querySQL := fmt.Sprintf("select methodID, obsInfo from obs_methods;")
	rows, err := sql.db.Query(querySQL)
	if err != nil {
		log.Warn("RonFi Mysql LoadObsMethods query data failed", "querySQL", querySQL, "err", err)
		return nil
	}
	defer func() {
		_ = rows.Close()
	}()

	obsMethodsMap := make(map[uint64]string)
	for rows.Next() {
		var (
			methodIDStr string
			obsInfo     string
		)

		if err = rows.Scan(&methodIDStr, &obsInfo); err == nil {
			methodID, _ := strconv.ParseUint(methodIDStr, 0, 64)
			obsMethodsMap[methodID] = obsInfo
		}
	}

	return obsMethodsMap
}

func (sql *Mysql) UpdateObsMethods(obsMethods map[uint64]string) int64 {
	startTime := mclock.Now()

	updateSQL := fmt.Sprintf("insert into obs_methods (methodID, obsInfo) values ")
	index := 0
	length := len(obsMethods)
	for methodID, obsInfo := range obsMethods {
		if index == length-1 {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\") ", updateSQL, fmt.Sprintf("0x%08x", methodID), obsInfo)
		} else {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\"), ", updateSQL, fmt.Sprintf("0x%08x", methodID), obsInfo)
		}
		index++
	}
	updateSQL = fmt.Sprintf("%s as new on duplicate key update obsInfo=new.obsInfo", updateSQL)

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			log.Info("RonFi Mysql UpdateObsMethods", "total methods", len(obsMethods), "elapsed", mclock.Since(startTime))
			return lastId
		}
	}

	log.Warn("RonFi Mysql UpdateObsMethods failed", "err", err)
	return -1
}

func (sql *Mysql) LoadLoops() []*DBLoop {
	records := sql.fetchLoopsInParallel()
	allLoops := make([]*DBLoop, 0, len(records))

	for _, record := range records {
		var (
			err       error
			pathArr   []string
			path      []common.Address
			poolFee   []uint64
			tokenFee  []uint64
			direction []uint64
			index     []uint64
		)

		if err = json.Unmarshal([]byte(record.Path), &pathArr); err != nil {
			log.Info("RonFi Mysql LoadLoops Unmarshal pathStr failed", "loopId", record.LoopsId, "err", err)
			continue
		}
		for _, addr := range pathArr {
			path = append(path, common.HexToAddress(addr))
		}

		if err = json.Unmarshal([]byte(record.PoolFee), &poolFee); err != nil {
			log.Info("RonFi Mysql LoadLoops Unmarshal poolFeeStr failed", "loopId", record.LoopsId, "err", err)
			continue
		}

		if err = json.Unmarshal([]byte(record.TokenFee), &tokenFee); err != nil {
			log.Info("RonFi Mysql LoadLoops Unmarshal tokenFeeStr failed", "loopId", record.LoopsId, "err", err)
			continue
		}

		if err = json.Unmarshal([]byte(record.Direction), &direction); err != nil {
			log.Info("RonFi Mysql LoadLoops Unmarshal directionStr failed", "loopId", record.LoopsId, "err", err)
			continue
		}

		if err = json.Unmarshal([]byte(record.Index), &index); err != nil {
			log.Info("RonFi Mysql LoadLoops Unmarshal indexStr failed", "loopId", record.LoopsId, "err", err)
			continue
		}

		allLoops = append(allLoops, &DBLoop{
			LoopId:    common.HexToHash(record.LoopsId),
			Path:      path,
			PoolFee:   poolFee,
			TokenFee:  tokenFee,
			Direction: direction,
			Index:     index,
			Count:     record.Counts,
			Cancel:    record.Cancel,
		})
	}

	return allLoops
}

func (sql *Mysql) fetchLoopsInParallel() []*LoopRecord {
	ids := make([]int, 0)
	querySQL := fmt.Sprintf("select id from loops;")
	if rows, err := sql.db.Query(querySQL); err == nil {
		for rows.Next() {
			var id int
			if err = rows.Scan(&id); err == nil {
				ids = append(ids, id)
			}
		}
	}
	sort.Ints(ids)
	count := len(ids)
	records := make([]*LoopRecord, 0, count)
	bucketSize := 1000
	resultCount := 0
	resultChannel := make(chan []*LoopRecord, 0)
	groups := count / bucketSize
	if count%bucketSize != 0 {
		groups += 1
	}
	for i := 0; i < groups; i++ {
		begin := i * bucketSize
		end := (i + 1) * bucketSize
		if end >= count {
			end = count - 1
		}
		beginID := ids[begin]
		endId := ids[end]
		go func(beginId int, endId int) {
			querySQL := fmt.Sprintf("select loopsId, path, poolFee, tokenFee, direction, indexes, counts, canceled from loops where id between %d and %d;", beginId, endId)
			rows, err := sql.db.Query(querySQL)
			if err != nil {
				log.Warn("RonFi Mysql fetchLoopsInParallel query data failed", "querySQL", querySQL, "err", err)
				return
			}
			defer func() {
				_ = rows.Close()
			}()

			currentRecords := make([]*LoopRecord, 0, bucketSize)
			for rows.Next() {
				var (
					loopsId   string
					path      string
					poolFee   string
					tokenFee  string
					direction string
					index     string
					counts    uint64
					cancel    uint64
					hasV3     uint64
				)

				if err = rows.Scan(&loopsId, &path, &poolFee, &tokenFee, &direction, &index, &counts, &cancel, &hasV3); err == nil {
					currentRecords = append(
						currentRecords,
						&LoopRecord{
							loopsId,
							path,
							poolFee,
							tokenFee,
							direction,
							index,
							counts,
							cancel == 1,
						})
				}
			}

			resultChannel <- currentRecords
		}(beginID, endId)
		resultCount += 1
	}
	for i := 0; i < resultCount; i++ {
		currentRecords := <-resultChannel
		records = append(records, currentRecords...)
	}

	return records
}

func (sql *Mysql) InsertLoop(loopRecord *LoopRecord) int64 {
	updateSQL := fmt.Sprintf("insert ignore into loops (loopsId, path, poolFee, tokenFee, direction, indexes, counts, canceled) values ")
	updateSQL = fmt.Sprintf("%s ('%s', '%s', '%s', '%s', '%s', '%s', %d, %d)",
		updateSQL,
		loopRecord.LoopsId,
		loopRecord.Path,
		loopRecord.PoolFee,
		loopRecord.TokenFee,
		loopRecord.Direction,
		loopRecord.Index,
		loopRecord.Counts,
		0)

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			return lastId
		}
	}

	log.Warn("RonFi Mysql InsertLoop failed", "err", err)
	return -1
}

func (sql *Mysql) CancelLoop(loopsId common.Hash) int64 {
	updateSQL := fmt.Sprintf("update loops set canceled=1 where loopsId='%s'", loopsId.String())
	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			return lastId
		}
	}

	log.Warn("RonFi Mysql CancelLoop failed", "err", err)
	return -1
}

func (sql *Mysql) RestoreLoop(loopsId common.Hash) int64 {
	updateSQL := fmt.Sprintf("update loops set canceled=0 where loopsId='%s'", loopsId.String())
	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			return lastId
		}
	}

	log.Warn("RonFi Mysql RestoreLoop failed", "err", err)
	return -1
}

func (sql *Mysql) LoadLoopById(loopsId common.Hash) *LoopRecord {
	querySQL := fmt.Sprintf("select loopsId, path, poolFee, tokenFee, direction, indexes, counts, canceled from loops where loopsId='%s';", loopsId.String())
	rows, err := sql.db.Query(querySQL)
	if err != nil {
		log.Warn("RonFi Mysql LoadLoopById query data failed", "querySQL", querySQL, "err", err)
		return nil
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var (
			loopsId   string
			path      string
			poolFee   string
			tokenFee  string
			direction string
			index     string
			counts    uint64
			cancel    uint64
		)

		if err = rows.Scan(&loopsId, &path, &poolFee, &tokenFee, &direction, &index, &counts, &cancel); err == nil {
			return &LoopRecord{
				loopsId,
				path,
				poolFee,
				tokenFee,
				direction,
				index,
				counts,
				cancel == 1,
			}
		}
	}

	return nil
}

type PairInfoRecord struct {
	Pair         string
	Name         string
	Index        uint64
	BothBriToken bool
	KeyToken     string
	Token0       string
	Token1       string
	Factory      string
}

func (sql *Mysql) LoadPairsInfo() []*PairInfoRecord {
	ids := make([]int, 0)
	if rows, err := sql.db.Query("select id from pairs;"); err == nil {
		for rows.Next() {
			var id int
			if err = rows.Scan(&id); err == nil {
				ids = append(ids, id)
			}
		}
	}
	sort.Ints(ids)
	count := len(ids)
	records := make([]*PairInfoRecord, 0, count)
	bucketSize := 1000
	resultCount := 0
	resultChannel := make(chan []*PairInfoRecord, 0)
	groups := count / bucketSize
	if count%bucketSize != 0 {
		groups += 1
	}
	for i := 0; i < groups; i++ {
		begin := i * bucketSize
		end := (i + 1) * bucketSize
		if end >= count {
			end = count - 1
		}
		beginID := ids[begin]
		endId := ids[end]

		go func(beginId int, endId int) {
			querySQL := fmt.Sprintf("select pair, name, pairIndex, bothBriToken, keyToken, token0, token1, factory from pairs where id between %d and %d;", beginId, endId)
			rows, err := sql.db.Query(querySQL)
			if err != nil {
				log.Warn("RonFi Mysql LoadPairsInfo query data failed", "querySQL", querySQL, "err", err)
				return
			}
			defer func() {
				_ = rows.Close()
			}()

			currentRecords := make([]*PairInfoRecord, 0, bucketSize)
			for rows.Next() {
				var (
					pair         string
					name         string
					index        int
					bothBriToken bool
					keyToken     string
					token0       string
					token1       string
					factory      string
				)

				if err = rows.Scan(&pair, &name, &index, &bothBriToken, &keyToken, &token0, &token1, &factory); err == nil {
					if pair == "" {
						continue
					}
					currentRecords = append(
						currentRecords,
						&PairInfoRecord{
							Pair:         pair,
							Name:         name,
							Index:        uint64(index),
							BothBriToken: bothBriToken,
							KeyToken:     keyToken,
							Token0:       token0,
							Token1:       token1,
							Factory:      factory,
						})
				}
			}

			resultChannel <- currentRecords
		}(beginID, endId)
		resultCount += 1
	}
	for i := 0; i < resultCount; i++ {
		currentRecords := <-resultChannel
		records = append(records, currentRecords...)
	}

	return records
}

func (sql *Mysql) InsertPairsInfo(newPairsInfo []*PairInfoRecord) int64 {
	updateSQL := fmt.Sprintf("insert ignore into pairs (pair, name, pairIndex, bothBriToken, keyToken, token0, token1, factory) values ")

	length := len(newPairsInfo)
	for index, info := range newPairsInfo {
		bothBriToken := 0
		if info.BothBriToken {
			bothBriToken = 1
		}
		if index == length-1 {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\", %d, %d, \"%s\", \"%s\", \"%s\", \"%s\") ",
				updateSQL,
				info.Pair,
				info.Name,
				info.Index,
				bothBriToken,
				info.KeyToken,
				info.Token0,
				info.Token1,
				info.Factory)
		} else {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\", %d, %d, \"%s\", \"%s\", \"%s\", \"%s\"), ",
				updateSQL,
				info.Pair,
				info.Name,
				info.Index,
				bothBriToken,
				info.KeyToken,
				info.Token0,
				info.Token1,
				info.Factory)
		}
		index++
	}

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			//log.Info("RonFi Mysql InsertPairsInfo done!", "num", length)
			return lastId
		}
	}

	log.Warn("RonFi Mysql InsertPairsInfo failed", "err", err)
	return -1
}

type PoolInfoRecord struct {
	Pool        string
	Name        string
	Token0      string
	Token1      string
	Fee         int
	TickSpacing int
	Factory     string
}

func (sql *Mysql) LoadPoolsInfo() []*PoolInfoRecord {
	ids := make([]int, 0)
	if rows, err := sql.db.Query("select id from pools;"); err == nil {
		for rows.Next() {
			var id int
			if err = rows.Scan(&id); err == nil {
				ids = append(ids, id)
			}
		}
	}
	sort.Ints(ids)
	count := len(ids)
	records := make([]*PoolInfoRecord, 0, count)
	bucketSize := 1000
	resultCount := 0
	resultChannel := make(chan []*PoolInfoRecord, 0)
	groups := count / bucketSize
	if count%bucketSize != 0 {
		groups += 1
	}
	for i := 0; i < groups; i++ {
		begin := i * bucketSize
		end := (i + 1) * bucketSize
		if end >= count {
			end = count - 1
		}
		beginID := ids[begin]
		endId := ids[end]

		go func(beginId int, endId int) {
			querySQL := fmt.Sprintf("select pool, name, token0, token1, fee, tickSpacing, factory from pools where id between %d and %d;", beginId, endId)
			rows, err := sql.db.Query(querySQL)
			if err != nil {
				log.Warn("RonFi Mysql LoadPoolsInfo query data failed", "querySQL", querySQL, "err", err)
				return
			}
			defer func() {
				_ = rows.Close()
			}()

			currentRecords := make([]*PoolInfoRecord, 0, bucketSize)
			for rows.Next() {
				var (
					pool        string
					name        string
					token0      string
					token1      string
					fee         int
					tickSpacing int
					factory     string
				)

				if err = rows.Scan(&pool, &name, &token0, &token1, &fee, &tickSpacing, &factory); err == nil {
					currentRecords = append(
						currentRecords,
						&PoolInfoRecord{
							Pool:        pool,
							Name:        name,
							Token0:      token0,
							Token1:      token1,
							Fee:         fee,
							TickSpacing: tickSpacing,
							Factory:     factory,
						})
				}
			}

			resultChannel <- currentRecords
		}(beginID, endId)
		resultCount += 1
	}
	for i := 0; i < resultCount; i++ {
		currentRecords := <-resultChannel
		records = append(records, currentRecords...)
	}

	return records
}

func (sql *Mysql) InsertPoolsInfo(newPoolsInfo []*PoolInfoRecord) int64 {
	updateSQL := fmt.Sprintf("insert ignore into pools (pool, name, token0, token1, fee, tickSpacing, factory) values ")

	length := len(newPoolsInfo)
	for index, info := range newPoolsInfo {
		if index == length-1 {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\", \"%s\", \"%s\", %d, %d, \"%s\") ",
				updateSQL,
				info.Pool,
				info.Name,
				info.Token0,
				info.Token1,
				info.Fee,
				info.TickSpacing,
				info.Factory)
		} else {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\", \"%s\", \"%s\", %d, %d, \"%s\"), ",
				updateSQL,
				info.Pool,
				info.Name,
				info.Token0,
				info.Token1,
				info.Fee,
				info.TickSpacing,
				info.Factory)
		}
		index++
	}

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			//log.Info("RonFi Mysql InsertPairsInfo done!", "num", length)
			return lastId
		}
	}

	log.Warn("RonFi Mysql InsertPoolsInfo failed", "err", err)
	return -1
}

type TokenInfoRecord struct {
	Token    string
	Symbol   string
	Decimals uint64
}

func (sql *Mysql) LoadTokensInfo() []*TokenInfoRecord {
	querySQL := fmt.Sprintf("select token, symbol, decimals from tokens;")
	rows, err := sql.db.Query(querySQL)
	if err != nil {
		log.Warn("RonFi Mysql LoadTokensInfo query data failed", "querySQL", querySQL, "err", err)
		return nil
	}
	defer func() {
		_ = rows.Close()
	}()

	tokensInfo := make([]*TokenInfoRecord, 0, 5000)
	for rows.Next() {
		var (
			token    string
			symbol   string
			decimals uint64
		)

		if err = rows.Scan(&token, &symbol, &decimals); err == nil {
			tokensInfo = append(tokensInfo, &TokenInfoRecord{
				Token:    token,
				Symbol:   symbol,
				Decimals: decimals,
			})
		}
	}

	return tokensInfo
}

func (sql *Mysql) InsertTokensInfo(newTokensInfo []*TokenInfoRecord) int64 {
	updateSQL := fmt.Sprintf("insert ignore into tokens (token, symbol, decimals) values ")

	length := len(newTokensInfo)
	for index, info := range newTokensInfo {
		if index == length-1 {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\", %d) ",
				updateSQL,
				info.Token,
				info.Symbol,
				info.Decimals)
		} else {
			updateSQL = fmt.Sprintf("%s (\"%s\", \"%s\", %d), ",
				updateSQL,
				info.Token,
				info.Symbol,
				info.Decimals)
		}
		index++
	}

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			//log.Info("RonFi Mysql InsertTokensInfo done!", "num", length)
			return lastId
		}
	}

	log.Warn("RonFi Mysql InsertTokensInfo failed", "err", err)
	return -1
}

type LoopObsRecord struct {
	Tx    string
	ObsId string
	Loops string
}

func (sql *Mysql) InsertObsAll(record *LoopObsRecord) int64 {
	updateSQL := fmt.Sprintf("insert ignore into obsall (tx, obsId, loops) values ")
	updateSQL = fmt.Sprintf("%s ('%s', '%s', '%s')",
		updateSQL,
		record.Tx,
		record.ObsId,
		record.Loops,
	)

	rows, err := sql.db.Exec(updateSQL)
	if err == nil {
		if lastId, err := rows.LastInsertId(); err == nil {
			return lastId
		}
	}

	log.Warn("RonFi Mysql InsertObsAll failed", "err", err)
	return -1
}

func (sql *Mysql) GetObsRecordByTxHash(hash string) (string, bool) {
	querySQL := fmt.Sprintf("select loops from obsall where tx='%s';", hash)

	rows, err := sql.db.Query(querySQL)
	if err != nil {
		log.Warn("RonFi Mysql GetObsRecordByTxHash failed", "querySQL", querySQL, "err", err)
		return "", false
	}
	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var loops string

		if err = rows.Scan(&loops); err == nil {
			return loops, true
		}
	}

	return "", false
}
