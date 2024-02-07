package domain

import (
	"Store_RestAPI/pkg/models"
	"gorm.io/gorm"
)

// ICartService is an interface for Cart service
type ICartService interface {
	GetCartByUserID(userID uint) ([]models.Cart, error)
	GetCartByCartID(cartID uint) (*models.Cart, error)
	CreateCart(cart *models.Cart) error
	UpdateCart(cart *models.Cart) error
	DeleteCart(cart *models.Cart) error
	GetAllCarts(*gorm.Model) ([]models.Cart, error)
}

// ICartRepository is an interface for Cart repository
type ICartRepository interface {
	GetCartByUserID(userID uint) ([]models.Cart, error)
	GetCartByCartID(cartID uint) (*models.Cart, error)
	CreateCart(cart *models.Cart) error
	UpdateCart(cart *models.Cart) error
	DeleteCart(cart *models.Cart) error
	GetAllCarts(*gorm.Model) ([]models.Cart, error)
}
