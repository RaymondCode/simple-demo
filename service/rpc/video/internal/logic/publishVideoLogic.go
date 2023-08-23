package logic

import (
	"context"

	"github.com/RaymondCode/simple-demo/service/rpc/video/internal/svc"
	"github.com/RaymondCode/simple-demo/service/rpc/video/video"

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
	// todo: add your logic here and delete this line

	return &video.Empty{}, nil
}
