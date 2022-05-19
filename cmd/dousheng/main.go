package main

import (
	"github.com/fitenne/youthcampus-dousheng/internal/common/jwt"
	"github.com/fitenne/youthcampus-dousheng/internal/common/settings"
	"github.com/fitenne/youthcampus-dousheng/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func init() {
	//配置文件初始化
	if err := settings.Init(); err != nil {
		panic("config init failed...")
	}

	jwt.Init([]byte(viper.GetString("app.secret")))

	err := repository.Init(repository.DBConfig{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.pass"),
		DBname:   viper.GetString("db.database"),
	})
	if err != nil {
		panic(err.Error())
	}
}
