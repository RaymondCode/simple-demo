package model

import (
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error

	if dto.Conf.Env.IsDebug {
		// 开发环境
		DB, err = gorm.Open(
			mysql.Open(
				dto.Conf.MySQL.Local.Username+":"+
					dto.Conf.MySQL.Local.Password+"@tcp("+
					dto.Conf.MySQL.Local.Host+":"+
					dto.Conf.MySQL.Local.Port+")/"+
					dto.Conf.MySQL.Local.Database+"?charset=utf8&parseTime=True&loc=Local"),
			&gorm.Config{
				PrepareStmt:            true,
				SkipDefaultTransaction: true,
			},
		)
	} else {
		// 生产环境
		DB, err = gorm.Open(
			mysql.Open(
				dto.Conf.MySQL.Default.Username+":"+
					dto.Conf.MySQL.Default.Password+"@tcp("+
					dto.Conf.MySQL.Default.Host+":"+
					dto.Conf.MySQL.Default.Port+")/"+
					dto.Conf.MySQL.Default.Database+"?charset=utf8&parseTime=True&loc=Local"),
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
