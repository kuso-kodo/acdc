package user

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/config"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
	"strconv"
)

// @Summary List all tickets.
// @Security ApiKeyAuth
// @Param page_size query integer false "Page Size""
// @Param offset query integer false "Page Count"
// @Produce json
// @Success 200 {array} model.Ticket
// @Failure 401 {object} schema.CommonFailureSchema
// @Failure 403 {object} schema.CommonFailureSchema
// @Router /user/ticket/all [get]
func GetAllTicket(c *gin.Context) {
	user, err := service.GetUserFromClaims(c)
	if err != nil {
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.GetConfig().TicketConfig.PageSize)))
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, "Wrong query type.")
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", strconv.Itoa(1)))
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, "Wrong query type.")
		return
	}
	var tickets []model.Ticket
	err = db.GetDataBase().Where("user_refer = ?", user.UserID).Limit(pageSize).Offset((offset - 1) * pageSize).Find(&tickets).Error
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, tickets)
}

// @Summary List all tickets.
// @Security ApiKeyAuth
// @Param page_size query integer false "Page Size""
// @Param offset query integer false "Page Count"
// @Produce json
// @Success 200 {array} model.Ticket
// @Failure 401 {object} schema.CommonFailureSchema
// @Failure 403 {object} schema.CommonFailureSchema
// @Router /user/ticket/unpaid [get]
func GetUnpaidTicket(c *gin.Context) {
	user, err := service.GetUserFromClaims(c)
	if err != nil {
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.GetConfig().TicketConfig.PageSize)))
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, "Wrong query type.")
		return
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", strconv.Itoa(1)))
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, "Wrong query type.")
		return
	}
	var tickets []model.Ticket
	err = db.GetDataBase().Where("user_refer = ?", user.UserID).Where("paid = ?", false).Limit(pageSize).Offset((offset - 1) * pageSize).Find(&tickets).Error
	if err != nil {
		schema.NewCommonFailureSchema(c, http.StatusForbidden, err.Error())
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func BindTicketRouters(router *gin.RouterGroup) {
	ticketGroup := router.Group("/ticket")
	ticketGroup.Use(middleware.JWTUserAuthenticator().MiddlewareFunc())
	ticketGroup.GET("/all", GetAllTicket)
	ticketGroup.GET("/unpaid", GetUnpaidTicket)
}
