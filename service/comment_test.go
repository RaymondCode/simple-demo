package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
)

func TestComment(t *testing.T) {
	db.Init()
	commentModel := &model.Comment{
		UserId:  1,
		VideoId: 1,
		Content: "nihao",
	}
	model.CreateComment(context.Background(), commentModel)
	res, total, _ := model.QueryComment(context.Background(), 1, 10, 0)
	fmt.Println(res[0].Content, total)
}
