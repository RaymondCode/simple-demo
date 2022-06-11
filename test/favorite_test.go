package service

import (
	"context"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"testing"
)

func TestFavorite(t *testing.T) {
	dto.InitConfigForTest()
	db.Init()
	service.GetFavoriteList(context.Background(), 5)
}
