package service

import (
	"errors"
	"fmt"
	"github.com/life-studied/douyin-simple/dao"
	"github.com/life-studied/douyin-simple/global"
	"github.com/life-studied/douyin-simple/model"
	"github.com/life-studied/douyin-simple/response"
	"gorm.io/gorm"
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

// GetFollowByUserId 检查user1是否follow了user2
func GetFollowByUserId(userId1 int64, userId2 int64) bool {
	relationship := model.Relation{}
	if userId1 == userId2 || userId1 == 0 {
		return false
	}
	if err := global.DB.Model(&model.Relation{}).Where("host_id=? And guest_id=?", userId1, userId2).First(&relationship).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func CreateComment(userID int64, videoID int64, commentText string) (response.CommentActionResponse, error) {
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

		commentActionResponse := response.CommentActionResponse{
			Response: response.Response{StatusCode: http.StatusInternalServerError, StatusMsg: "添加评论异常"},
		}
		return commentActionResponse, err

	}
	// 更新视频表评论总数+1
	err = dao.UpdateVideoCommentCount(videoID, 1)
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
		Comment: response.Comment_Response{
			ID:         comment.ID,
			Content:    comment.Content,
			CreateDate: createDate,
			Userresponse: response.User_Response{
				ID:            uint(commenter.Id),
				Name:          commenter.Name,
				FollowCount:   uint(commenter.FollowCount),
				FollowerCount: uint(commenter.FollowerCount),
				IsFollow:      false, // 待确定自己与自己的关注状态
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
		err = dao.UpdateVideoCommentCount(videoID, -1)
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
			Comment:  response.Comment_Response{},
		}

	} else {
		commentActionResponse = response.CommentActionResponse{
			Response: response.Response{StatusCode: 0, StatusMsg: "无删除权限"},
			Comment:  response.Comment_Response{},
		}
	}
	return commentActionResponse, err
}

func GetCommentList(videoID int64, userID int64) ([]response.Comment_Response, error) {

	// 从数据库中获取id为video_id的全部评论
	comments, err := dao.GetCommentByIdListById(videoID)
	if err != nil {
		err = errors.New("视频数据错误")
		return nil, err
	}

	// 将获取到的评论添加到commentList列表中
	// 将model.comment解析为response.Comment_Response格式
	// 将获取到的评论添加到commentList列表中
	var commentList []response.Comment_Response
	// 将model.comment解析为response.Comment_Response格式
	for _, comment := range comments {

		// 根据评论者id构建user_response
		commenter, err := dao.GetUserById(comment.UserId)
		if err != nil {
			// 处理获取用户信息错误
			continue // 继续处理下一个评论
		}

		// 构建Comment_Response中嵌套的User_Response字段
		userResponse := response.User_Response{
			ID:            uint(commenter.Id),
			FollowCount:   uint(commenter.FollowCount),
			FollowerCount: uint(commenter.FollowerCount),
			Name:          commenter.Name,
		}
		//查询该用户是否被关注

		userResponse.IsFollow = GetFollowByUserId(userID, commenter.Id)

		commentResponse := response.Comment_Response{
			ID:           uint(comment.Id),
			Content:      comment.Content,
			CreateDate:   IntTime2StrTime(comment.CreateDate),
			Userresponse: userResponse,
		}
		commentList = append(commentList, commentResponse)
	}
	return commentList, err
}
