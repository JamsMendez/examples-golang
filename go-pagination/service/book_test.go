package service

import (
	"context"
	"testing"
	"time"

	"go-pagination/domain"
	"go-pagination/repository"
)

func TestGetBooksByPagination(t *testing.T) {
	bookRepo := repository.NewBookRepository(testDB)
	bookServ := NewBookService(bookRepo)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	params := domain.ListBookParams{
		Page:    2,
		PerPage: 5,
	}

	pagination, err := bookServ.List(ctx, params)
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
