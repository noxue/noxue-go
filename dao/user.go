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

func (UserDaoType) GroupSelect(condition ormgo.M, sorts []string, page, size int) (userGroups []model.UserGroup, err error) {

	query := ormgo.Query{
		Condition:  condition,
		SortFields: sorts,
		Limit:      size,
		Skip:       (page - 1) * size,
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

// 统计用户组个数
func (UserDaoType) GroupCount(conditions ormgo.M) (n int, err error) {
	g := &model.UserGroup{}
	g.SetDoc(g)
	n, err = g.Count(ormgo.Query{
		Condition: conditions,
	})
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

func (UserDaoType) UserCount(conditions ormgo.M, containType ormgo.ContainType) (n int, err error) {
	u := &model.User{}
	u.SetDoc(u)

	n, err = u.Count(ormgo.Query{
		Condition: conditions,
		Contain:   containType,
	})
	return
}

// 把用户添加到指定的用户组中
func (UserDaoType) UserAddToGroups(uid string, groupIds []string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	u := model.User{}
	u.SetDoc(u)

	var ids []bson.ObjectId
	for _, v := range groupIds {
		ids = append(ids, bson.ObjectIdHex(v))
	}

	err = u.UpdateId(uid, ormgo.M{
		"$addToSet": ormgo.M{
			"groups": ormgo.M{
				"$each": ids,
			},
		},
	})

	utils.CheckErr(err)

	return
}

// =================================================================================================================

func (this *UserDaoType) AuthInsert(auth *model.Auth) (err error) {
	auth.CreatedAt = time.Now().UTC()
	auth.UpdatedAt = time.Now().UTC()
	if !bson.IsObjectIdHex(auth.User.Hex()) {
		err = errors.New("Auth指定的user Id格式不正确")
		return
	}
	err = ormgo.Save(auth)
	return
}

func (this *UserDaoType) AuthFind(authType model.AuthType, name string, isThird bool) (auth model.Auth, err error) {
	cond := ormgo.M{
		"name":  name,
		"third": isThird,
	}
	if authType != 0 {
		cond["type"] = authType
	}
	err = ormgo.FindOne(
		cond,
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

func (UserDaoType) AuthCount(conditions ormgo.M) (n int, err error) {
	auth := &model.Auth{}
	auth.SetDoc(auth)

	n, err = auth.Count(ormgo.Query{
		Condition: conditions,
	})
	return
}

//=========================================================================================================

func (UserDaoType) ResourceInsert(r model.Resource) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	r.SetDoc(r)
	var n int
	n, err = r.Count(ormgo.Query{
		Condition: ormgo.M{
			"api":   r.Api,
			"group": r.Group,
		},
	})
	utils.CheckErr(err)
	if n > 0 {
		utils.CheckErr(errors.New("该规则已存在"))
	}

	err = ormgo.Save(r)
	return
}

func (UserDaoType) ResourceSelect(condition ormgo.M, selector map[string]bool, sorts []string, page, size int) (rs []model.Resource, total int, err error) {
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

	p := &model.Resource{}
	p.SetDoc(p)
	total, err = p.Count(query)
	utils.CheckErr(err)

	err = ormgo.FindAll(query, &rs)
	return
}

func (UserDaoType) ResourceFindById(id string) (r model.Resource, err error) {
	err = ormgo.FindById(id, nil, &r)
	return
}

func (UserDaoType) ResourceRemoveById(id string) (err error) {
	p := &model.Resource{}
	p.SetDoc(p)
	err = p.RemoveById(id)
	return
}

func (UserDaoType) ResourceEditById(id string, r model.Resource) (err error) {
	err = r.UpdateId(id, ormgo.M{"api": r.Api, "group": r.Group})
	return
}

func (UserDaoType) ResourceCount(conditions ormgo.M) (n int, err error) {
	query := ormgo.Query{
		Condition: conditions,
	}

	p := &model.Resource{}
	p.SetDoc(p)
	n, err = p.Count(query)
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
