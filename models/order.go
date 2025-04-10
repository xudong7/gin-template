package models

import "gorm.io/gorm"

type Order struct {
	/*
		- id: int, primary key, auto increment
		- order_no: string, unique, not null
		- buyer_id: int, foreign key to users(id), not null
		- seller_id: int, foreign key to users(id), not null
		- item_id: int, foreign key to items(id), not null
		- item_snapshot: text, not null
		- price: decimal(10, 2), not null
		- shipping_fee: decimal(10, 2), default 0.00
		- total_price: decimal(10, 2), not null
		- address_id: int, foreign key to user_addresses(id), not null
		- address_snapshot: text, not null
		- status: uint, default 0 (0: unpaid, 1: paid, 2: shipped, 3: completed, 4: canceled, 5: refunded) not null
		- payment_method: uint, default 0 (0: alipay, 1: wechat, 2: bank transfer) not null
		- payment_time: datetime, default null
		- shipping_time: datetime, default null
		- delivery_company: string, default null
		- tracking_number: string, default null
		- completed_time: datetime, default null
		- cancel_reason: string, default null
		- created_at: datetime, default current timestamp
		- updated_at: datetime, default current timestamp on update current timestamp
	*/
	gorm.Model
	OrderNo         string `json:"order_no" gorm:"unique"`
	BuyerId         int    `json:"buyer_id" binding:"required" gorm:"default:-1"`
	SellerId        int    `json:"seller_id" binding:"required" gorm:"default:-1"`
	ItemId          int    `json:"item_id" binding:"required" gorm:"default:-1"`
	ItemSnapshot    string `json:"item_snapshot"`
	Price           int    `json:"price"`
	ShippingFee     int    `json:"shipping_fee" gorm:"default:0"`
	TotalPrice      int    `json:"total_price"`
	AddressId       int    `json:"address_id"`
	AddressSnapshot string `json:"address_snapshot"`
	Status          uint   `json:"status" gorm:"default:0"`
	PaymentMethod   uint   `json:"payment_method" gorm:"default:0"`
	PaymentTime     string `json:"payment_time"`
	ShippingTime    string `json:"shipping_time"`
	DeliveryCompany string `json:"delivery_company"`
	TrackingNumber  string `json:"tracking_number"`
	CompletedTime   string `json:"completed_time"`
	CancelReason    string `json:"cancel_reason"`
}
