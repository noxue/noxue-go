package model

import "gopkg.in/mgo.v2/bson"

// 标签
type Tag struct {
	Id    bson.ObjectId
	Name  string // 标签名称
	Count int    // 标签被引用次数
	Time
}
