package gql

import (
	"go-pagination/domain"
	"go-pagination/gql/mutation"
	"go-pagination/gql/query"

	"github.com/graphql-go/graphql"
)

func NewSchema(bookServ domain.BookService) (graphql.Schema, error) {
	config := graphql.SchemaConfig{
		Query:    query.NewRootQuery(bookServ),
		Mutation: mutation.NewRootMutation(),
	}

	return graphql.NewSchema(config)
}
