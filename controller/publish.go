package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	//token := c.PostForm("token")
	//验证token
	title := c.PostForm("title")
	const DatabaseAddress string = "root:root@tcp(localhost:3306)/momotok"
	db, err := sql.Open("mysql", DatabaseAddress)
	if err != nil {
		fmt.Println("Database connected failed: ", err)
	}
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.SaveUploadedFile(file, "./data/"+file.Filename)
	_, err = db.Exec("INSERT INTO video (author_id,title) VALUES (?,?)", 1, title) //从token里面获取之后更新
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "success",
	})
	return
}

// PublishList shows user's published videos
func PublishList(c *gin.Context) {
	if !checkToken(c.Query("token")) {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Unauthorized request"})
		return
	}
	db, err := sql.Open("mysql", DatabaseAddress)
	userId := c.Query("user_id")
	parseIntID, _ := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("Database connected failed: ", err)
	}
	var user = User{Id: parseIntID}
	rows, _ := db.Query("select video.id, play_url, cover_url, favourite_count, comment_count, title, publish_time, user.username FROM video JOIN user ON video.author_id = user.id where author_id = %d ", parseIntID)
	videoList := make([]Video, 0)
	for rows.Next() {
		var videoId int64
		var playUrl string
		var coverUrl string
		var favoriteCount int64
		var commentCount int64
		var title string
		var publishTime int64
		err := rows.Scan(&videoId, &playUrl, &coverUrl, &favoriteCount, &commentCount, &title, &publishTime, &user.Name)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			continue
		} //find published videos
		var likedID int
		isFavourite := false
		db.QueryRow("select id FROM likes where user_id = %d AND video_id = %d", userId, videoId).Scan(&likedID)
		if likedID != 0 {
			isFavourite = true
		}
		video := Video{ //载入视频结构
			Id:            videoId,
			Author:        user,
			PlayUrl:       playUrl,
			CoverUrl:      coverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    isFavourite,
		}
		videoList = append(videoList, video) //视频切片加入视频列表
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
