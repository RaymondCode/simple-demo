package controller

import (
	"strconv"

	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  Response{StatusCode: 0},
	//	VideoList: DemoVideos,
	//	NextTime:  time.Now().Unix(),
	//})
	var latestTime int64
	if c.Query("latest_time") != "" {
		t, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
		if err != nil {
			response.FailWithMessage("无效的latest_time参数", c)
		}
		latestTime = t

	} else {
		latestTime = -1
	}

	token := c.Query("token")
	videos, err := service.GroupApp.FeedService.QueryFeed(latestTime, token)
	if err != nil {
		response.FailWithMessage("无法返回有效的videos", c)
	}
	response.OkWithVideoList(videos, "success", c)
}
