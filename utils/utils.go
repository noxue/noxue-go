/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 11:25
 */
package utils

import "golang.org/x/crypto/bcrypt"

// 加密密码，调用者需要用recover捕获错误信息
func EncodePassword(password string) (str string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	CheckErr(err)
	str = string(hash)
	return
}
