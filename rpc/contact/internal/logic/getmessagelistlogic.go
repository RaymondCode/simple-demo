package logic

import (
	"context"
	"tikstart/common/model"
	"tikstart/rpc/contact/contact"
	"tikstart/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageListLogic) GetMessageList(in *contact.GetMessageListRequest) (*contact.GetMessageListResponse, error) {
	var messages []model.Message
	err := l.svcCtx.Mysql.Where("from_id = ?", in.FromId).Where("to_user_id = ?", in.ToId).Find(&messages).Error
	if err != nil {
		return nil, err
	}

	var messageList []*contact.Message
	for _, message := range messages {
		messageList = append(messageList, &contact.Message{
			Id:         int64(message.ID),
			Content:    message.Content,
			CreateTime: message.CreatedAt.Unix(),
			FromId:     message.FromId,
			ToId:       message.ToUserId,
		})
	}
	return &contact.GetMessageListResponse{
		Messages: messageList,
	}, nil
}
