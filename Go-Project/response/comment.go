package response

import "github.com/life-studied/douyin-simple/model"

// User_Response 用户信息的响应结构体
type User_Response struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment model.Comment `json:"comment,omitempty"`
}
