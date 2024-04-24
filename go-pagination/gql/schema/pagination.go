package schema

import "github.com/graphql-go/graphql"

type Pagination[T any] struct {
	Items    []T  `json:"items"`
	NextPage *int `json:"nextPage"`
	Page     int  `json:"page"`
	Pages    int  `json:"pages"`
	Total    int  `json:"total"`
}

func NewPagination[T graphql.Type](item T) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: KeyPaginationType,
			Fields: graphql.Fields{
				KeyItems:    &graphql.Field{Type: graphql.NewNonNull(graphql.NewList(item))},
				KeyNextPage: &graphql.Field{Type: graphql.Int},
				KeyPage:     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
				KeyPages:    &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
				KeyTotal:    &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			},
		},
	)
}
