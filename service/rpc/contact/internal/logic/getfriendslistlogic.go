package logic

import (
	"context"

	"tiktok_startup/service/rpc/contact/contact"
	"tiktok_startup/service/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsListLogic {
	return &GetFriendsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendsListLogic) GetFriendsList(in *contact.GetFriendsListRequest) (*contact.GetFriendsListResponse, error) {
	// todo: add your logic here and delete this line

	return &contact.GetFriendsListResponse{}, nil
}
