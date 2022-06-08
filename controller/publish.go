package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/service"
	"log"
	"net/http"
)


func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "用户不存在"})
		return
	}
	header,err:= c.FormFile("data")
	user := usersLoginInfo[token]
	userid := user.Id
	if err != nil {
		log.Print(err)
	}
	responses,err := service.PublishVideo(header,userid,c)
	if err != nil {
		fmt.Print("上传失败")
	}
	c.JSON(http.StatusOK,responses)

}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.PostForm("token")
	user := usersLoginInfo[token]
	userid := user.Id
	videoList ,err := service.PublishList(userid)
	videoLists := videoList.VideoLists
	if err != nil {

	}
	c.JSON(http.StatusOK, api.PublishList{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoLists: videoLists,
	})
}
