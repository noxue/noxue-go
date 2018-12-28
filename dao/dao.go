/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/28 7:19
 */
package dao

import (
	"gopkg.in/noxue/ormgo.v1"
	"noxue/config"
	"noxue/utils"
	"time"
)

func init() {
	err := ormgo.Init(config.Config.Db.Url, config.Config.Db.DbName, false, time.Second*30)
	utils.CheckErr(err)
}
