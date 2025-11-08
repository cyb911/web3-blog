package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"username" gorm:"size:225;not null"`
	Content string `json:"content" gorm:"type:longtext;not null"`
	UserId  uint   `json:"user_id" gorm:"index;not null"`
	Author  User   `json:"author" gorm:"foreignKey:UserId;references:ID"`
}
