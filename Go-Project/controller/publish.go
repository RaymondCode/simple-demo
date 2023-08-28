package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/service"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory and mysql
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%d_%s", user.Id, time.Now().Unix(), filename)
	saveFilePath := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFilePath); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	newId := len(DemoVideos) + 1
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	playUrl := "http://" + string(ip) + ":8080" + "/public" + finalName
	DemoVideos = append(DemoVideos, Video{
		Id:            int64(newId),
		Author:        user,
		PlayUrl:       playUrl,
		CoverUrl:      "",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	})
	title := c.PostForm("title")
	err = service.SaveVideo(service.Video{
		ID:       int64(newId),
		AuthorID: user.Id,
		PlayURL:  playUrl,
		CoverURL: "",
		Title:    title,
	})
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideoListResponse{
			Response:  Response{1, err.Error()},
			VideoList: nil,
		})
		return
	}
	var userPublishVideos []Video
	for _, video := range DemoVideos {
		if video.Author.Id == userId {
			userPublishVideos = append(userPublishVideos, video)
		}
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:  Response{0, "response successfully"},
		VideoList: userPublishVideos,
	})
}
