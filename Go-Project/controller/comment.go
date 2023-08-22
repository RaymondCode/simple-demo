package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/model"
	"github.com/life-studied/douyin-simple/service"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64) //把前端请求中的视频id查出
	if err != nil {
		c.JSON(http.StatusOK, Response{
			-1,
			"video_id参数错误",
		})
		return
	}
	token := c.Query("token") //把前端请求中的token查出

	cs := &service.CommentService{}
	DemoComments, err := cs.QueryComment(video_id, token)

	if err != nil {
		c.JSON(http.StatusOK, Response{
			-1,
			"获取评论失败",
		})
		return
	} //查询评论后判断查询是否成功

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})

}
