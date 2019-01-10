/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 11:41
 */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noxue/api/v1"
)


// 不需要授权的路由组
var r1 *gin.RouterGroup
// 需要授权的路由组
var ar1 *gin.RouterGroup

func init() {
	Router = gin.Default()
	Router.Use(gin.Recovery())

	Router.Use(Cors())
	r1 = Router.Group("/v1")
	initJwt()

	r1.GET("/captcha", v1.ApiCaptcha.Create) // 获取图片验证码
	r1.GET("/code", v1.ApiCode.Create)       // 获取邮箱或手机验证码

	r1.POST("/users", v1.ApiUser.Register)
	//r1.POST("/token", v1.ApiUser.Login)  // 创建一个token，登陆
	//r1.PUT("/token", v1.ApiUser.Refresh) // 更新token，刷新token信息
	ar1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"test": 1})
	})

}

// 用户登陆
type UserLogin struct {
	Type   int    // 登陆类型
	Name   string // 账号
	Secret string // 密码
}

type User struct {
	Id   string
	Nick string
	Name string
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "360000")
		//放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{"status": 0})
			c.Abort()
			return
		}
		c.Next()
	}
}
