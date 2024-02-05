package types


import (
	validate "github.com/go-ozzo/ozzo-validation"
)

// Response struct || marshalled into JSON format from struct
type ProductRequest struct {
	Title string `json:"title"`
	Price float64 `json:"price"`
	Category string `json:"category"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
}

func (product ProductRequest) Validate() error{
	return validate.ValidateStruct(&product,
		validate.Field(&product.Title, validate.Required, validate.Length(3, 50)),
		validate.Field(&product.Price, validate.Required),
		validate.Field(&product.Category, validate.Required),
	)
}