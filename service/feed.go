package service

import (
	"context"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"time"
)

func QueryFeedResponse(useId int64, nextTime string) ([]dto.Video, time.Time) {
	var videoList []dto.Video = make([]dto.Video, 50)
	var isFavorite bool
	var isFollow bool
	//query video list for feed
	if nextTime == "" {
		now := time.Now()
		nextTime = now.Format("2006-01-02 15:04:05")
	}
	_, res := model.QueryVideoList(context.Background(), nextTime)
	at := res[len(res)-1].CreatedAt
	//constitute FeedResponse struct ([]dto.video)
	for index, value := range res {
		//user is favorite
		if favoriteInfo, _ := model.QueryIsFavorite(context.Background(), useId, value.ID); favoriteInfo.Status == 1 {
			isFavorite = true
		} else {
			isFavorite = false
		}

		authorInfo, _ := model.QueryUserById(context.Background(), value.AuthorID)
		//user is follow
		if followInfo, _ := model.QueryIsFavorite(context.Background(), useId, value.AuthorID); followInfo.Status == 1 {
			isFollow = true
		} else {
			isFollow = false
		}
		fmt.Println(isFollow, isFavorite)
		videoList[index] = dto.Video{
			Id: value.ID,
			Author: dto.User{
				Id:            authorInfo.ID,
				Name:          authorInfo.Name,
				FollowCount:   authorInfo.FollowCount,
				FollowerCount: authorInfo.FollowerCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       value.PlayUrl,
			CoverUrl:      value.CoverUrl,
			FavoriteCount: value.FavoriteCount,
			CommentCount:  value.CommentCount,
			IsFavorite:    isFavorite,
		}
	}
	return videoList[0:len(res)], at
}
