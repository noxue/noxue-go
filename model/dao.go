package model

import (
	"github.com/noxue/mgodb"
	"noxue/config"
	"noxue/utils"
	"time"
)

func init() {
	var dbm *mgodb.Dbm
	err := dbm.Init(config.Config.Db.Url, config.Config.Db.DbName, time.Second*30)
	utils.CheckErr(err)
}

