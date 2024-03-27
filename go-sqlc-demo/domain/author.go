package domain

import "context"

type Author struct {
	ID   int
	Name string
	Bio  string
}

type CreateAuthorParams struct {
	Name string
	Bio  string
}

type UpdateAuthorParams struct {
	ID   int
	Name string
	Bio  string
}

type AuthorRepository interface {
	FindAll(ctx context.Context) ([]Author, error)
	FindOne(ctx context.Context, id int) (*Author, error)
	Create(ctx context.Context, params CreateAuthorParams) (*Author, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, params UpdateAuthorParams) (*Author, error)
}
