package book

import (
	"strings"
	"time"

	repo "github.com/eminetto/clean-architecture-go-v2/domain/repository/book"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//Service book usecase
type Service struct {
	repo repo.Repository
}

//NewService create new service
func NewService(r repo.Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create a book
func (s *Service) CreateBook(e *entity.Book) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

//Get a book
func (s *Service) GetBook(id entity.ID) (*entity.Book, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

//Search books
func (s *Service) SearchBooks(query string) ([]*entity.Book, error) {
	books, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, domain.ErrNotFound
	}
	return books, nil
}

//List books
func (s *Service) ListBooks() ([]*entity.Book, error) {
	books, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, domain.ErrNotFound
	}
	return books, nil
}

//Delete a book
func (s *Service) DeleteBook(id entity.ID) error {
	_, err := s.GetBook(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update a book
func (s *Service) UpdateBook(e *entity.Book) error {
	return s.repo.Update(e)
}
