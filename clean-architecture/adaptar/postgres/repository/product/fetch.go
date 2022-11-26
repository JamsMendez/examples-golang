package product

import (
	"clean-architecture/core/domain"
	"clean-architecture/core/dto"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	// ctx := context.Background{}

	products := []domain.Product{}
	var total uint32 = 0

	// Process SQL
	var err error
	if err != nil {
		return nil, err
	}

	paginationProducts := domain.Pagination[[]domain.Product]{
		Items: products,
		Total: total,
	}

	return &paginationProducts, nil
}
