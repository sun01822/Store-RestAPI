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
	user.DELETE("/user/:id", us.userController.DeleteUser)
	user.PUT("/user/:id", us.userController.UpdateUser)
	user.POST("/login", us.userController.LoginUser)
}



