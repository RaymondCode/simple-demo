package service

import (
	"github.com/life-studied/douyin-simple/dao"
	"github.com/life-studied/douyin-simple/model"
)

type CommentService struct{}

func (cs *CommentService) QueryComment(videoId int64, token string) ([]model.Comment, error) {
	//返回值定义
	var comments []model.Comment
	var err error
	//dao层操作
	c := &dao.Comments{}
	comments, err = c.QueryCommentsByVideoId(videoId)
	if err != nil {
		return nil, err
	}
	return comments, nil

}
