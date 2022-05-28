package config

type Redis struct {
	DB       int    `mapstructure:"db" yaml:"db"`             // redis的哪个数据库
	Addr     string `mapstructure:"addr" yaml:"addr"`         // 服务器地址:端口
	Password string `mapstructure:"password" yaml:"password"` // 密码
}
