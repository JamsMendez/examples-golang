package mutation

import (
	"go-pagination/domain"
	"go-pagination/gql/resolve"
	"go-pagination/gql/schema"

	"github.com/graphql-go/graphql"
)

func NewRootMutation(bookServ domain.BookService) *graphql.Object {
	resolveBook := resolve.NewBookResolve(bookServ)

	config := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			schema.KeyMutationCreateBook: MutationCreateBook(resolveBook.CreateBook),
		},
	}

	return graphql.NewObject(config)
}
