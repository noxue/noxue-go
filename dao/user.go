package dao

import (
	"github.com/noxue/mgodb"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"noxue/model"
	"noxue/utils"
)

var UserDao *UserDaoType

func init() {
	UserDao = &UserDaoType{}
}

type UserDaoType struct {
}

// 添加用户组
func (UserDaoType) AddGroup(name string, icon string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()
	ug := model.NewUserGroup()
	ug.Name = name
	ug.Icon = icon
	n, err := ug.Count(&model.UserGroup{
		Name: name,
	})

	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该用户组已存在"))
	}
	return ug.Save()
}

func (UserDaoType) Select(userGroup *model.UserGroup, sorts []string, page, size int) (userGroups []model.UserGroup, err error) {

	query := mgodb.Query{
		Limit:    size,
		Skip:     (page - 1) * size,
		Selector: map[string]bool{"name": true, "icon":true},
	}

	// 排序
	query.SetSortFields(sorts)

	if userGroup != nil {
		query.QueryDoc = userGroup
	} else {
		query.QueryDoc = bson.M{}
	}

	ug := model.NewUserGroup()
	err = ug.FindAll(query, &userGroups)
	return
}

// 编辑用户组
func (UserDaoType) Edit(Id string, userGroup model.UserGroup) (err error) {
	return nil
}
