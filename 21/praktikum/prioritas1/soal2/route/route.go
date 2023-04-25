package route

import (
	"alterra/golang/21/praktikum/prioritas1/soal2/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Route / to handler function
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	return e
}
