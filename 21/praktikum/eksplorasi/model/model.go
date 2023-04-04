package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  int    `json:"user_id" form:"user_id"`
	User    User   
}

// type Blog struct {
// 	gorm.Model
// 	Title   string `json:"title" form:"title"`
// 	Content string `json:"content" form:"content"`
// 	UserID  int    `json:"user_id" form:"user_id" gorm:"unique; not null; foreignKey:UserID; refernces:ID" `
// }
