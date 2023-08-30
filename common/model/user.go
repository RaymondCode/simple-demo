package model

import (
	"gorm.io/gorm"
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
	UserId   int64  `gorm:"not null;primarykey;autoIncrement"`
	Username string `gorm:"type:varchar(24);not null;uniqueIndex"`
	Password []byte `gorm:"type:VARBINARY(60);not null"`

	FollowingCount int64
	FollowerCount  int64

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt

	// 添加关注关系
	Following []*User `gorm:"many2many:follow;foreignKey:UserId;joinForeignKey:FollowerID;References:UserId;JoinReferences:FollowedID"`
	Followers []*User `gorm:"many2many:follow;foreignKey:UserId;joinForeignKey:FollowedID;References:UserId;JoinReferences:FollowerID"`
}

type Follow struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	FollowerID int64     `gorm:"not null"`
	FollowedID int64     `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`

	// 添加外键约束
	Follower User `gorm:"foreignKey:FollowerID"`
	Followed User `gorm:"foreignKey:FollowedID"`
}
