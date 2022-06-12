package controller

import (
	"fmt"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"

	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []dto.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	//需要指定最大上传尺寸，此处无法指定
	fmt.Println("进入publish")
	file, err := c.FormFile("data")
	if err != nil {
		return
	}
	token := c.PostForm("token")
	title := c.PostForm("title")
	//上传视频,并添加一个video到数据库
	userIdFromC, _ := c.Get("user_id")
	userId, _ := userIdFromC.(int64)
	service.UploadVideoAliyun(file, token, title, userId)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "test" + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
