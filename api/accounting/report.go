package accounting

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
)

// @Summary Get report.
// @Security ApiKeyAuth
// @Param start query string false "Start date."
// @Param end query string false "End date."
// @Produce json
// @Success 200 {array} model.Report
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /accounting/report [get]
func GetReport(c *gin.Context) {
	startTime, endTime, err := service.GetQueryTime(c)
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
		return
	}
	result, err := service.GetResultsByTime(startTime, endTime)
	if err != nil {
		schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func BindRouters(router *gin.RouterGroup) {
	accountingGroup := router.Group("/accounting")
	//accountingGroup.Use(middleware.JWTAccountingAuthenticator().MiddlewareFunc())
	accountingGroup.GET("/report", GetReport)
}