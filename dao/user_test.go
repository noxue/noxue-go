package dao

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/model"
	"testing"
	"time"
)

func init() {
	ormgo.UseSoftDelete(model.User{})
}

func TestUserDaoType_GroupInsert(t *testing.T) {
	err := UserDao.GroupInsert(fmt.Sprint("不学网", time.Now().UnixNano()), fmt.Sprint("icon", time.Now().UnixNano()))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDaoType_GroupSelect(t *testing.T) {
	groups, err := UserDao.GroupSelect(nil, nil, 1, 10,0)
	if err != nil {
		t.Fatal(err)
	}

	if len(groups) == 0 {
		t.Fatal("请先添加点数据")
	}

	// 测试编辑
	err = UserDao.GroupEditById(groups[0].Id.Hex(),
		&model.UserGroup{
			Name: fmt.Sprint(groups[0].Name,"哈哈"),
			Icon: fmt.Sprint("icon", time.Now().UnixNano()),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// 测试删除
	err = UserDao.GroupRemove(groups[0].Id.Hex())
	if err != nil {
		panic(err)
	}
}

func TestUserDaoType_UserInsert(t *testing.T) {
	err := UserDao.UserInsert(&model.User{
		Name:  fmt.Sprint("刘荣飞", time.Now().UnixNano()),
		Sex:   model.SexMale,
		Group: bson.ObjectIdHex("5c23a1d3311f62802e8aacd6"),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDaoType_UserSelect(t *testing.T) {
	users, total, err := UserDao.UserSelect(nil, nil, nil, 1, 10,ormgo.ContainTypeDefault)
	fmt.Println(users, total, err)
}

func TestUserDaoType_UserEditById(t *testing.T) {
	err := UserDao.UserEditById("5c25805a265e2a43fb2af7ce",ormgo.M{"name":"1111111111111222222222222"})
	if err != nil {
		t.Error(err)
	}
}