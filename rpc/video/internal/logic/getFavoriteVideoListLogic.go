package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tikstart/common/model"
	"tikstart/rpc/video/internal/svc"
	"tikstart/rpc/video/video"
)

type GetFavoriteVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoListLogic {
	return &GetFavoriteVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteVideoListLogic) GetFavoriteVideoList(in *video.GetFavoriteVideoListRequest) (*video.GetFavoriteVideoListResponse, error) {
	var favoriteVideoList []*model.Favorite

	if err := l.svcCtx.Mysql.
		Where("user_id = ?", in.UserId).
		Preload("Video").
		Order("created_at desc").
		Find(&favoriteVideoList).Error; err != nil {
		return nil, err
	}

	videoList := make([]*video.VideoInfo, 0, len(favoriteVideoList))
	for _, v := range favoriteVideoList {
		if v.Video.ID == 0 {
			continue
		}
		videoInfo := &video.VideoInfo{
			Id:            int64(v.Video.ID),
			AuthorId:      v.Video.AuthorId,
			Title:         v.Video.Title,
			PlayUrl:       v.Video.PlayUrl,
			CoverUrl:      v.Video.CoverUrl,
			FavoriteCount: v.Video.FavoriteCount,
			CommentCount:  v.Video.CommentCount,
		}
		videoList = append(videoList, videoInfo)
	}
	return &video.GetFavoriteVideoListResponse{
		VideoList: videoList,
	}, nil
}
