package dao

import (
	"errors"
	"fmt"
	"github.com/life-studied/douyin-simple/global"
	"github.com/life-studied/douyin-simple/model"
	"gorm.io/gorm"
)

// CommentRepository QueryCommentsByVideoId 根据视频id查询该视频的评论列表
type Comments struct{}

func (c *Comments) QueryCommentsByVideoId(videoId int64) ([]model.Comment, error) {
	var comments []model.Comment
	if err := global.DB.Preload("User").Where("video_id = ?", videoId).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("没有找到%d这个视频的评论！", videoId)
		}
		return nil, fmt.Errorf("查询评论失败：%w", err)
	}
	return comments, nil
}
