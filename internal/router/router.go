package router

import (
	"log"
	"web-blog/internal/handlers"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
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
	r.POST("/api/login", func(c *gin.Context) {
		utils.OkMsg(c, "登录功能开发中。。。。。。")
	})

	// 用户管理模块

	return r
}
