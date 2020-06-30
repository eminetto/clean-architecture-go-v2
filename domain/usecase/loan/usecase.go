package loan

import (
	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/book"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity/user"
)

type usecase struct {
	uManager user.Manager
	bManager book.Manager
}

//NewUseCase create new use case
func NewUseCase(u user.Manager, b book.Manager) *usecase {
	return &usecase{
		uManager: u,
		bManager: b,
	}
}

//Borrow borrow a book to an user
func (s *usecase) Borrow(u *user.User, b *book.Book) error {
	u, err := s.uManager.Get(u.ID)
	if err != nil {
		return err
	}
	b, err = s.bManager.Get(b.ID)
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
	err = s.uManager.Update(u)
	if err != nil {
		return err
	}
	b.Quantity--
	err = s.bManager.Update(b)
	if err != nil {
		return err
	}
	return nil
}

//Return return a book
func (s *usecase) Return(b *book.Book) error {
	b, err := s.bManager.Get(b.ID)
	if err != nil {
		return err
	}

	all, err := s.uManager.List()
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
	u, err := s.uManager.Get(borrowedBy)
	if err != nil {
		return err
	}
	for i, j := range u.Books {
		if j == b.ID {
			u.Books = append(u.Books[:i], u.Books[i+1:]...)
			err = s.uManager.Update(u)
			if err != nil {
				return err
			}
			break
		}
	}
	b.Quantity++
	err = s.bManager.Update(b)
	if err != nil {
		return err
	}

	return nil
}
