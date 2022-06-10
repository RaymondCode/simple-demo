package model

import (
	"context"
	"time"
)

type Video struct {
	ID            int64 `gorm:"primarykey"`
	AuthorID      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	CreatedAt     time.Time
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, video *Video) error {
	if err := DB.Table("video").WithContext(ctx).Create(video).Error; err != nil {
		return err
	}
	return nil
}
