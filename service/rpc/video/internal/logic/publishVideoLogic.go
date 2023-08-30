package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"tikstart/service/rpc/video/common/model"
	"tikstart/service/rpc/video/internal/svc"
	"tikstart/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoLogic) PublishVideo(in *video.PublishVideoRequest) (*video.Empty, error) {
	newVideo := &model.Video{
		AuthorId: in.Video.AuthorId,
		Title:    in.Video.Title,
		PlayUrl:  in.Video.PlayUrl,
		CoverUrl: in.Video.CoverUrl,
	}
	if err := l.svcCtx.Mysql.Create(newVideo).Error; err != nil {
		return nil, status.Error(1000, err.Error())
	}
	return &video.Empty{}, nil
}
