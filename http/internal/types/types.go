// Code generated by goctl. DO NOT EDIT.
package types

type Empty struct {
}

type BasicResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type RegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterResponse struct {
	BasicResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResponse struct {
	BasicResponse
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	IsFollow bool   `json:"is_follow"`
}

type GetUserInfoRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type GetUserInfoResponse struct {
	BasicResponse
	User User `json:"user"`
}

type FollowRequest struct {
	ToUserId   int64  `form:"to_user_id"`
	Token      string `form:"token"`
	ActionType int32  `form:"action_type"`
}

type FollowResponse struct {
	BasicResponse
}

type Video struct {
	Id            int64  `json:"id"`
	Title         string `json:"title"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

type GetVideoListRequest struct {
	LatestTime int64  `form:"latest_time,optional"`
	Token      string `form:"token,optional"`
}

type GetVideoListResponse struct {
	BasicResponse
	Next_time int64   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}

type GetFriendListRequest struct {
	UserId int64  `form:"user_id"`
	Token  string `form:"token"`
}

type GetFriendListResponse struct {
	BasicResponse
	UserList []int64 `json:"user_list"`
}
