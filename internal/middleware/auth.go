package middleware

import (
	"net/http"
	"strings"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			utils.FailMsg(c, "P00001", "Authorization 缺失或无效")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "error": "invalid token"})
			c.Abort()
			return
		}
		c.Set(CtxUserIDKey, claims.UserID)
		c.Next()
	}
}

func MustGetUserID(c *gin.Context) uint {
	if v, ok := c.Get(CtxUserIDKey); ok {
		if id, ok := v.(uint); ok {
			return id
		}
	}
	return 0
}
