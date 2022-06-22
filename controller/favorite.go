package controller

import (
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	UserIdFromC, _ := c.Get("user_id")
	UserID, _ := UserIdFromC.(int64)
	VideoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err := service.FavoriteAction(c, UserID, VideoID, int(actionType)); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "操作成功",
		},
	})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	UserIdFromC, _ := c.Get("user_id")
	UserID, _ := UserIdFromC.(int64)
	videoList, _ := service.GetFavoriteList(c, UserID)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
