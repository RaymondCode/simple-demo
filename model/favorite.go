package model

type Favorite struct {
	Id      int64 `json:"id"  gorm:"primaryKey"`
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

func (Favorite) TableName() string {
	return "video_favorite"
}
