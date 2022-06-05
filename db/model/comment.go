package model

import "time"

type Comment struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	VideoId   int64
	Content   string
	CreatedAt time.Time
}
