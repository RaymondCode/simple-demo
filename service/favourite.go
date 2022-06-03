package service

import (
	"errors"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
)

// FavoriteActionInfo service层添加或者删除一条点赞记录
func FavoriteActionInfo(userId, videoId int64, actionType int32) error {
	return newFavoriteActionInfoFlow(userId, videoId, actionType).Do()
}

func newFavoriteActionInfoFlow(userId, videoId int64, actionType int32) *FavoriteActionInfoFlow {
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
	if f.actionType == api.FavoriteAction {
		if err := f.AddRecord(); err != nil {
			return err
		}
	} else if f.actionType == api.UnFavoriteAction {
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

type VideoList []api.Video

// FavoriteListInfo 获得用户点赞后的视频列表
func FavoriteListInfo(userId int64) (*VideoList, error) {
	return newFavoriteListInfoFlow(userId).Do()
}

func newFavoriteListInfoFlow(userId int64) *FavoriteListInfoFlow {
	return &FavoriteListInfoFlow{
		userId: userId,
	}
}

type FavoriteListInfoFlow struct {
	userId int64
}

func (f *FavoriteListInfoFlow) Do() (*VideoList, error) {
	return f.getFavoriteList()
}

func (f *FavoriteListInfoFlow) getFavoriteList() (*VideoList, error) {
	videoIds, err := dao.NewFavoriteDaoInstance().VideoIDListByUserID(f.userId)
	if err != nil {
		return nil, err
	}
	var videolist VideoList
	for _, videoId := range videoIds {
		user, err := videoService.getUserFromVideoId(videoId)
		if err != nil {
			return nil, err
		}
		video, err := dao.NewVideoDaoInstance().GetVideoFromId(videoId)

		videolist = append(videolist, api.Video{
			Id:            videoId,
			Author:        *user,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: int64(video.FavoriteCount),
			CommentCount:  int64(video.CommentCount),
			IsFavorite:    false,
		})
	}
	return &videolist, nil
}
