package main

import (
	"web-blog/internal/config"
	"web-blog/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	err := config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	if err != nil {
		panic(err)
	}

	cfg := config.Get()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.Any("/login", func(c *gin.Context) {
		c.String(200, "login")
	})
	err = router.Run(":" + cfg.AppPort)
	if err != nil {
		panic(err)
	}
}
