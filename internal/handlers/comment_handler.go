package handlers

import (
	"fmt"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

type CreateCommentReq struct {
	PostID  uint   `json:"post_id"  binding:"required"`
	Content string `json:"content"  binding:"required"`
}

func CreateComment(c *gin.Context) {
	utils.OkMsg(c, "待开发中")
}

func ListComments(c *gin.Context) {
	postID := c.Param("postId")
	fmt.Println(postID)
	utils.OkMsg(c, "待开发中")
}
