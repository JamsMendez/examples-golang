package repository

import (
	"context"
	"database/sql"

	"go-pagination/domain"
	psql "go-pagination/postgresql"
)

type bookRepo struct {
	queries *psql.Queries
}

func (br bookRepo) Count(ctx context.Context) (int, error) {
	count, err := br.queries.CountBooks(ctx)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (br bookRepo) FindBook(ctx context.Context, params domain.FindBookParams) ([]domain.Book, error) {
	findArgs := psql.FindBooksByOffsetParams{
		Limit:  int32(params.Limit),
		Offset: int32(params.Offset),
	}

	bb, err := br.queries.FindBooksByOffset(ctx, findArgs)
	if err != nil {
		return nil, err
	}

	var books []domain.Book

	for _, b := range bb {
		book := domain.Book{
			ID:        b.ID,
			Title:     b.Title,
			Author:    b.Author,
			CreatedAt: b.CreatedAt,
			UpdatedAt: b.UpdatedAt,
		}

		books = append(books, book)
	}

	return books, nil
}

func NewBookRepository(db *sql.DB) domain.BookRepository {
	return &bookRepo{
		queries: psql.New(db),
	}
}