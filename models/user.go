package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int    `json:"id" gorm:"primaryKey"`
	Nickname string `json:"nickname" default:"user"`
	Username string `json:"username" unique:"true"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}
