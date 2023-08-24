package service

import (
	"fmt"
	"github.com/life-studied/douyin-simple/controller"
	"github.com/life-studied/douyin-simple/dao"
	"github.com/life-studied/douyin-simple/model"
	"net/http"
	"time"
)

func IntTime2CommentTime(intTime int64) string {

	template := "01-02"
	return time.Unix(intTime, 0).Format(template)
}

func CreateComment(userID int64, videoID int64, commentText string) (controller.CommentActionResponse, error) {
	// 获取评论时间
	currentTime := time.Now().Unix()
	// 发布评论
	// 创建comment结构体
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

		commentActionResponse := controller.CommentActionResponse{
			Response: controller.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
		}
		return commentActionResponse, err

	}
	// 更新视频表评论总数+1
	//err = dao.UpdateVideoCommentCount(videoID, 1)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("更新评论总数异常：%s", err)

		commentActionResponse := controller.CommentActionResponse{
			Response: controller.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "更新视频评论数异常"},
		}
		return commentActionResponse, err

	}

	// 创建Comment_Response响应结构体
	createDate := IntTime2CommentTime(currentTime)
	//commenter, err := dao.GetUserById(userID)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("获取用户异常：%s", err)

		commentActionResponse := controller.CommentActionResponse{
			Response: controller.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
		}
		return commentActionResponse, err
	}
	// 返回响应
	commentActionResponse := controller.CommentActionResponse{
		Response: controller.Response{StatusCode: 0, StatusMsg: "OK"},
		Comment: controller.Comment{
			Id: 0,

			Content:    commentText,
			CreateDate: createDate,
		},
	}
	return commentActionResponse, err
}

func DeleteComment(userID int64, videoID int64, commentID int64) (controller.CommentActionResponse, error) {
	var commentActionResponse controller.CommentActionResponse
	// 删除评论
	//  根据commentID在数据库中找到待删除的评论

	//  判断是否有权限删除
	// 		通过commentID找到commenterID
	comment, err := dao.GetCommentById(commentID)
	if err != nil {
		fmt.Printf("获取评论异常：%s", err)

		commentActionResponse = controller.CommentActionResponse{
			Response: controller.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "获取评论异常"},
		}
		return commentActionResponse, err
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

			commentActionResponse = controller.CommentActionResponse{
				Response: controller.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "删除评论异常"},
			}

			return commentActionResponse, err
		}
		// 更新视频表评论总数-1
		//err = dao.UpdateVideoCommentCount(videoID, -1)
		if err != nil {
			// 如果发生错误，将数据库回滚到未删除评论的初始状态
			defer dao.RollbackTransaction(tx)
			fmt.Printf("更新视频评论数异常：%s", err)

			commentActionResponse = controller.CommentActionResponse{
				Response: controller.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "更新视频评论数异常"},
			}

			return commentActionResponse, err
		}

		commentActionResponse = controller.CommentActionResponse{
			Response: controller.Response{StatusCode: 0, StatusMsg: "删除成功"},
			Comment:  controller.Comment{},
		}

	} else {
		commentActionResponse = controller.CommentActionResponse{
			Response: controller.Response{StatusCode: 0, StatusMsg: "无删除权限"},
			Comment:  controller.Comment{},
		}
	}
	return commentActionResponse, err
}

/*func GetCommentList(videoID int64, userID int64) ([]controller.Comment, error) {
}*/
