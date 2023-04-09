package route

import (
	"alterra/golang/22_middleware/praktikum/controller"
	m "alterra/golang/22_middleware/praktikum/middleware"
	"alterra/golang/22_middleware/praktikum/constant"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	
	// all route using logger middleware
	m.LogMiddleware(e)

	
	// Route user
	e.POST("/login", controller.LoginUserController)
	e.POST("/users", controller.CreateUserController)
	// route above using JWT MIDDLEWARE
	eJWTUsers := e.Group("")
	eJWTUsers.Use(mid.JWT([]byte(constant.SECRET_JWT)))
	eJWTUsers.GET("/users", controller.GetUsersController)
	eJWTUsers.GET("/users/:id", controller.GetUserController)
	eJWTUsers.DELETE("/users/:id", controller.DeleteUserController)
	eJWTUsers.PUT("/users/:id", controller.UpdateUserController)
	
	// route book
	eJWTBooks := e.Group("")
	eJWTBooks.Use(mid.JWT([]byte(constant.SECRET_JWT)))
	eJWTBooks.GET("/books", controller.GetBooksController)
	eJWTBooks.GET("/books/:id", controller.GetBookController)
	eJWTBooks.DELETE("/books/:id", controller.DeleteBookController)
	eJWTBooks.PUT("/books/:id", controller.UpdateBookController)
	

	// Route blog
	e.GET("/blogs", controller.GetBlogsController)
	e.GET("/blogs/:id", controller.GetBlogController)
	e.POST("/blogs", controller.CreateBlogController)
	e.PUT("/blogs/:id", controller.UpdateBlogController)
	e.DELETE("/blogs/:id", controller.DeleteBlogController)
	
	// all route above using jwt middleware

	return e
}
