package main

import (
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initialize.Config()
	initialize.Mysql()
	initialize.Redis()
	initialize.Routers(r)

	r.Run(":7878") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
