package user

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
)

// @Summary List all tickets.
// @Security ApiKeyAuth
// @Param page_size query integer false "Page Size""
// @Param offset query integer false "Page Count"
// @Produce json
// @Success 200 {array} model.Ticket
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /user/ticket/all [get]
func GetAllTicket(c *gin.Context) {
	service.UserHandlerWrapper(c, func(c *gin.Context, user model.User) {
		pageSize, offset, err := service.GetQuerySizeAndOffset(c)
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, "Wrong query type.")
			return
		}
		service.GetAllTicketByUserID(c, user.UserID, pageSize, offset)
	})
}

// @Summary List all unpaid tickets.
// @Security ApiKeyAuth
// @Param page_size query integer false "Page Size"
// @Param offset query integer false "Page Count"
// @Produce json
// @Success 200 {array} model.Ticket
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /user/ticket/unpaid [get]
func GetUnpaidTicket(c *gin.Context) {
	service.UserHandlerWrapper(c, func(c *gin.Context, user model.User) {
		pageSize, offset, err := service.GetQuerySizeAndOffset(c)
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, "Wrong query type.")
			return
		}
		service.GetUnpaidTicketByUserID(c, user.UserID, pageSize, offset)
	})
}

// @Summary List total fee.
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} schema.UserTotalFeeResponse
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /user/ticket/fee [get]
func GetTotalFee(c *gin.Context) {
	service.UserHandlerWrapper(c, func(c *gin.Context, user model.User) {
		service.GetTotalFeeByUserID(c, user.UserID)
	})
}

func BindTicketRouters(router *gin.RouterGroup) {
	ticketGroup := router.Group("/ticket")
	ticketGroup.Use(middleware.JWTUserAuthenticator().MiddlewareFunc())
	ticketGroup.GET("/all", GetAllTicket)
	ticketGroup.GET("/unpaid", GetUnpaidTicket)
	ticketGroup.GET("/fee", GetTotalFee)
}
