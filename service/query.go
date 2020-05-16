package service

import (
	"github.com/gin-gonic/gin"
	"github.com/name1e5s/acdc/config"
	"strconv"
)

func GetQuerySizeAndOffset(c *gin.Context) (int, int, error) {
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", strconv.Itoa(config.GetConfig().TicketConfig.PageSize)))
	if err != nil {
		return 0, 0, err
	}
	offset, err := strconv.Atoi(c.DefaultQuery("offset", strconv.Itoa(1)))
	if err != nil {
		return 0, 0, err
	}
	return pageSize, offset, nil
}
