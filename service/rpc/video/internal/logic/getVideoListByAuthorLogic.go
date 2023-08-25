package logic

import (
	"context"

	"github.com/ZhouXiinlei/tiktok_startup/service/rpc/video/internal/svc"
	"github.com/ZhouXiinlei/tiktok_startup/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListByAuthorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListByAuthorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListByAuthorLogic {
	return &GetVideoListByAuthorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListByAuthorLogic) GetVideoListByAuthor(in *video.GetVideoListByAuthorRequest) (*video.GetVideoListByAuthorResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetVideoListByAuthorResponse{}, nil
}
