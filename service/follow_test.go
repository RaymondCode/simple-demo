package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
)

func TestFollow(t *testing.T) {
	db.Init()
	//followModel := &model.User{
	//	ID:       "223",
	//	Name:     "nyf123456",
	//	PassWord: "12345678",
	//}
	//model.CreateUser(context.Background(), followModel)
	//res, total, _ := model.QueryFollow(context.Background(), "223", 1, 1, 10)
	//fmt.Println(res, total)
	//re, _ := model.QueryUserById(context.Background(), "08dc2b99ef974d47a2554ed3dea73ea0")
	//for index, value := range re {
	//	fmt.Println("index=", index, "value=", value)
	//}
	//fmt.Println(*re[0])

	followModel := &model.Follow{
		UserId:       1,
		FollowedUser: 7,
		Status:       1,
	}
	model.CreateFollow(context.Background(), followModel)
	res, total, _ := model.QueryFollow(context.Background(), 1, 1, 10, 0)
	fmt.Println(len(res), total)
}
