package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title string `json:"title"`
	Price float64 `json:"price"`
	Category string `json:"category"`
	Description string `json:"description"`
	Image string `json:"image"`
}