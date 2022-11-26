package product

import "clean-architecture/core/domain"

type usecase struct {
	repository domain.ProductRepository
}

func New(repository domain.ProductRepository) domain.ProductUseCase {
	return &usecase{
		repository: repository,
	}
}
