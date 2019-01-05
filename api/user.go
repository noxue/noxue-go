/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 13:18
 */
package api

import (
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"noxue/config"
	"noxue/middleware/myjwt"
	"noxue/model"
	"noxue/srv"
	"noxue/utils"
	"regexp"
	"time"
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

// 用户登陆
func (UserApi) Login(c *gin.Context) {
	var loginReq model.UserLogin
	if c.BindJSON(&loginReq) == nil {
		user, auth, err := srv.SrvUser.UserLogin(&model.Auth{
			Name:   loginReq.Name,
			Secret: loginReq.Secret,
		})
		if err == nil {
			generateToken(c, &UserInfo{
				Id:   user.Id.Hex(),
				Nick: user.Name,
				Name: auth.Name,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败," + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "json 解析失败",
		})
	}
}

type UserInfo struct {
	Id   string
	Nick string
	Name string
}

type LoginResult struct {
	Token    string
	UserInfo `json:",inline"`
}

// 生成令牌
func generateToken(c *gin.Context, user *UserInfo) {
	j := &myjwt.JWT{
		[]byte(config.Config.AppKey),
	}
	claims := myjwt.CustomClaims{
		user.Nick,
		user.Name,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "noxue",            //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	glog.V(3).Info(token)

	data := LoginResult{
		UserInfo: *user,
		Token:    token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}
