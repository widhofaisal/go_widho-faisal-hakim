package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	gorm.Model
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  int    `json:"user_id" form:"user_id"`
	User    User
}

type Book struct {
	gorm.Model
	ID     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Year   int    `json:"year" form:"year"`
}

// model struct dibawah ini hanya menjadikan FK tanpa menerapkan behaviornya
// jangan di pake
// type Blog struct {
// 	gorm.Model
// 	Title   string `json:"title" form:"title"`
// 	Content string `json:"content" form:"content"`
// 	UserID  int    `json:"user_id" form:"user_id" gorm:"unique; not null; foreignKey:UserID; refernces:ID" `
// }
