package service

import (
	"context"
	"fmt"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
)

func AddComment(text string, users dto.User, videoId int64) dto.ResponseComment {
	newComment := &model.Comment{
		VideoId: videoId,
		UserId:  users.Id,
		Content: text,
		Status:  1,
	}
	//comment commit
	model.CreateComment(context.Background(), videoId, newComment)
	responseComment := dto.ResponseComment{
		ID:        newComment.ID,
		User:      users,
		Content:   text,
		CreatedAt: newComment.CreatedAt.Format("2006-01-02"),
	}

	return responseComment
}

func UpdatCacheCommentList(ctx context.Context, videoId int64, limit, offset int) {
	res, total, _ := GetCommentList(ctx, videoId, limit, offset)
	fmt.Printf("更新数据", total)
	videoList_id := string(videoId) + "_comment_list"
	db.CacheSetList(context.Background(), "default", videoList_id, res, time.Hour)
}

func GetCommentList(ctx context.Context, videoId int64, limit, offset int) ([]dto.ResponseComment, int64, error) {
	var responseComment []dto.ResponseComment
	videoList_id := string(videoId) + "_comment_list"
	if value, err := db.CacheGetList(context.Background(), "default", videoList_id, []dto.ResponseComment{}); err != nil {
		fmt.Printf("获取缓存数据", value)
		return value, int64(len(value)), err
	}
	res, total, _ := model.QueryComment(context.Background(), videoId, 10, 0)
	responseComment = make([]dto.ResponseComment, len(res))
	fmt.Printf("评论总数", total)
	for index, v := range res {
		userInfo, _ := model.QueryUserById(context.Background(), v.UserId)
		fmt.Println("user:", userInfo.Name)
		users := dto.User{
			Id:            userInfo.ID,
			Name:          userInfo.Name,
			FollowCount:   userInfo.FollowCount,
			FollowerCount: userInfo.FollowerCount,
			IsFollow:      false,
		}
		responseComment[index] = dto.ResponseComment{
			ID:        v.ID,
			User:      users,
			Content:   v.Content,
			CreatedAt: v.CreatedAt.Format("2006-01-02"),
		}
	}
	err := db.CacheSetList(context.Background(), "default", videoList_id, responseComment, time.Hour)
	return responseComment, total, err
}
