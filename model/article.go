package model

import "gopkg.in/mgo.v2/bson"



type ArticleType int

// 类似二级分类（子版块）
const (
	_                  ArticleType = iota
	ArticleTypeDiscuss  //讨论
	ArticleTypeWork     // 单子发布
	ArticleTypeAsk      // 问答
	ArticleTypeJob      // 招聘
	ArticleTypeNotice   // 公告
)

type ArticleStatus int

// 文章状态
const (
	_                     ArticleStatus = iota
	ArticleStatusTopAll    // 全局置顶
	ArticleStatusTopType   // 主题分类置顶
	ArticleStatusTopClass  // 板块置顶
	ArticleStatusTopGood1  // 精华1
	ArticleStatusTopGood2  // 精华2
	ArticleStatusTopGood3  // 精华3
)

// 文章分类（通过不同语言分类，类似社区板块）
type Class struct {
	Id    bson.ObjectId
	Name  string // 分类名称
	Order int    // 排序编号，越小越靠前
}

// 文章
type Article struct {
	Id           bson.ObjectId
	Class        bson.ObjectId // 分类编号
	Title        string        // 文章标题
	Author       bson.ObjectId // 作者ID
	Content      string        // 文章内容
	Html         string        // 从markdown转换来的html内容
	Score        int           // 问答悬赏积分，最佳答案获取
	Views        int           // 浏览量
	Likes        int           // 喜欢数
	CommentCount int           // 评论次数
	Tags         []string      // 标签
	Status       ArticleStatus // 文章状态
	Type         ArticleType   // 文章分类
	Time
}

type CommentStatus int

const (
	_                    CommentStatus = iota
	CommentStatusDisable  // 被屏蔽
	CommentStatusBest     // 在问答中表示最佳答案
)

// 评论
type Comments struct {
	Id      bson.ObjectId
	Article bson.ObjectId // 文章ID
	User    bson.ObjectId // 评论者ID
	Content string        // 评论内容
	Good    int           // 点赞数
	Status  CommentStatus // 评论状态
	Time
}
