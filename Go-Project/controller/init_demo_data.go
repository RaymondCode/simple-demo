// Package controller -----------------------------
// @file      : init_demo_data.go
// @author    : Yunyin
// @contact   : yunyin_jayyi@qq.com
// @time      : 2023/8/28 23:42
// -------------------------------------------
package controller

import (
	"fmt"
	"github.com/life-studied/douyin-simple/dao"
	"os"
)

func InitCacheFromMysql() {
	InitUserFromMysql()
	InitVideoDataFromMysql()
}

func InitUserFromMysql() {
	users, err := dao.InitUserFromMysql()
	if err != nil {
		fmt.Println("Init fail:InitUser function failed")
		os.Exit(1)
	}
	for i := 0; i < len(users); i++ {
		token := users[i].Name + users[i].Password
		usersLoginInfo[token] = User{
			Id:            users[i].ID,
			Name:          users[i].Name,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
	}
}

func InitVideoDataFromMysql() {
	videos, err := dao.InitVideoDataFromMysql()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for i := 0; i < len(videos); i++ {
		name, password, err := dao.InitUserByUserIdFromMysql(videos[i].AuthorID)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		token := name + password
		user, exists := usersLoginInfo[token]
		if !exists {
			fmt.Println("init fail:usersLoginInfo don't have user from mysql, please check your initUser function")
			os.Exit(1)
		}

		favoriteCount, err := dao.InitVideoFavoriteCountFromMysql(videos[i].ID)
		if err != nil {
			fmt.Println("Init fail: Video favorite count init failed")
			os.Exit(1)
		}

		commentCount, err := dao.InitVideoCommentCountFromMysql(videos[i].ID)
		if err != nil {
			fmt.Println("Init fail: Video favorite count init failed")
			os.Exit(1)
		}
		DemoVideos = append(DemoVideos, Video{
			Id:            videos[i].ID,
			Author:        user,
			PlayUrl:       videos[i].PlayURL,
			CoverUrl:      videos[i].CoverURL,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    false,
		})
	}
}
