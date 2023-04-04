package main

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var DB *sql.DB

func init() {
	initDB()
}

func main() {
	e := echo.New()

	e.GET("/users", GetUserHandler)

	e.Start(":10000")
}

func initDB() {
	var err error
	DB, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/alterra")
	if err != nil {
		panic(err)
	}
}

func GetUserHandler(e echo.Context) error {
	ctx, cancel := context.WithTimeout((context.Background()), 2*time.Second)
	defer cancel()

	query := "SELECT id, username FROM users;"
	rows, err := DB.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		user := User{}

		err := rows.Scan(&user.Id, &user.Username)
		if err != nil {
			return err
		}

		users = append(users, user)
	}

	return e.JSON(http.StatusOK, users)
}
