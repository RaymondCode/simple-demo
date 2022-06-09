package model

import (
	"github.com/BaiZe1998/douyin-simple-demo/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true, // executes the given query in cached statement
			SkipDefaultTransaction: true, // disable default transaction
		},
	)
	if err != nil {
		panic(err)
	}
}
