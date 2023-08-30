package dao

import (
	"errors"
	"fmt"
	"github.com/life-studied/douyin-simple/global"
	"github.com/life-studied/douyin-simple/model"
	"gorm.io/gorm"
)

// QueryCommentsByVideoId 根据视频id查询该视频的评论列表
func QueryCommentsByVideoId(videoId int64) ([]model.Comment, error) {
	var comments []model.Comment
	if err := global.DB.Preload("User").Where("video_id = ?", videoId).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("没有找到%d这个视频的评论！", videoId)
		}
		return nil, fmt.Errorf("查询评论失败：%w", err)
	}
	return comments, nil
}

// GetUserById 根据user_id返回用户结构体
func GetUserById(userID int64) (model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetCommentById 通过commentID 返回comment结构体
func GetCommentById(commentID int64) (model.Comment, error) {
	var comment model.Comment
	err := global.DB.Where("id = ?", commentID).First(&comment).Error
	if err != nil {
		return model.Comment{}, err
	}
	return comment, nil
}

// CreateComment 创建评论
func CreateComment(comment *model.Comment) error {
	err := global.DB.Create(&comment).Error
	return err
}

// DeleteCommentById 根据id删除评论
func DeleteCommentById(commentID int64) error {
	err := global.DB.Where("id = ?", commentID).Delete(model.Comment{}).Error
	return err
}

// InCreCommentCount 增加评论数量
func InCreCommentCount(videoId int64, count int) error {
	return global.DB.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", count)).Error
}

// DeCreCommentCount 减少评论数量
func DeCreCommentCount(videoId int64, count int) error {
	return global.DB.Model(&model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", count)).Error
}

