package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	Zip     string `json:"zip"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
	Phone  string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}