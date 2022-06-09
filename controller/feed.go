package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"net/http"
	"time"
)

type FeedResponse struct {
	api.Response
	VideoList []api.Video `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

// Feed 推送视频流
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  api.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
