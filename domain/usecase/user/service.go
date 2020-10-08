package user

import (
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateUser Create an user
func (s *Service) CreateUser(email, password, firstName, lastName string) (entity.ID, error) {
	e, err := entity.NewUser(email, password, firstName, lastName)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetUser Get an user
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

//SearchUsers Search users
func (s *Service) SearchUsers(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

//ListUsers List users
func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}

//DeleteUser Delete an user
func (s *Service) DeleteUser(id entity.ID) error {
	u, err := s.GetUser(id)
	if u == nil {
		return domain.ErrNotFound
	}
	if err != nil {
		return err
	}
	if len(u.Books) > 0 {
		return domain.ErrCannotBeDeleted
	}
	return s.repo.Delete(id)
}

//UpdateUser Update an user
func (s *Service) UpdateUser(e *entity.User) error {
	err := e.Validate()
	if err != nil {
		return domain.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
