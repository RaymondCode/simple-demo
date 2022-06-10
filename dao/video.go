package dao

import (
	"errors"
	"sync"

	"github.com/warthecatalyst/douyin/model"
	"gorm.io/gorm"
)

const LIMITVIDEOLISTNUMS = 30

// VideoDao dao层执行与视频相关的数据库查询
type VideoDao struct{}

var (
	videoDao  *VideoDao
	videoOnce sync.Once
)

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(func() {
		videoDao = &VideoDao{}
	})
	return videoDao
}

// GetUserIdFromVideoId 从VideoId得到对应的UserId
func (v *VideoDao) GetUserIdFromVideoId(videoId int64) (int64, error) {
	var video model.Video
	err := db.Select("user_id").Where("video_id = ?", videoId).First(&video).Error
	if err != nil {
		return 0, err
	}
	return video.UserID, nil
}

func (*VideoDao) GetVideoFromId(videoId int64) (*model.Video, error) {
	video := &model.Video{}
	if err := db.Where("video_id = ?", videoId).First(video).Error; err != nil {
		return nil, err
	}

	return video, nil
}

//GetLatest 获取最新的30条视频数据
//限制数后期可以新增控制
func (*VideoDao) GetLatest(latestTime string) ([]model.Video, error) {

	var v []model.Video
	err := db.Model(&model.Video{}).Order("create_at desc").
		Select("*").
		Where("create_at < ?", latestTime).
		Find(&v).
		Limit(LIMITVIDEOLISTNUMS).
		Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return v, nil
}
