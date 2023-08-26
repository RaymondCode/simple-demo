package logic

import (
	"context"

	"tiktok_startup/service/rpc/contact/contact"
	"tiktok_startup/service/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoseFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoseFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoseFriendsLogic {
	return &LoseFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoseFriendsLogic) LoseFriends(in *contact.LoseFriendsRequest) (*contact.Empty, error) {
	// todo: add your logic here and delete this line

	return &contact.Empty{}, nil
}
