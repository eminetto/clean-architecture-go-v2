package loan

import (
	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"github.com/eminetto/clean-architecture-go-v2/domain/usecase/book"
	"github.com/eminetto/clean-architecture-go-v2/domain/usecase/user"
)

//Service loan usecase
type Service struct {
	userService user.UseCase
	bookService book.UseCase
}

//NewService create new use case
func NewService(u user.UseCase, b book.UseCase) *Service {
	return &Service{
		userService: u,
		bookService: b,
	}
}

//Borrow borrow a book to an user
func (s *Service) Borrow(u *entity.User, b *entity.Book) error {
	u, err := s.userService.GetUser(u.ID)
	if err != nil {
		return err
	}
	b, err = s.bookService.GetBook(b.ID)
	if err != nil {
		return err
	}
	if b.Quantity <= 0 {
		return domain.ErrNotEnoughBooks
	}
	for _, v := range u.Books {
		if v == b.ID {
			return domain.ErrBookAlreadyBorrowed
		}
	}
	u.Books = append(u.Books, b.ID)
	err = s.userService.UpdateUser(u)
	if err != nil {
		return err
	}
	b.Quantity--
	err = s.bookService.UpdateBook(b)
	if err != nil {
		return err
	}
	return nil
}

//Return return a book
func (s *Service) Return(b *entity.Book) error {
	b, err := s.bookService.GetBook(b.ID)
	if err != nil {
		return err
	}

	all, err := s.userService.ListUsers()
	if err != nil {
		return err
	}
	borrowed := false
	var borrowedBy entity.ID
	for _, u := range all {
		for _, bookID := range u.Books {
			if bookID == b.ID {
				borrowed = true
				borrowedBy = u.ID
				break
			}
		}
	}
	if !borrowed {
		return domain.ErrBookNotBorrowed
	}
	u, err := s.userService.GetUser(borrowedBy)
	if err != nil {
		return err
	}
	for i, j := range u.Books {
		if j == b.ID {
			u.Books = append(u.Books[:i], u.Books[i+1:]...)
			err = s.userService.UpdateUser(u)
			if err != nil {
				return err
			}
			break
		}
	}
	b.Quantity++
	err = s.bookService.UpdateBook(b)
	if err != nil {
		return err
	}

	return nil
}
