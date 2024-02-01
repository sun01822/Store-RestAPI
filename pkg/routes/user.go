package routes

import (
	"Store_RestAPI/pkg/controllers"
	"github.com/labstack/echo/v4"
)

type userRoutes struct {
	echo *echo.Echo
	userController controllers.IUserController
}

func UserRoutes(e *echo.Echo, userController controllers.IUserController) *userRoutes {
	return &userRoutes{
		echo:           	e,
		userController: 	userController,
	}
}

func (user *userRoutes) InitUserRoutes() {
	e := user.echo
	user.initUserRoutes(e)
}

func (us *userRoutes) initUserRoutes(e *echo.Echo) {
	// grouping route endpoints
	user := e.Group("/store")

	user.POST("/user", us.userController.CreateUser)
	user.GET("/user", us.userController.GetUsers)
	user.GET("/user/info", us.userController.GetUsersInfo)
	user.DELETE("/user/:id", us.userController.DeleteUser)
	user.PUT("/user/:id", us.userController.UpdateUser)
	user.POST("/login", us.userController.LoginUser)

	user.POST("/user/address/:id", us.userController.CreateAddress)
	user.PUT("/user/address/:id", us.userController.UpdateAddress)
	user.GET("/user/address/:id", us.userController.GetAddress)


	user.POST("/user/address/geolocation/:id", us.userController.CreateGeoLocation)
	user.PUT("/user/address/geolocation/:id", us.userController.UpdateGeoLocation)

	
	user.POST("/user/name/:id", us.userController.CreateName)
	user.PUT("/user/name/:id", us.userController.UpdateName)
	user.GET("/user/name/:id", us.userController.GetName)
}



