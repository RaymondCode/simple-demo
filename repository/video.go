package repository

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
)

type VideoRepository struct{}

// QueryByIds 根据id列表查询video集合，同时填充video的author
func (vr *VideoRepository) QueryByIds(ids []int64) ([]model.Video, error) {
	var videoList []model.Video
	if err := global.DB.Preload("Author").Find(&videoList, ids).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// UpdateFavoriteCount 更新video的favorite_count
func (vr *VideoRepository) UpdateFavoriteCount(videoId int64, count int64) error {
	return global.DB.Model(&model.Video{}).Where("id = ?", videoId).Update("favorite_count", count).Error
}
