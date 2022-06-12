package model

import (
	"context"
	"fmt"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"gorm.io/gorm"
)

type Comment struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	VideoId   int64
	Content   string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateComment Comment info
func CreateComment(ctx context.Context, videoId int64, comment *Comment) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Table("comment").WithContext(ctx).Create(comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	//addVidelist commentcount
	if err := tx.Table("video").WithContext(ctx).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count+?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// QueryComment query list of Comment info
func QueryComment(ctx context.Context, videoId int64, limit, offset int) ([]*dto.ResponeComment, int64, error) {
	var total int64
	var res []*Comment
	var conn *gorm.DB
	var responeComment []*dto.ResponeComment
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return responeComment, total, err
	}
	conn = tx.Table("comment").WithContext(ctx).Model(&Comment{}).Where("video_id = ? and status = 1 ", videoId)

	if err := conn.Count(&total).Error; err != nil {
		tx.Rollback()
		return responeComment, total, err
	}
	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		tx.Rollback()
		return responeComment, total, err
	}
	responeComment = make([]*dto.ResponeComment, len(res))
	for index, v := range res {
		userInfo, _ := QueryUserById(context.Background(), v.UserId)
		fmt.Println("user:", userInfo.Name)
		users := dto.User{
			Id:            userInfo.ID,
			Name:          userInfo.Name,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			IsFollow:      false,
		}
		responeComment[index] = &dto.ResponeComment{
			ID:        v.ID,
			User:      users,
			Content:   v.Content,
			CreatedAt: v.CreatedAt.Format("2006-01-02"),
		}
	}
	return responeComment, total, nil
}

// DeleteComment delete comment info
func DeleteCommnet(ctx context.Context, videoId int64, commentId int64) error {
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Table("comment").WithContext(ctx).Where("id = ?  ", commentId).Update("status", 2).Error; err != nil {
		tx.Rollback()
		return err
	}

	//addVidelist commentcount
	if err := tx.Table("video").WithContext(ctx).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count-?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}
