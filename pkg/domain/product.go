package domain

import (
	"Store_RestAPI/pkg/models"
	"gorm.io/gorm"
)

// for database Repository operation (call from service)
type IProductRepo interface {
	GetProducts(*gorm.Model)([]models.Product, error)
	GetProductByID(ID uint)(models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(*gorm.Model) error
}

// for service operation (response to contorller || call from controller)
type IProductService interface {
	GetProducts(*gorm.Model)([]models.Product, error)
	GetProductByID(ID uint)(models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(*gorm.Model) error
}