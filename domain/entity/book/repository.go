package book

import (
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

type repository struct {
	repo Repository
}

//NewRepository create new repository
func NewRepository(r Repository) *repository {
	return &repository{
		repo: r,
	}
}

//Create a book
func (s *repository) Create(e *Book) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

//Get a book
func (s *repository) Get(id entity.ID) (*Book, error) {
	return s.repo.Get(id)
}

//Search books
func (s *repository) Search(query string) ([]*Book, error) {
	return s.repo.Search(strings.ToLower(query))
}

//List books
func (s *repository) List() ([]*Book, error) {
	return s.repo.List()
}

//Delete a book
func (s *repository) Delete(id entity.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update a book
func (s *repository) Update(e *Book) error {
	return s.repo.Update(e)
}
