package model

import (
	"time"
)

// TODO: follow and follower
//type User struct {
//	gorm.Model
//	Username    string `gorm:"not null;unique;index"`
//	Password    string `gorm:"not null"`
//	FollowCount int64
//	FanCount    int64
//
//	// many to many
//	Follows []*User `gorm:"many2many:follows;"`                         // 关注列表
//	Fans    []*User `gorm:"many2many:follows;joinForeignKey:follow_id"` // 粉丝列表
//}

type User struct {
	UserId    int64     `gorm:"not null;primarykey;autoIncrement"`
	Username  string    `gorm:"type:varchar(24);not null;uniqueIndex"`
	Password  []byte    `gorm:"type:VARBINARY(60);not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
