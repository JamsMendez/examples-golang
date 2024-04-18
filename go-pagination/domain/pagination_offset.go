package domain

const (
	DefaultPage    = 1
	DefaultPerPage = 10

	KeyPage    = "page"
	KeyPerPage = "per_page"

	MaxPerPage = 1000
)

type PaginationOffset[T any] struct {
	Items    []T
	NextPage *int
	Page     int
	Pages    int
	Total    int
	size     int
	skip     int
}

func (p *PaginationOffset[T]) Skip() int {
	return p.skip
}

func (p *PaginationOffset[T]) Size() int {
	return p.size
}

type PaginationOffsetParams struct {
	Page    int
	PerPage int
}

func NewPaginationOffset[T any](page, size, total int) *PaginationOffset[T] {
	var nextPage *int

	pages := total / size
	mod := total % size
	if mod > 0 {
		pages++
	}

	if page < pages {
		next := page + 1
		nextPage = &next
	}

	skip := size * (page - 1)

	return &PaginationOffset[T]{
		skip:     skip,
		size:     size,
		NextPage: nextPage,
		Page:     page,
		Pages:    pages,
		Total:    total,
		Items:    []T{},
	}
}
