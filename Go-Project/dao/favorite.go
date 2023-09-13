package dao

import (
	"github.com/life-studied/douyin-simple/global"
	"time"
)

func InsertFavoriteVideo(user User, video Video) (err error) {
	var likeok []Like
	// 判断用户是否已经收藏过该视频
	global.DB.Where("user_id=? and video_id=?", user.ID, video.ID).Find(&likeok).First(&likeok)
	if len(likeok) > 0 {
		return nil
	}
	var saveUser []User
	var saveVideo []Video
	err = global.DB.Where("id=?", user.ID).Find(&saveUser).Error
	if err != nil {
		return err
	}
	err = global.DB.Where("id=?", video.ID).Find(&saveVideo).Error
	if err != nil {
		return err
	}
	like := &Like{
		ID:      time.Now().Unix(),
		UserID:  user.ID,
		VideoID: video.ID,
		User:    saveUser[0],
		Video:   saveVideo[0],
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
	for i := range likes {
		global.DB.Where("id = ?", likes[i].UserID).Find(&likes[i].User)
		global.DB.Where("id = ?", likes[i].VideoID).Find(&likes[i].Video)
	}
	return likes, err
}

func GetFavoriteUser(video Video) (likes []Like, err error) {
	err = global.DB.Where("video_id = ?", video.ID).Find(&likes).Error
	for i := range likes {
		global.DB.Where("id = ?", likes[i].UserID).Find(&likes[i].User)
		global.DB.Where("id = ?", likes[i].VideoID).Find(&likes[i].Video)
	}
	return likes, err
}
