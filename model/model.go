package model

import (
	"time"

	"gorm.io/gorm"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// Model 数据库关系实体的基类
type Model struct {
	ID       uint64    `gorm:"common:自增主键"`
	CreateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UpdateAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
}

// Video 视频：数据库实体
type Video struct {
	Model
	gorm.DeletedAt
	VideoID       int64  `gorm:"type:BIGINT;not null;UNIQUE" json:"video_id" validate:""`
	VideoName     string `gorm:"type:varchar(100);not null" json:"video_name" validate:""`
	UserID        int64  `gorm:"type:BIGINT;not null;index:idx_author_id" json:"user_id" validate:""`
	FavoriteCount int32  `gorm:"type:INT;not null;default:0" json:"favorite_count" validate:""`
	CommentCount  int32  `gorm:"type:INT;not null;default:0" json:"comment_count" validate:""`
	PlayURL       string `gorm:"type:varchar(100);not null" json:"play_url" validate:""`
	CoverURL      string `gorm:"type:varchar(100);not null" json:"cover_url" validate:""`
}

// User 用户:数据库实体
type User struct {
	Model
	UserID        int64  `gorm:"type:bigint;unsigned;not null;unique;uniqueIndex:idx_user_id" json:"user_id"`
	UserName      string `gorm:"type:varchar(50);not null;unique;uniqueIndex:idx_user_name" json:"name" validate:"min=6,max=32"`
	PassWord      string `gorm:"type:varchar(50);not null" json:"password" validate:"min=6,max=32"`
	FollowCount   int64  `gorm:"type:bigint;unsigned;not null;default:0" json:"follow_count"`
	FollowerCount int64  `gorm:"type:bigint;unsigned;not null;default:0" json:"follower_count"`
}

// Comment 评论：数据库实体
type Comment struct {
	Model
	UserID  int64  `gorm:"type:BIGINT;not null;index:idx_user_id;评论用户ID" json:"user_id"`
	VideoID int64  `gorm:"type:BIGINT;not null;index:idx_video_id;common:被评论视频ID" json:"video_id" validate:""`
	Content string `gorm:"type:varchar(300);not null;common:评论内容" json:"content"`
}

// Favourite 点赞：数据库实体
type Favourite struct {
	Model
	UserID  int64 `gorm:"type:BIGINT;not nul;index:idx_user_id;common:点赞用户ID" json:"user_id"`
	VideoID int64 `gorm:"type:BIGINT;not null;index:idx_video_id;common:被点赞视频ID" json:"video_id" `
}

// Follow 关注：数据库实体
type Follow struct {
	Model
	FromUserID int64 `gorm:"type:BIGINT;not null;index:idx_user_id;common:粉丝用户ID" json:"from_user_id" validate:""`
	ToUserID   int64 `gorm:"type:BIGINT;not null;index:idx_to_user_id;common:被关注用户ID" json:"to_user_id" validate:""`
}

// VideoQuery 主要提供给查询操作使用
type VideoQuery struct {
	VideoID       int64     `json:"id,omitempty"`
	Author        UserQuery `json:"author,omitempty"`
	FavoriteCount int32     `json:"favorite_count"`
	CommentCount  int32     `json:"comment_count"`
	PlayURL       string    `json:"play_url"`
	CoverURL      string    `json:"cover_url"`
	IsFavorite    bool      `json:"is_favorite"`
}

// UserQuery 主要提供给接口使用
type UserQuery struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
