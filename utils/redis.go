package utils

import (
	"context"
	"github.com/RaymondCode/simple-demo/global"
	"strconv"
)

const (
	// FavoriteVideoKey 以videoId为key, value为userId的set集合, 即存放点赞该视频的用户id集合
	FavoriteVideoKey = "favorite:video"

	// FavoriteUserKey  以userId为key, value为videoId的set集合, 即存放该用户点赞的视频id集合
	FavoriteUserKey = "favorite:user"
)

func GetFavoriteVideoKey(videoId int64) string {
	return FavoriteVideoKey + ":" + strconv.FormatInt(videoId, 10)
}

func GetFavoriteUserKey(userId int64) string {
	return FavoriteUserKey + ":" + strconv.FormatInt(userId, 10)
}

// IsFavorite 判断video是否被用户点赞
func IsFavorite(userId, videoId int64) bool {
	flag, err := global.RD.SIsMember(context.Background(), GetFavoriteUserKey(userId), videoId).Result()
	if err != nil {
		return false
	}
	return flag
}
