package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/service"
	"github.com/warthecatalyst/douyin/tokenx"
	"net/http"
)

// middleware
func checkLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		userId, err := tokenx.GetUserIdFromToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, service.Response{StatusCode: tokenInvalidErr, StatusMsg: "非法token！"})
			return
		}
		if exist, err := service.UserExist(userId); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, service.Response{StatusCode: innerErr, StatusMsg: fmt.Sprintf("service.UserExist error: %s", err)})
			return
		} else if !exist {
			c.AbortWithStatusJSON(http.StatusOK, service.Response{StatusCode: userNotExistErr, StatusMsg: fmt.Sprintf("用户 [%v] 不存在！", userId)})
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", Feed)
	apiRouter.GET("/user/", checkLogin(), UserInfo)
	apiRouter.POST("/user/register/", checkLogin(), Register)
	apiRouter.POST("/user/login/", Login)
	apiRouter.POST("/publish/action/", checkLogin(), Publish)
	apiRouter.GET("/publish/list/", checkLogin(), PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", checkLogin(), FavoriteAction)
	apiRouter.GET("/favorite/list/", checkLogin(), FavoriteList)
	apiRouter.POST("/comment/action/", checkLogin(), CommentAction)
	apiRouter.GET("/comment/list/", checkLogin(), CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", checkLogin(), RelationAction)
	apiRouter.GET("/relation/follow/list/", checkLogin(), FollowList)
	apiRouter.GET("/relation/follower/list/", checkLogin(), FollowerList)
}
