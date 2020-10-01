package user

import (
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/pkg/password"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//Service  interface
type Service struct {
	repo Repository
	pwd  password.Service
}

//NewService create new use case
func NewService(r Repository, pwd password.Service) *Service {
	return &Service{
		repo: r,
		pwd:  pwd,
	}
}

//CreateUser Create an user
func (s *Service) CreateUser(e *entity.User) (entity.ID, error) {
	e.ID = entity.NewID()
	e.CreatedAt = time.Now()
	pwd, err := s.pwd.Generate(e.Password)
	if err != nil {
		return e.ID, err
	}
	e.Password = pwd
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
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
