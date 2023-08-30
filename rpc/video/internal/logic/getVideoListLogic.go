package logic

import (
	"context"
	"tikstart/common/model"
	"tikstart/rpc/video/internal/svc"
	"tikstart/rpc/video/video"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListLogic) GetVideoList(in *video.GetVideoListRequest) (*video.GetVideoListResponse, error) {
	// todo: add your logic here and delete this line
	var videos []model.Video
	err := l.svcCtx.Mysql.
		Model(&model.Video{}).
		Where("created_at < ?", time.Unix(in.LatestTime, 0)).
		Order("created_at desc").Limit(int(in.Num)).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	var videoList []*video.VideoInfo
	for _, v := range videos {
		videoList = append(videoList, &video.VideoInfo{
			Id:            int64(v.ID),
			AuthorId:      v.AuthorId,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			CreateTime:    v.CreatedAt.Unix(),
		})
	}
	return &video.GetVideoListResponse{VideoList: videoList}, nil
}
