package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname    string `json:"nickname" gorm:"default:user"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	CreditScore int    `json:"creditscore" gorm:"default:100"`
	Status      uint   `json:"status" gorm:"default:0"`
	LastLoginAt string `json:"lastloginat"`
}
