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

type JWTAuthenticator func(data interface{}, c *gin.Context) bool

func getJWTMiddleware(jwtAuthenticator JWTAuthenticator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       config.GetConfig().JWTConfig.Title,
		Key:         []byte(config.GetConfig().JWTConfig.Key),
		Timeout:     time.Hour * 300,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
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
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user := model.User{}
			err := json.Unmarshal([]byte(claims["payload"].(string)), &user)
			if err != nil {
				log.Panicln(err)
			}
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var authLoginRequest schema.AuthLoginRequestSchema
			if err := c.ShouldBind(&authLoginRequest); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			phone := authLoginRequest.Phone
			password := authLoginRequest.Password

			status, user := service.CheckAuth(phone, password)
			if status {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: jwtAuthenticator,
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

func JWTBaseAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTMiddleware(func(data interface{}, c *gin.Context) bool {
		return true
	})
}

func JWTReceptionistAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTMiddleware(func(data interface{}, c *gin.Context) bool {
		user := data.(model.User)
		return user.IsReceptionist()
	})
}

func JWMaintainerAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTMiddleware(func(data interface{}, c *gin.Context) bool {
		user := data.(model.User)
		return user.IsMaintainer()
	})
}

func JWTAccountingAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTMiddleware(func(data interface{}, c *gin.Context) bool {
		user := data.(model.User)
		return user.IsAccounting()
	})
}

func JWTSuperUserAuthenticator() *jwt.GinJWTMiddleware {
	return getJWTMiddleware(func(data interface{}, c *gin.Context) bool {
		user := data.(model.User)
		return user.IsSuperUser()
	})
}
