package controller

import (
	"github.com/warthecatalyst/douyin/logx"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/service"
)

type CommentListResponse struct {
	api.Response
	CommentList []api.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	api.Response
	Comment api.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	if err != nil {
		logx.DyLogger.Errorf("Can't get userId from token")
		c.JSON(http.StatusOK, api.Response{StatusCode: 2, StatusMsg: "Can't get userId from token"})
		return
	}
	vId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(vId, 10, 64)
	//global.DyLogger.Print(videoId)
	actp := c.Query("action_type")
	actionType, _ := strconv.ParseInt(actp, 10, 32)
	content := c.Query("comment_text")
	comId := c.DefaultQuery("comment_id", "0")
	commentId, _ := strconv.ParseInt(comId, 10, 64)

	err = service.CommentActionInfo(commentId, userId, videoId, int32(actionType), content)
	if err == nil {
		c.JSON(http.StatusOK, api.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "Something goes wrong"})
	}
}

// CommentList all videos have same demo comment list
// CommentList 传递给前端 某一个视频的所有评论
func CommentList(c *gin.Context) {
	vId := c.Query("video_id")
	videoId, _ := strconv.ParseInt(vId, 10, 64)

	commentlist, err := service.CommentListInfo(videoId)

	if err != nil {
		logx.DyLogger.Errorf("Can't get CommentList from videoId")
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "Can't get CommentList from videoId"})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		CommentList: *commentlist,
	})
}
