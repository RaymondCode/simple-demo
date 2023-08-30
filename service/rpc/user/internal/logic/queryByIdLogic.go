package logic

import (
	"context"

	"tikstart/service/rpc/user/internal/svc"
	"tikstart/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByIdLogic {
	return &QueryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryByIdLogic) QueryById(in *user.QueryByIdRequest) (*user.QueryResponse, error) {
	// todo: add your logic here and delete this line

	return &user.QueryResponse{}, nil
}
