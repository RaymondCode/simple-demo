package model

import (
	"context"
	"time"
)

type User struct {
	ID            string `gorm:"primarykey"`
	Name          string
	PassWord      string
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
func QueryUser(ctx context.Context, username string) ([]*User, error) {
	var userinfo []*User
	if err := DB.Table("user").WithContext(ctx).Where("name=?", username).Find(&userinfo).Error; err != nil {
		return userinfo, err
	}
	return userinfo, nil
}

func QueryUserById(ctx context.Context, id string) ([]*User, error) {
	var userinfo []*User
	if err := DB.Table("user").WithContext(ctx).Where("id=?", id).Find(&userinfo).Error; err != nil {
		return userinfo, err
	}
	return userinfo, nil
}
