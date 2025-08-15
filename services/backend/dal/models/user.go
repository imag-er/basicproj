package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" form:"username"`
	Password string `form:"password"`
}
