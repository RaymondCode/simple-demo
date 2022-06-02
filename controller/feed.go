package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"time"
)

type FeedResponse struct {
	service.Response
	VideoList []service.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  service.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
