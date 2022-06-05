package model

import "time"

type Comment struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	VideoId   int64
	Content   string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
