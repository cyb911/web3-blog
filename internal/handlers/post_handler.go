package handlers

import (
	"fmt"
	"strconv"
	"web-blog/internal/config"
	"web-blog/internal/middleware"
	"web-blog/internal/models"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreatePostReq struct {
	Title   string `json:"title" binding:"required,min=1,max=225"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var req CreatePostReq
	if err := c.ShouldBindJSON(&req); err != nil {
		if err.Error() == "EOF" {
			utils.FailMsg(c, "P00001", "请填写参数！")
		} else {
			utils.FailMsg(c, "P00001", err.Error())
		}
		return
	}
	userId := middleware.MustGetUserID(c)
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserId:  userId,
	}
	if err := config.DB.Create(&post).Error; err != nil {
		panic(err)
	}
	utils.Ok(c)
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
	id := c.Param("id")
	var post models.Post
	if err := config.DB.Preload("Author").Where("id = ?", id).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.FailMsg(c, "P000003", "文章不存在")
			return
		}
		panic(err)
	}
	utils.OkData(c, post)
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
