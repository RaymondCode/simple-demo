package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DRIVER = "mysql"

var SqlSession *gorm.DB

func InitMySql(conf MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DBName,
	)
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	SqlSession.AutoMigrate(&User{}, &Video{})
	return SqlSession.DB().Ping()
}

func CloseMySQL() {
	err := SqlSession.Close()
	if err != nil {
		panic(err)
	}
}
