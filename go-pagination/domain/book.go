package domain

import (
	"context"
	"time"
)

type Book struct {
	ID        int64
	Title     string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FindBookParams struct {
	Offset int
	Limit  int
}

type InsertBookParams struct {
	Title     string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookRepository interface {
	Count(ctx context.Context) (int, error)
	FindBook(ctx context.Context, params FindBookParams) ([]Book, error)
	FindBookByID(ctx context.Context, params FindBookParams) ([]Book, error)
	InsertBook(ctx context.Context, params InsertBookParams) (*Book, error)
}

type ListBookParams struct {
	Page    int
	PerPage int
}

type CreateBookParams struct {
	Title  string
	Author string
}

type BookService interface {
	List(ctx context.Context, params ListBookParams) (*PaginationOffset[Book], error)
	ListByID(ctx context.Context, params ListBookParams) (*PaginationCursor[Book], error)
	Create(ctx context.Context, params CreateBookParams) (*Book, error)
}
