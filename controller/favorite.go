package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/logx"
	"github.com/warthecatalyst/douyin/middleware"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"strconv"
)

// FavoriteAction 从前端传过来一条点赞或者取消点赞的记录
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	//通过token得到UserId，这边应该调用User的函数，此处仅为一个demo
	userId, err := middleware.GetUserIdFromToken(token)
	if err != nil {
		logx.DyLogger.Print("Can't get userId from token\n")
		c.JSON(http.StatusOK, api.Response{StatusCode: 2, StatusMsg: "Can't get userId from token"})

	}
	vId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(vId, 10, 64)
	actp := c.Query("action_type")
	actionType, _ := strconv.ParseInt(actp, 10, 32)
	err = service.FavoriteActionInfo(userId, videoId, int32(actionType))
	if err == nil {
		c.JSON(http.StatusOK, api.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "Something goes wrong"})
	}
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// FavoriteList 传递给前端被登录用户点赞的所有视频
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	//类似的，通过token获取userId
	_, err := middleware.GetUserIdFromToken(token)
	if err != nil {
		logx.DyLogger.Print("Can't get userId from token\n")
		c.JSON(http.StatusOK, api.Response{StatusCode: 2, StatusMsg: "Can't get userId from token"})
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
