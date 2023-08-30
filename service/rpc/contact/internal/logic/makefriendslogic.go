package logic

import (
	"context"
	"gorm.io/gorm"
	"tikstart/common/model"

	"tikstart/service/rpc/contact/contact"
	"tikstart/service/rpc/contact/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeFriendsLogic {
	return &MakeFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeFriendsLogic) MakeFriends(in *contact.MakeFriendsRequest) (*contact.Empty, error) {
	err := l.svcCtx.Mysql.Transaction(func(tx *gorm.DB) error {
		newFriendsA := model.Friend{
			UserId:   in.UserAId,
			FriendId: in.UserBId,
		}

		newFriendsB := model.Friend{
			UserId:   in.UserBId,
			FriendId: in.UserAId,
		}
		if err := tx.Create(&newFriendsA).Error; err != nil {
			return err
		}

		if err := tx.Create(&newFriendsB).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &contact.Empty{}, nil
}
