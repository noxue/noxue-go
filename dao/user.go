/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */
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
func (UserDaoType) GroupInsert(name string, icon string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()
	ug := model.NewUserGroup()
	ug.Name = name
	ug.Icon = icon
	n, err := ug.Count(bson.M{"name": name})
	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该用户组已存在"))
	}
	return ug.Save()
}

func (UserDaoType) GroupSelect(userGroup *model.UserGroup, sorts []string, page, size int) (userGroups []model.UserGroup, err error) {

	query := mgodb.Query{
		Limit:    size,
		Skip:     (page - 1) * size,
		Selector: map[string]bool{"name": true, "icon": true},
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
func (UserDaoType) GroupEditById(id string, userGroup *model.UserGroup) (err error) {
	if !bson.IsObjectIdHex(id) {
		err = errors.New("用户组ID不正确")
		return
	}
	userGroup.Id = bson.ObjectIdHex(id)
	ug := model.NewUserGroup()
	ug.ReloadDoc(userGroup)
	return ug.Save()
}

// 获取用户组
func (UserDaoType) GroupFindById(id string) (userGroup model.UserGroup, err error) {
	if !bson.IsObjectIdHex(id) {
		err = errors.New("用户组ID不正确")
		return
	}
	ug := model.NewUserGroup()
	err = ug.FindByPk(id, &userGroup)
	return
}

// 删除用户组
func (UserDaoType) GroupDelete(id string) (err error) {
	if !bson.IsObjectIdHex(id) {
		err = errors.New("用户组ID不正确")
		return
	}
	ug := model.NewUserGroup()
	ug.Id = bson.ObjectIdHex(id)
	err = ug.Delete()
	return
}
