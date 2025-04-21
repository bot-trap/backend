package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;size:20"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Posts    []Post
}
