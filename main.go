package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warthecatalyst/douyin/controller"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/middleware"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.CheckLogin(), controller.UserInfo)
	apiRouter.POST("/user/register/", middleware.CheckLogin(), controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middleware.CheckLogin(), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.CheckLogin(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.CheckLogin(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.CheckLogin(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.CheckLogin(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.CheckLogin(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", middleware.CheckLogin(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.CheckLogin(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.CheckLogin(), controller.FollowerList)
}

func init() {
	dao.InitDB()
}

func main() {
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
