package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/helper"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId   int64  `json:"user_id,omitempty"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

type UserResponse struct {
	Response
	UserId        uint   `json:"user_id,omitempty"`
	Token         string `json:"token"`
	Username      string `json:"username"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userInfo := model.User{UserName: username, Password: password}
	password = helper.GetMd5(password)
	token, _ := helper.GenerateToken(userInfo.UserName, userInfo.Password)

	_, err := service.GetUserByName(userInfo.UserName)
	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		userid, _ := service.CreateUser(&userInfo)

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userid,
			Token:    token,
			Username: username,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userInfo := model.User{UserName: username, Password: password}
	userid, err := service.UserLogin(&userInfo)

	password = helper.GetMd5(password)
	token, _ := helper.GenerateToken(userInfo.UserName, userInfo.Password)

	if err == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "User exist"},
			UserId:   userid,
			Token:    token,
			Username: username,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	fanid := c.Query("fan_id")
	token := c.Query("token")

	UID, _ := strconv.ParseUint(userid, 10, 32)
	FID, _ := strconv.ParseUint(fanid, 10, 32)

	fanInfo := model.User{UserID: uint(FID)}
	userInfo, err := service.GetUserByID(uint(UID))
	followercount, _ := service.GetFanCount(uint(UID))
	followcount, _ := service.GetFollowCount(uint(UID))
	isfollow, _ := service.IsFollow(fanInfo.UserID, userInfo.UserID)
	if err == nil && UID != 0 {
		c.JSON(http.StatusOK, UserResponse{
			Response:      Response{StatusCode: 0},
			UserId:        uint(UID),
			Token:         token,
			Username:      userInfo.UserName,
			FollowCount:   followcount,
			FollowerCount: followercount,
			IsFollow:      isfollow,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
