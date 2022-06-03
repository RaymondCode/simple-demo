package global

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	RD     *redis.Client
	CONFIG config.Server
)
