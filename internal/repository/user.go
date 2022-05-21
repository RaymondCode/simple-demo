package repository

/*
 * 此文件实现与 User 相关数据库的操作
 * 实现 UserCtl 接口
 * 对 User 表的操作应当通过 UserCtl 接口完成
 */

import (
	"time"

	"github.com/fitenne/youthcampus-dousheng/pkg/model"
	"gorm.io/gorm"
)

// user struct mapped to database
type User struct {
	ID            int64  `gorm:"primarykey"`
	UserName      string `gorm:"index:username,class:FULLTEXT,size:256"` // indexed for better authentication peformance
	Password      string `gorm:"size:256"`
	FollowCount   int64
	FollowerCount int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type userCtl struct{}

var ctl userCtl

// 指定 User 对应的数据库表名
func (User) TableName() string {
	return "users"
}

func GetUserCtl() model.UserCtl {
	return &ctl
}

func (ctl *userCtl) QueryUserByID(id int64) *model.User {
	var user []User
	dbProvider.GetDB().Limit(1).Find(&user)
	if len(user) == 0 {
		return nil
	}

	return &model.User{
		ID:            int64(user[0].ID),
		Name:          user[0].UserName,
		FollowCount:   user[0].FollowCount,
		FollowerCount: user[0].FollowerCount,
	}
}
