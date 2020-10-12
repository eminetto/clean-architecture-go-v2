package book

import (
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateBook create a book
func (s *Service) CreateBook(title string, author string, pages int, quantity int) (entity.ID, error) {
	b, err := entity.NewBook(title, author, pages, quantity)
	if err != nil {
		return b.ID, err
	}
	return s.repo.Create(b)
}

//GetBook get a book
func (s *Service) GetBook(id entity.ID) (*entity.Book, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SearchBooks search books
func (s *Service) SearchBooks(query string) ([]*entity.Book, error) {
	books, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, entity.ErrNotFound
	}
	return books, nil
}

//ListBooks list books
func (s *Service) ListBooks() ([]*entity.Book, error) {
	books, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, entity.ErrNotFound
	}
	return books, nil
}

//DeleteBook Delete a book
func (s *Service) DeleteBook(id entity.ID) error {
	_, err := s.GetBook(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//UpdateBook Update a book
func (s *Service) UpdateBook(e *entity.Book) error {
	err := e.Validate()
	if err != nil {
		return err
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
