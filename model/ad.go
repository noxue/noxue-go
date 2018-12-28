/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/noxue/ormgo.v1"
)

// 广告
type Ad struct {
	ormgo.Model `bson:",inline"`
	Id          bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Name        string // 广告名称，区分不同广告的标识，全局唯一
	Title       string // 广告标题，自定义给人看
	Content     string // 广告内容，html格式
	Visible     bool   // 广告是否显示，默认false不显示
	Time        `bson:",inline"`
}
