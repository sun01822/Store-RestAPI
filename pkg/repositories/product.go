package repositories

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"errors"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}


// interface binding
func ProductDBInstance(db *gorm.DB) domain.IProductRepo {
	return &productRepo{
		db: db,
	}
}

// GetProducts implements domain.IProductRepo interface
func (p *productRepo) GetProducts(model *gorm.Model) ([]models.Product, error) {
	var products []models.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

// GetProductByID implements domain.IProductRepo interface
func (p *productRepo) GetProductByID(ID uint) (models.Product, error) {
	var product models.Product
	err := p.db.Where("id = ?", ID).First(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}



// CreateProduct implements domain.IProductRepo interface
func (p *productRepo) CreateProduct(product *models.Product) error {
	productTitle := product.Title

	// Check if the product already exists
	var existingProduct models.Product
	if err := p.db.Where("title = ?", productTitle).First(&existingProduct).Error; err == nil {
		return errors.New("Product title already exists")
	}
	err := p.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteProduct implements domain.IProductRepo interface
func (p *productRepo) DeleteProduct(model *gorm.Model) error {
	var product models.Product
	err := p.db.Where("id = ?", model.ID).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateProduct implements domain.IProductRepo interface
func (p *productRepo) UpdateProduct(product *models.Product) error {
	err := p.db.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}
