package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	e.GET("/students", get_students_handler)
	e.Logger.Fatal(e.Start(":1323"))
}

func get_students_handler(c echo.Context) error {
	print(http.StatusOK)
	return c.String(http.StatusOK, "Hello, World!")
}
