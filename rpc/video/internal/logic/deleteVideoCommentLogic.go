package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tikstart/common/model"
	"tikstart/rpc/video/internal/svc"
	"tikstart/rpc/video/video"
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
	if err := l.svcCtx.Mysql.Transaction(func(tx *gorm.DB) error {
		var comment model.Comment
		if err := tx.Where("id = ?", in.CommentId).First(&comment).Error; err != nil {
			return err
		}

		if err := tx.
			Where("id = ?", in.CommentId).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Delete(&model.Comment{}).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Video{}).
			Where("id = ?", comment.VideoId).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Update("comment_count", gorm.Expr("comment_count - ?", 1)).
			Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return &video.Empty{}, nil
}
