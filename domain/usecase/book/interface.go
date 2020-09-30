package book

import (
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//UseCase interface
type UseCase interface {
	GetBook(id entity.ID) (*entity.Book, error)
	SearchBooks(query string) ([]*entity.Book, error)
	ListBooks() ([]*entity.Book, error)
	CreateBook(e *entity.Book) (entity.ID, error)
	UpdateBook(e *entity.Book) error
	DeleteBook(id entity.ID) error
}
