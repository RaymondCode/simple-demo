package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/service"
	"net/http"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	strLatestTime, exists := c.GetQuery("latest_time")
	if !exists {
		c.JSON(http.StatusBadRequest, FeedResponse{Response: Response{StatusCode: 1, StatusMsg: "latest_time is empty"}})
		return
	}
	nextTime, startId, err := service.GetNextTime(strLatestTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, FeedResponse{Response: Response{StatusCode: 1, StatusMsg: err.Error()}})
		return
	}
	var endId int64
	if startId+30 > int64(len(DemoVideos)) {
		endId = int64(len(DemoVideos))
	} else {
		endId = startId + 30
	}
	c.JSON(http.StatusOK, FeedResponse{Response: Response{StatusCode: 0},
		VideoList: DemoVideos[startId:endId],
		NextTime:  nextTime,
	})
}
