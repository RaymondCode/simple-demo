package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/service"
)

// Feed 推送视频流
func Feed(c *gin.Context) {
	latestTime := toTimeString(c.Query("latest_time"))
	resp, err := service.Feed(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, resp.Response)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func toTimeString(sec string) string {
	location, _ := time.LoadLocation("Asia/Shanghai")
	if len(sec) == 0 {
		return time.Now().In(location).Format("2006-01-02 15:04:05")
	}
	t, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		return time.Now().In(location).Format("2006-01-02 15:04:05")
	}

	return time.Unix(t, 0).In(location).Format("2006-01-02 15:04:05")
}
