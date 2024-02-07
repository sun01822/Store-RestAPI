package controllers

import (
	"Store_RestAPI/pkg/domain"
	"Store_RestAPI/pkg/models"
	"Store_RestAPI/pkg/types"
	"net/http"
	"strconv"
	"time"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ICartController is an interface for Cart controller
type ICartController interface {
	GetCartByUserID(e echo.Context) error
	GetCartByCartID(e echo.Context) error
	CreateCart(e echo.Context) error
	UpdateCart(e echo.Context) error
	DeleteCart(e echo.Context) error
	GetAllCarts(e echo.Context) error
}

type cartController struct {
	cartsvc domain.ICartService
}

// CartInstance is a function to create an instance of Cart controller
func NewCartController(cartSvc domain.ICartService) ICartController {
	return &cartController{
		cartsvc: cartSvc,
	}
}

// CreateCart implements ICartController.
func (controller *cartController) CreateCart(e echo.Context) error {
	reqCart := &types.CartRequest{}
	if err := e.Bind(reqCart); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	if err := reqCart.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	Cart := &models.Cart{
		UserID: reqCart.UserID,
		Date:   time.Now(),
		Product: models.ProductsInCart{
			ProductID: reqCart.ProductID,
			Quantity:  reqCart.Quantity,
		},
	}
	if err := controller.cartsvc.CreateCart(Cart); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Cart created successfully")
}

// GetCartByUserID implements ICartController.
func (controller *cartController) GetCartByUserID(e echo.Context) error {
	userID, _ := strconv.Atoi(e.Param("id"))
	Cart, err := controller.cartsvc.GetCartByUserID(uint(userID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, Cart)
}

// GetCartByCartID implements ICartController.
func (controller *cartController) GetCartByCartID(e echo.Context) error {
	cartID, _ := strconv.Atoi(e.Param("id"))
	Cart, err := controller.cartsvc.GetCartByCartID(uint(cartID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, Cart)
}


// UpdateCart implements ICartController.
func (controller *cartController) UpdateCart(e echo.Context) error {
	cartID, _ := strconv.Atoi(e.Param("id"))
	existingCart, err := controller.cartsvc.GetCartByCartID(uint(cartID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	reqCart := &types.CartUpdateRequest{}
	if err := e.Bind(reqCart); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid data")
	}
	cart := &models.Cart{
		Model: gorm.Model{ID: uint(cartID), CreatedAt: existingCart.CreatedAt, UpdatedAt: time.Now(), DeletedAt: existingCart.DeletedAt},
		UserID: reqCart.UserID,
		Date:   time.Now(),
		Product: models.ProductsInCart{
			ProductID: reqCart.ProductID,
			Quantity:  reqCart.Quantity,
		},
	}
	if cart.UserID == 0 {
		cart.UserID = existingCart.UserID
	}
	if cart.Product.ProductID == 0 {
		cart.Product.ProductID = existingCart.Product.ProductID
	}
	if cart.Product.Quantity == 0 {
		cart.Product.Quantity = existingCart.Product.Quantity
	}
	if err := controller.cartsvc.UpdateCart(cart); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Cart updated successfully")
}

// DeleteCart implements ICartController.
func (controller *cartController) DeleteCart(e echo.Context) error {
	cartID, _ := strconv.Atoi(e.Param("id"))
	existingCart, err := controller.cartsvc.GetCartByCartID(uint(cartID))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	if err := controller.cartsvc.DeleteCart(existingCart); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, "Cart deleted successfully")
}

// GetAllCarts implements ICartController.
func (controller *cartController) GetAllCarts(e echo.Context) error {
	Carts, err := controller.cartsvc.GetAllCarts(&gorm.Model{})
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, Carts)
}
