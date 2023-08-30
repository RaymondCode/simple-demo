package logic

import (
	"context"
	"tikstart/rpc/user/internal/svc"
	"tikstart/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user.PingRequest) (*user.PingResponse, error) {
	return &user.PingResponse{
		Pong: "pong",
	}, nil
}
