package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tikstart/http/internal/config"
	"tikstart/http/internal/middleware"
	"tikstart/rpc/user/userClient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userClient.User
	JwtAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userClient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		JwtAuth: middleware.NewJwtAuthMiddleware(c).Handle,
	}
}
