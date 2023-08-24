package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/dao"
	"github.com/life-studied/douyin-simple/model"
	"net/http
	"strconv"
	"time"

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
	// 获取请求参数
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// 处理videoID解析错误
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response: Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的video_id"},
		})
		return
	}
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentID, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	if err != nil {
		// 处理commentID解析错误
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response: Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的comment_id"},
		})
		return
	}

	value, _ := c.Get("userid")
	userID, ok := value.(int64)
	if !ok {
		// 处理userid类型断言失败的情况
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response: Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的userid"},
		})
		return
	}

	// 获取评论时间
	currentTime := time.Now().Unix()

	// 判断操作类型
	if actionType == "1" {
		// 发布评论
		//  创建comment结构体
		comment := model.Comment{
			UserId:     userID,
			VideoId:    videoID,
			Content:    commentText,
			CreateDate: currentTime,
		}
		//  将comment增添到数据库中
		tx := dao.BeginTransaction()
		err := dao.CreateComment(&comment)
		if err != nil {
			// 如果发生错误，将数据库回滚到未添加评论的初始状态
			defer dao.RollbackTransaction(tx)
			fmt.Printf("添加评论异常：%s", err)

			c.JSON(http.StatusInternalServerError, CommentActionResponse{
				Response: Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
			})
			return
		}
		//  创建Comment_Response响应结构体

		//commenter, err := dao.GetUserById(userID)
		if err != nil {
			// 如果发生错误，将数据库回滚到未添加评论的初始状态
			defer dao.RollbackTransaction(tx)
			fmt.Printf("获取用户异常：%s", err)

			c.JSON(http.StatusInternalServerError, CommentActionResponse{
				Response: Response{StatusCode: http.StatusInternalServerError, StatusMsg: "获取用户异常"},
			})
			return
		}
		// 返回响应
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: Response{StatusCode: 0, StatusMsg: "OK"},
			Comment: Comment{
				Id: 0,
				//////
			},
		})

	} else if actionType == "2" {
		// 删除评论
		//  根据commentID在数据库中找到待删除的评论
		// 判断是否有权限删除
		comment, err := dao.GetCommentById(commentID)
		if err != nil {
			fmt.Printf("获取评论异常：%s", err)

			c.JSON(http.StatusInternalServerError, CommentActionResponse{
				Response: Response{StatusCode: http.StatusInternalServerError, StatusMsg: "获取评论异常"},
			})
			return
		}
		commenterID := comment.UserId
		//  若有权限，则删除id为commentID评论;若无权限，则拒绝删除
		if commenterID == userID {
			tx := dao.BeginTransaction()
			err = dao.DeleteCommentById(commentID)
			if err != nil {
				// 如果发生错误，将数据库回滚到未删除评论的初始状态
				defer dao.RollbackTransaction(tx)
				fmt.Printf("删除评论异常：%s", err)

				c.JSON(http.StatusInternalServerError, CommentActionResponse{
					Response: Response{StatusCode: http.StatusInternalServerError, StatusMsg: "删除评论异常"},
				})
				return
			}
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "删除成功"},
				Comment:  Comment{},
			})
		} else {
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{StatusCode: 0, StatusMsg: "无删除权限"},
				Comment:  Comment{},
			})
		}
	}

	return
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		// 处理videoID解析错误
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response: Response{StatusCode: http.StatusBadRequest, StatusMsg: "无效的video_id"},
		})
		return
	} // 从数据库中获取id为video_id的全部评论
	comments, err := dao.GetCommentByIdListById(videoID)
	if err != nil {
		// 处理数据库查询错误
		c.JSON(http.StatusInternalServerError, CommentListResponse{
			Response: Response{
				StatusCode: http.StatusInternalServerError,
				StatusMsg:  "获取评论失败",
			},
			CommentList: nil,
		})
		return
	}

	// 将获取到的评论添加到commentList列表中
	var commentList []Comment

	for _, comment := range comments {

		commentResponse := Comment{
			Id:      comment.Id,
			Content: comment.Content,
		}
		commentList = append(commentList, commentResponse)
	}

	// 返回response
	c.JSON(http.StatusOK,
		CommentListResponse{
			Response:    Response{StatusCode: 0, StatusMsg: "OK"},
			CommentList: commentList,
		})
	return

