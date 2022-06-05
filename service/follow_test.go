package service

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/db"
	"github.com/RaymondCode/simple-demo/db/model"
	"testing"
)

func TestFollow(t *testing.T) {
	db.Init()
	followModel := &model.Follow{
		UserId:       1,
		FollowedUser: 2,
		Status:       1,
	}
	model.CreateFollow(context.Background(), followModel)
	res, total, _ := model.QueryFollow(context.Background(), 1, 1, 1, 10)
	fmt.Println(res, total)
}
