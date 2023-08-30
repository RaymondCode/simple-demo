package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorId      int64  `gorm:"not null;index"`
	Title         string `gorm:"not null;index"`
	PlayUrl       string `gorm:"not null"`
	CoverUrl      string `gorm:"not null"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64

	// has many
	Comments  []Comment
	Favorites []Favorite
}

const (
	PopularVideoStandard = 1000 // 拥有超过 1000 个赞或 1000 个评论的视频成为热门视频，有特殊处理
)

func IsPopularVideo(favoriteCount, commentCount int64) bool {
	return favoriteCount >= PopularVideoStandard || commentCount >= PopularVideoStandard
}

type Comment struct {
	gorm.Model
	UserId  int64
	VideoId int64
	Content string `gorm:"not null"`
}

type Favorite struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;"`
	VideoId int64 `gorm:"column:video_id;"`

	// belongs to
	Video Video
}
