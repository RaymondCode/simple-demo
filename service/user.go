package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
)

type UserService struct{}

func (us *UserService) QueryUser(userId int64, token string) (*model.User, error) {
	rawUser, err := repository.GroupApp.UserRepository.QueryUserById(userId)
	if err != nil {
		return nil, err
	}
	return rawUser, nil
}
