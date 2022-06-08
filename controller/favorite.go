package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/logx"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"strconv"
)

// FavoriteAction 从前端传过来一条点赞或者取消点赞的记录
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	//通过token得到UserId，这边应该调用User的函数，此处仅为一个demo
	//userId, err := tokenx.GetUserIdFromToken(token)
	userId, err := strconv.ParseInt(token, 10, 64) //used for test
	if err != nil {
		logx.DyLogger.Errorf("Can't get userId from token")
		c.JSON(http.StatusOK, api.Response{StatusCode: 2, StatusMsg: "Can't get userId from token"})
		return
	}
	vId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(vId, 10, 64)
	//tokenx.DyLogger.Print(videoId)
	actp := c.Query("action_type")
	actionType, _ := strconv.ParseInt(actp, 10, 32)
	err = service.FavoriteActionInfo(userId, videoId, int32(actionType))
	if err == nil {
		c.JSON(http.StatusOK, api.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "Something goes wrong"})
	}
}

// FavoriteList 传递给前端被登录用户点赞的所有视频
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	//类似的，通过token获取userId
	// userId, err := tokenx.GetUserIdFromToken(token)
	userId, err := strconv.ParseInt(token, 10, 64) //used for test
	if err != nil {
		logx.DyLogger.Errorf("Can't get userId from token")
		c.JSON(http.StatusOK, api.Response{StatusCode: 2, StatusMsg: "Can't get userId from token"})
		return
	}
	videoList, err := service.FavoriteListInfo(userId)
	if err != nil {
		logx.DyLogger.Errorf("Can't get videoList from userId")
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "Can't get videoList from userId"})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: *videoList,
	})
}
