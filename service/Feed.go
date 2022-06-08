package service

import (
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/erromsg"
	"time"
)

func Feed(latestTime string) (api.Feed, error) {
	handleErr := func(errType *erromsg.Eros) api.Feed {
		return api.Feed{
			Response: api.Response{
				StatusCode:    errType.Code,
				StatusMsg: errType.Message,
			},
		}
	}

	videos, err := dao.NewVideoDao().GetLatest(latestTime)
	nextTime, err := time.Parse(time.RFC3339, videos[len(videos)-1].CreateAt)
	if err != nil {
		return handleErr(erromsg.ErrQueryVideosFail), err
	}
	v := newVideoList(videos)
	for i := 0; i < len(v); i++ {
		//查询视频作者信息
		resp, err := UserInfo(videos[i].UserID)
		if err != nil {
			return handleErr(erromsg.ErrQueryUserInfoFail), err
		}
		v[i].Author = resp.User //作者信息
	}
	return api.Feed{VideoLists: v, Response: api.OK, NextTime: nextTime.Unix()}, nil
}
