package test

import (
	"testing"
	"web-blog/internal/config"
	"web-blog/internal/models"
)

// run automatic database migrations
func TestAutoMigrate(t *testing.T) {
	config.InitDB()
	defer config.CloseDB()
	_ = config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
