// Package service -----------------------------
// @file      : feed.go
// @author    : Yunyin
// @contact   : yunyin_jayyi@qq.com
// @time      : 2023/8/22 23:39
// -------------------------------------------
package service

import (
	"github.com/life-studied/douyin-simple/dao"
	"strconv"
	"time"
)

func GetNextTime(latest_time string) (int64, int64, error) {
	i64LatestTime, err := strconv.ParseInt(latest_time, 10, 64)
	i64LatestTime /= 1000
	if err != nil {
		return 0, 0, err
	}
	tmLatestTime := time.Unix(i64LatestTime, 0)
	nextTime, startId, err := dao.QueryNextTimeByLatestTime(tmLatestTime)
	if err != nil {
		return 0, 0, err
	}
	return nextTime, startId, nil
}
