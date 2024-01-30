package containers

import (
	"Store_RestAPI/pkg/config"
	"Store_RestAPI/pkg/connection"
	"fmt"
	"log"
	"github.com/labstack/echo/v4"
)

func Serve(e *echo.Echo) {
	
	// config initialization
	config.SetConfig()

	// database initialization
	db := connection.GetDB()
	if db == nil {
		fmt.Println("Database not connected")
	}
	fmt.Println("Database connected")

	// Starting Server 
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}