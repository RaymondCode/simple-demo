package config

import (
	G "github.com/NoCLin/douyin-backend-go/config/global"
	"github.com/NoCLin/douyin-backend-go/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initGorm() error {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Video{},
		&model.Comment{},
		&model.Follow{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
	G.DB = db
	return nil

}
