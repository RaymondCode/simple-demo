package model

type User struct {
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// 对数据库的修改应通过 model.UserCtl 完成
type UserCtl interface {
	QueryUserByID(id int64) *User
}
