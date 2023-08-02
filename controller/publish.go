package controller

import (
	"database/sql"
	"fmt"
	"path"
	"strconv"

	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	token := c.PostForm("token")
	//验证token

	ischeak := checkToken(token)
	uid, err := getUID(token)
	println(uid)
	if !ischeak {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "token is useless",
		})
		return
	}
	username, err := getUsername(token)

	id := 0
	title := c.PostForm("title")
	const DatabaseAddress string = "root:root@tcp(localhost:3306)/momotok"
	db, err := sql.Open("mysql", DatabaseAddress)
	if err != nil {
		fmt.Println("Database connected failed: ", err)
	}
	file, err := c.FormFile("data")
	c.SaveUploadedFile(file, "./public/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filepath := path.Join("https://" + "/" + username + "/" + strconv.Itoa(id))
	println(filepath)

	_, err = db.Exec("INSERT INTO video (author_id,title,favourite_count,comment_count,play_url) VALUES (?,?,?,?,?)", uid, title, 0, 0, filepath) //,favourite_count,comment_count,play_url
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
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

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
