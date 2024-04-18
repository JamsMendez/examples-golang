package domain

type PaginationCursor[T any] struct {
	Items    []T
	NextPage *int
}

type PaginationCursorParams struct {
	Page    int
	PerPage int
}

func NewPaginationCursor[T any](next int) *PaginationCursor[T] {
	var nextPage *int

	if next > 0 {
		nextPage = &next
	}

	return &PaginationCursor[T]{
		NextPage: nextPage,
		Items:    []T{},
	}
}
