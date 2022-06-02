package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"strconv"
)

const (
	followTyp       = 1
	cancelFollowTyp = 2
)

type UserListResponse struct {
	service.Response
	UserList []*service.User `json:"user_list"`
}

func getUserId(c *gin.Context) (int64, error) {
	userIdInterface, _ := c.Get("user_id")
	userIdFromQueryStr := c.Query("user_id")
	userId, ok := userIdInterface.(int64)
	if !ok {
		errMsg := fmt.Sprintf("user_id(%v) from context is not int", userIdInterface)
		c.JSON(http.StatusOK, service.Response{
			StatusCode: innerErr,
			StatusMsg:  errMsg})
		return -1, errors.New(errMsg)
	}
	userIdFromQuery, err := strconv.Atoi(userIdFromQueryStr)
	if err != nil {
		errMsg := fmt.Sprintf("strconv.Atoi error: %s", err)
		c.JSON(http.StatusOK, service.Response{
			StatusCode: innerErr,
			StatusMsg:  errMsg})
		return -1, errors.New(errMsg)
	}
	if userId != int64(userIdFromQuery) {
		errMsg := "请求参数中UID和token解析得到的UID不一致！"
		c.JSON(http.StatusOK, service.Response{
			StatusCode: userIdNotMatchErr,
			StatusMsg:  errMsg})
		return -1, errors.New(errMsg)
	}

	return userId, nil
}

func RelationAction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	actTyp := c.Query("action_type")
	actTypInt, err := strconv.Atoi(actTyp)
	if err != nil {
		c.JSON(http.StatusOK, service.Response{
			StatusCode: innerErr,
			StatusMsg:  fmt.Sprintf("strconv.Atoi error: %s", err)})
		return
	}
	toUserIdStr := c.Query("to_user_id")
	toUserId, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		c.JSON(http.StatusOK, service.Response{
			StatusCode: innerErr,
			StatusMsg:  fmt.Sprintf("strconv.Atoi error: %s", err)})
		return
	}

	if actTypInt == followTyp {
		if err := service.Follow(userId, int64(toUserId)); err != nil {
			c.JSON(http.StatusOK, service.Response{
				StatusCode: innerErr,
				StatusMsg:  fmt.Sprintf("service.Follow error: %s", err)})
			return
		}
		c.JSON(http.StatusOK, service.Response{StatusCode: 0, StatusMsg: ""})
	} else if actTypInt == cancelFollowTyp {
		if err := service.UnFollow(userId, int64(toUserId)); err != nil {
			c.JSON(http.StatusOK, service.Response{
				StatusCode: innerErr,
				StatusMsg:  fmt.Sprintf("service.UnFollow error: %s", err)})
			return
		}
		c.JSON(http.StatusOK, service.Response{StatusCode: 0, StatusMsg: ""})
	} else {
		c.JSON(http.StatusOK, service.Response{
			StatusCode: unknownFollowActionType,
			StatusMsg:  "未知关注相关操作类型！"})
		return
	}
}

func FollowList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	users, err := service.GetFollowList(userId)
	if err != nil {
		c.JSON(http.StatusOK, service.Response{
			StatusCode: innerErr,
			StatusMsg:  fmt.Sprintf("service.GetFollowList error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: service.Response{
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
		c.JSON(http.StatusOK, service.Response{
			StatusCode: innerErr,
			StatusMsg:  fmt.Sprintf("service.GetFollowerList error: %s", err)})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
