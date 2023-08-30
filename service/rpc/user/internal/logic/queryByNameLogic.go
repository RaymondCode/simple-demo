package logic

import (
	"context"

	"tikstart/service/rpc/user/internal/svc"
	"tikstart/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByNameLogic {
	return &QueryByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryByNameLogic) QueryByName(in *user.QueryByNameRequest) (*user.QueryResponse, error) {
	// todo: add your logic here and delete this line

	return &user.QueryResponse{}, nil
}
