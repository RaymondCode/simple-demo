package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

const (
	ERROR   = -1
	SUCCESS = 0
)

func Result(code int32, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, c)
}
