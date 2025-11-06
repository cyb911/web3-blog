package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"size:64"`
	Email    string `json:"email" gorm:"size:128;uniqueIndex"`
	Password string `json:"-" gorm:"size:255"`
}
