/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 11:25
 */
package utils

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// 加密密码，调用者需要用recover捕获错误信息
func EncodePassword(password string) (str string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	CheckErr(err)
	str = string(hash)
	return
}

// 生成并返回一个四位数的随机数
// pre 做为key的前缀，用于区分不同号码，防止a号码发送的验证码，提交的时候换成b号码也能注册的问题
func GenerateVerifyCode(pre string) (key, code string) {
	code = fmt.Sprint(rand.Intn(8999) + 1000)
	key = Uuid()
	// 保存到redis
	Redis.Do("SET", pre+key, code, "EX", 300)
	return
}

// 检测验证码
func CheckVerifyCode(pre, key, code string) error {
	retCode, err := redis.String(Redis.Do("GET", pre+key))
	if err != nil {
		return err
	}
	if retCode != code {
		return errors.New("验证码不正确，请重新输入")
	}

	// 验证码正确，删除redis中的记录
	Redis.Do("DEL", key)
	return nil
}

func Uuid() string {
	unix32bits := uint32(time.Now().UTC().Unix())

	buff := make([]byte, 12)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		//panic(err)
		glog.Error(err)
		return ""
	}

	return fmt.Sprintf("%x%x%x%x%x%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}
