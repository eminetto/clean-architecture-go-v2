package user

import (
	"strings"
	"time"

	repo "github.com/eminetto/clean-architecture-go-v2/domain/repository/user"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/pkg/password"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//Service  interface
type Service struct {
	repo repo.Repository
	pwd  password.Service
}

//NewService create new use case
func NewService(r repo.Repository, pwd password.Service) *Service {
	return &Service{
		repo: r,
		pwd:  pwd,
	}
}

//Create an user
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

//Get an user
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

//Search users
func (s *Service) SearchUsers(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

//List users
func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}

//Delete an user
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

//Update an user
func (s *Service) UpdateUser(e *entity.User) error {
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
