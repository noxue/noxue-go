package model

import "gopkg.in/mgo.v2/bson"

// 用户组
type UserGroup struct {
	Id   bson.ObjectId
	Name string // 用户组名称
	Icon string // 用户组图标
}

type Sex int

const (
	SexUnknow Sex = iota // 未知
	SexMale              // 男
	SexFemale            // 女
)

// 用户
type User struct {
	Id        bson.ObjectId
	Group     bson.ObjectId // 所属用户组ID
	Name      string        // 昵称
	RealName  string        // 真实姓名
	Sex       Sex           // 性别
	Avatar    string        // 头像地址
	Email     string        // 邮箱地址
	Phone     string        // 手机号
	Summary   string        // 用户简介
	Followers int           // 用户关注的人数
	Fans      int           // 粉丝个数
	Extends   []UserExtend  // 扩展信息
	Disable   int           // 禁言时长，-1表示永久禁言
	Time
}

// 用户扩展信息
type UserExtend struct {
	Name  string // 用来区分的标题
	Title string // 用于显示的名称
	Value string // 内容
	Desc  string // 描述该信息的内容
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
	User   bson.ObjectId // 用户ID
	Type   AuthType      // 授权登陆类型
	Id     string        // 账号
	Secret string        // 密码
}

// 授权信息,记录哪个用户组可以访问哪些api
type Permission struct {
	Id    bson.ObjectId
	Api   string        // 允许访问的api地址
	Group bson.ObjectId // 用户组
}

// 登陆日志
type LoginLog struct {
	Id bson.ObjectId
	Ip string // 登陆IP
	Ua string // 浏览器类型
	Time
}

// 用户登陆
type UserLogin struct {
	Type   int    // 登陆类型
	Id     string // 账号
	Secret string // 密码
	Code   string // 验证码
}

// 用户注册
type UserReg struct {
	Type   int    // 登陆类型
	Name   string // 名字
	Id     string // 账号
	Secret string // 密码
	Code   string // 验证码
}
