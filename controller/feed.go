package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList    []Video `json:"video_list,omitempty"`
	NextTime     int64   `json:"next_time,omitempty"`
	CommentCount int64   `json:"comment_count,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0,
			StatusMsg: "",
			NextTime:  time.Now().Unix()},
		VideoList: DemoVideos,
	})
}
