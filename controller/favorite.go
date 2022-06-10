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
	userId, err := getUserId(c)
	if err != nil {
		logx.DyLogger.Errorf("Can't get userId from token")
		c.JSON(http.StatusOK, api.Response{StatusCode: api.TokenInvalidErr, StatusMsg: api.ErrorCodeToMsg[api.TokenInvalidErr]})
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
		if err.Error() == "actionType Error" {
			c.JSON(http.StatusOK, api.Response{
				StatusCode: api.UnKnownActionType,
				StatusMsg:  api.ErrorCodeToMsg[api.UnKnownActionType],
			})
		} else if err.Error() == "no Such Record" {
			c.JSON(http.StatusOK, api.Response{
				StatusCode: api.RecordNotExistErr,
				StatusMsg:  api.ErrorCodeToMsg[api.RecordNotExistErr],
			})
		} else {
			c.JSON(http.StatusOK, api.Response{StatusCode: api.InnerErr, StatusMsg: api.ErrorCodeToMsg[api.InnerErr] + ":" + err.Error()})
		}
	}
}

// FavoriteList 传递给前端被登录用户点赞的所有视频
func FavoriteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		logx.DyLogger.Errorf("Can't get userId from token")
		c.JSON(http.StatusOK, api.Response{StatusCode: api.TokenInvalidErr, StatusMsg: api.ErrorCodeToMsg[api.TokenInvalidErr]})
		return
	}
	videoList, err := service.FavoriteListInfo(userId)
	if err != nil {
		logx.DyLogger.Errorf("Can't get videoList from userId")
		c.JSON(http.StatusOK, api.Response{StatusCode: api.RecordNotExistErr, StatusMsg: api.ErrorCodeToMsg[api.RecordNotExistErr]})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: *videoList,
	})
}
