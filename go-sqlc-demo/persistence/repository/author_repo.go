package repository

import (
	"context"
	"database/sql"

	"go-sqlc-demo/domain"
	"go-sqlc-demo/persistence/postgres"
)

type authorRepo struct {
	*postgres.Queries
}

// Create implements domain.AuthorRepository.
func (a *authorRepo) Create(ctx context.Context, params domain.CreateAuthorParams) (*domain.Author, error) {
	createParams := postgres.CreateAuthorParams{
		Name: params.Name,
		Bio: sql.NullString{
			String: params.Bio,
			Valid:  true,
		},
	}

	au, err := a.CreateAuthor(ctx, createParams)
	if err != nil {
		return nil, err
	}

	var bio string
	if au.Bio.Valid {
		bio = au.Bio.String
	}

	author := &domain.Author{
		ID:   int(au.ID),
		Name: au.Name,
		Bio:  bio,
	}

	return author, nil
}

// Delete implements domain.AuthorRepository.
func (a *authorRepo) Delete(ctx context.Context, id int) error {
	return a.DeleteAuthor(ctx, int64(id))
}

// FindAll implements domain.AuthorRepository.
func (a *authorRepo) FindAll(ctx context.Context) ([]domain.Author, error) {
	aa, err := a.ListAuthors(ctx)
	if err != nil {
		return nil, err
	}

	authors := make([]domain.Author, len(aa))

	for i, a := range aa {
		var bio string
		if a.Bio.Valid {
			bio = a.Bio.String
		}

		authors[i] = domain.Author{
			ID:   int(a.ID),
			Name: a.Name,
			Bio:  bio,
		}
	}

	return authors, nil
}

// FindOne implements domain.AuthorRepository.
func (a *authorRepo) FindOne(ctx context.Context, id int) (*domain.Author, error) {
	au, err := a.GetAuthor(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	var bio string
	if au.Bio.Valid {
		bio = au.Bio.String
	}

	author := domain.Author{
		ID:   int(au.ID),
		Name: au.Name,
		Bio:  bio,
	}

	return &author, nil
}

// Update implements domain.AuthorRepository.
func (a *authorRepo) Update(ctx context.Context, params domain.UpdateAuthorParams) (*domain.Author, error) {
	updateParams := postgres.UpdateAuthorParams{
		ID:   int64(params.ID),
		Name: params.Name,
		Bio: sql.NullString{
			String: params.Bio,
			Valid:  true,
		},
	}

	au, err := a.UpdateAuthor(ctx, updateParams)
	if err != nil {
		return nil, err
	}

	var bio string
	if au.Bio.Valid {
		bio = au.Bio.String
	}

	author := domain.Author{
		ID:   int(au.ID),
		Name: au.Name,
		Bio:  bio,
	}

	return &author, nil
}

func NewAuthorRepo(db *sql.DB) domain.AuthorRepository {
	return &authorRepo{
		Queries: postgres.New(db),
	}
}
