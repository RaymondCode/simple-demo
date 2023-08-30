package logic

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"tikstart/service/rpc/video/common/model"

	"tikstart/service/rpc/video/internal/svc"
	"tikstart/service/rpc/video/video"

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
	comment := model.Comment{
		VideoId: in.VideoId,
		UserId:  in.UserId,
		Content: in.Content,
	}

	err := l.svcCtx.Mysql.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&comment).Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Video{}).
			Where("id = ?", in.VideoId).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Update("comment_count", gorm.Expr("comment_count + ?", 1)).
			Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &video.CommentVideoResponse{
		Id:          int64(comment.ID),
		UserId:      comment.UserId,
		Content:     comment.Content,
		CreatedTime: comment.CreatedAt.Unix(),
	}, nil
}
