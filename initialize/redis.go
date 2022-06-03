package initialize

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("redis connect ping reRDonse:", zap.String("pong", pong))
		global.RD = client
	}
}
