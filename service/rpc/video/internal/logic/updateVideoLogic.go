package logic

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tiktok_startup/common/model"

	"tiktok_startup/service/rpc/video/internal/svc"
	"tiktok_startup/service/rpc/video/video"

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

	db := l.svcCtx.Mysql
	err := db.Transaction(func(tx *gorm.DB) error {
		var newVideo model.Video
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", in.Video.Id).First(&newVideo).Error
		if err != nil {
			return err
		}

		newVideo.CommentCount = in.Video.CommentCount
		newVideo.FavoriteCount = in.Video.FavoriteCount

		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&newVideo).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &video.Empty{}, nil
}
