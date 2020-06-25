package loan

import (
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"
)

//UseCase use case interface
type UseCase interface {
	Borrow(u *user.User, b *book.Book) error
	Return(b *book.Book) error
}
