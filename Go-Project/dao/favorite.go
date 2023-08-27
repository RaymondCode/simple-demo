package dao

import (
	"github.com/life-studied/douyin-simple/global"
	"time"
)

func InsertFavoriteVideo(user User, video Video) (err error) {
	like := &Like{
		ID:      time.Now().Unix(),
		UserID:  user.ID,
		VideoID: video.ID,
		User:    user,
		Video:   video,
	}
	return global.DB.Create(like).Error

}

func DeleteFavoriteVideo(user User, video Video) (err error) {
	var like Like
	err = global.DB.Where("user_id = ? and video_id = ?", user.ID, video.ID).Find(&like).Error
	if err != nil {
		return err
	}
	return global.DB.Delete(like).Error
}

func GetFavoriteVideo(user User) (likes []Like, err error) {
	err = global.DB.Where("user_id = ?", user.ID).Find(&likes).Error
	return likes, err
}

func GetFavoriteUser(video Video) (likes []Like, err error) {
	err = global.DB.Where("video_id = ?", video.ID).Find(&likes).Error
	return likes, err
}
