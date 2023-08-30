package logic

import (
	"context"
	"tikstart/common/model"

	"tikstart/service/rpc/contact/contact"
	"tikstart/service/rpc/contact/internal/svc"

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
	var result []model.Friend

	err := l.svcCtx.Mysql.Where("user_id = ?", in.UserId).Select("friend_id").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &contact.GetFriendsListResponse{
		FriendsId: func() []int64 {
			var friendsId []int64
			for _, v := range result {
				friendsId = append(friendsId, v.FriendId)
			}
			return friendsId
		}(),
	}, nil
}
