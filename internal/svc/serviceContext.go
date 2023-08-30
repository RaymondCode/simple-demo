package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"tikstart/internal/config"
	"tikstart/internal/middleware"
)

type ServiceContext struct {
	Config  config.Config
	JwtAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		JwtAuth: middleware.NewJwtAuthMiddleware().Handle,
	}
}
