/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

package model

import (
	"gopkg.in/noxue/ormgo.v1"
	"gopkg.in/mgo.v2/bson"
)

// 教程列表
type Course struct {
	ormgo.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Author   bson.ObjectId // 作者
	Finished bool          // 是否完结，false表示更新中
	Chapters []struct {
		Order int    // 排序编号
		Title string // 章节名称
		Desc  string // 章节简介描述信息
		Items []struct {
			Order      int           // 排序编号
			Price      int           // 价格，单位为分
			CourseItem bson.ObjectId // 教程内容ID
			Title      string        // 教程标题
			IsVideo    bool          // 是否是视频教程
		}
	}
}


// 教程内容
type CourseItem struct {
	Article `bson:",inline"` // 继承文章所有字段
	Price   int             // 教程价格
	Video   string          // 视频地址
}



// 订单类型
type OrderType int

const (
	_ OrderType = iota
	OrderTypeCourse
)

// 订单记录
type Order struct {
	ormgo.Model `bson:",inline"`
	Id bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	User     bson.ObjectId // 用户
	Name     string        // 订单名称
	Type     OrderType     // 订单名称
	Finished bool          // 是否完成订单
	Closed   bool          // 订单是否关闭，默认false
	Time     `bson:",inline"`
}

