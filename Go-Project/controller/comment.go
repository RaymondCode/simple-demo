package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/response"
	"github.com/life-studied/douyin-simple/service"
	"net/http"
	"strconv"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	// 获取请求参数
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// 处理videoID解析错误
		c.JSON(http.StatusBadRequest, response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的video_id"},
		})
		return
	}
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")

	value, _ := c.Get("userid")
	userID, ok := value.(int64)

	if !ok {
		// 处理userid类型断言失败的情况
		c.JSON(http.StatusBadRequest, response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的userid"},
		})
		return
	}

	// 判断操作类型
	if actionType == "1" {
		commentActionResponse, _ := service.CreateComment(userID, videoID, commentText)
		c.JSON(http.StatusOK, commentActionResponse)

	} else if actionType == "2" {
		commentID, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			// 处理commentID解析错误
			c.JSON(http.StatusBadRequest, response.CommentActionResponse{
				Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的comment_id"},
			})
			return
		}
		commentActionResponse, _ := service.DeleteComment(userID, videoID, commentID)
		c.JSON(http.StatusOK, commentActionResponse)
	}

	return
}

func CommentList(c *gin.Context) {
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// 处理videoID解析错误
		c.JSON(http.StatusBadRequest, response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的video_id"},
		})
		return
	}

	// 将获取到的评论添加到commentList列表中

	comments, err := service.GetCommentList(videoID)
	// 返回response
	c.JSON(http.StatusOK, response.CommentListResponse{
		Response:    response.Response{StatusCode: 0, StatusMsg: "OK"},
		CommentList: comments,
	})
	return
}
