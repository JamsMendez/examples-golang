package query

import (
	"go-pagination/domain"
	"go-pagination/gql/resolve"
	"go-pagination/gql/schema"

	"github.com/graphql-go/graphql"
)

func NewRootQuery(bookServ domain.BookService) *graphql.Object {
	resolveBook := resolve.NewBookResolve(bookServ)

	config := graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			schema.KeyQueryBooks: QueryBooks(resolveBook.FindBooks),
		},
	}

	return graphql.NewObject(config)
}
