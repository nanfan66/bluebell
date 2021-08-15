package router

import (
	"bluebell01/controller"
	"bluebell01/logger"
	"bluebell01/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter(mode string) *gin.Engine  {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	// 使用中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册路由组
	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)
	}
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.CreatePostHandler)
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}