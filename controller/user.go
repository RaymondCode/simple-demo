package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

var userIdSequence = int64(1)

type UserIdTokenResponse struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserRegisterLoginResponse struct {
	Response
	UserIdTokenResponse
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userIdTokenResponse, err := RegisterService(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserRegisterLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	} else {
		c.JSON(http.StatusOK, UserRegisterLoginResponse{
			Response:            Response{StatusCode: 0},
			UserIdTokenResponse: userIdTokenResponse,
		})
	}
}

func RegisterService(username string, password string) (UserIdTokenResponse, error) {
	var userIdTokenResponse = UserIdTokenResponse{}

	err := service.IsUserLegal(username, password)
	if err != nil {
		return userIdTokenResponse, err
	}

	userId, err := service.CreateRegisterUser(username, password)
	if err != nil {
		return userIdTokenResponse, err
	}

	token, err := GenerateToken(username)
	if err != nil {
		return userIdTokenResponse, err
	}

	userIdTokenResponse = UserIdTokenResponse{
		UserId: userId,
		Token:  token,
	}
	return userIdTokenResponse, nil
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	userIdTokenResponse, err := LoginService(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserRegisterLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	} else {
		c.JSON(http.StatusOK, UserRegisterLoginResponse{
			Response:            Response{StatusCode: 0},
			UserIdTokenResponse: userIdTokenResponse,
		})
	}
}

func LoginService(username string, password string) (UserIdTokenResponse, error) {
	var userIdTokenResponse = UserIdTokenResponse{}

	err := service.IsUserLegal(username, password)
	if err != nil {
		return userIdTokenResponse, err
	}

	userId, err := service.FindLoginUser(username, password)
	if err != nil {
		return userIdTokenResponse, err
	}

	token, err := GenerateToken(username)
	if err != nil {
		return userIdTokenResponse, err
	}

	userIdTokenResponse = UserIdTokenResponse{
		UserId: userId,
		Token:  token,
	}
	return userIdTokenResponse, nil
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
