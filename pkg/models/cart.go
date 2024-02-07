package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID   uint         `json:"user_id"`
	Date     time.Time      `json:"date"`
	Product ProductsInCart `json:"product" gorm:"embedded"`
}

type ProductsInCart struct {
	ProductID uint `json:"product_id"`
	Quantity uint `json:"quantity"`
}