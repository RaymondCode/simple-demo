package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"strconv"
	"sync/atomic"
)

func getUserId(c *gin.Context) (int64, error) {
	userIdInterface, _ := c.Get("user_id")
	userIdFromQueryStr := c.Query("user_id")
	userId, ok := userIdInterface.(int64)
	if !ok {
		errMsg := fmt.Sprintf("user_id(%v) from context is not int", userIdInterface)
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  errMsg})
		return -1, errors.New(errMsg)
	}
	userIdFromQuery, err := strconv.Atoi(userIdFromQueryStr)
	if err != nil {
		errMsg := fmt.Sprintf("strconv.Atoi error: %s", err)
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  errMsg})
		return -1, errors.New(errMsg)
	}
	if userId != int64(userIdFromQuery) {
		errMsg := "请求参数中UID和token解析得到的UID不一致！"
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.UserIdNotMatchErr,
			StatusMsg:  errMsg})
		return -1, errors.New(errMsg)
	}

	return userId, nil
}

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]api.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	api.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	api.Response
	User api.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := api.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId, err := strconv.ParseInt(token, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: api.Response{StatusCode: 2, StatusMsg: "Can't get UserId from token"},
		})
		return
	}

	if user, err := service.NewUserServiceInstance().GetUserFromUserId(userId); err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "Something goes wrong"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: api.Response{StatusCode: 0},
			User:     *user,
		})
	}

	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: api.Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
