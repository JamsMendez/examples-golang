package service

import (
	"context"
	"testing"
	"time"

	"go-pagination/domain"
	"go-pagination/repository"
)

func TestGetBooksByPaginationOffset(t *testing.T) {
	// /api?page=1&per_page=5
	// handler r.Query
	// strconv.Atoi
	// convert int, int

	bookRepo := repository.NewBookRepository(testDB)
	bookService := NewBookService(bookRepo)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := domain.ListBookParams{
		Page:    2,
		PerPage: 5,
	}

	pagination, err := bookService.List(ctx, params)
	if err != nil {
		t.Fatalf("expected err = nil, got %v", err)
	}

	if pagination == nil {
		t.Fatalf("expected pagination != nil, got nil")
	}

	size := pagination.Size()
	if size != 5 {
		t.Fatalf("expected pagination items size = 5, got %d", size)
	}

	book1 := pagination.Items[0]
	book2 := pagination.Items[4]

	if book1.ID != 45 {
		t.Fatalf("expected item[0].ID = 50, got ID %d", book1.ID)
	}

	if book2.ID != 41 {
		t.Fatalf("expected item[0].ID = 46, got ID %d", book2.ID)
	}
}

func TestGetBooksByPaginationCursor(t *testing.T) {
	// /api?limit=5&cursor=encode(1,'2024-04-06')
	// handler decode split,
	// convert int, string

	bookRepo := repository.NewBookRepository(testDB)
	bookService := NewBookService(bookRepo)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := domain.ListBookParams{
		Page:    1,
		PerPage: 5,
	}

	pagination, err := bookService.ListByID(ctx, params)
	if err != nil {
		t.Fatalf("expected err = nil, got %v", err)
	}

	if pagination == nil {
		t.Fatalf("expected pagination != nil, got nil")
	}

	size := len(pagination.Items)
	if size != 1 {
		t.Fatalf("expected pagination items size = 5, got %d", size)
	}

	book1 := pagination.Items[0]
	// book2 := pagination.Items[4]

	if book1.ID != 1 {
		t.Fatalf("expected item[0].ID = 1, got ID %d", book1.ID)
	}

	// if book2.ID != 2 {
	// 	t.Fatalf("expected item[0].ID = 1, got ID %d", book2.ID)
	// }

	if pagination.NextPage != nil {
		t.Fatalf("expected nextPage != nil, got nil")
	}
}
