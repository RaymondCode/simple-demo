package model

import "time"

type User struct {
	ID            uint `gorm:"primarykey"`
	Name          string
	FollowCount   int64
	FollowerCount int64
	CreatedAt     time.Time
}
