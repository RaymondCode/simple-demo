package controller

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/global"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	api.Response
	VideoList []api.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.Query("token")
	userId, err := strconv.ParseInt(token, 10, 64) //得到用户ID
	if err != nil {
		global.DyLogger.Print("Can't get userId from token")
		c.JSON(http.StatusOK, api.Response{StatusCode: api.TokenInvalidErr, StatusMsg: api.ErrorCodeToMsg[api.TokenInvalidErr]})
		return
	}
	global.DyLogger.Print(userId)
	data, err := c.FormFile("data")
	if err != nil {
		global.DyLogger.Print("Can't form file")
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.InnerErr,
			StatusMsg:  api.ErrorCodeToMsg[api.InnerErr],
		})
		return
	}
	filename := filepath.Base(data.Filename)
	//title := c.Query("title")
	finalName := strconv.FormatInt(userId, 10) + "_" + filename
	//将文件先保存到本地
	savePath := filepath.Join("./public/" + finalName)
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		global.DyLogger.Print("Error in SaveUploadedFile")
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.UploadFailErr,
			StatusMsg:  api.ErrorCodeToMsg[api.UploadFailErr],
		})
		return
	}
	//将文件上传到阿里云上
	err = uploadtoServer(data, userId)
	if err != nil {
		global.DyLogger.Print("Error in upload to Aliyun: " + err.Error())
		c.JSON(http.StatusOK, api.Response{
			StatusCode: api.UploadFailErr,
			StatusMsg:  api.ErrorCodeToMsg[api.UploadFailErr],
		})
		return
	}
	c.JSON(http.StatusOK, api.Response{
		StatusCode: 0,
		StatusMsg:  "Upload Success : " + finalName,
	})

}

func uploadtoServer(file *multipart.FileHeader, userId int64) error {
	Endpoint := global.UploadServerURL
	AccessKeyID := global.UploadAccessKeyID         // AccessKeyId
	AccessKeySecret := global.UploadAccessKeySecret // AccessKeySecret

	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		global.DyLogger.Print("Error In upload1:", err)
		return err
	}

	// 指定bucket
	bucket, err := client.Bucket(global.UploadServerBucket) // 根据自己的填写
	if err != nil {
		global.DyLogger.Print("Error In upload2:", err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		global.DyLogger.Print("Error in upload3:", err)
		return err
	}
	defer src.Close()

	// 先将文件流上传至douyin_test目录下
	path := "douyin_test/" + strconv.FormatInt(userId, 10) + "/" + file.Filename
	err = bucket.PutObject(path, src)
	if err != nil {
		global.DyLogger.Print("Error in upload4:", err)
		os.Exit(-1)
	}

	global.DyLogger.Print("file upload success")
	return nil
}

// PublishList 返回用户发布的所有视频列表
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
