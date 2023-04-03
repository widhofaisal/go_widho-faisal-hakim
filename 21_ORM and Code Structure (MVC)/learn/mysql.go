package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var DB *gorm.DB

func init() {
	initDB()
	DB.AutoMigrate()
}

func main() {
	e := echo.New()

	e.GET("/users", GetUserHandler)
	e.POST("/users", CreateUserHandler)

	e.Start(":10000")
}

func initDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/alterra?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetUserHandler(e echo.Context) error {
	var users []User
	DB.Find(&users)
	return e.JSON(http.StatusOK, users)
}

func CreateUserHandler(c echo.Context) error {
	var user User
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	result := DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	data := map[string]interface{}{
		"message": "success",
	}

	return c.JSON(http.StatusOK, data)
}
