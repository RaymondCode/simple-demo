package controller

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/service"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	// "zhangleidouyin": {
	// 	Id:            6,
	// 	Name:          "zhanglei",
	// 	FollowCount:   10,
	// 	FollowerCount: 5,
	// 	IsFollow:      true,
	// },
}

var userIdSequence = int64(0)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var enToken string

	//合法性校验
	err := service.IsUserLegal(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	//获取数据库所有数据
	users, err := service.RequireAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取数据失败"},
		})
		return
	}
	// 遍历查询结果,存入映射
	var idmax int64 = 0
	for _, user := range users {
		idmax = user.ID
		enToken = service.Encryption(user.Name, user.Password)
		Info := User{
			Id:   user.ID,
			Name: user.Name,
		}
		usersLoginInfo[enToken] = Info
	}
	userIdSequence = idmax

	//判断用户是否重复
	flag := service.IsUsernameExists(username)
	if flag {
		enToken = service.Encryption(username, password) //生成token并进行加密
	} else {
		c.JSON(http.StatusConflict, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "该用户名已存在"},
		})
		return
	}

	//将id加一后注册用户存入映射中
	atomic.AddInt64(&userIdSequence, 1)
	newUser := User{
		Id:   userIdSequence,
		Name: username,
	}
	usersLoginInfo[enToken] = newUser

	//存入数据库
	err = service.CreateInfo(userIdSequence, username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserListResponse{
			Response: Response{StatusCode: 1, StatusMsg: "存储用户信息失败"},
		})
		return
	}

	//返回正确响应
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   userIdSequence,
		Token:    enToken,
	})

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user, err := service.IsUserLegal(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "登录失败！请检查用户名和密码。"},
		})
		return
	}

	flag := service.IsTokenExists(token)
	if flag {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "登录失败！请检查用户名和密码。"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
		return
	}
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
