package db

import (
	"context"
	"time"

	"github.com/BaiZe1998/douyin-simple-demo/pkg/constants"
	"github.com/go-redis/redis/v8"
)

var RedisCaches map[string]*redis.Client

func InitRedisPools() error {
	RedisCaches := make(map[string]*redis.Client)
	for k, v := range constants.RedisDBList {
		RedisCaches[k] = redis.NewClient(&redis.Options{
			Addr:     constants.RedisDefaultDSN,
			Password: constants.RedisDefaultPWD,
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
