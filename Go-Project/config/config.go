package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
	//接下来Redis这些也在这相似地定义
}
