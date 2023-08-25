package model

import (
	"time"
)

type User struct {
	UserId    uint64    `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	Username  string    `gorm:"column:username;type:varchar(24);not null;uniqueIndex"`
	Email     string    `gorm:"column:email;type:varchar(255);not null"`
	Signature string    `gorm:"column:signature;type:varchar(255)"`
}
