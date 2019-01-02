/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 12:58
 */
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"noxue/srv"
	"noxue/utils"
	"regexp"
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
	numberType := ""                // 号码类型

	if id == "" {
		utils.CheckApiError(422, errors.New("缺少参数id"))
	}

	if code == "" {
		utils.CheckApiError(422, errors.New("验证码不能为空，请输入您收到的验证码"))
	}

	if number == "" {
		utils.CheckApiError(422, errors.New("账号不能为空，请输入正确的邮箱或手机号"))
	}

	if !srv.SrvCaptcha.Verfiy(id, code) {
		utils.CheckApiError(422, errors.New("验证码不正确，请输入您收到的正确验证码"))
	}

	if m, _ := regexp.MatchString("^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+", number); m {
		numberType = "email"
	} else if m, _ := regexp.MatchString(`^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\d{8}$`, number); m {
		numberType = "phone"
	} else {
		utils.CheckApiError(422, errors.New("只支持邮箱或手机注册,请确认账号格式是否正确"))
	}

	key, err := srv.ApiCode.SendReg(number, numberType)
	if err != nil {
		utils.CheckApiError(422, err)
	}

	c.JSON(200, gin.H{"id": key})
}
