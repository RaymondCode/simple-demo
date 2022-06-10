package rdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/warthecatalyst/douyin/common"
	"github.com/warthecatalyst/douyin/config"
	"github.com/warthecatalyst/douyin/logx"
	"github.com/warthecatalyst/douyin/util"
)

var rdb *redis.Client

func InitRdb() {
	logx.DyLogger.Infof("start init redis...")
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:3306", config.DbHost),
		Password: "zxy19991031",
		DB:       0,
		PoolSize: 100,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		logx.DyLogger.Panicf("[InitRdb] connect redis error, err=%+v", err)
	}

	setSalts()
	return
}

func setSalts() {
	salts := rdb.SMembers(common.KeySalt).Val()
	if len(salts) != 0 {
		logx.DyLogger.Infof("[setSalts] salts = %v", salts)
		return
	}
	err := rdb.SAdd(common.KeySalt, util.CreateRandomString(10)).Err()
	if err != nil {
		logx.DyLogger.Panicf("[setSalts] set salts error, err=%+v", err)
	}
	return
}

func GetRdb() *redis.Client {
	return rdb
}
