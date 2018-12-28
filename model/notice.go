/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

 package model

import (
	"gopkg.in/noxue/ormgo.v1"
	"gopkg.in/mgo.v2/bson"
)

type NoticeType int

const (
	_                NoticeType = iota
	NoticeTypeSystem  // 系统通知
	NoticeTypeMsg     // 私信
	NoticeTypeUpdate  // 动态
	NoticeTypeCourse  // 来自教程讨论帖子的动态
)

// 记录 系统通知/私信/动态
type Notice struct {
	ormgo.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	User         bson.ObjectId // 通知给谁
	Article      bson.ObjectId // 帖子Id
	Comment      bson.ObjectId // 回复Id
	FromUser     bson.ObjectId // 来自谁的通知
	FromUserName string        // 通知者的用户名，冗余设计，减少查询次数
	Type         NoticeType    // 通知类型
	Content      string        // 通知内容
	Time         `bson:",inline"`
}

