package controller

import (
	"github.com/NoCLin/douyin-backend-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, model.CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
