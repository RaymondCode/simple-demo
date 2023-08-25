package logic

import (
	"context"

	"github.com/ZhouXiinlei/tiktok_startup/service/rpc/video/internal/svc"
	"github.com/ZhouXiinlei/tiktok_startup/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentInfoLogic {
	return &GetCommentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentInfoLogic) GetCommentInfo(in *video.GetCommentInfoRequest) (*video.GetCommentInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetCommentInfoResponse{}, nil
}
