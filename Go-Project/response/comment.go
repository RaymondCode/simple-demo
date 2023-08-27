package response

// User_Response 用户信息的响应结构体
type User_Response struct {
	ID            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   uint   `json:"follow_count,omitempty"`
	FollowerCount uint   `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// Comment_Response 评论信息的响应结构体
type Comment_Response struct {
	ID           uint          `json:"id,omitempty"`
	Content      string        `json:"content,omitempty"`
	CreateDate   string        `json:"create_date,omitempty"`
	Userresponse User_Response `json:"user,omitempty"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment_Response `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment_Response `json:"comment,omitempty"`
}
