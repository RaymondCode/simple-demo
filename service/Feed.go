package service

import (
	"time"

	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
)

func Feed(latestTime string) (api.Feed, error) {
	videos, err := dao.NewVideoDaoInstance().GetLatest(latestTime)

	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	nextTime, err := time.ParseInLocation("2006-01-02 15:04:05", videos[len(videos)-1].CreateAt.Format("2006-01-02 15:04:05"), LOC)
	if err != nil {
		return api.Feed{
			Response: api.Response{
				StatusCode: api.InnerErr,
				StatusMsg:  api.ErrorCodeToMsg[api.InnerErr],
			},
		}, err
	}
	v := newVideoList(videos)
	for i := 0; i < len(v); i++ {
		//查询视频作者信息
		resp, err := UserInfo(videos[i].UserID)
		if err != nil {
			return api.Feed{
				Response: api.Response{
					StatusCode: api.InnerErr,
					StatusMsg:  api.ErrorCodeToMsg[api.InnerErr],
				},
			}, err
		}
		v[i].Author = resp.User //作者信息
	}
	return api.Feed{VideoLists: v, Response: api.OK, NextTime: nextTime.Unix()}, nil
}
