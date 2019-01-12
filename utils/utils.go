/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 11:25
 */
package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/noxue/ormgo.v1"
	"math/rand"
	"strconv"
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
		if "redigo: nil returned" == err.Error() {
			return errors.New("验证码不正确，请确认与收到的验证码是否一致")
		}
		return err
	}
	if retCode != code {
		return errors.New("验证码不正确，请确认与收到的验证码是否一致")
	}

	// 验证码正确，删除redis中的记录
	Redis.Do("DEL", pre+key)
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

// 从url种解析出查询参数
func ParseSelectParam(c *gin.Context) (sort []string, field map[string]bool, filter map[string]interface{}, ids []string,page int, size int, err error) {
	defer func() {
		if e := recover(); e != nil {
			glog.Error(e)
			err = errors.New("携带的参数格式不正确，请仔细检查")
		}
	}()
	sortParam := c.Query("sort")
	fieldParam := c.Query("field")
	pageParam := c.Query("page")
	sizeParam := c.Query("size")
	filterParam := c.Query("filter")
	idsParam := c.Query("ids")
	fmt.Println(sortParam, fieldParam, filterParam)
	if sortParam != "" {
		err = json.Unmarshal([]byte(sortParam), &sort)
		CheckErr(err)
	}

	if filterParam != "" {
		err = json.Unmarshal([]byte(filterParam), &filter)
		CheckErr(err)
		// 让查询支持正则过滤
		for k, v := range filter {
			filter[k] = ormgo.M{"$regex": v}
		}
	}
	if fieldParam != "" {
		err = json.Unmarshal([]byte(fieldParam), &field)
		CheckErr(err)
	}
	if idsParam != "" {
		err = json.Unmarshal([]byte(idsParam), &ids)
		CheckErr(err)
	}
	if pageParam!=""{
		page, err = strconv.Atoi(pageParam)
		CheckErr(err)
	}
	if sizeParam!=""{
		size, err = strconv.Atoi(sizeParam)
		CheckErr(err)
	}
	return
}
