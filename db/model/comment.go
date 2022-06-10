package model

import (
	"context"
	"fmt"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"gorm.io/gorm"
)

type Comment struct {
	ID      int64 `gorm:"primarykey"`
	UserId  int64
	VideoId int64
	Content string
	// Status    int
	CreatedAt time.Time
	// UpdatedAt time.Time
}

type ResponeComment struct {
	ID        int64     `json:"id,omitempty"`
	User      dto.User  `json:"user,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"create_date,omitempty"`
}

// CreateComment Comment info
func CreateComment(ctx context.Context, comment *Comment) error {
	if err := DB.Table("comment").WithContext(ctx).Create(comment).Error; err != nil {
		return err
	}
	return nil
}

// QueryComment query list of Comment info
func QueryComment(ctx context.Context, videoId int64, limit, offset int) ([]*ResponeComment, int64, error) {
	var total int64
	var res []*Comment
	var conn *gorm.DB
	var responeComment []*ResponeComment
	i := 0
	conn = DB.Table("comment").WithContext(ctx).Model(&Comment{}).Where("video_id = ? ", videoId)
	if err := conn.Count(&total).Error; err != nil {
		return responeComment, total, err
	}
	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return responeComment, total, err
	}
	responeComment = make([]*ResponeComment, len(res))
	for _, v := range res {
		userInfo, _ := QueryUserById(context.Background(), v.UserId)
		fmt.Println("user:", userInfo.Name)
		users := dto.User{
			Id:            userInfo.ID,
			Name:          userInfo.Name,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			IsFollow:      false,
		}
		responeComment[i] = &ResponeComment{
			ID:        v.ID,
			User:      users,
			Content:   v.Content,
			CreatedAt: v.CreatedAt,
		}
		i++
	}
	return responeComment, total, nil
}

// DeleteComment delete comment info
func DeleteCommnet(ctx context.Context, commentId int64) error {
	return DB.Table("comment").WithContext(ctx).Where("id = ?  ", commentId).Delete(&Comment{}).Error
}
