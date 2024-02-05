package services

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"errors"
	"gorm.io/gorm"
)

// Parent stuct to implement IProductService interface binding
type productService struct {
	repo domain.IProductRepo
}


// Interface binding
func ProductInstance(productRepo domain.IProductRepo) domain.IProductService {
	return &productService{
		repo: productRepo,
	}
}

// All the methods of IProductService interface are implemented here
// GetProducts implements domain.IProductService.
func (service *productService) GetProducts(model *gorm.Model) ([]models.Product, error) {
	var allProducts []models.Product
	Product, _ := service.repo.GetProducts(model)
	if len(Product) == 0 {
		return nil, errors.New("No Products found")
	}
	allProducts = append(allProducts, Product...)
	return allProducts, nil
}

// GetProductsByID implements domain.IProductService.
func (service *productService) GetProductByID(ID uint) (models.Product, error) {
	Product, err := service.repo.GetProductByID(ID)
	if err != nil {
		return models.Product{}, errors.New("Product not found")
	}
	return Product, nil
}

// CreateProduct implements domain.IProductService.
func (service *productService) CreateProduct(Product *models.Product) error {
	if err := service.repo.CreateProduct(Product); err != nil {
		return err
	}
	return nil
}

// DeleteProduct implements domain.IProductService.
func (service *productService) DeleteProduct(model *gorm.Model) error {
	if err := service.repo.DeleteProduct(model); err != nil {
		return errors.New("Product is not deleted")
	}
	return nil
}

// UpdateProduct implements domain.IProductService.
func (service *productService) UpdateProduct(Product *models.Product) error {
	if err := service.repo.UpdateProduct(Product); err != nil {
		return errors.New("Product is not updated")
	}
	return nil
}
