package dto

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	NextTime   int64  `json:"next_time,omitempty"`
}
