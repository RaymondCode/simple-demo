package logic

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tikstart/common/model"
	"tikstart/service/rpc/video/internal/svc"
	"tikstart/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoLogic {
	return &FavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteVideoLogic) FavoriteVideo(in *video.FavoriteVideoRequest) (*video.Empty, error) {
	err := l.svcCtx.Mysql.Transaction(func(tx *gorm.DB) error {
		NewFavorite := model.Favorite{}
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ? AND video_id = ?", in.UserId, in.VideoId).
			First(&NewFavorite).
			Error
		if err == nil {
			return nil
		}
		if err != gorm.ErrRecordNotFound {
			return err
		}
		NewFavorite.VideoId = in.VideoId
		NewFavorite.UserId = in.UserId
		err = tx.Create(&NewFavorite).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.Video{}).
			Where("id = ?", in.VideoId).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).
			Error
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
