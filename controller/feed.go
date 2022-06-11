package controller

import (
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, dto.FeedResponse{
		Response: dto.Response{StatusCode: 0,
			StatusMsg: "",
			NextTime:  time.Now().Unix()},
		VideoList: DemoVideos,
	})
}
