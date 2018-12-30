/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/29 20:26
 */
package srv

import (
	"fmt"
	"noxue/dao"
	"noxue/model"
	"testing"
	"time"
)

func TestUserService_GroupAdd(t *testing.T) {
	err := dao.UserDao.GroupInsert(fmt.Sprint("用户组", time.Now().Unix()), fmt.Sprint("icon", time.Now().Unix()))
	if err != nil {
		t.Error(err)
	}
}

func TestUserService_UserRegister(t *testing.T) {
	user := &model.User{
		Name:     "noxue",
		RealName: "不学网",
	}
	auth := &model.Auth{
		Type:   model.AuthTypeEmail,
		Name:   "no@noxue.com",
		Secret: "admin",
	}

	err := UserSrv.UserRegister(user, auth)
	if err != nil {
		t.Error(err)
	}
}

func TestUserService_UserLogin(t *testing.T) {
	auth := &model.Auth{
		Name:   "yes@noxue.com",
		Secret: "admin",
	}

	u, a, err := UserSrv.UserLogin(auth)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u, a)
}

func TestUserService_UserAddToGroups(t *testing.T) {
	err := UserSrv.UserAddToGroups("5c28140e8609ab00903e72f5", []string{"5c2813571d3930cf6a79f931"})
	if err != nil {
		t.Error(err)
	}
}

func TestUserService_UserRemoveFromGroup(t *testing.T) {
	err := UserSrv.UserRemoveFromGroup("5c28140e8609ab00903e72f5", "5c2813571d3930cf6a79f931")
	if err != nil {
		t.Error(err)
	}
}
