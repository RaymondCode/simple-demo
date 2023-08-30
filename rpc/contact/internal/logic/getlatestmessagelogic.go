package logic

import (
	"context"
	"tikstart/common/model"
	"tikstart/rpc/contact/contact"
	"tikstart/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLatestMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLatestMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestMessageLogic {
	return &GetLatestMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLatestMessageLogic) GetLatestMessage(in *contact.GetLatestMessageRequest) (*contact.GetLatestMessageResponse, error) {
	result := model.Message{}

	l.svcCtx.Mysql.
		Where("from_id = ? and to_user_id = ?", in.UserAId, in.UserBId).
		Or("from_id = ? and to_user_id = ?", in.UserBId, in.UserAId).
		Order("created_at desc").
		First(&result)

	l.Logger.Info("GetLatestMessage", result)

	return &contact.GetLatestMessageResponse{
		Message: &contact.Message{
			Id:         int64(result.ID),
			Content:    result.Content,
			CreateTime: result.CreatedAt.Unix(),
			FromId:     result.FromId,
			ToId:       result.ToUserId,
		},
	}, nil
}
