package service

import (
	"github.com/warthecatalyst/douyin/model"
)

type PublishService struct {
}

var publishService = &PublishService{}

// 构造 Video 切片
func newVideoList(videos []model.Video) []model.VideoQuery {
	var v []model.VideoQuery
	for _, i := range videos {
		v = append(v, model.VideoQuery{
			VideoID:       i.VideoID,
			PlayURL:       i.PlayURL,
			CoverURL:      i.CoverURL,
			CommentCount:  i.CommentCount,
			FavoriteCount: i.FavoriteCount,
		})
	}

	return v
}
