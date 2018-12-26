package dao

import (
	"fmt"
	"noxue/model"
	"testing"
	"time"
)

func TestUserDaoType_GroupInsert(t *testing.T) {
	err := UserDao.GroupInsert(fmt.Sprint("不学网", time.Now().UnixNano()), fmt.Sprint("icon", time.Now().UnixNano()))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserDaoType_GroupSelect(t *testing.T) {
	groups, err := UserDao.GroupSelect(nil, []string{"-name"}, 1, 10)
	if err != nil {
		t.Fatal(err)
	}

	if len(groups) == 0 {
		t.Fatal("请先添加点数据")
	}

	// 测试编辑
	err = UserDao.GroupEditById(groups[0].Id.Hex(),
		&model.UserGroup{
			Name: fmt.Sprint("不学网", time.Now().UnixNano()),
			Icon: fmt.Sprint("icon", time.Now().UnixNano()),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	// 测试删除
	err = UserDao.GroupDelete(groups[0].Id.Hex())
	if err != nil {
		t.Fatal(err)
	}
}
