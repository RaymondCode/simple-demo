// Package service -----------------------------
// @file      : publish.go
// @author    : Yunyin
// @contact   : yunyin_jayyi@qq.com
// @time      : 2023/8/24 16:56
// -------------------------------------------
package service

import (
	"github.com/life-studied/douyin-simple/dao"
	"time"
)

type Video struct {
	AuthorID int64  `gorm:"column:author_id;not null;comment:'作者ID'" json:"author_id"`
	PlayURL  string `gorm:"column:play_url;not null;comment:'播放链接'" json:"play_url"`
	CoverURL string `gorm:"column:cover_url;not null;comment:'封面链接'" json:"cover_url"`
	Title    string `gorm:"column:title;not null;comment:'标题'" json:"title"`
}

func SaveVideo(video Video) error {
	err := dao.SaveVideoToMysql(dao.Video{
		AuthorID:    video.AuthorID,
		PlayURL:     video.PlayURL,
		CoverURL:    video.CoverURL,
		Title:       video.Title,
		PublishTime: time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}
