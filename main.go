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

	if err := dto.InitLogger(); err != nil {
		return
	}

	r := gin.Default()
	r.Use(dto.GinLogger(), dto.GinRecovery(true))

	initRouter(r)

	r.Run(":" + cfg.Server.Port)
}
