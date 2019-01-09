/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 11:41
 */
package main

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"noxue/api/v1"
	"noxue/config"
	"noxue/model"
	"noxue/srv"
	"time"
)

const identityKey = "id"

// 不需要授权的路由组
var r1 *gin.RouterGroup
// 需要授权的路由组
var ar1 *gin.RouterGroup

func init() {
	Router = gin.Default()
	Router.Use(gin.Recovery())

	Router.Use(Cors())
	r1 = Router.Group("/v1")
	initJwt()

	r1.GET("/captcha", v1.ApiCaptcha.Create) // 获取图片验证码
	r1.GET("/code", v1.ApiCode.Create)       // 获取邮箱或手机验证码

	r1.POST("/users", v1.ApiUser.Register)
	//r1.POST("/token", v1.ApiUser.Login)  // 创建一个token，登陆
	//r1.PUT("/token", v1.ApiUser.Refresh) // 更新token，刷新token信息
	ar1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"test": 1})
	})

}

// 用户登陆
type UserLogin struct {
	Type   int    // 登陆类型
	Name   string // 账号
	Secret string // 密码
}

type User struct {
	Id   string
	Nick string
	Name string
}

// 初始化jwt
func initJwt() {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "web",
		Key:         []byte(config.Config.AppKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour * 2,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Id,
					"nick":      v.Nick,
					"name":      v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Id:   claims["id"].(string),
				Nick: claims["nick"].(string),
				Name: claims["name"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals UserLogin
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Name
			password := loginVals.Secret
			user, auth, err := srv.SrvUser.UserLogin(&model.Auth{
				Name:   userID,
				Secret: password,
			})
			if (err == nil) {
				u := &User{
					Id:   user.Id.Hex(),
					Nick: user.Name,
					Name: auth.Name,
				}
				c.Set("user", u)
				return u, nil
			}

			return nil, err
		},
		LoginResponse: func(c *gin.Context, code int, s string, t time.Time) {
			u, ok := c.Get("user")
			if ok {
				c.JSON(code, gin.H{"expire": t, "token": s, "user": u})
			} else {
				c.JSON(401, gin.H{"status": 1000, "msg": "登陆失败"})
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.Name == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"status": code,
				"msg":    message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}

	r1.POST("/token", authMiddleware.LoginHandler)
	ar1 = Router.Group("/v1")
	ar1.PUT("/token", authMiddleware.RefreshHandler)
	ar1.Use(authMiddleware.MiddlewareFunc())

	Router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		glog.V(2).Info("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{"code": 0})
			return
		}
		c.Next()
	}
}
