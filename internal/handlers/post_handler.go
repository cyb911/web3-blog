package handlers

import (
	"fmt"
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

func ListPost(c *gin.Context) {
	utils.OkMsg(c, "待开发")
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
