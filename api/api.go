package api

import (
	"github.com/warthecatalyst/douyin/model"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var OK = Response{
	StatusCode: 0,
	StatusMsg:  "success",
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

const (
	FavoriteAction   = 1
	UnFavoriteAction = 2
	FollowAction     = 1
	UnfollowAction   = 2
)
//一个自己视频的返回
//type PublishList VideoList1
//type VideoList1 struct {
//	Response
//	VideoLists []model.VideoQuery `json:"video_list"`
//}
//一个点赞视频的返回
//type AllFavoriteVideos AllFavoriteVideo
//type AllFavoriteVideo struct {
//	Response
//	FavoriteVideos []Video  `json:"unknown"`
//}
//type VideoList2 []Video


//Feed流
type Feed struct {
	Response
	NextTime   int64            `json:"next_time"`
	VideoLists []model.VideoQuery `json:"video_list"`
}
//主体video list的结构体
type VideoList struct {
	Response
	VideoLists []model.VideoQuery `json:"video_list"`
}
//返回一个点赞的视频列表
type FavoriteList VideoList
//返回一个最近的发布
type PublishList VideoList

type UserInfo struct {
	Response
	User model.UserQuery
}