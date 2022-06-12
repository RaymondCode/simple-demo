package service

import (
	"context"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"testing"
)

func TestFavorite(t *testing.T) {
	dto.InitConfig()
	db.Init()
	//userList := make([]dto.User, 0)
	//userList = append(userList, dto.User{Id: 1, Name: "xiaoming"})
	//userList = append(userList, dto.User{Id: 2, Name: "xiaohong"})
	//db.CacheSetList(context.Background(), "default", "user_list", userList, 0)
	//value, _ := db.CacheGetList(context.Background(), "default", "user_list", []dto.User{})
	//op, ok := value.([]dto.User)
	//fmt.Println("------------", op, ok)
	videoList := make([]dto.Video, 0)
	videoList = append(videoList, dto.Video{Id: 1})
	videoList = append(videoList, dto.Video{Id: 2})
	db.CacheSetList(context.Background(), "default", "video_list", videoList, 0)
	value, _ := db.CacheGetList(context.Background(), "default", "video_list", []dto.Video{})
	fmt.Println("------------", value)
}
