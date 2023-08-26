package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql MysqlConf
}
type MysqlConf struct {
	Address     string
	Username    string
	Password    string
	DBName      string
	TablePrefix string
}
