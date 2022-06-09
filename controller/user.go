package controller

import (
	"context"
	"github.com/RaymondCode/simple-demo/db/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"tasdgfasdg123456": {
		Id:            "08dc2b99ef974d47a2554ed3dea73ea0",
		Name:          "tasdgfasdg",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId string `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	//Password encrypted with salt
	password, _ = service.Encryption(password)

	//QueryUser QueryUser By Name for judged user is exit or not
	user, _ := model.QueryUser(context.Background(), username)

	//judege user exit or not
	if exist := len(user) > 0; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := &model.User{
			Name:     username,
			PassWord: password,
		}
		//userinfo register
		model.CreateUser(context.Background(), newUser)
		//Query Userinfo for get id
		userinfo, _ := model.QueryUser(context.Background(), username)
		userid := userinfo[0].ID
		//token
		token := util.GenerateToken(&util.UserClaims{ID: userid, Name: username, PassWord: password})
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userid,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//Password encrypted with salt
	encryptionPassWord, _ := service.Encryption(password)

	user, _ := model.QueryUser(context.Background(), username)
	token := util.GenerateToken(&util.UserClaims{ID: user[0].ID, Name: username, PassWord: encryptionPassWord})

	if exist := len(user) > 0; exist {
		//judge password
		if service.ComparePasswords(user[0].PassWord, password) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user[0].ID,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "password wrong"},
			})
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	userinfo := util.ParseToken(token)
	users, _ := model.QueryUserById(context.Background(), userinfo.ID)
	user1 := users[0]

	user := User{
		Id:            userinfo.ID,
		Name:          userinfo.Name,
		FollowCount:   user1.FollowCount,
		FollowerCount: user1.FollowerCount,
		IsFollow:      false}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})

	//log.Printf(token)
	//log.Printf(userinfo.ID)
}
