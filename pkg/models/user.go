package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string  `json:"email"`
	Username  string  `json:"username"`
	Password  string  `json:"password"`
	NameID    uint    `json:"name_id"`
	AddressID uint    `json:"address_id"`
	Phone     string  `json:"phone"`
}

type Name struct {
	ID        	uint   `json:"name_id"`
	FirstName 	string `json:"first_name"`
	LastName  	string `json:"last_name"`
}

type Address struct {
	ID          uint       `gorm:"primaryKey;auto_increment"`
	City        string     `json:"city"`
	Street      string     `json:"street"`
	Number      string     `json:"number"`
	Zip         string     `json:"zip"`
	GeoLocationID uint     `json:"geo_location_id"`
}

type GeoLocation struct {
	ID  uint   `gorm:"primaryKey;auto_increment"`
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
