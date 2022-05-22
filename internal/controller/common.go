package controller

import "github.com/fitenne/youthcampus-dousheng/pkg/model"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Comment struct {
	Id         int64      `json:"id,omitempty"`
	User       model.User `json:"user"`
	Content    string     `json:"content,omitempty"`
	CreateDate string     `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
