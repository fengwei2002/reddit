package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"reddit/controller"
	"reddit/logger"
	"reddit/middleware"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
		// call this line directly to enter release mode
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// use logger middleware

	// register service routing
	v1 := r.Group("/api/v1")
	v1.POST("/sign_up", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	// v1 group use jwtAuthMiddleware
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler) // 创建帖子
		// {
		// 	"title": "学习 golang",
		// 	"content": "必须学 mysql",
		// 	"community_id": 1
		// }
		v1.POST("/post/:id", controller.PostDetailHandler)
		v1.GET("/posts/", controller.PostListHandler) // 分页展示帖子列表
		// v1.GET("/posts2", controller.PostList2Handler) // 根据时间或者分数排序分页展示帖子列表
	}

	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		// if the current user is a sign in user
		// determine whether there is a valid jwt token in the request header
		// if jwt token is valid then send the pong
		c.String(http.StatusOK, "pong")
	})

	// change noRoute to 404 page
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404 Not Found",
		})
	})
	return r
}
