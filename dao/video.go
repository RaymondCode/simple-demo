package dao

import (
	"errors"
	"github.com/warthecatalyst/douyin/global"
	"github.com/warthecatalyst/douyin/model"
	"gorm.io/gorm"
	"sync"
)


const LIMITVIDEOLISTNUMS = 30

type VideoDao struct{}

func NewVideoDao() *VideoDao {
	return &VideoDao{}
}

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

//通过GetVideoFromId 使用VideoId 查询相关视频
func (*VideoDao) GetVideoFromId(videoId int64) (*model.Video, error) {
	video := &model.Video{}
	if err := db.Where("video_id = ?", videoId).First(video).Error; err != nil {
		return nil, err
	}

	return video, nil
}
//creat 新增数据到数据库中
func (*VideoDao) Create(values map[string]interface{}) error {
	err := db.Model(&model.Video{}).Create(values).Error
	if err != nil {
		global.DyLogger.Print("新增视频失败")
	}
	return err
}
//GetVideos 根据传入参数查询对应的视频
func (*VideoDao) GetVideos(conditions map[string]interface{},fields ...string) ([]model.Video,error){
	var v []model.Video
	err := db.Model(&model.Video{}).
		Select(fields).
		Where(conditions).
		Find(&v).Error
	if err != nil {
		return nil,err
	}
	return v,nil
}

//GetLatest 获取最新的30条视频数据
//限制数后期可以新增控制
func (*VideoDao) GetLatest(latestTime string) ([]model.Video,error) {
	var v []model.Video
	err := db.Model(&model.Video{}).Order("create_at desc").
		Select([]string{"author_id","play_url","favorite_count","comment_count","create_at"}).
		Where("create_at < ?", latestTime).
		Find(&v).
		Limit(LIMITVIDEOLISTNUMS).
		Error
	if err != nil && errors.Is(err,gorm.ErrRecordNotFound) {
		return nil,err
	}
	return v,nil
}
//AddFavorite 增加 count 数量的点赞数
func (*VideoDao) AddFavorite(videoID int64, count int) error {
	err := db.Model(model.Video{}).
		Where("video_id = ?", videoID).
		Update("favorite_count", gorm.Expr("favorite_count + ?", count)).
		Error

	if err != nil {
		return err
	}
	return nil
}

// SubFavorite  减少 count 数量的点赞数
func (*VideoDao) SubFavorite(videoID int64, count int) error {
	err := db.Model(model.Video{}).
		Where("video_id = ?" , videoID).
		Update("favorite_count", gorm.Expr("favorite_count + ?", count)).
		Error

	if err != nil {
		return err
	}
	return nil
}

// GetVideosByID 通过 ID 列表批量查询视频
func (*VideoDao) GetVideosByID(IDList []int64, fields ...string) ([]model.Video, error) {
	var v []model.Video
	err := db.Model(&model.Video{}).
		Select(fields).
		Where("video_id IN ? ", IDList).
		Find(&v).
		Error
	// 处理错误和未查询到记录的err
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return v, nil
}
