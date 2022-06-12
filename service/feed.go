package service

import (
	"context"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"log"
	"strconv"
	"time"
)

func QueryFeedResponse(useId int64, lastTime string) ([]dto.Video, time.Time) {
	var videoList []dto.Video = make([]dto.Video, 10)
	var isFavorite bool
	var isFollow bool
	//query video list for feed
	if lastTime == "0" {
		now := time.Now()
		lastTime = now.Format("2006-01-02 15:04:05")
	} else {
		parseInt, err := strconv.ParseInt(lastTime, 10, 64)
		parseInt = parseInt / 1000
		log.Println(parseInt)
		if err != nil {
			return nil, time.Time{}
		}
		lastTime = time.Unix(parseInt, 0).Format("2006-01-02 15:04:05")
		log.Println(lastTime)
	}
	_, res := model.QueryVideoList(context.Background(), lastTime)
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
			Title:         value.Title,
		}
	}
	return videoList[0:len(res)], at
}
