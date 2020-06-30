package book

import (
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

type manager struct {
	repo Repository
}

//NewManager create new manager
func NewManager(r Repository) *manager {
	return &manager{
		repo: r,
	}
}

//Create a book
func (s *manager) Create(e *Book) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

//Get a book
func (s *manager) Get(id entity.ID) (*Book, error) {
	return s.repo.Get(id)
}

//Search books
func (s *manager) Search(query string) ([]*Book, error) {
	return s.repo.Search(strings.ToLower(query))
}

//List books
func (s *manager) List() ([]*Book, error) {
	return s.repo.List()
}

//Delete a book
func (s *manager) Delete(id entity.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update a book
func (s *manager) Update(e *Book) error {
	return s.repo.Update(e)
}
