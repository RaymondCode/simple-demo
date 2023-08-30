package logic

import (
	"context"

	"tikstart/service/rpc/user/internal/svc"
	"tikstart/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateRequest) (*user.CreateResponse, error) {
	// todo: add your logic here and delete this line

	return &user.CreateResponse{}, nil
}
