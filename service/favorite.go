package service

import (
	"context"
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"strconv"
)

type FavoriteService struct{}

// FavoriteAction 点赞操作
func (fs *FavoriteService) FavoriteAction(actionType string, userId, videoId int64) error {
	ctx := context.Background()
	videoKey := utils.GetFavoriteVideoKey(videoId)
	userKey := utils.GetFavoriteUserKey(userId)
	switch actionType {
	case "1":
		// 点赞
		pipe := global.RD.TxPipeline()
		defer pipe.Close()
		pipe.SAdd(ctx, videoKey, userId)
		pipe.SAdd(ctx, userKey, videoId)
		if _, err := pipe.Exec(ctx); err != nil {
			pipe.Discard() // 取消提交
			return err
		}
		count, _ := global.RD.SCard(ctx, videoKey).Result()
		if err := repository.GroupApp.VideoRepository.UpdateFavoriteCount(videoId, count); err != nil {
			pipe.Discard()
		}
	case "2":
		// 取消点赞
		pipe := global.RD.TxPipeline()
		defer pipe.Close()
		pipe.SRem(ctx, videoKey, userId)
		pipe.SRem(ctx, userKey, videoId)
		if _, err := pipe.Exec(ctx); err != nil {
			pipe.Discard() // 取消提交
			return err
		}
		count, _ := global.RD.SCard(ctx, videoKey).Result()
		if err := repository.GroupApp.VideoRepository.UpdateFavoriteCount(videoId, count); err != nil {
			pipe.Discard()
			return err
		}
	default:
		return errors.New("点赞类型action_type的值不是0或1）")
	}
	return nil
}

// FavoriteList 获取用户的点赞视频列表
func (fs *FavoriteService) FavoriteList(userId int64) ([]model.Video, error) {
	userKey := utils.GetFavoriteUserKey(userId)
	videoIdStrList, err := global.RD.SMembers(context.Background(), userKey).Result()
	if err != nil {
		return nil, err
	}
	length := len(videoIdStrList)
	videoIds := make([]int64, length)
	for i, v := range videoIdStrList {
		videoIds[i], _ = strconv.ParseInt(v, 10, 64)
	}
	videoList, err := repository.GroupApp.VideoRepository.QueryByIds(videoIds)
	for _, v := range videoList {
		v.IsFavorite = true
	}
	return videoList, err
}
