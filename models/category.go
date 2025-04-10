package models

import "gorm.io/gorm"

type Category struct {
	/*
		- id: int, primary key, auto increment
		- parent_id: int, foreign key to categories(id), default 0 (0: root category)
		- name: string, not null
		- icon: string, not null
		- sort_order: int, default 0 comment 'smaller number means higher priority'
		- status: uint, default 0 (0: normal, 1: hidden) not null
		- created_at: datetime, default current timestamp
		- updated_at: datetime, default current timestamp on update current timestamp
	*/
	gorm.Model
	ParentId  int    `json:"parent_id" gorm:"default:0"`
	Name      string `json:"name" gorm:"not null"`
	Icon      string `json:"icon" gorm:"not null"`
	SortOrder int    `json:"sort_order" gorm:"default:0"`
	Status    uint   `json:"status" gorm:"default:0"`
}
