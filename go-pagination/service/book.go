package service

import (
	"context"
	"time"

	"go-pagination/domain"
)

type bookServ struct {
	repo domain.BookRepository
}

// Create implements domain.BookService.
func (bs *bookServ) Create(ctx context.Context, params domain.CreateBookParams) (*domain.Book, error) {
	now := time.Now().UTC()

	insertParams := domain.InsertBookParams{
		Title:     params.Title,
		Author:    params.Author,
		CreatedAt: now,
		UpdatedAt: now,
	}

	b, err := bs.repo.InsertBook(ctx, insertParams)
	if err != nil {
		return nil, err
	}

	book := &domain.Book{
		ID:        b.ID,
		Title:     b.Title,
		Author:    b.Author,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}

	return book, nil
}

// ListByID implements domain.BookService.
func (bs *bookServ) ListByID(ctx context.Context, params domain.ListBookParams) (*domain.PaginationCursor[domain.Book], error) {
	findParams := domain.FindBookParams{
		Offset: params.Page,
		Limit:  params.PerPage,
	}

	books, err := bs.repo.FindBookByID(ctx, findParams)
	if err != nil {
		return nil, err
	}

	var nextCursor int
	perPage := params.PerPage
	size := len(books)
	if size > params.PerPage {
		nextCursor = int(books[size-1].ID)
	} else {
		perPage = size
	}

	pagination := domain.NewPaginationCursor[domain.Book](nextCursor)

	pagination.Items = append(pagination.Items, books[:perPage]...)

	return pagination, nil
}

func (bs bookServ) List(ctx context.Context, params domain.ListBookParams) (*domain.PaginationOffset[domain.Book], error) {
	count, err := bs.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	pagination := domain.NewPaginationOffset[domain.Book](
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
