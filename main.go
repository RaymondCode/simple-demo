package main

import (
	"github.com/BaiZe1998/douyin-simple-demo/db"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/gin-gonic/gin"
)

func Init() {
	dto.InitConfig()
	db.Init()
}

func main() {
	Init()

	cfg := dto.GetConfig()

	r := gin.Default()

	initRouter(r)

	r.Run(":" + cfg.Server.Port)
}
