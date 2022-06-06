package db

import (
	"github.com/warthecatalyst/douyin/common"
	"github.com/warthecatalyst/douyin/config"
	"time"
)

type FollowInfo struct {
	UserId       int64     `gorm:"user_id"`
	FollowUserId int64     `gorm:"follow_user_id"`
	Status       int       `gorm:"status"`
	CreateTime   time.Time `gorm:"column:create_time;default:null"`
	UpdateTime   time.Time `gorm:"column:update_time;default:null"`
}

func GetFollowCount(userId int64) int64 {
	var count int64
	config.Db.Table(common.FollowInfoTable).Where("user_id = ? and status = ?", userId, common.FollowOn).Count(&count)
	return count
}

func GetFollowerCount(followUserId int64) int64 {
	var count int64
	config.Db.Table(common.FollowInfoTable).Where("follow_user_id = ? and status = ?", followUserId, common.FollowOn).Count(&count)
	return count
}

func IsFollowedRelation(userId int64, anotherUserId int64) bool {
	var count int64
	config.Db.Table(common.FollowInfoTable).Where("user_id = ? and follow_user_id = ? and status = ?", anotherUserId, userId, common.FollowOn).Count(&count)
	return count >= 1
}

func GetFollowInfo(userId int64, followUserId int64) *FollowInfo {
	var followInfo FollowInfo
	config.Db.Table(common.FollowInfoTable).
		Where("user_id = ? and follow_user_id = ?", userId, followUserId).
		First(&followInfo)
	if followInfo.UserId == 0 || followInfo.FollowUserId == 0 {
		return nil
	}
	return &followInfo
}

func AddFollowInfo(userId int64, followUserId int64, actionType int) {
	followInfo := &FollowInfo{
		UserId:       userId,
		FollowUserId: followUserId,
		Status:       actionType,
	}
	config.Db.Table(common.FollowInfoTable).Create(followInfo)
}

func UpdateFollowInfo(userId int64, followUserId int64, actionType int) {
	//config.Db.Table(common.FollowInfoTable).Where("user_id = ? and follow_user_id = ?", userId, followUserId).
	//Update(map[string]interface{}{
	//	"status": actionType,
	//})
}

func GetAllFollowUser(userId int64) []FollowInfo {
	var followInfos []FollowInfo
	config.Db.Table(common.FollowInfoTable).Where("user_id = ? and status = ?", userId, common.FollowOn).Find(&followInfos)
	return followInfos
}

func GetAllFollowerUser(userId int64) []FollowInfo {
	var followerInfos []FollowInfo
	config.Db.Table(common.FollowInfoTable).Where("follow_user_id = ? and status = ?", userId, common.FollowOn).Find(&followerInfos)
	return followerInfos
}
