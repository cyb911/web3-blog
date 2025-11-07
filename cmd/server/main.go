package main

import (
	"log"
	"web-blog/internal/config"
	"web-blog/internal/models"
	"web-blog/internal/router"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	err := config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

	if err != nil {
		panic(err)
	}

	cfg := config.Get()

	// 设置路由
	r := router.SetupRouter()

	log.Printf("Server listening on :%s", cfg.AppPort)

	err = r.Run(":" + cfg.AppPort)
	if err != nil {
		panic(err)
	}
}
