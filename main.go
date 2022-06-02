package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/controller"
	"github.com/warthecatalyst/douyin/dao"
)

func init() {
	dao.InitDB()
}

func main() {

	r := gin.Default()

	controller.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
