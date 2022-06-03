package dao

import (
	"github.com/warthecatalyst/douyin/model"
	"sync"
)

type UserDao struct{}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (u *UserDao) GetUserById(userId int64) (*model.User, error) {
	user := &model.User{}
	if err := db.Where("user_id = ?", userId).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserDao) UpdateUser(user *model.User) error {
	return db.Updates(user).Error
}
