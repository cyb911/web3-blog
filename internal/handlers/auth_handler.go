package handlers

import (
	"log"
	"web-blog/internal/config"
	"web-blog/internal/models"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required,min=3,max=64"`
	Email    string `json:"email"    binding:"required,email,max=128"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

type LoginReq struct {
	Email    string `json:"email"    binding:"required,email,max=128"`
	Password string `json:"password" binding:"required,min=6,max=64"`
}

func Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailMsg(c, "P00001", err.Error())
		return
	}
	var count int64
	config.DB.Model(&models.User{}).Where("email = ? ", req.Email).Count(&count)

	if count > 0 {
		utils.FailMsg(c, "P00002", "此邮箱已经被注册了！")
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	u := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hash),
	}

	err := config.DB.Create(&u).Error
	if err != nil {
		log.Printf("注册用户失败！原因：%v", err)
		utils.FailMsg(c, "P00003", "注册用户失败！")
	}

	utils.OkMsg(c, "注册成功！")

}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailMsg(c, "P00001", err.Error())
		return
	}
	var user models.User
	if err := config.DB.Where("email = ? ", req.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.FailMsg(c, "P00003", "此邮箱号未注册！")
			return
		}
		panic(err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		utils.FailMsg(c, "P00003", "invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		panic(err)
	}

	utils.OkData(c, "Bearer "+token)

}
