package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"type:text;not null"`
	UserId  uint   `json:"user_id" gorm:"index;not null"`
	PostId  uint   `json:"post_id" gorm:"index;not null"`
}
