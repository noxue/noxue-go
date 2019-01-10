/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2019/1/9 13:33
 */
package main

import (
	"errors"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"noxue/api/v1"
	"noxue/config"
	"noxue/model"
	"noxue/srv"
	"noxue/utils"
	"time"
)

const identityKey = "id"
// 初始化jwt
func initJwt() {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "web",
		Key:         []byte(config.Config.AppKey),
		Timeout:     time.Minute,
		MaxRefresh:  time.Second * 20,
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
				c.JSON(200, gin.H{"status":code, "expire": t, "token": s, "user": u})
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

	r1.POST("/token", func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				v1.CheckError(c, e)
			}
		}()

		id := c.Query("captcha_id")     // 验证码id
		code := c.Query("captcha_code") // 验证码值

		if id == "" {
			utils.CheckApiError(422, -1, errors.New("缺少参数id"))
		}

		if code == "" {
			utils.CheckApiError(422, -1, errors.New("验证码不能为空，请输入正确的验证码"))
		}
		if !srv.SrvCaptcha.Verfiy(id, code) {
			utils.CheckApiError(422, -1, errors.New("验证码不正确，请输入正确验证码"))
		}
		authMiddleware.LoginHandler(c)
	})
	ar1 = Router.Group("/v1")
	ar1.PUT("/token", authMiddleware.RefreshHandler)
	ar1.Use(authMiddleware.MiddlewareFunc())

	Router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		glog.V(2).Info("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

}
