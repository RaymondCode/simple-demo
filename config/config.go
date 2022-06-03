package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" yaml:"redis"`
	JWT   JWT   `mapstructure:"jwt" yaml:"jwt"`
}
