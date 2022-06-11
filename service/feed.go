package service

import (
	"context"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
)

func QueryFeedResponse(useId int64) []dto.Video {
	var videoList []dto.Video = make([]dto.Video, 50)
	var isFavorite bool
	var isFollow bool
	//query video list for feed
	_, res := model.QueryVideoList(context.Background())
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
	return videoList
}
