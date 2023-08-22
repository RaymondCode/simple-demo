// Package dao -----------------------------
// @file      : common.go
// @author    : Yunyin
// @contact   : yunyin_jayyi@qq.com
// @time      : 2023/8/19 19:45
// -------------------------------------------
package dao

import "time"

type User struct {
	ID       int64  `gorm:"primary_key;column:id;comment:'用户ID'" json:"id"`
	Name     string `gorm:"column:name;not null;comment:'用户名'" json:"name"`
	Password string `gorm:"column:password;not null;comment:'密码'" json:"password"`
}

type Video struct {
	ID          int64     `gorm:"primary_key;column:id;comment:'视频ID'" json:"id"`
	AuthorID    int64     `gorm:"column:author_id;not null;comment:'作者ID'" json:"author_id"`
	PlayURL     string    `gorm:"column:play_url;not null;comment:'播放链接'" json:"play_url"`
	CoverURL    string    `gorm:"column:cover_url;not null;comment:'封面链接'" json:"cover_url"`
	Title       string    `gorm:"column:title;not null;comment:'标题'" json:"title"`
	PublishTime time.Time `gorm:"column:publish_time;not null;comment:'发布时间戳'" json:"publish_time"`
}

type Follow struct {
	UserID       int64 `gorm:"column:user_id;not null;comment:'用户ID'" json:"user_id"`
	FollowUserID int64 `gorm:"column:follow_user_id;not null;comment:'被关注的用户ID'" json:"follow_user_id"`

	// Foreign key references
	User       User `gorm:"foreignkey:UserID" json:"-"`
	FollowUser User `gorm:"foreignkey:FollowUserID" json:"-"`
}

type Comment struct {
	ID         int64     `gorm:"primary_key;column:id;comment:'评论ID'" json:"id"`
	UserID     int64     `gorm:"column:user_id;not null;comment:'用户ID'" json:"user_id"`
	VideoID    int64     `gorm:"column:video_id;not null;comment:'视频ID'" json:"video_id"`
	Content    string    `gorm:"column:content;not null;comment:'评论内容'" json:"content"`
	CreateDate time.Time `gorm:"column:create_date;not null;comment:'创建日期'" json:"create_date"`

	// Foreign key references
	User  User  `json:"user" gorm:"foreignKey:user_id;references:id;"`
	Video Video `gorm:"foreignkey:VideoID" json:"-"`
}

type Like struct {
	ID      int64 `gorm:"column:id;not null;comment:'主键ID'" json:"id"`
	UserID  int64 `gorm:"column:user_id;not null;comment:'点赞者ID'" json:"user_id"`
	VideoID int64 `gorm:"column:video_id;not null;comment:'视频ID'" json:"video_id"`

	// Foreign key references
	User  User  `gorm:"foreignkey:UserID" json:"user"`
	Video Video `gorm:"foreignkey:VideoID" json:"-"`
}
