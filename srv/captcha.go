/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 10:04
 */
package srv

import "github.com/noxue/base64Captcha"

var SrvCaptcha CaptchaService

type CaptchaService struct {
}

func (this *CaptchaService) Create() (id, data string) {
	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 4,
	}

	captchaId, digitCap := base64Captcha.GenerateCaptcha("", configD)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)
	return captchaId, base64Png
}

func (this *CaptchaService) Verfiy(idkey, verifyValue string) bool {
	return base64Captcha.VerifyCaptcha(idkey, verifyValue)
}
