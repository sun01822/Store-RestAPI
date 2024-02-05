package controllers

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"Store_RestAPI/pkg/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// IproductController is an interface for Product controller
type IProductController interface {
	GetProducts(e echo.Context) error
	GetProductByID(e echo.Context) error
	CreateProduct(e echo.Context) error
	DeleteProduct(e echo.Context) error
	UpdateProduct(e echo.Context) error
}

type productController struct {
	productsvc domain.IProductService
}

// ProductInstance is a function to create an instance of Product controller
func NewProductController(productSvc domain.IProductService) IProductController {
	return &productController{
		productsvc: productSvc,
	}
}

// CreateProduct implements IproductController.
func (controller *productController) CreateProduct(e echo.Context) error {
	reqProduct := &types.ProductRequest{}
	if err := e.Bind(reqProduct); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := reqProduct.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	Product := &models.Product{
		Title:       reqProduct.Title,
		Price:       reqProduct.Price,
		Category:    reqProduct.Category,
		Description: reqProduct.Description,
		Image:       reqProduct.Image,
	}
	if err := controller.productsvc.CreateProduct(Product); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Product created successfully")
}

// GetProducts implements IproductController.
func (controller *productController) GetProducts(e echo.Context) error {
	Products, err := controller.productsvc.GetProducts(&gorm.Model{})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, Products)
}

// GetProductsByID implements IproductController.
func (controller *productController) GetProductByID(e echo.Context) error {
	tempProductID := e.Param("id")
	ProductID, err := strconv.ParseInt(tempProductID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Product ID")
	}
	Product, err := controller.productsvc.GetProductByID(uint(ProductID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, Product)
}

// DeleteProduct implements IproductController.
func (controller *productController) DeleteProduct(e echo.Context) error {
	tempProductID := e.Param("id")
	ProductID, err := strconv.ParseInt(tempProductID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Product ID")
	}
	_, err = controller.productsvc.GetProducts(&gorm.Model{ID: uint(ProductID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.productsvc.DeleteProduct(&gorm.Model{ID: uint(ProductID)}); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Product deleted successfully")
}

// UpdateProduct implements IproductController.
func (controller *productController) UpdateProduct(e echo.Context) error {
	reqProduct := &types.ProductRequest{}
	if err := e.Bind(reqProduct); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	tempProductID := e.Param("id")
	ProductID, err := strconv.ParseInt(tempProductID, 0, 0)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Enter a valid Product ID")
	}
	existingProduct, err := controller.productsvc.GetProducts(&gorm.Model{ID: uint(ProductID)})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	updateProduct := &models.Product{
		Model:       gorm.Model{ID: uint(ProductID), CreatedAt: existingProduct[0].CreatedAt, UpdatedAt: time.Now(), DeletedAt: existingProduct[0].DeletedAt},
		Title:       reqProduct.Title,
		Price:       reqProduct.Price,
		Category:    reqProduct.Category,
		Description: reqProduct.Description,
		Image:       reqProduct.Image,
	}
	if updateProduct.Title == "" {
		updateProduct.Title = existingProduct[0].Title
	}
	if updateProduct.Price == 0 {
		updateProduct.Price = existingProduct[0].Price
	}
	if updateProduct.Category == "" {
		updateProduct.Category = existingProduct[0].Category
	}
	if updateProduct.Description == "" {
		updateProduct.Description = existingProduct[0].Description
	}
	if updateProduct.Image == "" {
		updateProduct.Image = existingProduct[0].Image
	}
	if err := controller.productsvc.UpdateProduct(updateProduct); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Product updated successfully")
}
