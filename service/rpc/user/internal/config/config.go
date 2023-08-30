package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MySQL struct {
		Host        string
		Port        int
		User        string
		Password    string
		Database    string
		TablePrefix string
	}
}
