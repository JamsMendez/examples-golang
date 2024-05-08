package api

import (
	"database/sql"
	"log"
	"net/http"

	"go-pagination/repository"
	"go-pagination/service"
)

func Start(testDB *sql.DB) error {
	bookRepo := repository.NewBookRepository(testDB)
	bookService := service.NewBookService(bookRepo)

	handlerGQL, err := NewHandlerGQL(bookService)
	if err != nil {
		log.Fatalf("expected handler gql error nil, got %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/graphql/", handlerGQL.HandlerQueries)
	mux.HandleFunc("POST /api/graphql/", handlerGQL.HandlerMutations)

	return http.ListenAndServe(":3000", mux)
}
