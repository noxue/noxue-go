package model

import (
	"github.com/noxue/mgodb"
	"gopkg.in/mgo.v2/bson"
)

// 标签
type Tag struct {
	mgodb.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Name  string // 标签名称
	Count int    // 标签被引用次数
	Time  `bson:",inline"`
}

func (this *Tag) GetCName() string {
	return "tag"
}
