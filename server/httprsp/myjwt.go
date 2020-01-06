package httprsp

import (
	"errors"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

var authMiddleware *jwt.GinJWTMiddleware
var identityKey = "id"

// My Errors
var (
	errMatchIDPass = errors.New("failLogin")
	errEmailAuth   = errors.New("failAuth")
)

type JwtAuthorizator func(data interface{}, c *gin.Context) bool

func initJWT(jwtAuthorizator JwtAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	var err error
	authMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(ourKey),
		MaxRefresh:  time.Hour * 6,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*user); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &user{
				ID: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals loginInfo
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.ID
			password := loginVals.Password

			if loginSuc(userID, password) {
				if userAuthValid(userID) {
					return &user{
						ID: userID,
					}, nil
				}
				return nil, errEmailAuth
			}
			return nil, errMatchIDPass
		},
		Authorizator: jwtAuthorizator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":   code,
				"status": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "INUser",

		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

func loginUserAuthorizator(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*user); ok {
		return true
	}
	return false
}
