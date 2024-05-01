package resolve

import (
	"time"

	"go-pagination/domain"
	"go-pagination/gql/schema"

	"github.com/graphql-go/graphql"
)

func (br *bookResolve) FindBooks(params graphql.ResolveParams) (any, error) {
	var page, perPage int

	vPage, ok := params.Args[schema.KeyPage]
	if ok {
		page = vPage.(int)
	}

	vPerPage, ok := params.Args[schema.KeyPerPage]
	if ok {
		perPage = vPerPage.(int)
	}

	if page <= 0 {
		page = domain.DefaultPage
	}

	if perPage <= 0 {
		perPage = domain.DefaultPerPage
	}

	listParams := domain.ListBookParams{
		Page:    page,
		PerPage: perPage,
	}

	p, err := br.bookService.List(params.Context, listParams)
	if err != nil {
		return nil, err
	}

	books := make([]schema.Book, 0, len(p.Items))

	for _, item := range p.Items {
		b := schema.Book{
			ID:        item.ID,
			Title:     item.Title,
			Author:    item.Author,
			CreatedAt: item.CreatedAt.Format(time.DateTime),
			UpdatedAt: item.UpdatedAt.Format(time.DateTime),
		}

		books = append(books, b)
	}

	pagination := schema.Pagination[schema.Book]{
		Items:    books,
		Page:     p.Page,
		Pages:    p.Pages,
		Total:    p.Total,
		NextPage: p.NextPage,
	}

	return pagination, nil
}
