/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

package main

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	Router.Run() // listen and serve on 0.0.0.0:8080
}
