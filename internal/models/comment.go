package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content" gorm:"type:text;not null"`
	UserId  uint   `json:"userId" gorm:"index;not null"`
	PostId  uint   `json:"postId" gorm:"index;not null"`
	User    User   `json:"user" gorm:"foreignKey:UserId;references:ID"`
}
