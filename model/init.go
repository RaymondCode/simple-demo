package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-demo/define"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := define.MysqlDNS + "/simple_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	return db
}