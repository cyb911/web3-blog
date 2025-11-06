package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success"})
}

func OkMsg(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "Fail"})
}

func FailMsg(c *gin.Context, status string, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": status, "msg": msg})
}
