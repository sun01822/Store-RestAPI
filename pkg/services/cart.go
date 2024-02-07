package services

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"errors"
	"gorm.io/gorm"
)

// Parent stuct to implement ICartService interface binding
type cartService struct {
	repo domain.ICartRepository
}

// Interface binding
func CartInstance(cartRepo domain.ICartRepository) domain.ICartService {
	return &cartService{
		repo: cartRepo,
	}
}

// All the methods of ICartService interface are implemented here

// GetCartByCartID implements domain.ICartService.
func (service *cartService) GetCartByCartID(cartID uint) (*models.Cart, error) {
	Cart, err := service.repo.GetCartByCartID(cartID)
	if err != nil {
		return nil, errors.New("Cart not found")
	}
	return Cart, nil
}

// GetCartByUserID implements domain.ICartService.
func (service *cartService) GetCartByUserID(userID uint) ([]models.Cart, error) {
	Cart, err := service.repo.GetCartByUserID(userID)
	if err != nil {
		return nil, errors.New("Cart not found")
	}
	return Cart, nil
}

// CreateCart implements domain.ICartService.
func (service *cartService) CreateCart(Cart *models.Cart) error {
	if err := service.repo.CreateCart(Cart); err != nil {
		return err
	}
	return nil
}

// UpdateCart implements domain.ICartService.
func (service *cartService) UpdateCart(Cart *models.Cart) error {
	if err := service.repo.UpdateCart(Cart); err != nil {
		return err
	}
	return nil
}

// DeleteCart implements domain.ICartService.
func (service *cartService) DeleteCart(Cart *models.Cart) error {
	if err := service.repo.DeleteCart(Cart); err != nil {
		return errors.New("Cart is not deleted")
	}
	return nil
}

// GetAllCarts implements domain.ICartService.
func (service *cartService) GetAllCarts(model *gorm.Model) ([]models.Cart, error) {
	var allCarts []models.Cart
	Cart, _ := service.repo.GetAllCarts(model)
	if len(Cart) == 0 {
		return nil, errors.New("No Carts found")
	}
	allCarts = append(allCarts, Cart...)
	return allCarts, nil
}
