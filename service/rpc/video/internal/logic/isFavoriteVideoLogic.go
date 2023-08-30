package logic

import (
	"context"
	"gorm.io/gorm"
	"tikstart/service/rpc/video/common/model"

	"tikstart/service/rpc/video/internal/svc"
	"tikstart/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavoriteVideoLogic {
	return &IsFavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFavoriteVideoLogic) IsFavoriteVideo(in *video.IsFavoriteVideoRequest) (*video.IsFavoriteVideoResponse, error) {
	// todo: add your logic here and delete this line
	db := l.svcCtx.Mysql
	err := db.Where("user_id = ? AND video_id = ?", in.UserId, in.VideoId).First(&model.Favorite{}).Error
	if err == gorm.ErrRecordNotFound {
		return &video.IsFavoriteVideoResponse{
			IsFavorite: false,
		}, nil
	}
	return &video.IsFavoriteVideoResponse{
		IsFavorite: true,
	}, nil
}
