package service

import (
	"context"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"strconv"
)

func QueryPublishList(userId string) []dto.Video {
	var videoList []dto.Video = make([]dto.Video, 10)
	var isFavorite bool
	var isFollow bool
	//query video list for feed

	_, res := model.QueryPublishList(context.Background(), userId)
	if len(res) == 0 {
		return videoList[0:0]
	}
	//constitute FeedResponse struct ([]dto.video)
	useId, _ := strconv.ParseInt(userId, 10, 64)
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
	return videoList[0:len(res)]
}
