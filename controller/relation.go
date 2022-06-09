package controller

import (
	"net/http"
	"strconv"

	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	// token := c.Query("token")

	user_id, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)

	err := service.FollowAction(c, uint(user_id), uint(to_user_id), int(action_type))

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusAccepted, Response{
			StatusCode: 1,
			StatusMsg:  "操作成功",
		})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
