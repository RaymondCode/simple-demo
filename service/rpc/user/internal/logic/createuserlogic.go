package logic

import (
	"context"
	"fmt"
	"user/internal/model"

	"user/internal/svc"
	"user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	username := in.Username
	email := in.Email
	signature := in.Signature

	fmt.Printf("in: %v\n", in)

	userRecord := model.User{
		Username:  username,
		Email:     email,
		Signature: signature,
	}

	err := l.svcCtx.DB.Create(&userRecord).Error
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		UserId: userRecord.UserId,
	}, nil
}
