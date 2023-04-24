package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"alterra/golang/22_middleware/praktikum/config"
	"alterra/golang/22_middleware/praktikum/model"
	"alterra/golang/22_middleware/praktikum/middleware"
)

func LoginUserController(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Where("email=? AND password=?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    userResponse,
	})
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User

	err := config.DB.First(&user, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get user, id not found",
			"id":      id,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"data":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := config.DB.First(&user, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get user, id not found",
			"id":      id,
		})
	}
	config.DB.Delete(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by id",
		"data":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	err := config.DB.First(&user, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get user, id not found",
			"id":      id,
		})
	}

	err2 := c.Bind(&user)
	if err2 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	config.DB.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"data":    user,
	})
}
