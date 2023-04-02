package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json: "id"`
	Age   int    `json: "age"`
	Email string `json: "email"`
	Name  string `json: "name"`
}

func main() {
	e := echo.New()

	// routing
	e.GET("/hello", HelloController)
	e.GET("/user/:id/:age", UserController)
	e.POST("/user", RegisterController)
	e.POST("/bindingjson", BindingController)

	e.Start(":8000")
}

// response hello world
func HelloController(e echo.Context)error{
	return e.JSON(http.StatusOK, "mantap")
}

func UserController(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	age, _ := strconv.Atoi(e.Param("age"))
	search := e.QueryParam("search")

	user := User{id, age, "joko@gmail.com", "Joko"}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"user":   user,
		"search": search,
	})
	// return e.String(200, "Hello World")	// same
}

func RegisterController(e echo.Context) error {
	email := e.FormValue("email")
	name := e.FormValue("name")

	return e.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
		"name":  name,
	})
}

func BindingController(e echo.Context) error {
	user := User{}
	e.Bind(&user)

	return e.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}
