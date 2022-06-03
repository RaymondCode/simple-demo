package dao

import (
	"errors"
	"github.com/warthecatalyst/douyin/model"
	"gorm.io/gorm"
	"sync"
)

type FavouriteDao struct{}

var (
	favoriteDao  *FavouriteDao
	favoriteOnce sync.Once
)

func NewFavoriteDaoInstance() *FavouriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavouriteDao{}
		})
	return favoriteDao
}

//QueryCountOfVideo 方法 从Video表中查询点赞的数据
func (*FavouriteDao) QueryCountOfVideo(conditions map[string]interface{}) (int32, error) {
	var video model.Video
	err := db.Where(conditions).First(&video).Error
	if err != nil {
		return 0, err
	}
	return video.FavoriteCount, err
}

//IsFavourite 查询 userID的用户是否对videoID的视频进行点赞
func (*FavouriteDao) IsFavourite(userID, videoID int64) bool {
	var fav model.Favourite
	result := db.Where("user_id = ? AND video_id = ?", userID, videoID).First(&fav)
	return result.RowsAffected != 0
}

//Add 向数据库中增加一条点赞记录
func (*FavouriteDao) Add(userID, videoID int64) error {
	f := model.Favourite{
		UserID:  userID,
		VideoID: videoID,
	}
	err := db.Model(&model.Favourite{}).Create(&f).Error
	if err != nil {
		return err
	}

	//不要忘记在Video表中更新点赞记录
	var video model.Video
	err = db.Where("video_id = ?", videoID).First(&video).Error
	if err != nil {
		return err
	}
	video.FavoriteCount++ //可能会引发并发问题
	db.Save(&video)
	return nil
}

//Del 从数据库中删除一条点赞记录
func (*FavouriteDao) Del(userID, videoID int64) error {
	f := model.Favourite{
		UserID:  userID,
		VideoID: videoID,
	}

	err := db.Model(&model.Favourite{}).Delete(&f).Error

	if err != nil {
		return err
	}
	var video model.Video
	err = db.Where("video_id = ?", videoID).First(&video).Error
	if err != nil {
		return err
	}
	video.FavoriteCount-- //可能会引发并发问题
	db.Save(&video)
	return nil
}

//VideoIDListByUserID 获取某用户点赞的所有视频的ID列表
func (*FavouriteDao) VideoIDListByUserID(userID int64) ([]int64, error) {
	var f []model.Favourite
	err := db.Model(&model.Favourite{}).
		Select("video_id").
		Where("user_id = ?", userID).
		Find(&f).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	var res []int64
	for _, i := range f {
		res = append(res, i.VideoID)
	}
	return res, nil
}
