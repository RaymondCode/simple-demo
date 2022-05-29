package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

// Register 用户注册功能，要求用户输入用户名、密码和昵称。其中，用户名要求是唯一的。
// 当请求没有携带必要的参数时，返回412。
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	if username == "" || password == "" {
		c.JSON(http.StatusPreconditionFailed, Response{StatusCode: 1, StatusMsg: "用户名或密码没有给出"})
		return
	}
	// 检查用户名存在的情况
	count := service.GroupApp.UserService.IfNameExist(username)
	if count > 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "用户名已存在"})
		return
	}
	// 如果没有携带昵称，就用用户名当昵称
	if nickname == "" {
		nickname = username
	}
	// 存db，返回用户Id
	userId := service.GroupApp.UserService.SaveUser(username, password, nickname)
	token, err := utils.GenerateToken(userId)
	// 如果无法生成token，打log记录问题并返回500
	if err != nil {
		log.Println(fmt.Sprintf("无法生成token, err: %s", err))
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "服务器异常，请联系管理员"})
		return
	}
	response.OkWithToken(userId, token, c)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusPreconditionFailed, Response{StatusCode: 1, StatusMsg: "用户名或密码没有给出"})
		return
	}
	userId, err := service.GroupApp.UserService.QueryUserByNameAndPassword(username, password)
	if err != nil || userId == 0 {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "用户名或密码有误"})
		return
	}
	token, err := utils.GenerateToken(userId)
	// 如果无法生成token，打log记录问题并返回500
	if err != nil {
		log.Println(fmt.Sprintf("无法生成token, err: %s", err))
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "服务器异常，请联系管理员"})
		return
	}
	response.OkWithToken(userId, token, c)
}

func UserInfo(c *gin.Context) {
	//token := c.Query("token")
	//
	//if user, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 0},
	//		User:     user,
	//	})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
	token := c.Query("token")
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		response.FailWithMessage("user_id参数无效", c)
	}
	user, err := service.GroupApp.UserService.QueryUser(user_id, token)
	if err != nil {
		response.FailWithMessage("获取user过程发生错误", c)
	}
	if user == nil {
		response.FailWithMessage("无法获取user_id对应的user", c)
	} else {
		response.OkWithUserInfo(*user, "success", c)
	}
}
