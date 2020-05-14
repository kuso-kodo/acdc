package service

import (
	"encoding/json"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"net/http"
)

func GetUserFromClaims(c *gin.Context) (model.User, error) {
	claims := jwt.ExtractClaims(c)
	user := model.User{}
	err := json.Unmarshal([]byte(claims["payload"].(string)), &user)
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusUnauthorized, "Invalid user token.")
	}
	return user, err
}
