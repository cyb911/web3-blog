package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": "S00000", "msg": "success"})
}

func OkMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": "S00000", "msg": msg})
}

func OkData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"code": "S00000", "msg": "success", "data": data})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "Fail"})
}

func FailMsg(c *gin.Context, errCode string, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": errCode, "msg": msg})
}
