package app

import (
	"context"
	"tikstart/http/internal/svc"
	"tikstart/http/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(req *types.Empty) (resp *types.BasicResponse, err error) {
	return &types.BasicResponse{
		StatusCode: 0,
		StatusMsg:  "pong",
	}, nil
}
