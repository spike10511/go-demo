package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

// User 定义了用户的结构体
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
