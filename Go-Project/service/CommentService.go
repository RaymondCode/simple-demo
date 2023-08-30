package service

import (
	"fmt"
	"github.com/life-studied/douyin-simple/dao"
	"github.com/life-studied/douyin-simple/model"
	"github.com/life-studied/douyin-simple/response"
	"net/http"
	"time"
)

func IntTime2CommentTime(intTime int64) string {

	template := "01-02"
	return time.Unix(intTime, 0).Format(template)
}

func IntTime2StrTime(intTime int64) string {
	template := "2006-01-02 15:04:05"
	return time.Unix(intTime, 0).Format(template)
}

func CreateComment(userID int64, videoID int64, commentText string) (response.CommentActionResponse, error) {
	// 获取评论时间
	currentTime := time.Now().Unix()
	user, err := dao.QueryUserById(userID)
	if err != nil {
		// 处理videoID解析错误
		commentActionResponse := response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "用户查询异常"},
		}
		return commentActionResponse, err
	}

	// 发布评论
	// 创建comment结构体
	comment := model.Comment{
		UserId:     userID,
		VideoId:    videoID,
		User:       *user,
		Content:    commentText,
		CreateDate: IntTime2StrTime(currentTime),
	}
	//  将comment增添到数据库中
	tx := dao.BeginTransaction()
	err = dao.CreateComment(&comment)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("添加评论异常：%s", err)

		commentActionResponse := response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
		}
		return commentActionResponse, err

	}
	// 更新视频表评论总数+1
	err = dao.InCreCommentCount(videoID, 1)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("更新评论总数异常：%s", err)

		commentActionResponse := response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "更新视频评论数异常"},
		}
		return commentActionResponse, err

	}

	// 创建Comment_Response响应结构体
	createDate := IntTime2CommentTime(currentTime)
	commenter, err := dao.GetUserById(userID)
	if err != nil {
		// 如果发生错误，将数据库回滚到未添加评论的初始状态
		defer dao.RollbackTransaction(tx)
		fmt.Printf("获取用户异常：%s", err)

		commentActionResponse := response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
		}
		return commentActionResponse, err
	}
	// 返回响应
	commentActionResponse := response.CommentActionResponse{
		Response: response.Response{StatusCode: 0, StatusMsg: "OK"},
		Comment: model.Comment{
			ID:         comment.ID,
			Content:    comment.Content,
			CreateDate: createDate,
			User: model.User{
				Id:   commenter.Id,
				Name: commenter.Name,
			},
		},
	}
	return commentActionResponse, err
}

func DeleteComment(userID int64, videoID int64, commentID int64) (response.CommentActionResponse, error) {
	var commentActionResponse response.CommentActionResponse
	// 删除评论
	//  根据commentID在数据库中找到待删除的评论

	//  判断是否有权限删除
	// 		通过commentID找到commenterID
	comment, err := dao.GetCommentById(commentID)
	if err != nil {
		fmt.Printf("获取评论异常：%s", err)

		commentActionResponse = response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "获取评论异常"},
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

			commentActionResponse = response.CommentActionResponse{
				Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "删除评论异常"},
			}

			return commentActionResponse, err
		}
		// 更新视频表评论总数-1
		err = dao.DeCreCommentCount(videoID, -1)
		if err != nil {
			// 如果发生错误，将数据库回滚到未删除评论的初始状态
			defer dao.RollbackTransaction(tx)
			fmt.Printf("更新视频评论数异常：%s", err)

			commentActionResponse = response.CommentActionResponse{
				Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "更新视频评论数异常"},
			}

			return commentActionResponse, err
		}

		commentActionResponse = response.CommentActionResponse{
			Response: response.Response{StatusCode: 0, StatusMsg: "删除成功"},
			Comment:  model.Comment{},
		}

	} else {
		commentActionResponse = response.CommentActionResponse{
			Response: response.Response{StatusCode: 0, StatusMsg: "无删除权限"},
			Comment:  model.Comment{},
		}
	}
	return commentActionResponse, err
}

func GetCommentList(videoID int64) ([]model.Comment, error) {

	var rawComments []model.Comment
	var err error
	rawComments, err = dao.QueryCommentsByVideoId(videoID)
	if err != nil {
		return nil, err
	}
	return rawComments, nil

}
