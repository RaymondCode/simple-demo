package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
	"time"
)

func UploadVideoAliyun(file *multipart.FileHeader, token string, title string, userid int64) {
	vodClient := util.OOSInit()
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = title
	request.FileName = file.Filename
	response, err := vodClient.CreateUploadVideo(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var videoId = response.VideoId
	var uploadAuthDTO UploadAuthDTO
	var uploadAddressDTO UploadAddressDTO
	var uploadAuthDecode, _ = base64.StdEncoding.DecodeString(response.UploadAuth)
	var uploadAddressDecode, _ = base64.StdEncoding.DecodeString(response.UploadAddress)
	json.Unmarshal(uploadAuthDecode, &uploadAuthDTO)
	json.Unmarshal(uploadAddressDecode, &uploadAddressDTO)
	// 使用UploadAuth和UploadAddress初始化OSS客户端
	var ossClient, _ = InitOssClient(uploadAuthDTO, uploadAddressDTO)
	// 上传文件，注意是同步上传会阻塞等待，耗时与文件大小和网络上行带宽有关
	open, err := file.Open()
	if err != nil {
		return
	}
	UploadLocalFile(ossClient, uploadAddressDTO, open)
	//MultipartUploadFile(ossClient, uploadAddressDTO, localFile)
	fmt.Println("Succeed, VideoId:", videoId)
	//https://video.liufei.fun/sv/2eaa675f-1815638c4ce/2eaa675f-1815638c4ce.mp4
	//https://video.liufei.fun/553d29da2b3844c3bc930e5888f4c5d4/snapshots/216c32eac6064241ada9704637c88101-00002.jpg
	var playurl = "https://video.liufei.fun" + "/" + uploadAddressDTO.FileName
	video := &model.Video{
		AuthorID:      userid,
		PlayUrl:       playurl,
		CoverUrl:      request.CoverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		CreatedAt:     time.Time{},
		Title:         title,
	}
	err = model.CreateVideo(context.Background(), video)
	if err != nil {
		return
	}
}

func InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}
func UploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, open multipart.File) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	err = bucket.PutObject(uploadAddressDTO.FileName, open)
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

type UploadAuthDTO struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}
type UploadAddressDTO struct {
	Endpoint string
	Bucket   string
	FileName string
}
