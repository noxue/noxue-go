package model

import "gopkg.in/mgo.v2/bson"

// 广告
type Ad struct {
	Id      bson.ObjectId
	Name    string // 广告名称，区分不同广告的标识，全局唯一
	Title   string // 广告标题，自定义给人看
	Content string // 广告内容，html格式
	Visible bool   // 广告是否显示，默认false不显示
	Time
}
