package service

import (
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"testing"
)

func TestVideo(t *testing.T) {
	db.Init()
	//_, res := model.QueryVideoList(context.Background())
	res := service.QueryFeedResponse(7)
	//constitute FeedResponse struct ([]dto.video)
	//for index, value := range res {
	//	fmt.Println("index", index, "value=", value)
	//}
	fmt.Println(res)
}
