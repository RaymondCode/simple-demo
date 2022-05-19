package main

import (
	"github.com/fitenne/youthcampus-dousheng/internal/common/settings"
	"github.com/fitenne/youthcampus-dousheng/internal/repository"
	"github.com/gin-gonic/gin"
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

	c := repository.DBConfig{
		Driver:   "mysql",
		Host:     "example.com",
		Port:     "3306",
		User:     "demo",
		Password: "secret",
		DBname:   "dousheng",
	}

	err := repository.Init(c)
	if err != nil {
		panic(err.Error())
	}
}
