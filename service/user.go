package service

import (
	"errors"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
	"gorm.io/gorm"
)

type UserService struct{}

var (
	userService = &UserService{}
)

func NewUserServiceInstance() *UserService {
	return userService
}

func (u *UserService) UserExist(userId int64) (bool, error) {
	user, err := dao.NewUserDaoInstance().GetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return user != nil, nil
}

func (u *UserService) GetUserFromUserId(userId int64) (*api.User, error) {
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
