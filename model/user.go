/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/utils"
)

// 用户组
type UserGroup struct {
	ormgo.Model `bson:",inline"`
	Id          bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Name        string // 用户组名称
	Icon        string // 用户组图标
}

func NewUserGroup() *UserGroup {
	ug := new(UserGroup)
	ug.SetDoc(ug)
	return ug
}

func NewUser() *User {
	u := new(User)
	u.SetDoc(u)
	return u
}

type Sex int

const (
	SexUnknow Sex = iota // 未知
	SexMale              // 男
	SexFemale            // 女
)

// 用户
type User struct {
	ormgo.Model `bson:",inline"`
	Id          bson.ObjectId   `bson:"_id,omitempty" json:"Id,omitempty"`
	Groups      []bson.ObjectId `json:",omitempty"` // 所属用户组ID
	Name        string          `json:",omitempty"` // 昵称
	RealName    string          `json:",omitempty"` // 真实姓名
	Sex         Sex             `json:",omitempty"` // 性别
	Avatar      string          `json:",omitempty"` // 头像地址
	Email       string          `json:",omitempty"` // 邮箱地址
	Phone       string          `json:",omitempty"` // 手机号
	Summary     string          `json:",omitempty"` // 用户简介
	Followers   int             `json:",omitempty"` // 用户关注的人数
	Fans        int             `json:",omitempty"` // 粉丝个数
	Disable     int             `json:",omitempty"` // 禁言时长，-1表示永久禁言
	Time        `bson:",inline" json:"Id,omitempty"`
}

// 授权类型
type AuthType int

const (
	_              AuthType = iota
	AuthTypeEmail   // 邮箱登陆
	AuthTypePhone   // 手机登陆
	AuthTypeQQ      // qq登陆
	AuthTypeGithub  // github登陆
)

// 授权登陆信息
type Auth struct {
	ormgo.Model `bson:",inline"`
	Id          bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	User        bson.ObjectId    // 用户ID
	Type        AuthType         // 授权登陆类型
	Name        string           // 账号
	Secret      string           // 密码
	Third       bool             // 是否时第三方登陆
	Time        `bson:",inline"` // 记录什么时候注册或绑定
}

// 在保存密码的时候自动加密密码
func (this *Auth) BeforeSave() {

	// 不是第三方登陆，就加密密码
	if !this.Third {
		this.Secret = utils.EncodePassword(this.Secret)
	}
}

// 授权信息,记录哪个用户组可以访问哪些api
type Resource struct {
	ormgo.Model `bson:",inline"`
	Id          bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Api         string        // 允许访问的api地址
	Group       bson.ObjectId // 用户组
}

// 登陆日志
type LoginLog struct {
	ormgo.Model `bson:",inline"`
	Id          bson.ObjectId `bson:"_id,omitempty" json:"Id,omitempty"`
	Ip          string // 登陆IP
	Ua          string // 浏览器类型
	Time        `bson:",inline"`
}

// 用户登陆
type UserLogin struct {
	Type     int    // 登陆类型
	Username string // 账号
	Password string // 密码
	Code     string // 验证码
}

// 用户注册
type UserReg struct {
	Type     int    // 登陆类型
	Name     string // 名字
	Username string // 账号
	Password string // 密码
	Code     string // 验证码
}
