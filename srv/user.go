/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

package srv

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/noxue/ormgo.v1"
	"noxue/dao"
	"noxue/model"
	"noxue/utils"
)

var UserSrv UserService

type UserService struct {
}

func (UserService) GroupExists(name string) (isExists bool, err error) {
	var n int
	n, err = dao.UserDao.GroupCount(map[string]interface{}{"name": name})
	if err != nil {
		return
	}
	isExists = n > 0
	return
}

func (UserService) GroupAdd(group model.UserGroup) (err error) {
	err = dao.UserDao.GroupInsert(group.Name, group.Icon)
	return
}

func (UserService) GroupFindById(id string) (group model.UserGroup, err error) {
	group, err = dao.UserDao.GroupFindById(id)
	return
}

func (UserService) GroupSelect(condition map[string]interface{}, sorts []string, page, size int) (groups []model.UserGroup, err error) {
	groups, err = dao.UserDao.GroupSelect(condition, sorts, page, size)
	return
}

// 获取指定api能被哪些用户组访问
func (this *UserService) GroupSelectByApi(api string) (groups []model.UserGroup, total int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	// 获取所有允许访问此api的groupId
	rs, err := this.ResourceSelectByApi(api)
	utils.CheckErr(err)
	var ids []bson.ObjectId
	for _, v := range rs {
		ids = append(ids, v.Group)
	}

	// 根据groupId数组查询出满足条件的group文档
	dao.UserDao.GroupSelect(ormgo.M{
		"_id": ormgo.M{"$in": ids},
	}, nil, 0, 0)
	return
}

func (UserService) GroupRemoveById(id string) (err error) {
	err = dao.UserDao.GroupRemove(id)
	return
}

// ====================================================================================================

// 注册用户，根据用户信息和授权信息添加用户
func (this *UserService) UserRegister(user *model.User, auth *model.Auth) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	// 检查用户是否存在
	exists, err := this.UserExists(user.Name)
	utils.CheckErr(err)
	if exists {
		utils.CheckErr(errors.New("用户名[" + user.Name + "]已被占用，请更换一个"))
	}

	// 检查授权信息是否存在
	exists, err = this.AuthExists(auth)
	utils.CheckErr(err)
	if exists {
		utils.CheckErr(errors.New("账号[" + auth.Name + "]已注册，请直接登陆"))
	}

	// 创建一个id，后面添加授权信息需要用到
	user.Id = bson.NewObjectId()
	// 添加用户信息
	err = dao.UserDao.UserInsert(user)
	utils.CheckErr(err)

	// 添加授权信息
	auth.User = user.Id
	err = dao.UserDao.AuthInsert(auth)
	if err != nil {
		// 如果添加授权失败，删除上面添加的用户信息
		// 防止用户名被占用缺无法登陆
		dao.UserDao.UserRemoveById(user.Id.Hex(), true)
	}

	return
}

// 根据授权信息登陆，登陆成功，返回用户信息和授权信息
func (UserService) UserLogin(auth *model.Auth) (user model.User, authRet model.Auth, err error) {

	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	// 查询授权信息
	authRet, err = dao.UserDao.AuthFind(auth.Type, auth.Name, auth.Third)
	if err != nil {
		utils.CheckErr(errors.New("账号不存在"))
	}

	// 不是第三方，就验证密码。验证密码
	if !auth.Third {
		// 密码在model.User.BeforeSave() hook方法中加密
		err = bcrypt.CompareHashAndPassword([]byte(authRet.Secret), []byte(auth.Secret))
		if err != nil {
			utils.CheckErr(errors.New("密码错误"))
		}
	}

	user, err = dao.UserDao.UserFindById(authRet.User.Hex())
	return
}

// 获取用户拥有的用户组列表
func (UserService) UserGetGroups(uid string) (groups []model.UserGroup, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(utils.Error)
		}
	}()

	u, err := dao.UserDao.UserFindById(uid)
	utils.CheckErr(err)

	groups, err = dao.UserDao.GroupSelect(ormgo.M{
		"_id": ormgo.M{
			"$in": u.Groups,
		},
	}, nil, 0, 0)
	
	return
}

// 检查用户名是否存在
func (UserService) UserExists(name string) (isExists bool, err error) {
	var n int
	n, err = dao.UserDao.UserCount(ormgo.M{"name": name}, ormgo.All)
	if err == nil {
		return
	}
	isExists = n > 0
	return
}

// 根据名称查找用户
func (UserService) UserFindByName(name string) (user model.User, err error) {
	return
}

// 编辑用户资料
// 会用整个user对象替代数据库中的数据
// 注意：如果只赋值了部分字段，其他值将丢失
func (UserService) UserUpdateById(user model.User) (err error) {
	return
}

// 编辑用户部分资料
func (UserService) UserUpdateFieldsById(id string, fields map[string]interface{}) (err error) {
	return
}

// 编辑用户部分资料
func (UserService) UserUpdateFields(conditions map[string]interface{}, fields map[string]interface{}) (err error) {
	return
}

//==============================================================================================================

// 判断授权信息是否存在
func (UserService) AuthExists(auth *model.Auth) (isExists bool, err error) {
	var n int
	n, err = dao.UserDao.AuthCount(ormgo.M{
		"type":     auth.Type,
		"username": auth.Name,
	})

	if err != nil {
		return
	}
	isExists = n > 0
	return
}

// 根据用户ID查询所有授权信息
func (UserService) AuthSelectByUid(uid string) (auths []model.Auth, err error) {
	return
}

// 添加授权信息
func (UserService) AuthAdd(auth *model.Auth) (err error) {

	return
}

// 修改授权信息
func (UserService) AuthEdit(conditions map[string]interface{}, auth *model.Auth) (err error) {
	return
}

// 判断账号密码是否正确
func (UserService) AuthCheck(auth *model.Auth) (err error) {
	return
}

// 根据用户编号修改所有非第三方登陆的密码，防止出现手机和邮箱登陆密码不一致的问题
func (UserService) AuthChangePassByUid(uid, password string) (err error) {
	return
}

// 删除指定用户的所有授权信息，用于删除用户的时候
func (UserService) AuthRemoveByUid(uid string) (err error) {
	return
}

// 根据id删除第三方授权信息，用于解绑第三方账号
func (UserService) AuthRemoveBId(id string) (err error) {
	return
}

//==============================================================================================================

// 授权给指定用户组
func (UserService) ResourceAdd(r *model.Resource) (err error) {
	return
}

// 根据用户组Id获取授权规则
func (UserService) ResourceSelectByGroupId(groupId string) (rs []model.Resource, err error) {
	return
}

// 根据api获取授权规则
func (UserService) ResourceSelectByApi(api string) (rs []model.Resource, err error) {
	return
}

// 根据用户Id获取拥有的授权规则
func (UserService) ResourceSelectByUserId(uid string) (rs []model.Resource, err error) {
	return
}

// 编辑授权规则
func (UserService) ResourceUpdate(conditions map[string]interface{}, r model.Resource) (err error) {
	return
}

func (UserService) ResourceUpdateById(id string, r model.Resource) (err error) {
	return
}

func (UserService) ResourceRemoveById(id string) (err error) {
	return
}
