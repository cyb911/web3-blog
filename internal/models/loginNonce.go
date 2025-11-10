package models

import (
	"time"

	"gorm.io/gorm"
)

type LoginNonce struct {
	gorm.Model
	Address   string
	Nonce     string
	ExpiresAt time.Time
	Used      bool
}
