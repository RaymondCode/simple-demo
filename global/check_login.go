package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/service"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		userId, err := GetUserIdFromToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, api.Response{StatusCode: api.TokenInvalidErr, StatusMsg: "非法token！"})
			return
		}
		if exist, err := service.UserExist(userId); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, api.Response{StatusCode: api.InnerErr, StatusMsg: fmt.Sprintf("service.UserExist error: %s", err)})
			return
		} else if !exist {
			c.AbortWithStatusJSON(http.StatusOK, api.Response{StatusCode: api.UserNotExistErr, StatusMsg: fmt.Sprintf("用户 [%v] 不存在！", userId)})
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
