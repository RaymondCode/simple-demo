package logic

import (
	"context"
	"tikstart/service/rpc/video/common/model"

	"tikstart/service/rpc/video/internal/svc"
	"tikstart/service/rpc/video/video"

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
	var comment model.Comment
	err := l.svcCtx.Mysql.Where("id = ?", in.CommentId).First(&comment).Error
	if err != nil {
		return nil, err
	}

	return &video.GetCommentInfoResponse{
		Id:          int64(comment.ID),
		UserId:      comment.UserId,
		Content:     comment.Content,
		CreatedTime: comment.CreatedAt.Unix(),
	}, nil
}
