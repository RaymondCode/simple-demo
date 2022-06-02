package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	api.Response
	UserList []*api.User `json:"user_list"`
}

func RelationAction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	actTyp := c.Query("action_type")
	actTypInt, err := strconv.Atoi(actTyp)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  fmt.Sprintf("strconv.Atoi error: %s", err)})
		return
	}
	toUserIdStr := c.Query("to_user_id")
	toUserId, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  fmt.Sprintf("strconv.Atoi error: %s", err)})
		return
	}
	if err := service.FollowAction(userId, int64(toUserId), actTypInt); err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  fmt.Sprintf("service.FollowAction error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, api.Response{
		StatusCode: 0,
		StatusMsg:  ""})
	return
}

func FollowList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	users, err := service.GetFollowList(userId)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  fmt.Sprintf("service.GetFollowList error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

func FollowerList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	users, err := service.GetFollowerList(userId)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  fmt.Sprintf("service.GetFollowerList error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
