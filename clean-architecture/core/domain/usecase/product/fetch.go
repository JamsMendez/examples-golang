package product

import (
	"clean-architecture/core/domain"
	"clean-architecture/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
  products, err := usecase.repository.Fetch(paginationRequest)

  if err != nil {
    return nil, err
  }

  return products, err
}
