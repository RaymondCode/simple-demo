package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"not null;unique;index"`
	Password    string `gorm:"not null"`
	FollowCount int64
	FanCount    int64

	// many to many
	Follows []*User `gorm:"many2many:follows;"`                         // 关注列表
	Fans    []*User `gorm:"many2many:follows;joinForeignKey:follow_id"` // 粉丝列表
}

const (
	PopularUserStandard = 1000 // 拥有超过 1000 个粉丝的用户成为大V，有特殊处理
)

func IsPopularUser(fanCount int64) bool {
	return fanCount >= PopularUserStandard
}
