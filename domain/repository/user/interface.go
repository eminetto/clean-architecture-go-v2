package user

import (
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List() ([]*entity.User, error)
}

//Writer user writer
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}
