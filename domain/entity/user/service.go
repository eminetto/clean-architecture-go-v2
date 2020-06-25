package user

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

//Create an user
func (s *Service) Create(e *User) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	return s.repo.Create(e)
}

//Get an user
func (s *Service) Get(id entity.ID) (*User, error) {
	return s.repo.Get(id)
}

//Search users
func (s *Service) Search(query string) ([]*User, error) {
	return s.repo.Search(strings.ToLower(query))
}

//List users
func (s *Service) List() ([]*User, error) {
	return s.repo.List()
}

//Delete an user
func (s *Service) Delete(id entity.ID) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Update an user
func (s *Service) Update(e *User) error {
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
