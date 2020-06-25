package book

import "github.com/eminetto/clean-architecture-go-v2/domain/entity"

//Reader interface
type Reader interface {
	Get(id entity.ID) (*Book, error)
	Search(query string) ([]*Book, error)
	List() ([]*Book, error)
}

//Writer book writer
type Writer interface {
	Create(e *Book) (entity.ID, error)
	Update(e *Book) error
	Delete(id entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}