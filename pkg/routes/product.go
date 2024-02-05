package routes

import (
	"Store_RestAPI/pkg/controllers"
	"github.com/labstack/echo/v4"
)

type productRoutes struct {
	echo *echo.Echo
	productController controllers.IProductController
}

func ProductRoutes(e *echo.Echo, productController controllers.IProductController) *productRoutes {
	return &productRoutes{
		echo:           	e,
		productController: 	productController,
	}
}

func (product *productRoutes) InitProductRoutes() {
	e := product.echo
	product.initProductRoutes(e)
}

func (pro *productRoutes) initProductRoutes(e *echo.Echo) {
	// grouping route endpoints
	product := e.Group("/store")

	product.POST("/product", pro.productController.CreateProduct)
	product.GET("/product", pro.productController.GetProducts)
	product.GET("/product/:id", pro.productController.GetProductByID)
	product.DELETE("/product/:id", pro.productController.DeleteProduct)
	product.PUT("/product/:id", pro.productController.UpdateProduct)
}



