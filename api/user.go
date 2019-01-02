/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 13:18
 */
package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"noxue/model"
	"noxue/srv"
	"noxue/utils"
	"regexp"
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
	utils.CheckApiError(422, err)

	// 判断code是否正确
	err = utils.CheckVerifyCode(userReg.Name, userReg.CodeId, userReg.Code)
	utils.CheckApiError(422, err)

	// 获取默认注册后的用户组
	group, err := srv.SrvUser.GroupFind("普通用户")
	utils.CheckApiError(422, err)

	var user model.User
	user.Groups = append(user.Groups, group.Id)
	user.Name = userReg.Nick

	var auth model.Auth
	auth.Name = userReg.Name
	auth.Secret = userReg.Secret
	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", auth.Name); m {
		auth.Type = model.AuthTypeEmail
		user.Email = auth.Name
	} else if m, _ := regexp.MatchString(`^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\d{8}$`, auth.Name); m {
		auth.Type = model.AuthTypePhone
		user.Phone = auth.Name
	} else {
		utils.CheckApiError(422, errors.New("只支持邮箱或手机注册,请确认账号是否正确"))
	}

	err = srv.SrvUser.UserRegister(&user, &auth)
	utils.CheckApiError(422, err)

	c.JSON(200, gin.H{})
}
