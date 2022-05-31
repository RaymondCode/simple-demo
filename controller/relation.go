package controller

import (
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//首先判断用户是否有效，获取请求
func RelationAction(c *gin.Context) {
	toUserIdStr := c.Query("to_user_id")
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("to_userId type is error", c)
		return
	}
	actionType := c.Query("action_type")
	if err := service.GroupApp.RelationService.RelationAction(utils.GetUserId(c), toUserId, actionType); err != nil {
		response.FailWithMessage("RelationAction failed", c)
		return
	}
	response.Ok(c)
}

//获取关注列表
func FollowList(c *gin.Context) {
	userList, err := service.GroupApp.RelationService.FollowList(utils.GetUserId(c))
	if err != nil {
		response.FailWithMessage("FollowList failed", c)
		return
	}
	response.OkWithUserList(userList, c)
}

//获取粉丝列表
func FollowerList(c *gin.Context) {
	userList, err := service.GroupApp.RelationService.FollowerList(utils.GetUserId(c))
	if err != nil {
		response.FailWithMessage("FollowerList failed", c)
		return
	}
	response.OkWithUserList(userList, c)
}
