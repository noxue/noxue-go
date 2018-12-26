package model

import (
	"github.com/noxue/mgodb"
	"gopkg.in/mgo.v2/bson"
)

// 广告
type Ad struct {
	mgodb.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Name    string // 广告名称，区分不同广告的标识，全局唯一
	Title   string // 广告标题，自定义给人看
	Content string // 广告内容，html格式
	Visible bool   // 广告是否显示，默认false不显示
	Time    `bson:",inline"`
}

func (this *Ad) GetCName() string {
	return "ad"
}
