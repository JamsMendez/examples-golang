package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-pagination/repository"
	"go-pagination/service"
)

func TestHandlerGQL(t *testing.T) {
	bookRepo := repository.NewBookRepository(testDB)
	bookService := service.NewBookService(bookRepo)

	handlerGQL, err := NewHandlerGQL(bookService)
	if err != nil {
		t.Fatalf("expected handler gql error nil, got %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/graphql/", handlerGQL.HandlerQueries)
	mux.HandleFunc("POST /api/graphql/", handlerGQL.HandlerMutations)

	server := httptest.NewServer(mux)

	const endpoint = "/api/graphql/"

	client := NewClient(server.URL, "")

	t.Run("test query book pagination", func(t *testing.T) {
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

		queries := map[string]string{
			KeyQuery: query,
		}

		var resp *http.Response
		resp, err = client.makeRequest(
			http.MethodGet,
			endpoint,
			queries,
			nil,
		)
		if err != nil {
			t.Fatalf("expected http request error nil, got %v", err)
		}

		var buffer bytes.Buffer
		_, err = io.Copy(&buffer, resp.Body)
		if err != nil {
			t.Fatalf("expected response read error nil, got %v", err)
		}

		defer func() {
			err = resp.Body.Close()
			if err != nil {
				log.Println("response body close error: ", err)
			}
		}()

		s := buffer.String()
		if s != "" {
			log.Fatalf("response body: %s", s)
		}
	})

	t.Run("test mutation create book", func(t *testing.T) {
		query := `
			mutation CreateBook($book: InputBook!) {
				createBook(input: $book) {
					id
					title
					author
					createdAt
					updatedAt
				}
			}
		`

		book := map[string]any{
			"title":  "Go development",
			"author": "JamsMendez",
		}

		variables := map[string]any{
			"book": book,
		}

		bodyGQL := BodyGQL{
			Query:     query,
			Variables: variables,
			Operation: "CreateBook",
		}

		var body bytes.Buffer

		err = json.NewEncoder(&body).Encode(bodyGQL)
		if err != nil {
			t.Fatalf("expected json encode error nil, got %v", err)
		}

		var queries map[string]string

		var resp *http.Response
		resp, err = client.makeRequest(
			http.MethodPost,
			endpoint,
			queries,
			&body,
		)
		if err != nil {
			t.Fatalf("expected http request error nil, got %v", err)
		}

		var buffer bytes.Buffer
		_, err = io.Copy(&buffer, resp.Body)
		if err != nil {
			t.Fatalf("expected response read error nil, got %v", err)
		}

		defer func() {
			err = resp.Body.Close()
			if err != nil {
				log.Println("response body close error: ", err)
			}
		}()

		s := buffer.String()
		if s != "" {
			log.Fatalf("response bodyGQL: %s", s)
		}
	})
}
