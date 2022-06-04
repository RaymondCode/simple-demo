package models

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	Id        int64 `gorm:"column:user_id"`
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
type VideoModel struct {
	Id        int64 `gorm:"column:video_id"`
	UserId    string
	Title     string
	PlayUrl   string
	CoverUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
type CommentModel struct {
	VideoId   int64
	UserId    int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
type FavoriteModel struct {
	UserId    int64
	VideoId   int64
	CreatedAt time.Time
}
type RelationModel struct {
	FollowerId int64
	UserId     int64
	CreatedAt  time.Time
}
