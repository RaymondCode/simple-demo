package dao

import (
	"errors"
	"github.com/warthecatalyst/douyin/model"
	"gorm.io/gorm"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func (u *UserDao) UpdateUser(user *model.User) error {
	return db.Updates(user).Error
}

func (u *UserDao) GetUserByUsername(userName string) (*model.User, error) {
	userInfo := &model.User{}
	if err := db.Where("user_name = ?", userName).First(userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}
	return userInfo, nil
}

func (u *UserDao) AddUser(user *model.User) error {
	return db.Create(user).Error
}
