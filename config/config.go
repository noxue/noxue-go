/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */
package config

import (
	"encoding/json"
	"io/ioutil"
)

var Config Conf // 保存所有的配置信息，全局可以访问

func init() {
	// 读取配置文件
	bs, err := ioutil.ReadFile(`./config.json`)
	if err != nil {
		panic(err)
	}

	// 解析json数据
	err = json.Unmarshal(bs, &Config)
	if err != nil {
		panic(err)
	}
}

type Conf struct {
	AppKey string // 加密jwt所需要
	Debug  bool   // 是否是调试模式
	Server Server // 服务器信息配置
	Db     Db     // 数据库信息
}

type Server struct {
	Port    int // API web服务监听的端口
	SeoPort int // Seo web服务监听端口
}

type Db struct {
	Url    string // mongodb连接字符串
	DbName string // 数据库名
}
