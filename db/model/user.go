package model

import (
	"context"
	"time"
)

type User struct {
	ID            int64 `gorm:"primarykey"`
	Name          string
	Password      string
	FollowCount   int64
	FollowerCount int64
	CreatedAt     time.Time
}

//CteateUser create user info
func CreateUser(ctx context.Context, user *User) error {
	if err := DB.Table("user").WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

//QueryUser Quert User By Name
func QueryUserByName(ctx context.Context, username string) (*User, error) {
	var userInfo *User
	if err := DB.Table("user").WithContext(ctx).Where("name=?", username).Find(&userInfo).Error; err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

func QueryUserById(ctx context.Context, id int64) (*User, error) {
	var userInfo *User
	if err := DB.Table("user").WithContext(ctx).Where("id=?", id).Find(&userInfo).Error; err != nil {
		return userInfo, err
	}
	return userInfo, nil
}
