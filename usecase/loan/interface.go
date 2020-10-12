package loan

import (
	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//UseCase use case interface
type UseCase interface {
	Borrow(u *entity.User, b *entity.Book) error
	Return(b *entity.Book) error
}
