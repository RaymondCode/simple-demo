package logic

import (
	"context"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFetchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchUserLogic {
	return &FetchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FetchUserLogic) FetchUser(in *user.FetchUserRequest) (*user.FetchUserResponse, error) {
	userId := in.UserId
	var userRecord model.User
	err := l.svcCtx.DB.Where("user_id = ?", userId).First(&userRecord).Error
	if err != nil {
		return nil, err
	}
	return &user.FetchUserResponse{
		UserId:    userRecord.UserId,
		Username:  userRecord.Username,
		Email:     userRecord.Email,
		Signature: userRecord.Signature,
		CreatedAt: userRecord.CreatedAt.Unix(),
		UpdatedAt: userRecord.UpdatedAt.Unix(),
	}, nil
}
