package resolve

import "go-pagination/domain"

type bookResolve struct {
	bookService domain.BookService
}

func NewBookResolve(bookServ domain.BookService) *bookResolve {
	return &bookResolve{
		bookService: bookServ,
	}
}
