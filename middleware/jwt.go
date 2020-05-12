package middleware

import (
	"encoding/json"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"log"
	"time"
)

type JWTAuthorizator func(data interface{}, c *gin.Context) bool
type JWTAuthenticator func(c *gin.Context) (interface{}, error)
type JWTPayloadFunc func(data interface{}) jwt.MapClaims
type JWTIdentityHandler func(c *gin.Context) interface{}

func getJWTMiddleware(authorizator JWTAuthorizator, authenticator JWTAuthenticator, payloadFunc JWTPayloadFunc, handler JWTIdentityHandler) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           config.GetConfig().JWTConfig.Title,
		Key:             []byte(config.GetConfig().JWTConfig.Key),
		Timeout:         time.Hour * 300,
		MaxRefresh:      time.Hour * 24,
		IdentityKey:     "id",
		PayloadFunc:     payloadFunc,
		IdentityHandler: handler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return
}

func getJWTAdminMiddleWare(authorizator JWTAuthorizator) *jwt.GinJWTMiddleware {
	adminAuthenticator := func(c *gin.Context) (interface{}, error) {
		var authLoginRequest schema.AuthLoginRequest
		if err := c.ShouldBind(&authLoginRequest); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		userName := authLoginRequest.UserName
		password := authLoginRequest.Password

		status, user := service.CheckAdminAuth(userName, password)
		if status {
			return user, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
	adminPayloadFunc := func(data interface{}) jwt.MapClaims {
		if v, ok := data.(model.Admin); ok {
			payload, err := json.Marshal(v)
			if err != nil {
				log.Panicln(err)
			}
			return jwt.MapClaims{
				"payload": string(payload),
			}
		}
		return jwt.MapClaims{}
	}
	adminHandler := func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		user := model.Admin{}
		err := json.Unmarshal([]byte(claims["payload"].(string)), &user)
		if err != nil {
			log.Panicln(err)
		}
		return user
	}
	return getJWTMiddleware(authorizator, adminAuthenticator, adminPayloadFunc, adminHandler)
}

func getJWTUserMiddleWare(authorizator JWTAuthorizator) *jwt.GinJWTMiddleware {
	adminAuthenticator := func(c *gin.Context) (interface{}, error) {
		var authLoginRequest schema.AuthLoginRequest
		if err := c.ShouldBind(&authLoginRequest); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		userName := authLoginRequest.UserName
		password := authLoginRequest.Password

		status, user := service.CheckUserAuth(userName, password)
		if status {
			return user, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
	adminPayloadFunc := func(data interface{}) jwt.MapClaims {
		if v, ok := data.(model.User); ok {
			payload, err := json.Marshal(v)
			if err != nil {
				log.Panicln(err)
			}
			return jwt.MapClaims{
				"payload": string(payload),
			}
		}
		return jwt.MapClaims{}
	}
	adminHandler := func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		user := model.User{}
		err := json.Unmarshal([]byte(claims["payload"].(string)), &user)
		if err != nil {
			log.Panicln(err)
		}
		return user
	}
	return getJWTMiddleware(authorizator, adminAuthenticator, adminPayloadFunc, adminHandler)
}

func JWTUserAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTUserMiddleWare(func(data interface{}, c *gin.Context) bool {
		return true
	})
}

func JWTBaseAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTAdminMiddleWare(func(data interface{}, c *gin.Context) bool {
		return true
	})
}

func JWTReceptionistAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTAdminMiddleWare(func(data interface{}, c *gin.Context) bool {
		user := data.(model.Admin)
		return user.IsReceptionist()
	})
}

func JWMaintainerAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTAdminMiddleWare(func(data interface{}, c *gin.Context) bool {
		user := data.(model.Admin)
		return user.IsMaintainer()
	})
}

func JWTAccountingAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTAdminMiddleWare(func(data interface{}, c *gin.Context) bool {
		user := data.(model.Admin)
		return user.IsAccounting()
	})
}

func JWTSuperUserAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTAdminMiddleWare(func(data interface{}, c *gin.Context) bool {
		user := data.(model.Admin)
		return user.IsSuperUser()
	})
}
