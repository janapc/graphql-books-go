package entity

type BookRepository interface {
	RegisterBook(book *Book) error
	ListBooks() ([]Book, error)
	FindBooksByAuthorId(authorID string) ([]Book, error)
	FindBookById(id string) (*Book, error)
	DeleteBook(id string) error
	UpdateBook(book *Book) error
	SearchBooksByTitle(title string) ([]Book, error)
}

type AuthorRepository interface {
	RegisterAuthor(author *Author) error
	ListAuthors() ([]Author, error)
	FindAuthorByBookId(bookId string) (*Author, error)
	FindAuthorById(id string) (*Author, error)
	SearchAuthorsByName(name string) ([]Author, error)
	UpdateAuthor(author *Author) error
	DeleteAuthor(id string) error
}
