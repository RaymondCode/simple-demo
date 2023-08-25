package logic

import (
	"context"

	"tiktok_startup/service/rpc/video/internal/svc"
	"tiktok_startup/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoCommentLogic {
	return &DeleteVideoCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteVideoCommentLogic) DeleteVideoComment(in *video.DeleteVideoCommentRequest) (*video.Empty, error) {
	// todo: add your logic here and delete this line

	return &video.Empty{}, nil
}
