package utils

import (
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/gin-gonic/gin"
)

//获取UserId
func GetUserId(c *gin.Context) int64 {
	userId, flag := c.Get("userId")
	if !flag {
		response.FailWithMessage("userId not exist ", c)
		return 0
	}
	id, _ := userId.(int64)
	return id
}
