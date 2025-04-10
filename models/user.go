package models

import "gorm.io/gorm"

type User struct {
	/*
		- id: int, primary key, auto increment
		- nickname: string, not null
		- username: string, unique, not null
		- password: string, not null
		- email: string, unique, not null
		- phone: string, unique
		- avatar: string, url
		- credit_score: int, default 100
		- status: uint, default 0 (0: normal, 1: locked, 2: deleted) not null
		- last_login_at: datetime, default current timestamp
		- created_at: datetime, default current timestamp
		- updated_at: datetime, default current timestamp on update current timestamp
	*/
	gorm.Model
	Nickname    string `json:"nickname" gorm:"default:user"`
	Username    string `json:"username" gorm:"unique"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	CreditScore int    `json:"credit_score" gorm:"default:100"`
	Status      uint   `json:"status" gorm:"default:0"`
	LastLoginAt string `json:"last_login_at"`
}
