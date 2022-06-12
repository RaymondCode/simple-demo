package controller

import (
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.Query("latest_time")
	token := c.Query("token")
	userInfo, _ := util.ParseToken(token)
	videoList, reTime := service.QueryFeedResponse(userInfo.ID, lastTime)
	c.JSON(http.StatusOK, dto.FeedResponse{
		Response: dto.Response{StatusCode: 0,
			StatusMsg: "",
			NextTime:  reTime.Unix()},
		VideoList: videoList,
	})
}
