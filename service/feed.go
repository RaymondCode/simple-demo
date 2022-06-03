package service

import (
	"time"

	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
)

type FeedService struct{}

func (fes *FeedService) QueryFeed(latestTime int64, token string) ([]model.Video, error) {
	var rawVideos []model.Video
	var err error
	if latestTime > 0 {
		tm := time.Unix(latestTime/1000, 0)
		timeLayout := "2006-01-02 15:04:05" //firm
		latestTimeStr := tm.Format(timeLayout)
		rawVideos, err = repository.GroupApp.VideoRepository.QueryVideosSince(latestTimeStr)
	} else {
		rawVideos, err = repository.GroupApp.VideoRepository.QueryAllVideos()
	}
	if err != nil {
		return nil, err
	}
	return rawVideos, nil
	// var videos []Video
	// for _, video := range rawVideos {
	// 	userId := video.UserId
	// 	rawAuthor, err := repository.GroupApp.UserRepository.QueryUserById(userId)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	author := User{
	// 		Id:            rawAuthor.Id,
	// 		Name:          rawAuthor.Username,
	// 		FollowCount:   rawAuthor.FollowCount,
	// 		FollowerCount: rawAuthor.FollowerCount,
	// 		IsFollow:      false,
	// 	}
	// 	videos = append(videos, Video{
	// 		Id:            video.Id,
	// 		Author:        author,
	// 		PlayUrl:       video.PlayUrl,
	// 		CoverUrl:      video.CoverUrl,
	// 		FavoriteCount: video.FavoriteCount,
	// 		CommentCount:  video.CommentCount,
	// 		IsFavorite:    false,
	// 	})
	// }
	// return videos, nil
}
