/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/28 7:19
 */
package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/config"
	"noxue/model"
	"noxue/utils"
	"time"
)

func init() {
	err := ormgo.Init(config.Config.Db.Url, config.Config.Db.DbName, false, time.Second*30)
	utils.CheckErr(err)

	ormgo.UseSoftDelete(&model.User{}, &model.Auth{}, &model.Article{}, &model.Blog{}, &model.Course{}, &model.CourseItem{})

	ormgo.SessionExec(func(database *mgo.Database) {
		index := mgo.Index{
			Key:        []string{"name"},
			Unique:     true, // 唯一索引 同mysql唯一索引
			Background: true, // 后台创建索引
		}
		database.C("User").EnsureIndex(index)
		database.C("UserGroup").EnsureIndex(index)
		database.C("Auth").EnsureIndexKey("name", "type")
	})

}
