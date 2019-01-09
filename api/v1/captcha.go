/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 11:30
 */
package v1

import (
	"github.com/gin-gonic/gin"
	"noxue/srv"
)

var ApiCaptcha CaptchaApi

type CaptchaApi struct {
}

func (CaptchaApi) Create(c *gin.Context) {
	id, data := srv.SrvCaptcha.Create()
	c.JSON(200, gin.H{"id": id, "data": data})
}
