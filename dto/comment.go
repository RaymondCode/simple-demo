package dto

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []*ResponeComment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment *ResponeComment `json:"comment,omitempty"`
}

type ResponeComment struct {
	ID        int64  `json:"id,omitempty"`
	User      User   `json:"user,omitempty"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"create_date,omitempty"`
}
