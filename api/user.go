/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 13:18
 */
package api

import (
	"github.com/gin-gonic/gin"
	"noxue/model"
	"noxue/srv"
	"noxue/utils"
	"strings"
)

var ApiUser UserApi

type UserApi struct {
}

// 注册用户
func (UserApi) Register(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			checkError(c, e)
		}
	}()

	var userReg model.UserReg
	err := c.BindJSON(&userReg)
	utils.CheckApiError(422, err.Error())

	// 判断code是否正确
	err = utils.CheckVerifyCode(userReg.Name, userReg.CodeId, userReg.Code)
	utils.CheckApiError(422, err.Error())

	// 获取默认注册后的用户组
	group, err := srv.SrvUser.GroupFind("普通用户")
	utils.CheckApiError(422, err.Error())

	var user model.User
	user.Groups = append(user.Groups, group.Id)
	user.Name = userReg.Nick

	var auth model.Auth
	auth.Name = userReg.Name
	auth.Secret = userReg.Secret
	if strings.Index(auth.Name, "@") == -1 {
		auth.Type = model.AuthTypePhone
	} else {
		auth.Type = model.AuthTypeEmail
	}

	err = srv.SrvUser.UserRegister(&user, &auth)
	utils.CheckApiError(422, err.Error())

	c.JSON(200, gin.H{})
}
