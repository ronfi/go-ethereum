package common

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func LoadDBConfig() *DBConfig {
	configFile := common.GetAppPath() + "/db_config.json"
	if jsonData, err := common.LoadJsonFile(configFile); err != nil {
		log.Error("RonFi LoadDBConfig file failed", "file", configFile)
	} else {
		var config DBConfig
		if err = json.Unmarshal(jsonData, &config); err != nil {
			log.Error("RonFi LoadDBConfig Unmarshal failed!", "err", err)
		} else {
			log.Info("RonFi LoadDBConfig success",
				"mysql host", config.MysqlConf.DbHost,
				"mysql port", config.MysqlConf.DbPort,
				"mysql user", config.MysqlConf.DbUser,
				"mysql data", config.MysqlConf.DbData,
				"redis host", config.RedisConf.RedisHost,
				"redis port", config.RedisConf.RedisPort)
			return &config
		}
	}

	return nil
}
