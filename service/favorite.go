package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
)

func IsFavorite(ctx context.Context, userID int64, videoID int64) (isExist bool, isFavorite bool) {
	favoriteRelation, _ := model.QueryFavorite(ctx, userID, videoID)
	if favoriteRelation.ID == 0 {
		return false, false
	}
	if favoriteRelation.Status == 1 {
		return true, true
	}
	return true, false
}

func FavoriteCountAction(ctx context.Context, videoID int64, actionType int) error {
	if actionType == 1 {
		// 视频点赞数+1
		model.UpdateVideoFavorite(ctx, videoID, 1)
	} else {
		// 视频点赞数-1
		model.UpdateVideoFavorite(ctx, videoID, -1)
	}
	return nil
}

func FavoriteAction(ctx context.Context, userID int64, videoID int64, actionType int) error {
	isExist, isFavorite := IsFavorite(ctx, userID, videoID)

	// TODO 修改video的点赞数据
	if !isExist {
		// 不存在的关系直接创建
		favorite := model.Favorite{
			UserId:  userID,
			VideoId: videoID,
			Status:  actionType,
		}
		if err := model.CreateFavorite(ctx, &favorite); err != nil {
			return err
		}
		FavoriteCountAction(ctx, videoID, actionType)
	} else {
		// 存在的关系进行更新
		if (actionType == 1 && !isFavorite) || (actionType == 2 && isFavorite) {
			if err := model.UpdateFavorite(ctx, userID, videoID, &actionType); err != nil {
				return err
			}
			FavoriteCountAction(ctx, videoID, actionType)
		}
	}
	return nil
}

func GetFavoriteList(ctx context.Context, userId int64) ([]dto.Video, error) {

	videoList, _, _ := model.QueryFavorites(ctx, userId, 0, 0)

	res := make([]dto.Video, 0)
	for i := range videoList {
		video := dto.Video{}
		video.Id = videoList[i].ID
		video.PlayUrl = videoList[i].PlayUrl
		video.CoverUrl = videoList[i].CoverUrl
		video.FavoriteCount = videoList[i].FavoriteCount
		video.CommentCount = videoList[i].CommentCount
		author, _ := model.QueryUserById(ctx, videoList[i].AuthorID)
		user := dto.User{}
		user.Id = author.ID
		user.FollowerCount = author.FollowCount
		user.Name = author.Name
		user.FollowerCount = author.FollowerCount
		isExist, isFollow := IsFollow(ctx, userId, author.ID)
		if isExist && isFollow {
			user.IsFollow = true
		} else {
			return nil, errors.New("get favirote list error")
		}
		video.Author = user
		video.IsFavorite = true
		res = append(res, video)
		fmt.Println(video)
	}

	return res, nil
}
