package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/controller"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/rdb"
	"github.com/warthecatalyst/douyin/service"
	"github.com/warthecatalyst/douyin/tokenx"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		userId, username := tokenx.ParseToken(token)
		if username == "" {
			// TODO: 端上应该重定向到登录界面
			c.AbortWithStatusJSON(http.StatusOK, api.Response{StatusCode: api.LogicErr, StatusMsg: "非法token"})
			return
		}
		if user, err := service.NewUserServiceInstance().GetUserByUserId(userId); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, api.Response{StatusCode: api.LogicErr, StatusMsg: "内部错误"})
			return
		} else if user == nil {
			c.AbortWithStatusJSON(http.StatusOK, api.Response{StatusCode: api.LogicErr, StatusMsg: "非法用户"})
			return
		}
		c.Set("user_id", userId)
		c.Next()
	}
}

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", CheckLogin(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", CheckLogin(), controller.Publish)
	apiRouter.GET("/publish/list/", CheckLogin(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", CheckLogin(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", CheckLogin(), controller.FavoriteList)
	apiRouter.POST("/common/action/", CheckLogin(), controller.CommentAction)
	apiRouter.GET("/common/list/", CheckLogin(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", CheckLogin(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", CheckLogin(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", CheckLogin(), controller.FollowerList)
}

func initAll() {
	dao.InitDB()
	rdb.InitRdb()
}

func main() {
	initAll()
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
