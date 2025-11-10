package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"time"
	"web-blog/internal/config"
	"web-blog/internal/models"
	"web-blog/internal/utils"

	"github.com/gin-gonic/gin"
)

type 

func GetNonce(c *gin.Context) {
	address := strings.ToLower(strings.TrimSpace(c.Query("address")))
	if len(address) != 42 || !strings.HasPrefix(address, "0x") {
		utils.FailMsg(c, "A00001", "invalid address")
		return
	}

	nonce, err := randNonce(16)
	if err != nil {
		utils.FailMsg(c, "A00002", "failed to generate nonce")
		return
	}

	ln := models.LoginNonce{
		Address:   address,
		Nonce:     nonce,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Used:      true,
	}

	if err := config.DB.Create(&ln).Error; err != nil {
		panic(err)
	}

	// SIWE 文本信息
	domain := "localhost:9010"
	uri := "https://localhost:9010/auth/callback"
	statement := "Sign in to Web3 Blog"
	chainId := 1 // 以太坊主网；测试网改对应ID
	version := "1"

}

func randNonce(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return strings.ToLower(hex.EncodeToString(b)), nil
}
