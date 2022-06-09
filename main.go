package main

import (
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/gin-gonic/gin"
)

func Init() {
	db.Init()
}

func main() {
	Init()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
