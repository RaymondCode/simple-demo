package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	ID           int64 `gorm:"primarykey"`
	UserId       int64
	FollowedUser int64
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CreateFollow create follow info
func CreateFollow(ctx context.Context, follow *Follow) error {
	if err := DB.Table("follow").WithContext(ctx).Create(follow).Error; err != nil {
		return err
	}
	return nil
}

// UpdateFollow update follow info
func UpdateFollow(ctx context.Context, userID, followedUser uint, status *int) error {
	params := map[string]interface{}{}
	if status != nil {
		params["status"] = *status
	}
	return DB.Table("follow").WithContext(ctx).Model(&Follow{}).Where("user_id = ? and followed_user = ?", userID, followedUser).
		Updates(params).Error
}

// DeleteFollow delete follow info
func DeleteFollow(ctx context.Context, userID uint, followedUser uint) error {
	return DB.Table("follow").WithContext(ctx).Where("user_id = ? and followed_user = ? ", userID, followedUser).Delete(&Follow{}).Error
}

// QueryNote query list of note info
func QueryFollow(ctx context.Context, userID int64, status, limit, offset int) ([]*Follow, int64, error) {
	var total int64
	var res []*Follow
	var conn *gorm.DB
	// query for followed users
	if status == 1 {
		conn = DB.Table("follow").WithContext(ctx).Model(&Follow{}).Where("user_id = ?", userID)
	} else { // query for followers
		conn = DB.Table("follow").WithContext(ctx).Model(&Follow{}).Where("followed_user = ?", userID)
	}
	if err := conn.Count(&total).Error; err != nil {
		return res, total, err
	}
	if err := conn.Limit(limit).Offset(offset).Find(&res).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}
