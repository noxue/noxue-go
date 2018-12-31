/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 12:58
 */
package api

import (
	"github.com/gin-gonic/gin"
	"noxue/srv"
	"noxue/utils"
)

var ApiCode CodeApi

type CodeApi struct {
}

func (CodeApi) Create(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			checkError(c, e)
		}
	}()

	id := c.Query("captcha_id")     // 验证码id
	code := c.Query("captcha_code") // 验证码值
	number := c.Query("number")     // 号码
	numberType := c.Query("type")   // 号码类型

	if id == "" {
		utils.CheckApiError(422, "缺少参数id")
	}

	if code == "" {
		utils.CheckApiError(422, "缺少参数code")
	}

	if number == "" {
		utils.CheckApiError(422, "缺少参数number")
	}

	if numberType == "" {
		utils.CheckApiError(422, "缺少参数type")
	}

	if !srv.SrvCaptcha.Verfiy(id, code) {
		utils.CheckApiError(422, "验证码错误")
	}

	key, err := srv.ApiCode.SendReg(number, numberType)
	if err != nil {
		utils.CheckApiError(422, err.Error())
	}

	c.JSON(200, gin.H{"id": key})
}
