package util

import "github.com/aliyun/alibaba-cloud-sdk-go/services/vod"

func OOSInit() *vod.Client {
	var accessKeyId string = "LTAI5tGcbULwMvTKqaSzgLQY"
	var accessKeySecret string = "zuLS8V0o65bSZBWMuOyeZtnQs4OTIG"
	client, err := InitVodClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil
	}
	return client
}
