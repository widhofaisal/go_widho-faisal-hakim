package controller

import (
	"alterra/golang/22_middleware/praktikum/config"
	"alterra/golang/22_middleware/praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateBlogController(c echo.Context) error {
	var blog model.Blog
	c.Bind(&blog)

	err := config.DB.Save(&blog).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create blog",
		"data":    blog,
	})
}

func GetBlogsController(c echo.Context) error {
	var blogs []model.Blog

	err := config.DB.Find(&blogs).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    blogs,
	})
}

func GetBlogController(c echo.Context) error {
	var blog model.Blog
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.First(&blog, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get blog, id not found",
			"id":      id,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    blog,
	})
}

func DeleteBlogController(c echo.Context) error {
	var blog model.Blog
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Find(&blog, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "failed delete blog, id not found",
			"id":      id,
		})
	}

	config.DB.Delete(&blog)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog by id",
		"data":    blog,
	})
}

func UpdateBlogController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var blog model.Blog
	err := config.DB.First(&blog, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "failed update blog, id not found",
			"id":      id,
		})
	}

	err2 := c.Bind(&blog)
	if err2 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	config.DB.Save(&blog)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog by id",
		"data":    blog,
	})
}
