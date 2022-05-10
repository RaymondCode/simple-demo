package model

import "gorm.io/gorm"

// TODO: https://gorm.io/zh_CN/docs/indexes.html

type Video struct {
	gorm.Model
	Id       int64 `json:"id,omitempty" `
	AuthorID int64
	Author   User   `json:"author"     `
	PlayUrl  string `json:"play_url"            `
	CoverUrl string `json:"cover_url,omitempty"   `
}

type Comment struct {
	Id         int64 `json:"id,omitempty"`
	UserID     int64
	User       User   `json:"user" `
	Content    string `json:"content,omitempty" `
	CreateDate string `json:"create_date,omitempty" `
}

type User struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty" `
	// TODO: ignore Password in JSON
	Password       string `json:"-" `
	PasswordHashed string `json:"-" `
}

type Follow struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	FollowerId int64  `json:"follower_id,omitempty" ` // 关注人
	FolloweeId int64  `json:"followee_id,omitempty" ` // 被关注人
	IsFollow   bool   `json:"is_follow,omitempty"`
}
