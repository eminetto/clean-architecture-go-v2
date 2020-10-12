package book

import (
	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Book, error)
	Search(query string) ([]*entity.Book, error)
	List() ([]*entity.Book, error)
}

//Writer book writer
type Writer interface {
	Create(e *entity.Book) (entity.ID, error)
	Update(e *entity.Book) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetBook(id entity.ID) (*entity.Book, error)
	SearchBooks(query string) ([]*entity.Book, error)
	ListBooks() ([]*entity.Book, error)
	CreateBook(title string, author string, pages int, quantity int) (entity.ID, error)
	UpdateBook(e *entity.Book) error
	DeleteBook(id entity.ID) error
}
