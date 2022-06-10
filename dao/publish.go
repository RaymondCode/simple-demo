package dao

import (
	"errors"
	"github.com/warthecatalyst/douyin/model"
	"gorm.io/gorm"
	"sync"
)

type PublishDao struct{}

var (
	publishDao  *PublishDao
	publishOnce sync.Once
)

func NewPublishDaoInstance() *PublishDao {
	publishOnce.Do(func() {
		publishDao = &PublishDao{}
	})
	return publishDao
}

func (p *PublishDao) AddVideo(video *model.Video) error {
	return db.Create(video).Error
}

func (p *PublishDao) GetVideoPublistList(userId int64) ([]int64, error) {
	var f []model.Video
	err := db.Model(&model.Video{}).
		Select("video_id").
		Where("user_id = ?", userId).
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
