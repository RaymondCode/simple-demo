package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   int64  `gorm:"index"`
	ToUserId int64  `gorm:"index"`
	Content  string `gorm:"not null"`
}

type Friend struct {
	gorm.Model
	UserId   int64 `gorm:"index"`
	FriendId int64 `gorm:"index"`
}
