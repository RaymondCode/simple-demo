package model

import (
	"fmt"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error

	cfg := dto.GetConfig()

	fmt.Println(cfg)

	if cfg.Env.IsDebug {
		// 开发环境
		DB, err = gorm.Open(
			mysql.Open(
				cfg.MySQL.Local.Username+":"+
					cfg.MySQL.Local.Password+"@tcp("+
					cfg.MySQL.Local.Host+":"+
					cfg.MySQL.Local.Port+")/"+
					cfg.MySQL.Local.Database+"?charset=utf8&parseTime=True&loc=Local"),
			&gorm.Config{
				PrepareStmt:            true,
				SkipDefaultTransaction: true,
			},
		)
	} else {
		// 生产环境
		DB, err = gorm.Open(
			mysql.Open(
				cfg.MySQL.Default.Username+":"+
					cfg.MySQL.Default.Password+"@tcp("+
					cfg.MySQL.Default.Host+":"+
					cfg.MySQL.Default.Port+")/"+
					cfg.MySQL.Default.Database+"?charset=utf8&parseTime=True&loc=Local"),
			&gorm.Config{
				PrepareStmt:            true,
				SkipDefaultTransaction: true,
			},
		)
	}
	if err != nil {
		panic(err)
	}
}
