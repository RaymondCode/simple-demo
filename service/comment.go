package service

import (
	"context"
	"fmt"
	"strconv"
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
	res, total := GetCommentList(ctx, videoId, limit, offset)
	fmt.Println("更新数据", total)
	videoList_id := strconv.FormatInt(videoId, 10) + "_comment_list"
	db.CacheSetList(context.Background(), "default", videoList_id, res, time.Minute)
}

func GetCacheCommentList(ctx context.Context, videoId int64, limit, offset int) ([]dto.ResponseComment, int64, error) {

	videoList_id := strconv.FormatInt(videoId, 10) + "_comment_list"

	if value, err := db.CacheGetList(context.Background(), "default", videoList_id, []dto.ResponseComment{}); len(value) > 0 {
		fmt.Println("获取缓存数据", value)
		return value, int64(len(value)), err
	} else {
		fmt.Println(videoList_id)
		fmt.Println(len(value))
	}
	responseComment, total := GetCommentList(ctx, videoId, 10, 0)
	err := db.CacheSetList(context.Background(), "default", videoList_id, responseComment, time.Minute)
	return responseComment, total, err
}

func GetCommentList(ctx context.Context, videoId int64, limit, offset int) ([]dto.ResponseComment, int64) {
	var responseComment []dto.ResponseComment
	res, total, _ := model.QueryComment(context.Background(), videoId, limit, offset)
	responseComment = make([]dto.ResponseComment, len(res))
	fmt.Println("评论总数", total)
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
	return responseComment, total
}
