package main

import (
	"github.com/gin-gonic/gin"
	"github.com/life-studied/douyin-simple/initialize"
	"github.com/life-studied/douyin-simple/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()
	initialize.Config()
	initialize.Mysql()
	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
