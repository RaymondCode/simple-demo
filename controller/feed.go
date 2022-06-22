package controller

import (
	"log"
	"net/http"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	lastTime := c.Query("latest_time")
	token := c.Query("token")
	userInfo, parseTokenErr := util.ParseToken(token)

	if parseTokenErr != nil {
		dto.WriteLog(
			"error", "解析token错误",
		)
	}

	videoList, reTime := service.QueryFeedResponse1(userInfo.ID, lastTime)
	log.Println(reTime.Unix())
	c.JSON(http.StatusOK, dto.FeedResponse{
		Response: dto.Response{StatusCode: 0,
			StatusMsg: "",
			NextTime:  reTime.Unix()},
		VideoList: videoList,
	})
}
