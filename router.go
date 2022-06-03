package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	auth := controller.AuthMiddleware(&controller.JwtAuth{}, false)
	optionalAuth := controller.AuthMiddleware(&controller.JwtAuth{}, true)

	// public directory is used to serve static resources
	r.Static("/"+config.Config.RemoteVideoPath, config.Config.LocalVideoPath)

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", optionalAuth, controller.Feed)
	apiRouter.GET("/user/", auth, controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", auth, controller.Publish)
	apiRouter.GET("/publish/list/", auth, controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", auth, controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", auth, controller.FavoriteList)
	apiRouter.POST("/comment/action/", auth, controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", auth, controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", auth, controller.FollowList)
	apiRouter.GET("/relation/follower/list/", auth, controller.FollowerList)
}
