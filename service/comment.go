package service

import (
	"context"

	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
)

func AddComment(text string, users dto.User, videoId int64) *dto.ResponeComment {

	newComment := &model.Comment{
		VideoId: videoId,
		UserId:  users.Id,
		Content: text,
		Status:  1,
	}
	//comment commit
	model.CreateComment(context.Background(), videoId, newComment)
	responseComment := &dto.ResponeComment{
		ID:        newComment.ID,
		User:      users,
		Content:   text,
		CreatedAt: newComment.CreatedAt,
	}
	return responseComment
}
