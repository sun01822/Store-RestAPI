package containers

import (
	"Store_RestAPI/pkg/config"
	"Store_RestAPI/pkg/connection"
	"Store_RestAPI/pkg/controllers"
	"Store_RestAPI/pkg/repositories"
	"Store_RestAPI/pkg/services"
	"fmt"
	"log"
	"Store_RestAPI/pkg/routes"
	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	
	// Config initializations
	config.SetConfig()

	// Database initializations
	db := connection.GetDB()
	if db == nil {
		fmt.Println("Database not connected")
	}
	fmt.Println("Database connected")

	// Repository initializations
	userRepository := repositories.UserDBInstance(db)
	productRepository := repositories.ProductDBInstance(db)
	cartRepository := repositories.CartDBInstance(db)

	// Service initializations
	userService := services.UserInstance(userRepository)
	productService := services.ProductInstance(productRepository)
	cartService := services.CartInstance(cartRepository)

	// Controller initializations
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	cartController := controllers.NewCartController(cartService)

	// Route initializations
	user := routes.UserRoutes(e, userController)
	user.InitUserRoutes()

	product := routes.ProductRoutes(e, productController)
	product.InitProductRoutes()

	cart := routes.CartRoutes(e, cartController)
	cart.InitCartRoutes()
	
	// Starting Server 
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}