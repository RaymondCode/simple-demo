package service

import (
	"context"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"log"
	"strconv"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
)

type videoListType struct {
	ID            int64
	AuthorId      int
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	CreatedAt     time.Time
	Title         string
	IsFavorite    int
	IsFollow      int

	AuthorName          string
	AuthorFollowCount   int64
	AuthorFollowerCount int64
}

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
	videoListErr, res := model.QueryVideoList(context.Background(), lastTime)

	if videoListErr != nil {
		dto.WriteLog(
			"error", "获取视频列表失败",
			"errInfo", videoListErr.Error(),
		)
	}

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

func QueryFeedResponse1(useId int64, lastTime string) ([]dto.Video, time.Time) {
	var videoList []dto.Video = make([]dto.Video, 10)
	//query video list for feed
	if lastTime == "0" {
		now := time.Now()
		lastTime = now.Format("2006-01-02 15:04:05")
	} else {
		parseInt, err := strconv.ParseInt(lastTime, 10, 64)
		if err != nil {
			return nil, time.Time{}
		}
		lastTime = time.Unix(parseInt, 0).Format("2006-01-02 15:04:05")
	}
	// _, res := model.QueryVideoList(context.Background(), lastTime)

	var res []videoListType
	sqlQuery := "SELECT video.*, IFNULL(favoriteList.status, 2) AS IsFavorite, IFNULL(followList.status, 2) as IsFollow, user.name AS AuthorName, user.follow_count AS AuthorFollowCount, user.follower_count AS AuthorFollowerCount FROM video LEFT JOIN (SELECT video_id, user_id, status FROM favorite WHERE user_id = 11) AS favoriteList ON video.id = favoriteList.video_id LEFT JOIN ( SELECT followed_user, status FROM follow WHERE user_id = 11) AS followList ON video.author_id = followList.followed_user LEFT JOIN user ON video.author_id=user.id LIMIT 10;"
	queryErr := model.DB.Raw(sqlQuery).Scan(&res).Error

	if queryErr != nil {
		fmt.Println(queryErr)
	}

	at := res[len(res)-1].CreatedAt

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
	return videoList[0:len(res)], at
}
