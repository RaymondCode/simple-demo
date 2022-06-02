package dao

import (
	"github.com/warthecatalyst/douyin/model"
	"gorm.io/gorm"
	"sync"
)

type FollowDao struct{}

var (
	followDao  *FollowDao
	followOnce sync.Once
)

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(
		func() {
			followDao = &FollowDao{}
		})
	return followDao
}

func (*FollowDao) AddFollow(tx *gorm.DB, userId, toUserId int64) error {
	follow := &model.Follow{
		FromUserID: userId,
		ToUserID:   toUserId,
	}
	return tx.Create(follow).Error
}

func (*FollowDao) DeleteFollow(tx *gorm.DB, userId, toUserId int64) error {
	follow := &model.Follow{}
	return tx.Where("from_user_id = ? AND to_user_id = ?", userId, toUserId).Delete(follow).Error
}

func (*FollowDao) GetFollowList(userId int64) ([]model.Follow, error) {
	var followUserIdList []model.Follow
	if err := db.Where("from_user_id = ?", userId).Find(&followUserIdList).Error; err != nil {
		return followUserIdList, err
	}

	return followUserIdList, nil
}

func (*FollowDao) GetFollowerList(userId int64) ([]model.Follow, error) {
	var followerUserIdList []model.Follow
	if err := db.Where("to_user_id = ?", userId).Find(&followerUserIdList).Error; err != nil {
		return followerUserIdList, err
	}

	return followerUserIdList, nil
}
