package logic

import (
	"context"
	"tikstart/common/model"

	"tikstart/service/rpc/contact/contact"
	"tikstart/service/rpc/contact/internal/svc"

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
	tx := l.svcCtx.Mysql.Begin()

	if err := tx.Where("user_id = ? AND friend_id = ?", in.UserAId, in.UserBId).Delete(&model.Friend{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Where("user_id = ? AND friend_id = ?", in.UserBId, in.UserAId).Delete(&model.Friend{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &contact.Empty{}, nil
}
