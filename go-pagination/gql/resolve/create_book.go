package resolve

import (
	"fmt"
	"time"

	"go-pagination/domain"
	"go-pagination/gql/schema"

	"github.com/graphql-go/graphql"
)

func (br *bookResolve) CreateBook(params graphql.ResolveParams) (any, error) {
	value, ok := params.Args[schema.KeyInput]
	if !ok {
		return nil, fmt.Errorf("param input not found, is required")
	}

	oJSON, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("param input is not json")
	}

	var title, author string

	valTitle, ok := oJSON[schema.KeyTitle]
	if !ok {
		return nil, fmt.Errorf("param input field title is required")
	}

	valAuthor, ok := oJSON[schema.KeyAuthor]
	if !ok {
		return nil, fmt.Errorf("param input field author is required")
	}

	title, ok = valTitle.(string)
	if !ok {
		return nil, fmt.Errorf("param input field title is not string")
	}

	author, ok = valAuthor.(string)
	if !ok {
		return nil, fmt.Errorf("param input field author is not string")
	}

	createParams := domain.CreateBookParams{
		Title:  title,
		Author: author,
	}

	b, err := br.bookService.Create(params.Context, createParams)
	if err != nil {
		return nil, err
	}

	book := schema.Book{
		ID:        b.ID,
		Title:     b.Title,
		Author:    b.Author,
		CreatedAt: b.CreatedAt.Format(time.DateTime),
		UpdatedAt: b.UpdatedAt.Format(time.DateTime),
	}

	return book, nil
}
