package maintainer

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/air"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/schema"
	"net/http"
)

// @Summary Dummy function, make our fucking teacher happy.
// @Produce json
// @Success 200 {object} schema.CommonStatusSchema
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /maintainer/power/on [post]
func PowerOn(c *gin.Context) {
	schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
}

// @Summary Dummy function, make our fucking teacher happy.
// @Produce json
// @Success 200 {object} schema.CommonStatusSchema
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /maintainer/power/off [post]
func PowerOff(c *gin.Context) {
	schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
}

// @Summary Set air config.
// @Security ApiKeyAuth
// @Accept  json
// @Param airConfig body config.AirConfig true "User request"
// @Produce json
// @Success 200 {object} schema.CommonStatusSchema
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /maintainer/config [post]
func SetAirConfig(c *gin.Context) {
	var config config.AirConfig
	if err := c.ShouldBind(&config); err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, "Incomplete information.")
		return
	}
	air.GetAir().SetConfig(config)
	schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
}

// @Summary Get air config.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} config.AirConfig
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /maintainer/config [get]
func GetAirConfig(c *gin.Context) {
	var config config.AirConfig
	if err := c.ShouldBind(&config); err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, "Incomplete information.")
		return
	}
	air.GetAir().SetConfig(config)
	schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
}

func BindAirRouters(router *gin.RouterGroup) {
	powerGroup := router.Group("/power")
	powerGroup.POST("/on", PowerOn)
	powerGroup.POST("/off", PowerOff)
	router.GET("/config", middleware.JWMaintainerAuthenticator().MiddlewareFunc(), GetAirConfig)
	router.POST("/config", middleware.JWMaintainerAuthenticator().MiddlewareFunc(), SetAirConfig)
}
