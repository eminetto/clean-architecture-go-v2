package book

import (
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//Service service interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create a book
func (s *Service) Create(e *Book) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

//Get a book
func (s *Service) Get(id entity.ID) (*Book, error) {
	return s.repo.Get(id)
}

//Search books
func (s *Service) Search(query string) ([]*Book, error) {
	return s.repo.Search(strings.ToLower(query))
}

//List books
func (s *Service) List() ([]*Book, error) {
	return s.repo.List()
}

//Delete a book
func (s *Service) Delete(id entity.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update a book
func (s *Service) Update(e *Book) error {
	return s.repo.Update(e)
}
