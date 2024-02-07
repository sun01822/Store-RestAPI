package types

import (
	validate "github.com/go-ozzo/ozzo-validation"
)

// Response struct || marshalled into JSON format from struct
type CartRequest struct {
	UserID   uint         `json:"user_id"`
	ProductID uint 		  `json:"product_id"`
	Quantity uint 		  `json:"quantity"`
}



func (cart CartRequest) Validate() error{
	return validate.ValidateStruct(&cart,
		validate.Field(&cart.UserID, validate.Required),
		validate.Field(&cart.ProductID, validate.Required),
		validate.Field(&cart.Quantity, validate.Required),
	)
}

type CartUpdateRequest struct {
	UserID  uint         `json:"user_id"`
	ProductID uint 		  `json:"product_id"`
	Quantity uint 		  `json:"quantity"`
}


