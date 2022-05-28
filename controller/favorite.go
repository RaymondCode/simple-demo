package controller

import (
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	videoIdStr := c.Query("video_id")
	actionType := c.Query("action_type")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("video_id类型错误，无法转换为int", c)
	}
	if err := service.GroupApp.FavoriteService.FavoriteAction(actionType, utils.GetUserId(c), videoId); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	videoList, err := service.GroupApp.FavoriteService.FavoriteList(utils.GetUserId(c))
	if err != nil {
		response.FailWithMessage("点赞列表查询失败", c)
	} else {
		response.OkWithVideoList(videoList, "查询成功", c)
	}
}
