package service

import (
	"errors"
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

func GetQueryUserPhone(c *gin.Context) (string, error) {
	phone := c.DefaultQuery("phone", "")
	if phone == "" {
		return phone, errors.New("invalid phone number")
	}
	return phone, nil
}

func GetQueryRoomName(c *gin.Context) (string, error) {
	room := c.DefaultQuery("room", "")
	if room == "" {
		return room, errors.New("invalid room name")
	}
	return room, nil
}

func GetQueryPriority(c *gin.Context) (int, error) {
	level, err := strconv.Atoi(c.DefaultQuery("level", strconv.Itoa(1)))
	if err != nil {
		return 0, err
	}
	if level < 0 || level > 3 {
		return 0, errors.New("unknown level")
	}
	return level, nil
}
