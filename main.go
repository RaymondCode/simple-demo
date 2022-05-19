package main

import (
	"github.com/RaymondCode/simple-demo/common/settings"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//配置文件初始化
	if err := settings.Init(); err != nil {
		panic("config init failed...")
	}
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
