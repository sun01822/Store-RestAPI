package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     Name   `json:"name" gorm:"embedded"`
	Address  Address `json:"address" gorm:"embedded"`
	Phone    string `json:"phone"`
}

type Name struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Address struct {
	City        string     `json:"city"`
	Street      string     `json:"street"`
	Number      string     `json:"number"`
	Zip         string     `json:"zip"`
	GeoLocation GeoLocation `json:"geo_location" gorm:"embedded"`
}

type GeoLocation struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}