package schema

import "github.com/gin-gonic/gin"

type CommonStatusSchema struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCommonStatusSchema(c *gin.Context, status int, message string) {
	response := CommonStatusSchema{
		Code:    status,
		Message: message,
	}
	c.JSON(status, response)
}
