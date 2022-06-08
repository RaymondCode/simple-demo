package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"net/http"
)

type CommentListResponse struct {
	api.Response
	CommentList []api.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	api.Response
	Comment api.Comment `json:"common,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	// TODO: ?
}

// CommentList all videos have same demo common list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    api.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
