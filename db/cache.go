package db

import (
	"context"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/go-redis/redis/v8"
)

var RedisCaches map[string]*redis.Client

func InitRedisPools() error {
	var RedisDSN string
	var RedisPWD string
	if dto.Conf.Env.IsDebug {
		// 开发环境
		RedisDSN = dto.Conf.Redis.Local.Host + ":" + dto.Conf.Redis.Local.Port
		RedisPWD = dto.Conf.Redis.Local.Password
	} else {
		// 生产环境
		RedisDSN = dto.Conf.Redis.Default.Host + ":" + dto.Conf.Redis.Default.Port
		RedisPWD = dto.Conf.Redis.Default.Password
	}

	RedisCaches := make(map[string]*redis.Client)
	for k, v := range dto.Conf.Redis.Databases {
		RedisCaches[k] = redis.NewClient(&redis.Options{
			Addr:     RedisDSN,
			Password: RedisPWD,
			DB:       v,
			PoolSize: 0,
		})
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := RedisCaches[k].Ping(ctx).Result()

		if err != nil {
			return err
		}
	}

	return nil
}

func CacheSet(ctx context.Context, dbName string, key string, value string, expire uint) error {
	err := RedisCaches[dbName].Set(ctx, key, value, time.Duration(expire)).Err()
	return err
}

func CacheGet(ctx context.Context, dbName string, key string) (string, error) {
	result, err := RedisCaches[dbName].Get(ctx, key).Result()
	return result, err
}
