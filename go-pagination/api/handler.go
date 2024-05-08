package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go-pagination/domain"
	"go-pagination/gql"

	"github.com/graphql-go/graphql"
)

const (
	KeyQuery = "query"
)

type BodyGQL struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
	Operation string         `json:"operation"`
}

type HandlerGQL struct {
	schema graphql.Schema
}

func NewHandlerGQL(bookService domain.BookService) (*HandlerGQL, error) {
	schema, err := gql.NewSchema(bookService)
	if err != nil {
		return nil, err
	}

	h := &HandlerGQL{
		schema: schema,
	}

	return h, nil
}

func (h *HandlerGQL) HandlerQueries(w http.ResponseWriter, r *http.Request) {
	var variables map[string]any
	var operation string
	var rootObject map[string]any

	q := r.URL.Query()
	query := q.Get(KeyQuery)

	params := graphql.Params{
		Schema:        h.schema,
		RequestString: query,
		// mutations
		VariableValues: variables,
		OperationName:  operation,
		Context:        r.Context(),
		RootObject:     rootObject,
	}

	result := graphql.Do(params)

	var buffer bytes.Buffer

	err := json.NewEncoder(&buffer).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(buffer.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *HandlerGQL) HandlerMutations(w http.ResponseWriter, r *http.Request) {
	var bGQL BodyGQL

	err := json.NewDecoder(r.Body).Decode(&bGQL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		err = r.Body.Close()
		if err != nil {
			log.Println("request body close error: ", err)
		}
	}()

	var rootObject map[string]any

	params := graphql.Params{
		Schema:        h.schema,
		RequestString: bGQL.Query,
		// mutations
		VariableValues: bGQL.Variables,
		OperationName:  bGQL.Operation,
		Context:        r.Context(),
		RootObject:     rootObject,
	}

	result := graphql.Do(params)

	var buffer bytes.Buffer

	err = json.NewEncoder(&buffer).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	_, err = io.Copy(w, &buffer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
