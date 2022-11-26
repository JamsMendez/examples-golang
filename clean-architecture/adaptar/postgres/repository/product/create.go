package product

import (
	"clean-architecture/core/domain"
	"clean-architecture/core/dto"
)

func (repository repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	// ctx := context.Background{}
	product := domain.Product{}

	// process SQL DB

	var err error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
