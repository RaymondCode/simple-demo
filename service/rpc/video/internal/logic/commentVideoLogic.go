package logic

import (
	"context"

	"github.com/ZhouXiinlei/tiktok_startup/service/rpc/video/internal/svc"
	"github.com/ZhouXiinlei/tiktok_startup/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentVideoLogic {
	return &CommentVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentVideoLogic) CommentVideo(in *video.CommentVideoRequest) (*video.CommentVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &video.CommentVideoResponse{}, nil
}
