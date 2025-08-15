package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `form:"title,required"`
	Content string `form:"content,required"`
	Preview string `form:"preview,required"`
	Likes   int    `form:"likes" gorm:"default:0"`
}
