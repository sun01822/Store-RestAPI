package routes

import (
	"Store_RestAPI/pkg/controllers"
	"github.com/labstack/echo/v4"
)

type cartRoutes struct {
	echo          *echo.Echo
	cartController controllers.ICartController
}

func CartRoutes(e *echo.Echo, cartController controllers.ICartController) *cartRoutes {
	return &cartRoutes{
		echo:           e,
		cartController: cartController,
	}
}

func (cart *cartRoutes) InitCartRoutes() {
	e := cart.echo
	cart.initCartRoutes(e)
}

func (cart *cartRoutes) initCartRoutes(e *echo.Echo) {
	c := e.Group("/store")
	
	c.POST("/cart", cart.cartController.CreateCart)
	c.GET("/cart/user/:id", cart.cartController.GetCartByUserID)
	c.PUT("/cart/:id", cart.cartController.UpdateCart)
	c.DELETE("/cart/:id", cart.cartController.DeleteCart)
	c.GET("/carts", cart.cartController.GetAllCarts)
	c.GET("/cart/:id", cart.cartController.GetCartByCartID)
}
