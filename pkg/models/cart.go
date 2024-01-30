package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID   uint         `json:"user_id"`
	Date     string       `json:"date"`
	Products []ProductsInCart `gorm:"foreignKey:CartID"`
}

type ProductsInCart struct {
	ID 	  uint   `json:"id"`
	ProductID uint `json:"product_id"`
	CartID uint `json:"cart_id"`
	Quantity uint `json:"quantity"`
}