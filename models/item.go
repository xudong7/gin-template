package models

import "gorm.io/gorm"

type Item struct {
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
