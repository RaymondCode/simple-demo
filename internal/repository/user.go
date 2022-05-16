package repository

import (
	"github.com/fitenne/youthcampus-dousheng/pkg/model"
	"gorm.io/gorm"
)

// user struct mapped to database
type User struct {
	gorm.Model
	UserName      string `gorm:"index:username,class:FULLTEXT"` // indexed for better authentication peformance
	Password      string
	FollowCount   int64
	FollowerCount int64
}

type userCtl struct{}

var ctl userCtl

func (User) TableName() string {
	return "user"
}

func GetUserCtl() model.UserCtl {
	return &ctl
}

func (ctl *userCtl) QueryUserByID(id uint) *model.User {
	var user []User
	dbProvider.GetDB().Limit(1).Find(&user)
	if len(user) == 0 {
		return nil
	}

	return &model.User{
		ID:            user[0].ID,
		UserName:      user[0].UserName,
		FollowCount:   user[0].FollowCount,
		FollowerCount: user[0].FollowerCount,
	}
}
