package handlers

import (
	"fmt"
	"strconv"
	"web-blog/internal/config"
	"web-blog/internal/models"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Title   string `json:"title" binding:"required,min=1,max=225"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	utils.OkMsg(c, "待开发")
}

func ListPosts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []models.Post
	err := config.DB.Preload("Author").Order("id desc").Limit(limit).Offset(offset).
		Find(&posts).Error

	if err != nil {
		panic(err)
	}

	utils.OkData(c, posts)
}

func GetPost(c *gin.Context) {
	utils.OkMsg(c, "待开发")
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	utils.OkMsg(c, "待开发")
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	utils.OkMsg(c, "待开发")
}
