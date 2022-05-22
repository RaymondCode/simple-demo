package model_test

import (
	errorcode "github.com/fitenne/youthcampus-dousheng/internal/common/error"
	"github.com/fitenne/youthcampus-dousheng/internal/common/settings"
	"github.com/fitenne/youthcampus-dousheng/internal/repository"
	"github.com/fitenne/youthcampus-dousheng/pkg/model"
	"github.com/spf13/viper"
	"testing"
)

func init() {
	if err := settings.Init("../../config.yaml"); err != nil {
		panic(err.Error())
	}

	repository.Init(repository.DBConfig{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.pass"),
		DBname:   viper.GetString("db.database"),
	})
}

//创建视频测试一：正常创建视频
func TestVideoCreate1(t *testing.T) {
	//创建视频
	video := &model.Video{
		AuthorID: 1,
		PlayUrl:  "http://www.baidu.com1",
		CoverUrl: "http://www.baidu.com1",
		Author: &model.User{
			ID:            1,
			FollowCount:   0,
			FollowerCount: 0,
		},
	}
	id, err := repository.GetVideoCtl().Create(video)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

//创建视频测试二：创建视频时，authorID、author.ID不统一
func TestVideoCreate2(t *testing.T) {
	//创建视频
	video := &model.Video{
		AuthorID: 1,
		PlayUrl:  "http://www.baidu.com2",
		CoverUrl: "http://www.baidu.com2",
		Author: &model.User{
			ID:            -999,
			FollowCount:   0,
			FollowerCount: 0,
		},
	}
	_, err := repository.GetVideoCtl().Create(video)
	if err == nil {
		t.Error("'创建视频时，authorID、author.ID不统一'测试用例失败")
	} else {
		if err.Error() == errorcode.VideoCreateForeignKeyNotUnified.Message() {
			t.Log("'创建视频时，authorID、author.ID不统一'测试用例通过")
		} else {
			t.Error(err)
		}
	}
}

//创建视频测试三：创建视频时，authorID不存在
func TestVideoCreate3(t *testing.T) {
	//创建视频
	video := &model.Video{
		AuthorID: -999,
		PlayUrl:  "http://www.baidu.com3",
		CoverUrl: "http://www.baidu.com3",
		Author: &model.User{
			ID:            -999,
			FollowCount:   0,
			FollowerCount: 0,
		},
	}
	_, err := repository.GetVideoCtl().Create(video)
	if err == nil {
		t.Error("'创建视频时，authorID不存在'测试用例失败")
	} else {
		if err.Error() == errorcode.VideoCreateForeignKeyNotExist.Message() {
			t.Log("'创建视频时，authorID不存在'测试用例通过")
		} else {
			t.Error(err)
		}
	}
}

//删除视频测试一：正常删除视频
func TestVideoDelete1(t *testing.T) {
	video := &model.Video{
		ID:       1,
		AuthorID: 1,
		PlayUrl:  "http://www.baidu.com3",
		CoverUrl: "http://www.baidu.com3",
		Author: &model.User{
			ID: 1,
		},
	}
	err := repository.GetVideoCtl().Delete(video)
	if err != nil {
		t.Error(err)
	}
}

//查询视频测试一：根据id查询
func TestVideoGet1(t *testing.T) {
	video, err := repository.GetVideoCtl().GetVideoById(3)
	if err != nil {
		t.Error(err)
	}
	t.Log(video)
}
