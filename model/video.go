package model

import "time"

type Video struct {
	Id            int64     `json:"id,omitempty" gorm:"primaryKey"`
	UserId        int64     `json:"user_id"`
	Author        User      `json:"author" gorm:"foreignKey:user_id;references:id;"`
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	Title         string    `json:"title,omitempty"`
	CreateTime    time.Time `json:"create_time,omitempty"`
	IsFavorite    bool      `json:"is_favorite" gorm:"-"`
}

func (Video) TableName() string {
	return "video_info"
}
