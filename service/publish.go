package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"sort"
)

type PublishService struct{}

// VideoPublishList 获取用户上传的视频列表, 按上传时间倒序展示
func (vs *PublishService) VideoPublishList(userId int64) ([]model.Video, error) {
	//repository.GroupApp.VideoRepository.QueryByIds()
	videoList, err := repository.GroupApp.VideoRepository.QueryVideosByUserId(userId)

	sort.Slice(videoList, func(i, j int) bool {
		return videoList[i].CreateTime.After(videoList[j].CreateTime)
	})

	return videoList, err
}
