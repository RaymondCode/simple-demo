package util

import (
	"github.com/gin-gonic/gin"
)

//自定义中间件，jwt验证
func VerifyJwt() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Query("token")
		if token == "" {
			panic("token not exist !")
		}
		context.Set("token", ParseToken(token))
		context.Next()
	}
}
