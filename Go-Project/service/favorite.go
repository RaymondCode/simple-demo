package service

import (
	"fmt"
	"github.com/life-studied/douyin-simple/dao"
	"strconv"
)

type FavoriteVideoID struct {
	VideoID       string
	Author        FavoriteUserID
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
}

type FavoriteUserID struct {
	UserName      string
	UserID        int64
	FollowCount   int64
	FollowerCount int64
}

type FavoriteVideoList struct {
	Id            int64
	Author        FavoriteUserID
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
}

func FavoriteVideo(video FavoriteVideoID, user FavoriteUserID) error {
	videoID, err := strconv.ParseInt(video.VideoID, 10, 64)
	if err != nil {
		return err
	}
	err = dao.InsertFavoriteVideo(dao.User{ID: user.UserID}, dao.Video{ID: videoID})
	return err
}

func UnfavoriteVideo(video FavoriteVideoID, user FavoriteUserID) error {
	videoID, err := strconv.ParseInt(video.VideoID, 10, 64)
	if err != nil {
		fmt.Println("videoID is not int64")
		return err
	}
	err = dao.DeleteFavoriteVideo(dao.User{ID: user.UserID}, dao.Video{ID: videoID})
	return err
}

func ReadFavoriteVideo(user FavoriteUserID) (videoList []FavoriteVideoList, err error) {
	likes, err := dao.GetFavoriteVideo(dao.User{ID: user.UserID})
	if err != nil {
		return nil, err
	}
	videoList = make([]FavoriteVideoList, 0, len(likes))
	for _, like := range likes {
		node := like.Video
		author := like.User
		videoList = append(videoList, FavoriteVideoList{
			Id: node.ID,
			Author: FavoriteUserID{
				UserName: author.Name,
				UserID:   author.ID,
			},
			PlayUrl:  node.PlayURL,
			CoverUrl: node.CoverURL,
		})
	}
	return videoList, nil
}
