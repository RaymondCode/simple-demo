package model

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"-"`
}
