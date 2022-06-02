package dao

import (
	"github.com/warthecatalyst/douyin/model"
)

func GetUserById(userId int64) (*model.User, error) {
	user := &model.User{}
	if err := db.Where("user_id = ?", userId).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user *model.User) error {
	return db.Updates(user).Error
}
