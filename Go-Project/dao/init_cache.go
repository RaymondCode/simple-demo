// Package dao -----------------------------
// @file      : init_cache.go
// @author    : Yunyin
// @contact   : yunyin_jayyi@qq.com
// @time      : 2023/8/29 0:06
// -------------------------------------------
package dao

import (
	"github.com/life-studied/douyin-simple/global"
)

func InitVideoDataFromMysql() ([]Video, error) {
	var videos []Video
	err := global.DB.Select("*").Find(&videos).Error
	return videos, err
}

func InitUserByUserIdFromMysql(userId int64) (string, string, error) {
	var user User
	err := global.DB.Where("id=?", userId).Find(&user).Error
	return user.Name, user.Password, err
}

func InitUserFromMysql() ([]User, error) {
	var users []User
	err := global.DB.Select("*").Find(&users).Error
	return users, err
}

func InitVideoFavoriteCountFromMysql(videoId int64) (int64, error) {
	var count int64
	err := global.DB.Model(&Like{}).Where("video_id=?", videoId).Count(&count).Error
	return count, err
}

func InitVideoCommentCountFromMysql(videoId int64) (int64, error) {
	var count int64
	err := global.DB.Model(&Comment{}).Where("video_id=?", videoId).Count(&count).Error
	return count, err
}
