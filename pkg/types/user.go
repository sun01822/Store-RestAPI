package types

import (
	validate "github.com/go-ozzo/ozzo-validation"
)

// Response struct || marshalled into JSON format from struct
type UserRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password"`
	NameID    uint   `json:"name_id,omitempty"`
	AddressID uint   `json:"address_id,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

func (user UserRequest) Validate() error{
	return validate.ValidateStruct(&user,
		validate.Field(&user.Email, validate.Required, validate.Length(1, 50)),
		validate.Field(&user.Password, validate.Required, validate.Length(1, 50)),
	)
}


// Response struct || marshalled into JSON format from struct
type UserAddressRequest struct {
	City        string     `json:"city"`
	Street      string     `json:"street"`
	Number      string     `json:"number"`
	Zip         string     `json:"zip"`
}

func (address UserAddressRequest) Validate() error{
	return validate.ValidateStruct(&address,
		validate.Field(&address.City, validate.Required, validate.Length(1, 50)),
	)
}

type UserGeoLocationRequest struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func (geoLocation UserGeoLocationRequest) Validate() error{
	return validate.ValidateStruct(&geoLocation,
		validate.Field(&geoLocation.Lat, validate.Required, validate.Length(1, 50)),
		validate.Field(&geoLocation.Lng, validate.Required, validate.Length(1, 50)),
	)
}

type NameRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (name NameRequest) Validate() error{
	return validate.ValidateStruct(&name,
		validate.Field(&name.FirstName, validate.Required, validate.Length(1, 50)),
		validate.Field(&name.LastName, validate.Required, validate.Length(1, 50)),
	)
}




