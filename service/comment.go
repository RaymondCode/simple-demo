package service

import (
	"errors"

	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
)

// CommentActionInfo service层添加或者删除一条评论记录
func CommentActionInfo(commentID, userID, videoID int64, actionType int32, content string) error {
	return newCommentActionInfoFlow(commentID, userID, videoID, actionType, content).Do()
}

func newCommentActionInfoFlow(commentID, userID, videoID int64, actionType int32, content string) *CommentActionInfoFlow {
	return &CommentActionInfoFlow{
		commentId:  commentID,
		userId:     userID,
		videoId:    videoID,
		actionType: actionType,
		content:    content,
	}
}

type CommentActionInfoFlow struct {
	commentId  int64
	userId     int64
	videoId    int64
	actionType int32
	content    string
}

func (c *CommentActionInfoFlow) Do() error {
	if c.actionType == api.CommentAction {
		if err := c.AddRecord(); err != nil {
			return err
		}
	} else if c.actionType == api.UnCommentAction {
		if err := c.checkRecord(); err != nil {
			return err
		}
		if err := c.DelRecord(); err != nil {
			return err
		}
	} else {
		return errors.New("actionType must be 1 or 2")
	}
	return nil
}

func (c *CommentActionInfoFlow) checkRecord() error {
	if flag := dao.NewCommentDaoInstance().IsComment(c.commentId); !flag {
		return errors.New("there's no such record")
	}
	return nil
}

func (c *CommentActionInfoFlow) AddRecord() error {
	if err := dao.NewCommentDaoInstance().Add(c.userId, c.videoId, c.content); err != nil {
		return err
	}
	return nil
}

func (c *CommentActionInfoFlow) DelRecord() error {
	if err := dao.NewCommentDaoInstance().Del(c.commentId, c.videoId); err != nil {
		return err
	}
	return nil
}

type CommentList []api.Comment

// CommentListInfo 获取视频评论列表
func CommentListInfo(videoId int64) (*CommentList, error) {
	return newCommentListInfoFlow(videoId).Do()
}

func newCommentListInfoFlow(videoId int64) *CommentListInfoFlow {
	return &CommentListInfoFlow{
		videoId: videoId,
	}
}

type CommentListInfoFlow struct {
	videoId int64
}

func (c *CommentListInfoFlow) Do() (*CommentList, error) {
	return c.getCommentList()
}

func getUserFromCommentId(commentId int64) (*api.User, error) {
	userId, err := dao.NewCommentDaoInstance().GetUserFromCommentId(commentId)
	if err != nil {
		return nil, err
	}
	userModel, err := dao.NewUserDaoInstance().GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return &api.User{
		Id:            userId,
		Name:          userModel.UserName,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,
		IsFollow:      false,
	}, nil
}

func (c *CommentListInfoFlow) getCommentList() (*CommentList, error) {
	commentIds, err := dao.NewCommentDaoInstance().CommentListByVideoID(c.videoId)
	if err != nil {
		return nil, err
	}

	var commentlist CommentList
	for _, commentId := range commentIds {
		user, err := getUserFromCommentId(commentId)
		if err != nil {
			return nil, err
		}
		comment, err := dao.NewCommentDaoInstance().GetCommentFromId(commentId)

		commentlist = append(commentlist, api.Comment{
			Id:         int64(comment.ID),
			User:       *user,
			Content:    comment.Content,
			CreateDate: comment.CreateAt.Format("2006-01-02 15:04:05"),
		})

	}
	return &commentlist, nil
}
