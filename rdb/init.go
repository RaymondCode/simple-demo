package rdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/warthecatalyst/douyin/common"
	"github.com/warthecatalyst/douyin/config"
	"github.com/warthecatalyst/douyin/util"
)

var rdb *redis.Client

func InitRdb() {
	logrus.Infof("start init redis...")
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", config.DbHost),
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Panicf("[InitRdb] connect redis error, err=%+v", err)
	}

	setSalts()
	return
}

func setSalts() {
	salts := rdb.SMembers(common.KeySalt).Val()
	if len(salts) != 0 {
		logrus.Infof("[setSalts] salts = %v", salts)
		return
	}
	err := rdb.SAdd(common.KeySalt, util.CreateRandomString(10)).Err()
	if err != nil {
		logrus.Panicf("[setSalts] set salts error, err=%+v", err)
	}
	return
}

func GetRdb() *redis.Client {
	return rdb
}
