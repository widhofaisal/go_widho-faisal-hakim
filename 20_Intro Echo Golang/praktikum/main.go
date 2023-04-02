package main

import (
	_ "fmt"
	"net/http"
	"strconv"
	_ "strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	for i, value := range users {
		if value.Id == id {
			return c.JSON(http.StatusOK, users[i])
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user by id not found",
		"id":      id,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))

	for i, value := range users {
		if value.Id == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "success delete user",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user by id not found",
		"id":      id,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{}
	c.Bind(&user)

	for i, value := range users {
		if value.Id == id {
			users[i] = user
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message":"success update user",
				"user":users[i],
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user by id not found",
		"id":      id,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
