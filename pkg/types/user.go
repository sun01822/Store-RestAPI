package types

import (
	validate "github.com/go-ozzo/ozzo-validation"
)

// Response struct || marshalled into JSON format from struct
type UserRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password"`
	Phone     string `json:"phone,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	City      string `json:"city,omitempty"`
	Street    string `json:"street,omitempty"`
	Number    string `json:"number,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Lat       string `json:"lat,omitempty"`
	Lng       string `json:"lng,omitempty"`
}

func (user UserRequest) Validate() error{
	return validate.ValidateStruct(&user,
		validate.Field(&user.Email, validate.Required, validate.Length(12, 50)),
		validate.Field(&user.Password, validate.Required, validate.Length(6, 20)),
	)
}
