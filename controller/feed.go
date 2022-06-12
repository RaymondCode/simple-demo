package controller

import (
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//nextTime:=c.Query("next_time")
	token := c.Query("token")
	userInfo, _ := util.ParseToken(token)
	videoList := service.QueryFeedResponse(userInfo.ID)
	c.JSON(http.StatusOK, dto.FeedResponse{
		Response: dto.Response{StatusCode: 0,
			StatusMsg: "",
			NextTime:  time.Now().Unix()},
		VideoList: videoList,
	})
}
