package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/util"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", util.VerifyJwt(), controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", util.VerifyJwt(), controller.Publish)
	apiRouter.GET("/publish/list/", util.VerifyJwt(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", util.VerifyJwt(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", util.VerifyJwt(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", util.VerifyJwt(), controller.CommentAction)
	apiRouter.GET("/comment/list/", util.VerifyJwt(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", util.VerifyJwt(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", util.VerifyJwt(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", util.VerifyJwt(), controller.FollowerList)
}
