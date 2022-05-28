package response

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

func OkWithUserInfo(user model.User, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{
			StatusCode: SUCCESS,
			StatusMsg:  msg,
		},
		User: user,
	})
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func OkWithToken(userId int64, token string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: SUCCESS,
		},
		UserId: userId,
		Token:  token,
	})
}
