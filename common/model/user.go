package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId   int64  `gorm:"not null;primarykey;autoIncrement"`
	Username string `gorm:"type:varchar(24);not null;uniqueIndex"`
	Password []byte `gorm:"type:VARBINARY(60);not null"`

	FollowingCount int64
	FollowerCount  int64

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt gorm.DeletedAt

	Following []*User `gorm:"many2many:follow;foreignKey:UserId;joinForeignKey:FollowerID;References:UserId;JoinReferences:FollowedID"`
	Followers []*User `gorm:"many2many:follow;foreignKey:UserId;joinForeignKey:FollowedID;References:UserId;JoinReferences:FollowerID"`
}

type Follow struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	FollowerID int64     `gorm:"not null"`
	FollowedID int64     `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`

	Follower User `gorm:"foreignKey:FollowerID"`
	Followed User `gorm:"foreignKey:FollowedID"`
}
