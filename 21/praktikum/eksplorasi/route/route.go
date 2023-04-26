package route

import (
	"alterra/golang/21/praktikum/eksplorasi/controller"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Route user
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController)
	e.PUT("/users/:id", controller.UpdateUserController)

	// Route blog
	e.GET("/blogs", controller.GetBlogsController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.POST("/blogs", controller.CreateBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)

	return e
}
