package service

import (
	"context"
	"testing"

	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
)

func TestFavorite(t *testing.T) {
	dto.InitConfig()
	db.Init()
	GetFavoriteList(context.Background(), 5)
}
