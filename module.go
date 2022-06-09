package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type ModeTime time.Time

type Model struct {
	UUID uint	`gorm:"primaryKey"`
	//Time time.Time	`gorm:"column:my_time"`
}
type User struct {
	ID        uint     `gorm:"primaryKey;column:id"`
	UserName  string   `gorm:"column:username;unique;not null"`
	Password  string   `gorm:"column:password;not null"`
	CreatedAt ModeTime `gorm:"column:created_at"`
	UpdatedAt ModeTime `gorm:"column:updated_at"`
}

func testUserCreate(){
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/simple_demo?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	_ = db.AutoMigrate(&User{})
}