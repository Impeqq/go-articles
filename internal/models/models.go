package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"size:200;not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
}
