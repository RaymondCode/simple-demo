package main

import (
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()

	// 初始化配置，如MySQL等
	err := models.InitProject()
	if err != nil {
		panic(err)
	}
	defer models.Close()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
