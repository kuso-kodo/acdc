package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	var users []model.User
	db.GetDataBase().Find(&users)
	c.JSON(http.StatusOK, users)
}
