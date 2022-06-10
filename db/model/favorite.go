package model

import "time"

type Favorite struct {
	ID        int64 `gorm:"primarykey"`
	UserId    int64
	VideoId   int64
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
