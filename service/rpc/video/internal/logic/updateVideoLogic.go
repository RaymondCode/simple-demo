package logic

import (
	"context"
	"github.com/RaymondCode/simple-demo/common/model"
	"gorm.io/gorm/clause"

	"github.com/RaymondCode/simple-demo/service/rpc/video/internal/svc"
	"github.com/RaymondCode/simple-demo/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVideoLogic {
	return &UpdateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVideoLogic) UpdateVideo(in *video.UpdateVideoRequest) (*video.Empty, error) {

	tx := l.svcCtx.Mysql.Begin()
	var newVideo model.Video
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", in.Video.Id).First(&newVideo).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	newVideo.CommentCount = in.Video.CommentCount
	newVideo.FavoriteCount = in.Video.FavoriteCount

	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&newVideo).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &video.Empty{}, nil
}
