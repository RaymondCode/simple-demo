package dao

import (
	"github.com/warthecatalyst/douyin/model"
	"sync"
)

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
	err := db.Select("user_id").Where("video_id = ", videoId).First(&video).Error
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
