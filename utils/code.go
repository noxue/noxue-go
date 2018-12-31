/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 12:41
 */
package utils

func SendEmail(email, title, content string) (err error) {

	return
}

func SendRegCodeEmail(email, code string) (err error) {
	content := `您的验证码是 ` + code
	return SendEmail(email, "不学网用户注册验证码", content)
}

func SendRegCodePhone(phone, code string) (err error) {
	return
}
