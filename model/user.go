package model

import "time"

type ModeTime time.Time

const (
	timeFormat = "2022-06-09 17:04:05"
	zone       = "Asia/Shanghai"
)

type User struct {
	ID        uint     `gorm:"primaryKey;column:id"`
	UserName  string   `gorm:"column:username;unique;not null"`
	Password  string   `gorm:"column:password;not null"`
	CreatedAt ModeTime `gorm:"column:created_at"`
	UpdatedAt ModeTime `gorm:"column:updated_at"`
}

func (um User) TableName() string {
	return "users"
}

// UserFind 对外查询使用用户模型
type UserFind struct {
	User
	Password string `gorm:"column:password;not null"`
}