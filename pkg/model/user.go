package model

type User struct {
	ID            uint
	UserName      string
	FollowCount   int64
	FollowerCount int64
}

type UserCtl interface {
	QueryUserByID(id uint) *User
}
