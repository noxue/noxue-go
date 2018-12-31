/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 20:04
 */
package utils

import "github.com/gomodule/redigo/redis"

var Redis redis.Conn

func init() {
	var err error
	Redis, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
}
