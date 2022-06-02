package service

import (
	"errors"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/model"
)

// FavoriteActionInfo service层添加或者删除一条点赞记录
func FavoriteActionInfo(userId, videoId int64, actionType int32) error {
	return NewFavoriteActionInfoFlow(userId, videoId, actionType).Do()
}

func NewFavoriteActionInfoFlow(userId, videoId int64, actionType int32) *FavoriteActionInfoFlow {
	return &FavoriteActionInfoFlow{
		userId:     userId,
		videoId:    videoId,
		actionType: actionType,
	}
}

type FavoriteActionInfoFlow struct {
	userId     int64
	videoId    int64
	actionType int32
}

func (f *FavoriteActionInfoFlow) Do() error {
	if f.actionType == 1 { //1为点赞
		if err := f.AddRecord(); err != nil {
			return err
		}
	} else if f.actionType == 2 { //2为取消点赞
		if err := f.checkRecord(); err != nil {
			return err
		}
		if err := f.DelRecord(); err != nil {
			return err
		}
	} else {
		return errors.New("actionType must be 1 or 2")
	}
	return nil
}

func (f *FavoriteActionInfoFlow) checkRecord() error {
	if flag := dao.NewFavoriteDaoInstance().IsFavourite(f.userId, f.videoId); !flag {
		return errors.New("there's no such record")
	}
	return nil
}

func (f *FavoriteActionInfoFlow) AddRecord() error {
	if err := dao.NewFavoriteDaoInstance().Add(f.userId, f.videoId); err != nil {
		return err
	}
	return nil
}

func (f *FavoriteActionInfoFlow) DelRecord() error {
	if err := dao.NewFavoriteDaoInstance().Del(f.userId, f.videoId); err != nil {
		return err
	}
	return nil
}

type VideoListInfo struct {
	videoList []*model.Video
}

// FavoriteListInfo 获得用户点赞后的视频列表
func FavoriteListInfo(userId int64) (*VideoListInfo, error) {
	return &VideoListInfo{}, nil //还没写，先暂时写个demo
}
