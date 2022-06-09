package controller

import (
	"context"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
	"log"
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
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User dto.User `json:"user"`
}

func Register(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")

	//Password encrypted with salt
	password, _ = service.Encryption(password)

	//QueryUser QueryUser By Name for judged user is exit or not
	user, _ := model.QueryUserByName(context.Background(), username)

	//judege user exit or not
	if user.Name != "" {
		log.Printf(user.Name, user.ID)
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
		userInfo, _ := model.QueryUserByName(context.Background(), username)
		//token
		token, _ := util.GenerateToken(&util.UserClaims{ID: userInfo.ID, Name: username, PassWord: password})
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userInfo.ID,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//Password encrypted with salt
	encryptionPassWord, _ := service.Encryption(password)

	user, _ := model.QueryUserByName(context.Background(), username)
	token, _ := util.GenerateToken(&util.UserClaims{ID: user.ID, Name: username, PassWord: encryptionPassWord})

	if user != nil {
		//judge password
		if service.ComparePasswords(user.PassWord, password) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.ID,
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

	userClaims, _ := util.ParseToken(token)
	userModel, _ := model.QueryUserById(context.Background(), userClaims.ID)

	user := dto.User{
		Id:            userModel.ID,
		Name:          userModel.Name,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,
		IsFollow:      false,
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})

	//log.Printf(token)
	//log.Printf(userinfo.ID)
}
