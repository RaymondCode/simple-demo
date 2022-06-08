package service

import (
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
)

//与video有关的service层实现

type VideoService struct{}

var videoService = &VideoService{}

func (*VideoService) getUserFromVideoId(videoId int64) (*api.User, error) {
	userId, err := dao.NewVideoDaoInstance().GetUserIdFromVideoId(videoId)
	if err != nil {
		return nil, err
	}
	userModel, err := dao.NewUserDaoInstance().GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return &api.User{
		Id:            userId,
		Name:          userModel.UserName,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,
		IsFollow:      false,
	}, nil
}
