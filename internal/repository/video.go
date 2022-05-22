package repository

// VideoRepository interface

import (
	"errors"
	errorcode "github.com/fitenne/youthcampus-dousheng/internal/common/error"
	"github.com/fitenne/youthcampus-dousheng/pkg/model"
	"gorm.io/gorm"
	"time"
)

type videoCtl struct {
}

var vctl videoCtl
var db *gorm.DB

func GetVideoCtl() model.VideoCtl {
	db = dbProvider.GetDB()
	return &vctl
}

func (v *videoCtl) Create(video *model.Video) (int64, error) {
	if video.Author != nil {
		//authorID和author.ID要一致
		if video.AuthorID != 0 && video.AuthorID != video.Author.ID {
			return 0, errors.New(errorcode.VideoCreateForeignKeyNotUnified.Message())
		}
		video.AuthorID = video.Author.ID
		video.Author = nil
	}
	//TODO 判断外键是否存在
	//if _,err:=GetUserCtl().QueryUserByID(video.AuthorID); err!= nil {
	//	return 0, errors.New(errorcode.VideoCreateForeignKeyNotExist.Message())
	//}
	error := db.Create(video).Error
	if error == nil {
		return video.ID, nil
	} else {
		return 0, error
	}
}
func (v *videoCtl) Delete(video *model.Video) error {
	return db.Delete(&video).Error
}

func (v *videoCtl) GetVideoList(latestTime int64, size int) ([]*model.Video, error) {
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}
	var videos []*model.Video
	err := db.Preload("Author").Where("create_at < ?", latestTime).Order("create_at desc").Limit(size).Find(&videos).Error
	return videos, err
}

func (v *videoCtl) GetVideoById(id int) (*model.Video, error) {
	var video model.Video
	//根据id查询video,查询外键author
	err := db.Preload("Author").Where("id = ?", id).First(&video).Error
	return &video, err
}

func (v *videoCtl) GetVideoByAuthorId(authorID int) ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Preload("Author").Where("author_id = ?", authorID).Order("create_at desc").Find(&videos).Error
	return videos, err
}

func (v *videoCtl) FindAllByCreateAt() ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Preload("Author").Order("create_at desc").Find(&videos).Error
	return videos, err
}

func (v *videoCtl) FindAllByFavoriteCount() ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Preload("Author").Order("favorite_count desc").Find(&videos).Error
	return videos, err
}

func (v *videoCtl) FindAllByCommentCount() ([]*model.Video, error) {
	var videos []*model.Video
	err := db.Preload("Author").Order("comment_count desc").Find(&videos).Error
	return videos, err
}
func (v *videoCtl) GetFavoriteCount(authorID int) (int64, error) {
	var count int64
	err := db.Model(&model.Video{}).Where("author_id = ?", authorID).Count(&count).Error
	return count, err
}
