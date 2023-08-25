// Package dao -----------------------------
// @file      : feed.go
// @author    : Yunyin
// @contact   : yunyin_jayyi@qq.com
// @time      : 2023/8/22 23:47
// -------------------------------------------
package dao

import (
	"github.com/life-studied/douyin-simple/global"
	"time"
)

func QueryNextTimeByLatestTime(latestTime time.Time) (int64, int64, error) {
	var videos []Video
	err := global.DB.Where("publish_time <= ?", latestTime).Order("publish_time desc").Limit(30).Find(&videos).Error
	if err != nil {
		return 0, 0, err
	}
	var nextTime int64
	var ID int64
	if len(videos) < 30 {
		err = global.DB.Where("publish_time <= ?", latestTime).Order("publish_time asc").First(&videos[0]).Error
		if err != nil {
			return 0, 0, err
		}
		nextTime = videos[0].PublishTime.Unix()
		ID = videos[0].ID
	} else {
		nextTime = videos[len(videos)-1].PublishTime.Unix()
		ID = videos[len(videos)-1].ID
	}
	return nextTime, ID, nil
}
