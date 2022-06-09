package middleware

import (
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
)

func whiteList() map[string]string {
	return map[string]string{
		"/douyin/feed":           "GET",
		"/douyin/user/register/": "POST",
		"/douyin/user/login/":    "POST",
	}
}

func withinWhiteList(url *url.URL, method string) bool {
	target := whiteList()
	queryUrl := strings.Split(fmt.Sprint(url), "?")[0]
	if _, ok := target[queryUrl]; ok {
		if target[queryUrl] == method {
			return true
		}
		return false
	}
	return false
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		type QueryToken struct {
			Token string `binding:"required" form:"token"`
		}

		// 当路由不在白名单内时进行token检测
		if !withinWhiteList(c.Request.URL, c.Request.Method) {
			var queryToken QueryToken
			if c.ShouldBindQuery(&queryToken) != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status_code": 40001,
					"status_msg":  "请先登录",
				})
				return
			}
			// 验证token有效性
			userClaims, err := util.ParseToken(queryToken.Token)
			if err != nil {
				c.AbortWithStatusJSON(200, gin.H{
					"status_code": 40001,
					"status_msg":  "登陆过期",
				})
				return
			}
			c.Set("token", userClaims)
		}
		c.Next()
	}
}
