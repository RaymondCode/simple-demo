package model

import "time"

type ModeTime time.Time

type User struct {
	UserID    uint   `gorm:"primaryKey;column:id;"`
	UserName  string `gorm:"column:username;unique;not null"`
	Password  string `gorm:"column:password;not null"`
	Videos    []*Video
	Comments  []*Comment
	Likes     []*Video  `gorm:"many2many:like;joinForeignKey:user_id;joinReferences:video_id;"`
	Fans      []*User   `gorm:"many2many:follow;joinForeignKey:user_id;joinReferences:fan_id;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func (um User) TableName() string {
	return "user"
}

// UserFind 对外查询使用用户模型
type UserFind struct {
	User
	Password string `gorm:"column:password;not null"`
}
