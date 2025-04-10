package models

import "gorm.io/gorm"

type Item struct {
	/*
		- id: int, primary key, auto increment
		- seller_id: int, foreign key to users(id), not null
		- category_id: int, foreign key to categories(id), not null
		- title: string, not null
		- description: text, not null
		- price: decimal(10, 2), not null
		- original_price: decimal(10, 2), default 0.00
		- condition: uint, default 0 (0: new, 1: almost new, 2: slightly used, 3: used, 4: damaged) not null
		- location: string, not null
		- status: uint, default 0 (0: normal, 1: sold, 2: down, 3: checking, 4: invalid) not null
		- view_count: int, default 0
		- favorite_count: int, default 0
		- is_free_shipping: int, default 0 (0: no, 1: yes)
		- shipping_fee: decimal(10, 2), default 0.00
		- created_at: datetime, default current timestamp
		- updated_at: datetime, default current timestamp on update current timestamp
	*/
	gorm.Model
	SellerId       int    `json:"seller_id"`
	CategoryId     int    `json:"category_id"`
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	OriginalPrice  int    `json:"original_price"`
	Condition      uint   `json:"condition" gorm:"default:0"`
	Location       string `json:"location"`
	Status         uint   `json:"status" gorm:"default:0"`
	ViewCount      int    `json:"view_count" gorm:"default:0"`
	FavoriteCount  int    `json:"favorite_count" gorm:"default:0"`
	IsFreeShipping int    `json:"is_free_shipping" gorm:"default:1"`
	ShippingFee    int    `json:"shipping_fee" gorm:"default:0"`
}
