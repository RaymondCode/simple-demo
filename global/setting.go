package global

import "strconv"

//for upload settings and get URLs
const (
	UploadServerURL       = "https://oss-cn-hangzhou.aliyuncs.com"
	UploadAccessKeyID     = "LTAI5tQ8pkD4D4CQmznGxU1A"
	UploadAccessKeySecret = "oNFSW47dgshTlt1y9TMqPx5SRFBeN0"
	UploadServerBucket    = "tdouyinbuc"
	UploadBucketDirectory = "douyin_test"
)

// GetUserURL 得到对应的云端存储路径
func GetUserURL(userId int64) string {
	return "https://" + UploadServerBucket + ".oss-cn-hangzhou.aliyuncs.com/" + UploadBucketDirectory + "/" + strconv.FormatInt(userId, 10) + "/"
}
