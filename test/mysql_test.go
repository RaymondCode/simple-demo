package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-demo/model"
	"testing"
)

func TestGormTest(t *testing.T) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/simple_demo?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:admin#317421@tcp(127.0.0.1:3306)/Tiktok_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*model.User, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("Found ==> %v \n", v)
	}
}
