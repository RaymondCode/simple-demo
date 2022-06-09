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
	followModel := &model.Follow{
		UserId:       1,
		FollowedUser: 7,
		Status:       1,
	}
	model.CreateFollow(context.Background(), followModel)
	res, total, _ := model.QueryFollow(context.Background(), 1, 1, 10, 0)
	fmt.Println(len(res), total)
}
