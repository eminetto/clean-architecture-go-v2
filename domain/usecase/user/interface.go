package user

import (
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//UseCase interface
type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(e *entity.User) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
}
