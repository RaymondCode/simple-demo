package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	VideoId   int64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	if err := DB.Table("favorite").WithContext(ctx).Create(favorite).Error; err != nil {
		return err
	}
	return nil
}

func UpdateFavorite(ctx context.Context, userID, videoID int64, status *int) error {
	params := map[string]interface{}{}
	if status != nil {
		params["status"] = *status
	}
	return DB.Table("favorite").WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", userID, videoID).
		Updates(params).Error
}

func QueryFavorite(ctx context.Context, userID int64, videoID int64) (Favorite, error) {
	var favorite Favorite

	if err := DB.Table("favorite").WithContext(ctx).Where("user_id=? and video_id=? ", userID, videoID).First(&favorite).Error; err != nil {
		return Favorite{}, nil
	}
	return favorite, nil
}

func QueryFavorites(ctx context.Context, userID int64, limit, offset int) ([]Video, int64, error) {
	var videoList []Video
	var total int64
	var conn *gorm.DB
	conn = DB.Table("favorite").WithContext(ctx).Joins("inner join video on favorite.video_id = video.id").
		Select("video.id", "video.author_id", "video.play_url", "video.cover_url", "video.favorite_count", "video.comment_count").
		Where("favorite.user_id = ?", userID)

	if err := conn.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if limit == 0 {
		if err := conn.Scan(&videoList).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err := conn.Limit(limit).Offset(offset).Scan(&videoList).Error; err != nil {
			return nil, 0, err
		}
	}
	return videoList, total, nil
}

//QueryIsFavoritequerytheuserisornotfavoritethevideo
func QueryIsFavorite(ctx context.Context, userId int64, videoId int64) (Favorite, error) {
	var res Favorite
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return Favorite{}, err
	}
	if err := DB.Table("favorite").WithContext(ctx).Where("user_id=? AND video_id=?", userId, videoId).Find(&res).Error; err != nil {
		tx.Rollback()
		return res, err
	}
	return res, nil
}
