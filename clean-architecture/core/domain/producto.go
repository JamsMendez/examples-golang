package domain

import (
	"clean-architecture/core/dto"
	"net/http"
)

// Product is entity of table product database column
type Product struct {
	ID          uint64
	Name        string
	Price       float32
	Description string
}

// ProductService is a contract of http adapter layer
type ProductService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// ProductUseCase is a contract of business rule layer
type ProductUseCase interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRequest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}

// ProductRepository is contract of database connection adapter layer
type ProductRepository interface {
	Create(productRequest *dto.CreateProductRequest) (*Product, error)
	Fetch(paginationRquest *dto.PaginationRequestParams) (*Pagination[[]Product], error)
}
