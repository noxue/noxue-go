/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 12:41
 */
package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/noxue/aliyun-communicate"
	"noxue/config"
)

func SendEmail(email, title, content string) (err error) {

	return
}

func SendRegCodeEmail(email, code string) (err error) {
	content := `您的验证码是 ` + code
	return SendEmail(email, "不学网用户注册验证码", content)
}

func SendRegCodePhone(phone, code string) (err error) {
	err = sendPhoneCode(phone,code,"reg")
	return
}

func sendPhoneCode(phone, code, Type string) (err error) {
	sms := config.Config.Sms
	smsClient := aliyunsmsclient.New(sms.Url)
	result, err := smsClient.Execute(sms.Id, sms.Key, phone, sms.Sign, sms.Reg, fmt.Sprintf(`{"code":"%s"}`, code))
	if err != nil {
		return
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		return err
	}
	if !result.IsSuccessful() {
		return errors.New(fmt.Sprint("Failed to send a SMS:", resultJson))
	}
	return
}
