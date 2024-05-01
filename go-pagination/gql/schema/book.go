package schema

import "github.com/graphql-go/graphql"

type Book struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var BookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: KeyBookType,
		Fields: graphql.Fields{
			KeyID:        &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			KeyTitle:     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			KeyAuthor:    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			KeyCreatedAt: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			KeyUpdatedAt: &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)

var InputBookType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: KeyInputBookType,
		Fields: graphql.InputObjectConfigFieldMap{
			KeyTitle:  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			KeyAuthor: &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)
