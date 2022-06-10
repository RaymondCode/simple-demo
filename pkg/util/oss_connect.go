package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func InitVodClient(accessKeyId string, accessKeySecret string) (client *vod.Client, err error) {
	// 点播服务接入区域
	regionId := "cn-shanghai"
	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		accessKeyId,
		accessKeySecret,
	}
	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
	// 创建vodClient实例
	return vod.NewClientWithOptions(regionId, config, credential)
}

func MyCreateUploadVideo(client *vod.Client) (response *vod.CreateUploadVideoResponse, err error) {
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = "Sample Video Title"
	request.Description = "Sample Description"
	request.FileName = "/opt/video/sample/video_file.mp4"
	//request.CateId = "-1"
	//Cover URL示例：http://example.alicdn.com/tps/TB1qnJ1PVXXXXXCXXXXXXXXXXXX-700-****.png
	request.CoverURL = "<your CoverURL>"
	request.Tags = "tag1,tag2"
	request.AcceptFormat = "JSON"
	return client.CreateUploadVideo(request)
}

func InitOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}

func UploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, localFile string) {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 上传本地文件。
	err = bucket.PutObjectFromFile(uploadAddressDTO.FileName, localFile)
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

func main() {
	var accessKeyId string = "<Your AccessKeyId>"             // 您的AccessKeyId
	var accessKeySecret string = "<Your AccessKeySecret>"     // 您的AccessKeySecret
	var localFile string = "/opt/video/sample/video_file.mp4" // 需要上传到VOD的本地视频文件的完整路径
	// 初始化VOD客户端并获取上传地址和凭证
	var vodClient, initVodClientErr = InitVodClient(accessKeyId, accessKeySecret)
	if initVodClientErr != nil {
		fmt.Println("Error:", initVodClientErr)
		return
	}
	// 获取上传地址和凭证
	var response, createUploadVideoErr = MyCreateUploadVideo(vodClient)
	if createUploadVideoErr != nil {
		fmt.Println("Error:", createUploadVideoErr)
		return
	}
	// 执行成功会返回VideoId、UploadAddress和UploadAuth
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
	UploadLocalFile(ossClient, uploadAddressDTO, localFile)
	//MultipartUploadFile(ossClient, uploadAddressDTO, localFile)
	fmt.Println("Succeed, VideoId:", videoId)
}
