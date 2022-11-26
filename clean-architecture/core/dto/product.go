package dto

import (
	"encoding/json"
	"io"
)

// CreateProductRequest is an representation request body to create a new Product
type CreateProductRequest struct {
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
}

// ParseJSONToCreateProductRequest convert json body request to a CreateProductRequest struct
func ParseJSONToCreateProductRequest(body io.Reader) (*CreateProductRequest, error) {
	createProductRequest := CreateProductRequest{}

	if err := json.NewDecoder(body).Decode(&createProductRequest); err != nil {
		return nil, err
	}

	return &createProductRequest, nil
}
