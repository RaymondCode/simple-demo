package dao

import (
	"github.com/warthecatalyst/douyin/model"
	"gorm.io/gorm"
)

func AddFollow(tx *gorm.DB, userId, toUserId int64) error {
	follow := &model.Follow{
		FromUserID: userId,
		ToUserID:   toUserId,
	}
	return tx.Create(follow).Error
}

func DeleteFollow(tx *gorm.DB, userId, toUserId int64) error {
	follow := &model.Follow{
		FromUserID: userId,
		ToUserID:   toUserId,
	}
	return tx.Delete(follow).Error
}

func GetFollowList(userId int64) ([]*model.Follow, error) {
	var followUserIdList []*model.Follow
	if err := db.Where("from_user_id = ?", userId).Find(followUserIdList).Error; err != nil {
		return followUserIdList, err
	}

	return followUserIdList, nil
}

func GetFollowerList(userId int64) ([]*model.Follow, error) {
	var followerUserIdList []*model.Follow
	if err := db.Where("to_user_id = ?", userId).Find(followerUserIdList).Error; err != nil {
		return followerUserIdList, err
	}

	return followerUserIdList, nil
}
