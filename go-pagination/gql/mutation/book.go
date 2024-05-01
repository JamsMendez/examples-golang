package mutation

import (
	"go-pagination/gql/schema"

	"github.com/graphql-go/graphql"
)

func MutationCreateBook(resolve graphql.FieldResolveFn) *graphql.Field {
	return &graphql.Field{
		Type: schema.BookType,
		Args: graphql.FieldConfigArgument{
			schema.KeyInput: &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(schema.InputBookType),
			},
		},
		Resolve: resolve,
	}
}
