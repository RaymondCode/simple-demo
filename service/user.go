package service

import (
	"errors"
	"github.com/warthecatalyst/douyin/dao"
	"gorm.io/gorm"
)

func UserExist(userId int64) (bool, error) {
	user, err := dao.NewUserDaoInstance().GetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return user != nil, nil
}
