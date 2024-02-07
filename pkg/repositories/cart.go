package repositories

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"gorm.io/gorm"
)

type cartRepo struct {
	db *gorm.DB
}

// CartDBInstance is a function to create a new Cart repository
func CartDBInstance(db *gorm.DB) domain.ICartRepository {
	return &cartRepo{
		db: db,
	}
}

// GetCartByCartID is a function to get a cart by cart ID
func (c *cartRepo) GetCartByCartID(cartID uint) (*models.Cart, error) {
	var cart models.Cart
	err := c.db.Where("id = ?", cartID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

// GetCartByUserID is a function to get a cart by user ID
func (c *cartRepo) GetCartByUserID(userID uint) ([]models.Cart, error) {
	var cart []models.Cart
	err := c.db.Where("user_id = ?", userID).Find(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// CreateCart is a function to create a new cart
func (c *cartRepo) CreateCart(cart *models.Cart) error {
	err := c.db.Create(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateCart is a function to update a cart
func (c *cartRepo) UpdateCart(cart *models.Cart) error {
	err := c.db.Save(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCart is a function to delete a cart
func (c *cartRepo) DeleteCart(cart *models.Cart) error {
	err := c.db.Delete(&cart).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAllCarts is a function to get all carts
func (c *cartRepo) GetAllCarts(model *gorm.Model) ([]models.Cart, error) {
	var carts []models.Cart
	err := c.db.Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

// GetCartByID is a function to get a cart by ID
func (c *cartRepo) GetCartByID(ID uint) (*models.Cart, error) {
	var cart models.Cart
	err := c.db.Where("id = ?", ID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}
