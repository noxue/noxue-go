/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */
package dao

import (
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/model"
	"noxue/utils"
	"time"
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
	n, err := ug.Count(ormgo.Query{
		Condition: ormgo.M{"name": name},
	})
	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该用户组已存在"))
	}
	return ug.Save()
}

func (UserDaoType) GroupSelect(condition ormgo.M, sorts []string, page, size int, contain ormgo.ContainType) (userGroups []model.UserGroup, err error) {

	query := ormgo.Query{
		Condition:  condition,
		SortFields: sorts,
		Limit:      size,
		Skip:       (page - 1) * size,
		Contain:    contain,
	}

	err = ormgo.FindAll(query, &userGroups)
	return
}

// 编辑用户组
func (UserDaoType) GroupEditById(id string, userGroup *model.UserGroup) (err error) {

	if userGroup == nil {
		glog.Error("userGroup不能为空指针")
		err = errors.New("userGroup不能为空指针")
		return
	}

	userGroup.SetDoc(userGroup)
	err = userGroup.UpdateId(id, ormgo.M{
		"name": userGroup.Name,
		"icon": userGroup.Icon,
	})

	return
}

// 获取用户组
func (UserDaoType) GroupFindById(id string) (userGroup model.UserGroup, err error) {
	err = ormgo.FindById(id, nil, &userGroup)
	return
}

// 删除用户组
func (UserDaoType) GroupRemove(id string) (err error) {
	err = model.NewUserGroup().RemoveById(id)
	return
}

// =======================================================================================================

// 添加用户
func (UserDaoType) UserInsert(user *model.User) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	if user == nil {
		glog.Error("user不能是nil")
		err = errors.New("user不能是nil")
		return
	}

	user.SetDoc(user)

	// 判断用户是否已存在
	n, err := user.Count(ormgo.Query{
		Condition: ormgo.M{"name": user.Name},
	})
	utils.CheckErr(err)
	if n > 0 {
		err = errors.New("用户名 [" + user.Name + "] 已存在")
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return user.Save()
}

// 查询用户列表
func (UserDaoType) UserSelect(condition ormgo.M, selector map[string]bool, sorts []string, page, size int, contain ormgo.ContainType) (users []model.User, total int, err error) {

	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	u := model.NewUser()

	query := ormgo.Query{
		Condition:  condition,
		Limit:      size,
		Skip:       (page - 1) * size,
		Selector:   selector,
		SortFields: sorts,
		Contain:    contain,
	}

	total, err = u.Count(query)
	utils.CheckErr(err)

	err = ormgo.FindAll(query, &users)

	return
}

func (UserDaoType) UserEditById(id string, v ormgo.M) (err error) {
	u := &model.User{}
	u.SetDoc(u)
	v["updatedat"] = time.Now().UTC()
	err = u.UpdateId(id, v)
	return
}

func (UserDaoType) UserFindById(id string) (user model.User, err error) {
	err = ormgo.FindById(id, nil, &user)
	return
}

// really 表示是否真正删除，false表示软删除
func (UserDaoType) UserRemoveById(id string, really bool) (err error) {
	u := &model.User{}
	u.SetDoc(u)
	if really {
		err = u.RemoveTrueById(id)
	} else {
		err = u.RemoveById(id)
	}
	return
}

// =================================================================================================================

func (UserDaoType) AuthExists(auth *model.Auth) bool {
	auth.SetDoc(auth)
	n, err := auth.Count(ormgo.Query{
		Condition: ormgo.M{
			"type":     auth.Type,
			"username": auth.Username,
		},
	})

	if err != nil || n == 0 {
		return true
	}

	return false
}

func (this *UserDaoType) AuthInsert(auth *model.Auth) (err error) {
	if this.AuthExists(auth) {
		err = errors.New("用户已存在[" + auth.Username + "]")
		return
	}
	auth.CreatedAt = time.Now().UTC()
	err = ormgo.Save(auth)
	return
}

func (this *UserDaoType) AuthFind(typeName, username, password string) (auth model.Auth, err error) {
	err = ormgo.FindOne(
		ormgo.M{
			"type":     auth.Type,
			"username": auth.Username,
			"password": auth.Password},
		nil,
		&auth,
	)
	return
}

func (this *UserDaoType) AuthFindByUid(userId string) (auths []model.Auth, err error) {
	err = ormgo.FindAll(
		ormgo.Query{
			Condition: ormgo.M{
				"user": bson.ObjectIdHex(userId),
			},
		},
		&auths,
	)
	return
}

func (this *UserDaoType) AuthUpdate(condition ormgo.M, v ormgo.M) (err error) {
	auth := &model.Auth{}
	auth.SetDoc(auth)
	v["updatedat"] = time.Now().UTC()
	err = auth.Update(condition, v)
	return
}

func (this *UserDaoType) AuthRemoveById(id string) (err error) {
	auth := &model.Auth{}
	auth.SetDoc(auth)
	err = auth.RemoveById(id)
	return
}

//=========================================================================================================

func (UserDaoType) PermissionInsert(p model.Permission) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	p.SetDoc(p)
	var n int
	n, err = p.Count(ormgo.Query{
		Condition: ormgo.M{
			"api":   p.Api,
			"group": p.Group,
		},
	})
	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该规则已存在"))
	}

	err = ormgo.Save(p)
	return
}

func (UserDaoType) PermissionSelect(condition ormgo.M, selector map[string]bool, sorts []string, page, size int) (ps []model.Permission, total int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	query := ormgo.Query{
		Condition:  condition,
		Limit:      size,
		Skip:       (page - 1) * size,
		Selector:   selector,
		SortFields: sorts,
	}

	p := &model.Permission{}
	p.SetDoc(p)
	total, err = p.Count(query)
	utils.CheckErr(err)

	err = ormgo.FindAll(query, &ps)
	return
}

func (UserDaoType) PermissionFindById(id string) (p model.Permission, err error) {
	err = ormgo.FindById(id, nil, &p)
	return
}

func (UserDaoType) PermissionRemoveById(id string) (err error) {
	p := &model.Permission{}
	p.SetDoc(p)
	err = p.RemoveById(id)
	return
}

func (UserDaoType) PermissionEditById(id string, p model.Permission) (err error) {
	err = p.UpdateId(id, ormgo.M{"api": p.Api, "group": p.Group})
	return
}

//================================================================================================================

func (UserDaoType) LoginLogInsert(log model.LoginLog) (err error) {
	err = ormgo.Save(log)
	return
}

func (UserDaoType) LoginLogSelect(condition ormgo.M, selector map[string]bool, sorts []string, page, size int) (logs []model.LoginLog, total int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	query := ormgo.Query{
		Condition:  condition,
		Limit:      size,
		Skip:       (page - 1) * size,
		Selector:   selector,
		SortFields: sorts,
	}

	log := &model.LoginLog{}
	log.SetDoc(log)

	total, err = log.Count(query)
	utils.CheckErr(err)

	err = ormgo.FindAll(query, &logs)

	return
}

func (UserDaoType) LoginLogRemove(condition ormgo.M) (err error) {
	log := &model.LoginLog{}
	log.SetDoc(log)

	_, err = log.RemoveAll(condition)
	return
}
