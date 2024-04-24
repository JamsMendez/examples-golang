package mutation

import (
	"go-pagination/gql/schema"

	"github.com/graphql-go/graphql"
)

func NewRootMutation() *graphql.Object {
	config := graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			schema.KeyEmpty: &graphql.Field{
				Type: graphql.String,
				Resolve: func(params graphql.ResolveParams) (any, error) {
					return "...", nil
				},
			},
		},
	}

	return graphql.NewObject(config)
}
