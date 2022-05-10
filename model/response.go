package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type CommentResponse struct {
	Comment
}

type VideoResponse struct {
	Video
	FavoriteCount int64 `json:"favorite_count,omitempty"  `
	CommentCount  int64 `json:"comment_count,omitempty"    `
	IsFavorite    bool  `json:"is_favorite,omitempty" `
}

type UserInfo struct {
	User
	FollowCount   int64 `json:"follow_count,omitempty"`
	FollowerCount int64 `json:"follower_count,omitempty"`
	IsFollow      bool  `json:"is_follow,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User UserInfo `json:"user"`
}

type UserListResponse struct {
	Response
	UserList []UserInfo `json:"user_list"`
}

type VideoListResponse struct {
	Response
	VideoList []VideoResponse `json:"video_list"`
}

type FeedResponse struct {
	Response
	VideoList []VideoResponse `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []CommentResponse `json:"comment_list,omitempty"`
}
