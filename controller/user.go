package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/db"
	"github.com/warthecatalyst/douyin/model"
	"github.com/warthecatalyst/douyin/rdb"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
	"strconv"
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
	if username == "" || password == "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{
				StatusCode: 1,
				StatusMsg:  "用户名和密码不能为空，请重新输入",
			},
		})
		return
	}

	userInfo := db.GetUserInfoByUserName(username)
	if userInfo != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "用户已经存在，请重新输入"},
		})
		return
	}
	userInfo, token := service.CreateUser(username, password)
	if userInfo == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "create userInfo error"},
		})
		return
	}
	//logrus.Infof("[Register] userId=%+v, token=%+v", userInfo.UserId, token)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: api.Response{StatusCode: 0},
		UserId:   userInfo.UserId,
		Token:    token,
	})
	return
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "用户名和密码不能为空，请重新输入"},
		})
		return
	}

	userInfo := db.GetUserInfoByPassword(userName, password)
	if userInfo == nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: api.Response{StatusCode: 1, StatusMsg: "用户不存在，请重新输入"},
		})
		return
	}
	token := rdb.GetToken(userInfo.UserId)
	logrus.Infof("[Login] userId=%+v, token=%+v", userInfo.UserId, token)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: api.Response{StatusCode: 0},
		UserId:   userInfo.UserId,
		Token:    token,
	})
	return
}

func UserInfo(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	userIdFromToken, _ := service.ParseToken(token)
	if userIdFromToken != userId {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "check token failed",
		})
		return
	}

	user := service.GetUser(userId, userId)
	if user == nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户不存在",
		})
		return
	}
	logrus.Infof("[UserInfo] user=%+v", *user)
	c.JSON(http.StatusOK, UserResponse{
		Response: api.Response{StatusCode: 0},
		//User:     *user,
	})
	return
}
