package query

import (
	"go-pagination/domain"
	"go-pagination/gql/schema"

	"github.com/graphql-go/graphql"
)

func QueryBooks(resolve graphql.FieldResolveFn) *graphql.Field {
	pagination := schema.NewPagination(schema.BookType)

	return &graphql.Field{
		Type: pagination,
		Args: graphql.FieldConfigArgument{
			schema.KeyPage: &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: domain.DefaultPage,
			},
			schema.KeyPerPage: &graphql.ArgumentConfig{
				Type:         graphql.Int,
				DefaultValue: domain.DefaultPerPage,
			},
		},
		Resolve: resolve,
	}
}
