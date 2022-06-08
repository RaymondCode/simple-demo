package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/erromsg"
	"github.com/warthecatalyst/douyin/model"
	"github.com/yitter/idgenerator-go/idgen"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type PublishService struct {
}

var publishService = &PublishService{}

func (*PublishService) getUserFromVideoId(videoId int64) (*api.User, error) {
	userId, err := dao.NewVideoDaoInstance().GetUserIdFromVideoId(videoId)
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

// PublishVideo 上传数据到本地public文件夹,并在数据库里面更新伪地址数据
var helloURL = "https://cdn.cnbj1.fds.api.mi-img.com/product-images/redmik40sb4bd68/31.jpg"

func PublishVideo(header *multipart.FileHeader, userID int64, c *gin.Context) (api.Response, error) {
	//定义一个通用的Response用于返回
	handleErr := func(errorType *erromsg.Eros) api.Response {
		return api.Response{StatusCode: errorType.Code, StatusMsg: errorType.Message}
	}
	//token := c.PostForm("token")
	// 生成视频id
	id := idgen.NextId() // 暂时和用户ID使用同一个id生成器
	filename := filepath.Base(header.Filename)
	//user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", userID, filename)
	saveFile := filepath.Join("./public/", finalName)
	//保存视频到本地public文件夹下
	err := c.SaveUploadedFile(header, saveFile)
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return api.Response{}, err
	}

	c.JSON(http.StatusOK, api.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " 上传成功 ",
	})

	// 将文件保存的相对路径保存到数据库中
	//后期需要转换为OOS
	fakeUrl := fmt.Sprintf("%s_%s", "./public/", filename)
	err = dao.NewVideoDao().Create(map[string]interface{}{
		"author_id":  userID,
		"play_url":   fakeUrl,
		"video_name": filename,
		"cover_url":  helloURL,
		"video_id":   id})
	if err != nil {
		return handleErr(erromsg.ErrCreateVideoRecordFail), err
	}
	return api.OK, nil
}




// PublishList 返回用户发布的所有的视频，包括视频的点赞数和评论数等视频相关信息
func PublishList(userID int64) (api.PublishList, error) {
	//通用警告
	handleErr := func(errorType *erromsg.Eros) api.PublishList {
		return api.PublishList{Response: api.Response{StatusCode: errorType.Code, StatusMsg: errorType.Message}}
	}
	// 首先查询视频
	videos, err := dao.NewVideoDao().GetVideos(
		map[string]interface{}{"author_id": userID},
		"play_url", "cover_url", "favorite_count", "comment_count", "video_id")
	if err != nil {
		return handleErr(erromsg.ErrQueryVideosFail), err
	}

	v := newVideoList(videos) //构造数据

	for i, video := range v {
		// 作者自己是否点赞
		v[i].IsFavorite = dao.NewFavoriteDaoInstance().IsFavourite(userID,video.VideoID)
	}
	return api.PublishList{
		Response:   api.OK,
		VideoLists: v,
	}, err
}

// 构造 Video 切片
func newVideoList(videos []model.Video) []model.VideoQuery {
	var v []model.VideoQuery
	for _, i := range videos {
		v = append(v, model.VideoQuery{
			VideoID:       i.VideoID,
			PlayURL:       i.PlayURL,
			CoverURL:      i.CoverURL,
			CommentCount:  i.CommentCount,
			FavoriteCount: i.FavoriteCount,
		})
	}

	return v
}
