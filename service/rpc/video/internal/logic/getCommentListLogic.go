package logic

import (
	"context"
	"tikstart/common/model"
	"tikstart/service/rpc/video/internal/svc"
	"tikstart/service/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *video.GetCommentListRequest) (*video.GetCommentListResponse, error) {
	var comments []*model.Comment
	if err := l.svcCtx.Mysql.
		Where("video_id = ?", in.VideoId).
		Limit(model.PopularVideoStandard).
		Order("created_at").
		Find(&comments).Error; err != nil {
		return nil, err
	}

	commentList := make([]*video.Comment, 0, len(comments))
	for _, v := range comments {
		commentList = append(commentList, &video.Comment{
			Id:         int64(v.ID),
			AuthorId:   v.UserId,
			CreateTime: v.CreatedAt.Unix(),
			Content:    v.Content,
		})
	}

	return &video.GetCommentListResponse{
		CommentList: commentList,
	}, nil
}
