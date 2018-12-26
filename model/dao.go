/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

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
	mgodb.Debug = true
	utils.CheckErr(err)
}

