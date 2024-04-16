package service

import (
	"context"
	"go-pagination/domain"
)

type bookServ struct {
	repo domain.BookRepository
}

func (bs bookServ) List(ctx context.Context, params domain.ListBookParams) (*domain.Pagination[domain.Book], error) {

	count, err := bs.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	pagination := domain.NewPagination[domain.Book](
		params.Page,
		params.PerPage,
		count,
	)

	findParams := domain.FindBookParams{
		Offset: pagination.Skip(),
		Limit:  pagination.Size(),
	}

	books, err := bs.repo.FindBook(ctx, findParams)
	if err != nil {
		return nil, err
	}

	pagination.Items = append(pagination.Items, books...)

	return pagination, nil
}

func NewBookService(repo domain.BookRepository) domain.BookService {
	return &bookServ{repo: repo}
}
