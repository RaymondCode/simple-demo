package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc  zrpc.RpcClientConf
	VideoRpc zrpc.RpcClientConf
	JwtAuth  struct {
		Secret string
		Expire int64
	}
}
