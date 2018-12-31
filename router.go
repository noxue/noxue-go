/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 11:41
 */
package main

import (
	"github.com/gin-gonic/gin"
	"noxue/api"
)

func init() {
	Router = gin.Default()

	v1 := Router.Group("/v1")

	v1.GET("/captcha", api.ApiCaptcha.Create) // 获取图片验证码
	v1.GET("/code", api.ApiCode.Create) // 获取邮箱或手机验证码

	v1.POST("/users", api.ApiUser.Register)

}
