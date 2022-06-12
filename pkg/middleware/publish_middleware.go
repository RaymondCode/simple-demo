package middleware

import (
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
)

func PbulishMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		log.Printf(token)
		// 验证token有效性
		userClaims, err := util.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(200, gin.H{
				"status_code": 40001,
				"status_msg":  "请登录",
			})
			return
		}
		c.Set("token", userClaims)
		c.Set("user_id", userClaims.ID)
		c.Next()
	}

}
