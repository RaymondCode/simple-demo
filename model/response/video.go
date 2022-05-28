package response

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

func OkWithVideoList(videoList []model.Video, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: SUCCESS,
			StatusMsg:  msg,
		},
		VideoList: videoList,
	})
}
