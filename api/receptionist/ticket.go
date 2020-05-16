package receptionist

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/db"
	"github.com/name1e5s/acdc/middleware"
	"github.com/name1e5s/acdc/model"
	"github.com/name1e5s/acdc/schema"
	"github.com/name1e5s/acdc/service"
	"net/http"
)

// @Summary List all tickets by user phone.
// @Security ApiKeyAuth
// @Param phone query string false "User phone number"
// @Param page_size query integer false "Page Size""
// @Param offset query integer false "Page Count"
// @Produce json
// @Success 200 {array} model.Ticket
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /receptionist/ticket/all [get]
func GetAllTicketByPhone(c *gin.Context) {
	service.ReceptionistHandlerWrapper(c, func(c *gin.Context, userID uint) {
		pageSize, offset, err := service.GetQuerySizeAndOffset(c)
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, "Wrong query type.")
			return
		}
		service.GetAllTicketByUserID(c, userID, pageSize, offset)
	})
}

// @Summary List all tickets by user phone.
// @Security ApiKeyAuth
// @Param phone query string false "User phone number"
// @Param page_size query integer false "Page Size""
// @Param offset query integer false "Page Count"
// @Produce json
// @Success 200 {array} model.Ticket
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /receptionist/ticket/unpaid [get]
func GetUnpaidTicketByPhone(c *gin.Context) {
	service.ReceptionistHandlerWrapper(c, func(c *gin.Context, userID uint) {
		pageSize, offset, err := service.GetQuerySizeAndOffset(c)
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, "Wrong query type.")
			return
		}
		service.GetUnpaidTicketByUserID(c, userID, pageSize, offset)
	})
}

// @Summary List total fee.
// @Security ApiKeyAuth
// @Param phone query string false "User phone number"
// @Produce json
// @Success 200 {object} schema.UserTotalFeeResponse
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /receptionist/ticket/fee [get]
func GetTotalFeeByPhone(c *gin.Context) {
	service.ReceptionistHandlerWrapper(c, func(c *gin.Context, userID uint) {
		service.GetTotalFeeByUserID(c, userID)
	})
}

// @Summary Clear all unpaid tickets.
// @Security ApiKeyAuth
// @Param phone query string false "User phone number"
// @Produce json
// @Success 200 {object} schema.CommonStatusSchema
// @Failure 401 {object} schema.CommonStatusSchema
// @Failure 403 {object} schema.CommonStatusSchema
// @Router /receptionist/ticket/clear [POST]
func ClearUnpaidTicketByPhone(c *gin.Context) {
	service.ReceptionistHandlerWrapper(c, func(c *gin.Context, userID uint) {
		err := db.GetDataBase().Model(model.Ticket{}).Where("user_refer = ?", userID).Where("paid = ?", false).Update("paid", true).Error
		if err != nil {
			schema.NewCommonStatusSchema(c, http.StatusForbidden, err.Error())
			return
		}
		schema.NewCommonStatusSchema(c, http.StatusOK, "Done.")
	})
}

func BindTicketRouters(router *gin.RouterGroup) {
	ticketGroup := router.Group("/ticket")
	ticketGroup.Use(middleware.JWTReceptionistAuthenticator().MiddlewareFunc())
	ticketGroup.GET("/all", GetAllTicketByPhone)
	ticketGroup.GET("/unpaid", GetUnpaidTicketByPhone)
	ticketGroup.GET("/fee", GetTotalFeeByPhone)
	ticketGroup.POST("/clear", ClearUnpaidTicketByPhone)
}
