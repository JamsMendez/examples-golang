package gql

import (
	"context"
	"testing"
	"time"

	"go-pagination/repository"
	"go-pagination/service"

	"github.com/graphql-go/graphql"
)

func TestGetBooksPaginationOffset(t *testing.T) {
	bookRepo := repository.NewBookRepository(testDB)
	bookService := service.NewBookService(bookRepo)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	schema, err := NewSchema(bookService)
	if err != nil {
		t.Fatalf("expected error = nil, got %v", err)
	}

	query := `
		query {
			books(page: 1, perPage: 5) {
				nextPage
				page
				pages
				items {
					id
					title
				}
			}
		}
	`

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
		// mutations
		VariableValues: map[string]interface{}{},
		OperationName:  "",
		Context:        ctx,
		// values extras - ID, cookie, jwt, sessions
		RootObject: map[string]interface{}{},
	}

	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		t.Fatalf("expected error len 0, got %v", result.Errors)
	}

	if result.Data == nil {
		t.Fatalf("expected data != nil, got nil")
	}

	// if result.Data != nil {
	// 	b, err := json.Marshal(result.Data)
	// 	if err != nil {
	// 		t.Fatalf("expected marshalling error = nil, got %v", err)
	// 	}
	// 	t.Fatalf("expected data != nil, got %v", string(b))
	// }
}
