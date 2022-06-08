package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"strconv"
	"time"
)




// Feed 推送视频流
func  Feed(c *gin.Context) {
	latestTime := toTimeString(c.Query("latest_time"))
	resp, err := service.Feed(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, resp.Response)
		return
	}

	c.JSON(http.StatusOK, resp)
}


func toTimeString(sec string) string {
	if len(sec) == 0 {
		return time.Now().Format("2022-05-22 14:01:05")
	}
	t, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		return time.Now().Format("2022-05-22 14:01:05")
	}

	return time.Unix(t/1000, 0).Format("2022-05-22 14:01:05")
}
