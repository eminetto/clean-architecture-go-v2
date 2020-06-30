package user

import "github.com/eminetto/clean-architecture-go-v2/domain/entity"

//Reader interface
type Reader interface {
	Get(id entity.ID) (*User, error)
	Search(query string) ([]*User, error)
	List() ([]*User, error)
}

//Writer user writer
type Writer interface {
	Create(e *User) (entity.ID, error)
	Update(e *User) error
	Delete(id entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//Manager interface
type Manager interface {
	Repository
}
