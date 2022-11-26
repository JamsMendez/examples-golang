package product

import (
	"clean-architecture/core/domain"
	"clean-architecture/core/dto"
)

func (usecase usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
  product, err := usecase.repository.Create(productRequest)

  if err != nil {
    return nil, err
  }

  return product, nil
}
