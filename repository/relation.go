package repository

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

type RelationRepository struct {
}

//InsertUserId 向关注列表中添加一条用户数据
func (rr *RelationRepository) InsertUserId(userId, toUserId int64) error {
	f := model.Follower{UserId: userId, ToUserId: toUserId}
	return global.DB.Create(&f).Error
}

//DeleteUserId 根据UserId和toUserId向关注列表中删除一条用户数据
func (rr *RelationRepository) DeleteUserId(userId, toUserId int64) error {
	return global.DB.Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(model.Follower{}).Error
}

//GetFollowListByUserId  根据UserId找到关注列表
func (rr *RelationRepository) GetFollowListByUserId(userId int64) ([]*model.User, error) {
	var userList []*model.User
	if err := global.DB.Where("id in (?)",
		global.DB.Table("user_follower").Select("to_user_id").Where("userId = ?", userId)).Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

//GetFollowerListByToUserId 根据ToUserId找到粉丝列表
func (rr *RelationRepository) GetFollowerListByToUserId(toUserId int64) ([]*model.User, error) {
	var userList []*model.User
	if err := global.DB.Where("id in (?)",
		global.DB.Table("user_follower").Select("user_id").Where("to_user_id"), toUserId).Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

//GetFollowerByUserIdAndToUserId 根据自己的id和对方的id精确查找到粉丝
func (rr *RelationRepository) GetFollowerByUserIdAndToUserId(userId, toUserId int64) (*model.Follower, error) {
	var f *model.Follower
	if err := global.DB.Where("user_id = ? and to_user_id = ?", userId, toUserId).First(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

//IncreaseFollowCount 关注的一个步骤 增加用户属性中的关注数量
func (rr *RelationRepository) IncreaseFollowCount(userId int64, count int) error {
	return global.DB.Model(&model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", count)).Error
}

//DecreaseFollowCount 减少关注数量
func (rr *RelationRepository) DecreaseFollowCount(userId int64, count int) error {
	return global.DB.Model(&model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count - ?", count)).Error
}

//IncreaseFollowerCount  增加粉丝的数量
func (rr *RelationRepository) IncreaseFollowerCount(userId int64, count int) error {
	return global.DB.Model(&model.User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count + ?", count)).Error
}

//DecreaseFollowerCount 减少粉丝的数量
func (rr *RelationRepository) DecreaseFollowerCount(userId int64, count int) error {
	return global.DB.Model(&model.User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count - ?", count)).Error
}

//
