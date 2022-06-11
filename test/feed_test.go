package service

import (
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"testing"
)

func TestFeed(t *testing.T) {
	db.Init()
	var re []dto.Video
	//_, re := model.QueryVideoList(context.Background())
	re = service.QueryFeedResponse(7)
	fmt.Println(re[0], re[1])
}
