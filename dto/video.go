package dto

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList    []Video `json:"video_list,omitempty"`
	NextTime     int64   `json:"next_time,omitempty"`
	CommentCount int64   `json:"comment_count,omitempty"`
}
