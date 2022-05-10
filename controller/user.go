package controller

import (
	G "github.com/NoCLin/douyin-backend-go/config/global"
	"github.com/NoCLin/douyin-backend-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var userIdSequence = int64(1)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := model.User{
			Id:   userIdSequence,
			Name: username,
		}

		newUserInfo := model.UserInfo{
			User:          newUser,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}

		usersLoginInfo[token] = newUserInfo
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0},
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
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, model.UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, model.UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func Test(c *gin.Context) {

	user := model.User{Name: "Jinzhu"}

	result := G.DB.Create(&user)
	print(result)
	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	})
}
