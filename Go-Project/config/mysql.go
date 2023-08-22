package config

type Mysql struct {
	Path         string `mapstructure:"path" yaml:"path"`                     // 服务器地址
	Port         string `mapstructure:"port" yaml:"port"`                     // 端口
	Config       string `mapstructure:"config" yaml:"config"`                 // 高级配置
	Dbname       string `mapstructure:"db-name" yaml:"db-name"`               // 数据库名
	Username     string `mapstructure:"username" yaml:"username"`             // 数据库用户名
	Password     string `mapstructure:"password" yaml:"password"`             // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}
