package controller

import (
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"net/http"
	"strconv"

	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	// token := c.Query("token")

	user_id_from_c, _ := c.Get("user_id")
	user_id, _ := user_id_from_c.(int64)
	to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)

	err := service.FollowAction(c, user_id, to_user_id, int(action_type))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusAccepted, dto.Response{
			StatusCode: 1,
			StatusMsg:  "操作成功",
		})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	followList, err := service.GetFollowList(c, userId, 1)

	if err == nil {
		c.JSON(http.StatusAccepted, dto.UserListResponse{
			Response: dto.Response{
				StatusCode: 0,
				StatusMsg:  "查找成功",
			},
			UserList: followList,
		})
	} else {
		c.JSON(http.StatusBadRequest, dto.UserListResponse{
			Response: dto.Response{
				StatusCode: 1,
				StatusMsg:  "查找失败",
			},
			UserList: nil,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	toUserId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	followerList, err := service.GetFollowList(c, toUserId, 2)

	if err == nil {
		c.JSON(http.StatusAccepted, dto.UserListResponse{
			Response: dto.Response{
				StatusCode: 0,
				StatusMsg:  "查找成功",
			},
			UserList: followerList,
		})
	} else {
		c.JSON(http.StatusBadRequest, dto.UserListResponse{
			Response: dto.Response{
				StatusCode: 1,
				StatusMsg:  "查找失败",
			},
			UserList: nil,
		})
	}
}
