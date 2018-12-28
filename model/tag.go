/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

 package model

import (
	"gopkg.in/noxue/ormgo.v1"
	"gopkg.in/mgo.v2/bson"
)

// 标签
type Tag struct {
	ormgo.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Name  string // 标签名称
	Count int    // 标签被引用次数
	Time  `bson:",inline"`
}

