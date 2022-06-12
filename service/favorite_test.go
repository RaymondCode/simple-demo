package service

import (
	"context"
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"testing"
)

func TestFavorite(t *testing.T) {
	//dto.InitConfigForTest()
	db.Init()
	GetFavoriteList(context.Background(), 5)
}
