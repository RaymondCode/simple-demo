package global

import (
	"github.com/life-studied/douyin-simple/config"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	CONFIG config.Server
)
