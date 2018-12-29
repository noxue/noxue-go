/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/29 20:26
 */
package srv

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"noxue/model"
	"testing"
)

func TestUserService_UserRegister(t *testing.T) {
	user := &model.User{
		Name: "admin",
		Groups: []bson.ObjectId{
			bson.ObjectIdHex("5c258168265e2a43fb2af7ea"),
		},
		RealName: "不学网",
	}
	auth := &model.Auth{
		Type:   model.AuthTypeEmail,
		Name:   "yes@noxue.com",
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
