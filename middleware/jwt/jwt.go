package jwt

import (
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		// 注册时不需要token
		if i := strings.Index(uri, "register"); i != -1 {
			c.Next()
			return
		}
		// 登录时不需要token
		if i := strings.Index(uri, "login"); i != -1 {
			c.Next()
			return
		}
		token := ""
		// 投稿发布时, token用PostForm取, 其他的用Query取
		if i := strings.Index(uri, "publish/action"); i != -1 {
			token = c.PostForm("token")
		} else {
			token = c.Query("token")
		}
		if token == "" {
			c.JSON(http.StatusForbidden, "empty token")
			c.Abort()
			return
		}
		// 解析token
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusForbidden, "Invalid token")
			c.Abort()
			return
		}
		c.Set("userId", claims.UserId)
		c.Next()
	}
}
