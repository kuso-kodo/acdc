package schema

import "github.com/gin-gonic/gin"

type CommonFailureSchema struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCommonFailureSchema(c *gin.Context, status int, message string) {
	response := CommonFailureSchema{
		Code:    status,
		Message: message,
	}
	c.JSON(status, response)
}
