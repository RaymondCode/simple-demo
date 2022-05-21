package repository

// VideoRepository interface

import (
	"github.com/fitenne/youthcampus-dousheng/pkg/model"
	"time"
)

type videoCtl struct {
}

var vctl videoCtl
var db = dbProvider.GetDB()

func GetVideoCtl() model.VideoCtl {
	return &vctl
}

func (v *videoCtl) Create(video *model.Video) error {
	return db.Create(video).Error
}

func (v *videoCtl) Update(video *model.Video) error {
	return db.Save(&video).Error
}

func (v *videoCtl) Delete(video *model.Video) error {
	return db.Delete(&video).Error
}

func (v *videoCtl) GetVideoList(latestTime int64, size int) ([]*model.Video, error) {
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	var videos []*model.Video
	err := db.Where("create_at < ?", latestTime).Order("create_at desc").Limit(size).Find(&videos).Error
	return videos, err
}

func (v *videoCtl) GetVideoById(id int) (*model.Video, error) {
	var video model.Video
	err := db.Where("id = ?", id).First(&video).Error
	return &video, err
}

func (v *videoCtl) GetVideoByAuthorId(authorID int) ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Where("author_id = ?", authorID).Order("create_at desc").Find(&videos).Error
	return videos, err
}

func (v *videoCtl) FindAllByCreateAt() ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Order("create_at desc").Find(&videos).Error
	return videos, err
}

func (v *videoCtl) FindAllByFavoriteCount() ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Order("favorite_count desc").Find(&videos).Error
	return videos, err
}

func (v *videoCtl) FindAllByCommentCount() ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Order("comment_count desc").Find(&videos).Error
	return videos, err
}
func (v *videoCtl) GetFavoriteCount(authorID int) (int64, error) {
	var count int64
	err := db.Model(&model.Video{}).Where("author_id = ?", authorID).Count(&count).Error
	return count, err
}
