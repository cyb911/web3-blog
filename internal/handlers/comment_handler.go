package handlers

import (
	"strconv"
	"web-blog/internal/config"
	"web-blog/internal/middleware"
	"web-blog/internal/models"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

type CreateCommentReq struct {
	Content string `json:"content"  binding:"required"`
}

func CreateComment(c *gin.Context) {
	postId64, err := strconv.ParseUint(c.Param("postId"), 10, 64)
	if err != nil {
		panic(err)
	}
	postId := uint(postId64)

	var req CreateCommentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		if err.Error() == "EOF" {
			utils.FailMsg(c, "P00001", "请填写参数！")
		} else {
			utils.FailMsg(c, "P00001", err.Error())
		}
		return
	}
	userId := middleware.MustGetUserID(c)
	comment := models.Comment{
		Content: req.Content,
		UserId:  userId,
		PostId:  postId,
	}
	if err := config.DB.Create(&comment).Error; err != nil {
		panic(err)
	}
	utils.Ok(c)
}

func ListComments(c *gin.Context) {
	postID := c.Param("id")
	var list []models.Comment
	if err := config.DB.Preload("User").
		Where("post_id = ?", postID).
		Order("id ASC").Find(&list).Error; err != nil {
		panic(err)
	}
	utils.OkData(c, list)
}
