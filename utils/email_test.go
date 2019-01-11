/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2019/1/11 22:04
 */
package utils

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	err:=SendEmail("173126019@qq.com","用户注册验证码","验证码：1023","")
	if err!=nil {
		t.Fatal(err)
	}
}
