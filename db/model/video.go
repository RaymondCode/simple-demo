package model

import (
	"context"
	"gorm.io/gorm"
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
	Title         string
}

// CreateVideo create video info
func CreateVideo(ctx context.Context, video *Video) error {
	if err := DB.Table("video").WithContext(ctx).Create(video).Error; err != nil {
		return err
	}
	return nil
}

//QueryVideoListqueryvideolist
func QueryVideoList(ctx context.Context, nextTime string) (error, []Video) {
	var videoList []Video
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err, nil
	}
	if err := tx.Table("video").WithContext(ctx).Where("created_at BETWEEN(\"2022-06-01 21:42:52\") and (\"" + nextTime + "\")").Order("video.created_at desc").Limit(50).Find(&videoList).Error; err != nil {
		tx.Rollback()
		return err, videoList
	}
	return nil, videoList
}

func UpdateVideoFavorite(ctx context.Context, videoID int64, action int) error {
	tx := DB.Begin()
	if err := tx.Table("video").WithContext(ctx).Where("id = ?", videoID).Update("favorite_count", gorm.Expr("favorite_count+?", action)).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
