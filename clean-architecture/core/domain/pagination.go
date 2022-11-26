package domain

type Pagination[T any] struct {
	Items T      `json:"items"`
	Total uint32 `json:"total"`
}
