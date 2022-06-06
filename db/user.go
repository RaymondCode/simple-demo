package db

import (
	"github.com/warthecatalyst/douyin/common"
	"github.com/warthecatalyst/douyin/config"
	"time"
)

type UserInfo struct {
	UserId     int64     `gorm:"user_id"`
	UserName   string    `gorm:"user_name"`
	Password   string    `gorm:"password"`
	Status     int       `gorm:"column:status;default:null"`
	CreateTime time.Time `gorm:"column:create_time;default:null"`
	UpdateTime time.Time `gorm:"column:update_time;default:null"`
}

func AddUserInfo(userId int64, userName, password string) *UserInfo {
	userInfo := UserInfo{
		UserId:   userId,
		UserName: userName,
		Password: password,
	}
	db := config.GetDB().Table(common.UserInfoTable)
	db.Create(&userInfo)
	return &userInfo
}

func GetUserInfoByUserName(userName string) *UserInfo {
	userInfo := &UserInfo{}
	db := config.GetDB().Table(common.UserInfoTable)
	db.Where("user_name = ? and status = ?", userName, common.NameOn).First(userInfo)
	if userInfo.UserId == 0 {
		return nil
	}
	return userInfo
}

func GetUserInfoByUserId(userId int64) *UserInfo {
	userInfo := &UserInfo{}
	db := config.GetDB().Table(common.UserInfoTable)
	db.Where("user_id = ? and status = ?", userId, common.NameOn).First(userInfo)
	if userInfo.UserId == 0 {
		return nil
	}
	return userInfo
}

func GetUserInfoByPassword(userName string, password string) *UserInfo {
	userInfo := &UserInfo{}
	db := config.GetDB().Table(common.UserInfoTable)
	db.Where("user_name = ? and password = ? and status = ?", userName, password, common.NameOn).First(userInfo)
	if userInfo.UserId == 0 {
		return nil
	}
	return userInfo
}

func GetUserNameByUserId(userId int64) string {
	userInfo := UserInfo{}
	db := config.GetDB().Table(common.UserInfoTable)
	db.Where("user_id = ? and status = ?", userId, common.NameOn).First(&userInfo)
	return userInfo.UserName
}
