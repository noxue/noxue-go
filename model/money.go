/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

 package model

import (
	"gopkg.in/noxue/ormgo.v1"
	"gopkg.in/mgo.v2/bson"
)

type MoneyType int

// 虚拟货币类型
const (
	_                 MoneyType = iota
	MoneyTypeDevotion  // 贡献点
	MoneyTypeScore     // 积分
)

// 虚拟货币
type Money struct {
	ormgo.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	User bson.ObjectId // 用户ID
	Type MoneyType     // 虚拟货币类型
}



// 虚拟货币变动日志
type MoneyLog struct {
	ormgo.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	User   bson.ObjectId // 用户ID
	Type   MoneyType     // 货币类型
	Change int           // 改变的货币数，正数为增，负数为减
	Reason string        // 改变的原因
	Time `bson:",inline"`
}

