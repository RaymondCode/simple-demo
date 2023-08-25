package common

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MYSQL struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
}
