package main

import (
	"fmt"
	"github.com/NoCLin/douyin-backend-go/config"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	initRouter(r)

	addr := "127.0.0.1:8080"
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err = r.Run(addr)
	if err != nil {
		fmt.Println("启动服务器失败", err)
		panic(err)
	}
}
