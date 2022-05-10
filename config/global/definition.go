package global

type Configuration struct {
	Server   `yaml:"server"`
	Redis    `yaml:"redis"`
	Database `yaml:"database"`
	Url      `yaml:"url"`
}

type Server struct {
	Port      int    `yaml:"port"`
	Mode      string `yaml:"mode"`
	LimitNum  int    `yaml:"limitNum"`
	UserMongo bool   `yaml:"useMongo"`
	UserRedis bool   `yaml:"useRedis"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Database struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Path         string `yaml:"path"`
	Database     string `yaml:"database"`
	Config       string `yaml:"config"`
	Driver       string `yaml:"driver"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	Log          bool   `yaml:"log"`
	AutoMigrate  bool   `yaml:"autoMigrate"`
}

type Url struct {
	Prefix string `yaml:"prefix"`
}
