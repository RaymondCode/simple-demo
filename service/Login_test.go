package service

import (
	"fmt"
	"gorm.io/gorm"
	"simple-demo/model"
	"testing"
)

func TestLogin(t *testing.T) {
	model.Init()
	userInfo := model.User{UserName: "text", Password: "text"}
	err := model.DB.Where("username = ? AND password = ? ", "Test", "Test").Find(&userInfo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("用户名或密码错误")
			return
		}
	} else {
		fmt.Println("用户名和密码正确")
	}
}
