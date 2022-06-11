package service

import (
	"context"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"testing"
)

func TestFeed(t *testing.T) {
	db.Init()
	_, re := model.QueryVideoList(context.Background())

	fmt.Println(re[0], re[1])
}
