package router

import (
	"log"
	"web-blog/internal/handlers"
	"web-blog/internal/middleware"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(func(c *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic: %v", rec)
				utils.FailMsg(c, "E99999", "系统内部错误！")
			}
		}()
		c.Next()
	})

	// 测试
	r.GET("/health", func(c *gin.Context) {
		utils.OkMsg(c, "health!")
	})

	// 注册
	r.POST("/api/register", handlers.Register)
	// 登录操作
	r.POST("/api/login", handlers.Login)

	// 文章查询
	r.GET("/api/posts", handlers.ListPost)
	r.GET("/api/posts/:id", handlers.GetPost)

	// 评论信息
	r.GET("/api/posts/:id/comments", handlers.ListComments)

	// 注册 Auth 中间件
	authGroup := r.Group("/api", middleware.AuthRequired())
	{
		//用户管理

		// 文章管理
		authGroup.POST("/posts", handlers.CreatePost)
		authGroup.PUT("/posts/:id", handlers.UpdatePost)
		authGroup.DELETE("/posts/:id", handlers.DeletePost)

		// 登录用户评论
		authGroup.POST("/comments", handlers.CreateComment)
	}

	return r
}
