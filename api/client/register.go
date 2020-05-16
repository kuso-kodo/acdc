package client

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"net/http"
	"time"
)

// @Summary Register client.
// @Accept  json
// @Param clientRequest body schema.ClientRegisterRequest true "Client request"
// @Produce json
// @Success 200 {object} schema.ClientRegisterResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var registerRequest schema.ClientRegisterRequest
	if err := c.BindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusOK, schema.ClientRegisterResponse{
			Status:  false,
			Period:  config.GetConfig().AirConfig.Period,
			Message: "Invalid request format.",
		})
		return
	}
	err := db.GetDataBase().Create(&model.Room{
		RoomName:   registerRequest.Room,
		LastOnTime: time.Now(),
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, schema.ClientRegisterResponse{
			Status:  false,
			Period:  config.GetConfig().AirConfig.Period,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, schema.ClientRegisterResponse{
		Status:  true,
		Period:  config.GetConfig().AirConfig.Period,
		Message: "Done.",
	})
}
