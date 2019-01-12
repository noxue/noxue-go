/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 13:18
 */
package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"noxue/model"
	"noxue/srv"
	"noxue/utils"
	"regexp"
)

var ApiUser UserApi

type UserApi struct {
}

// ==========================================================================


// 用户注册
type UserReg struct {
	Type   int    // 登陆类型
	Nick   string // 名字
	Name   string // 账号
	Secret string // 密码
	CodeId string // 验证码Id
	Code   string // 验证码
}

// 注册用户
func (UserApi) Register(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()

	var userReg UserReg
	err := c.BindJSON(&userReg)
	utils.CheckApiError(422, -1, err)

	// 判断code是否正确
	err = utils.CheckVerifyCode(userReg.Name, userReg.CodeId, userReg.Code)
	utils.CheckApiError(422, -1, err)

	// 获取默认注册后的用户组
	group, err := srv.SrvUser.GroupFind("普通用户")
	utils.CheckApiError(422, -1, err)

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
		utils.CheckApiError(422, -1, errors.New("只支持邮箱或手机注册,请确认账号是否正确"))
	}

	err = srv.SrvUser.UserRegister(&user, &auth)
	utils.CheckApiError(422, -1, err)

	c.JSON(200, gin.H{})
}


func  (UserApi)GroupList(c *gin.Context){
	defer func() {
		if e := recover(); e != nil {
			CheckError(c, e)
		}
	}()

	fmt.Println(utils.ParseSelectParam(c))

	groups,err := srv.SrvUser.GroupSelect(nil,nil,nil)
	utils.CheckErr(err)
	c.JSON(200,gin.H{"data":groups})

}