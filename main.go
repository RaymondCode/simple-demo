package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/controller"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/global"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", global.CheckLogin(), controller.UserInfo)
	apiRouter.POST("/user/register/", global.CheckLogin(), controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", global.CheckLogin(), controller.Publish)
	apiRouter.GET("/publish/list/", global.CheckLogin(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", global.CheckLogin(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", global.CheckLogin(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", global.CheckLogin(), controller.CommentAction)
	apiRouter.GET("/comment/list/", global.CheckLogin(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", global.CheckLogin(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", global.CheckLogin(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", global.CheckLogin(), controller.FollowerList)
}

func init() {
	dao.InitDB()
}

func main() {
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
