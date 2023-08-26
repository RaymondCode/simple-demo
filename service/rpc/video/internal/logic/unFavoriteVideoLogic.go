package logic

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tiktok_startup/service/rpc/video/common/model"
	"tiktok_startup/service/rpc/video/internal/svc"
	"tiktok_startup/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFavoriteVideoLogic {
	return &UnFavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnFavoriteVideoLogic) UnFavoriteVideo(in *video.UnFavoriteVideoRequest) (*video.Empty, error) {

	err := l.svcCtx.Mysql.Transaction(func(tx *gorm.DB) error {
		NewFavorite := model.Favorite{}
		err := tx.Clauses(clause.
			Locking{Strength: "UPDATE"}).Where("user_id = ? AND video_id = ?", in.UserId, in.VideoId).
			First(&NewFavorite).
			Error
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		err = tx.Where("user_id = ? And video_id = ?", in.UserId, in.VideoId).Delete(&NewFavorite).Error
		if err != nil {
			return err
		}
		err = tx.Model(&model.Video{}).Where("id = ?", in.VideoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
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
