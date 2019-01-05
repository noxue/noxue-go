/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 11:41
 */
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noxue/api"
)

type login struct {
	Type     string `form:"type" json:"type" binding:"required"`
	Id string `form:"id" json:"id" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type User struct {
	UserName string
}

const identityKey = "id"

func init() {
	Router = gin.Default()
	Router.Use(gin.Recovery())
	Router.Use(Cors())



	v1 := Router.Group("/v1")

	v1.GET("/captcha", api.ApiCaptcha.Create) // 获取图片验证码
	v1.GET("/code", api.ApiCode.Create) // 获取邮箱或手机验证码

	v1.POST("/users", api.ApiUser.Register)
	v1.POST("/login", api.ApiUser.Login)

}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{"code": 0})
			return
		}
		c.Next()
	}
}
