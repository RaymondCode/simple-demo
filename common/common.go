package common

import (
	"bytes"
	"os"
	"path"
	"strings"
)

const (
	UserId   = "user_id"
	Username = "user_name"
)

const (
	KeySalt = "salt"
)

//for upload settings
const (
	UploadServerURL       = "https://oss-cn-hangzhou.aliyuncs.com"
	UploadAccessKeyID     = "LTAI5tQ8pkD4D4CQmznGxU1A"
	UploadAccessKeySecret = "oNFSW47dgshTlt1y9TMqPx5SRFBeN0"
	UploadServerBucket    = "tdouyinbuc"
	UploadBucketDirectory = "douyin_test"
	UploadVideoMaxSize    = 1024 * 1024 * 1024 //设置最多传1GB文件
)

var (
	GlobalSavePath   = "../userdata/"
	AllowedVideoExts = []string{
		".mp4",
		".wmv",
		".avi",
	}
)

// PathExists 判断是否存在路径
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CheckFileExt 检查文件的拓展名
func CheckFileExt(fileName string) bool {
	//检查文件的扩展名
	ext := path.Ext(fileName)
	ext = string(bytes.ToLower([]byte(ext)))
	for _, legalExt := range AllowedVideoExts {
		if ext == legalExt {
			return true
		}
	}
	return false
}

// CheckFileMaxSize 检查文件的大小
func CheckFileMaxSize(videoSize int64) bool {
	return videoSize <= UploadVideoMaxSize
}

func GetFileNameWithOutExt(fileName string) string {
	ext := path.Ext(fileName)
	return strings.TrimSuffix(fileName, ext)
}
