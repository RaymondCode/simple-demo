package repository

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
)

type UserRepository struct{}

func (ur *UserRepository) QueryUserById(id int64) (*model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
