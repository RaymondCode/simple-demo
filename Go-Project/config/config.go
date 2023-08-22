package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
}
