package service

import (
	"fmt"
	"simple-demo/model"
	"testing"
)

func TestCreate(t *testing.T) {
	model.Init()

	userInfo := &model.User{UserName: "user_text", Password: "user_text"}
	err := model.DB.Create(userInfo).Error

	temp := make([]*model.User, 0)
	err = model.DB.Find(&temp).Error
	if err != nil {
		t.Fatal(err)
	}
	//遍历取出每一行数据

	for _, v := range temp {
		fmt.Printf("Found ==> %v \n", v)
	}
}
