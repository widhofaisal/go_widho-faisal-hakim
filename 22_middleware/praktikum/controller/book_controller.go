package controller

import (
	"alterra/golang/22_middleware/praktikum/config"
	"alterra/golang/22_middleware/praktikum/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book
	err := config.DB.Find(&books).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"data":    books,
	})
}

func GetBookController(c echo.Context) error {
	var book model.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Find(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by id",
		"data":    book,
	})

}

func CreateBookController(c echo.Context) error {
	var book model.Book
	c.Bind(&book)

	err := config.DB.Save(&book).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"data":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	var book model.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Find(&book, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	config.DB.Delete(&book)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
		"data":    book,
	})
}

func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var book model.Book

	err := config.DB.First(&book, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err2 := c.Bind(&book)
	if err2 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	config.DB.Save(&book)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"data":    book,
	})
}
