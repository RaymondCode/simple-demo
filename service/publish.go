package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
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

func QueryPublishList1(useId int64) []dto.Video {
	var videoList []dto.Video = make([]dto.Video, 10)
	//query video list for feed
	// _, res := model.QueryVideoList(context.Background(), lastTime)

	var res []videoListType
	sqlQuery := "SELECT video.*, IFNULL(favoriteList.status, 2) AS IsFavorite, IFNULL(followList.status, 2) as IsFollow, user.name AS AuthorName, user.follow_count AS AuthorFollowCount, user.follower_count AS AuthorFollowerCount FROM video LEFT JOIN (SELECT video_id, user_id, status FROM favorite WHERE user_id = 11) AS favoriteList ON video.id = favoriteList.video_id LEFT JOIN ( SELECT followed_user, status FROM follow WHERE user_id = 11) AS followList ON video.author_id = followList.followed_user LEFT JOIN user ON video.author_id=user.id LIMIT 10;"
	queryErr := model.DB.Raw(sqlQuery).Scan(&res).Error

	if queryErr != nil {
		fmt.Println(queryErr)
	}

	for index, value := range res {

		var is_follow bool
		var is_fav bool

		if value.IsFollow == 1 {
			is_follow = true
		} else {
			is_follow = false
		}

		if value.IsFavorite == 1 {
			is_fav = true
		} else {
			is_fav = false
		}

		videoList[index] = dto.Video{
			Id: value.ID,
			Author: dto.User{
				Id:            int64(value.AuthorId),
				Name:          value.AuthorName,
				FollowCount:   value.AuthorFollowCount,
				FollowerCount: value.AuthorFollowerCount,
				IsFollow:      is_follow,
			},
			PlayUrl:       value.PlayUrl,
			CoverUrl:      value.CoverUrl,
			FavoriteCount: value.FavoriteCount,
			CommentCount:  value.CommentCount,
			IsFavorite:    is_fav,
			Title:         value.Title,
		}
	}
	return videoList[0:len(res)]
}
