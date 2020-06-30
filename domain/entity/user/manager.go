package user

import (
	"strings"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/pkg/password"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//manager  interface
type manager struct {
	repo Repository
	pwd  password.UseCase
}

//NewManager create new repository
func NewManager(r Repository, pwd password.UseCase) *manager {
	return &manager{
		repo: r,
		pwd:  pwd,
	}
}

//Create an user
func (s *manager) Create(e *User) (entity.ID, error) {
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
func (s *manager) Get(id entity.ID) (*User, error) {
	return s.repo.Get(id)
}

//Search users
func (s *manager) Search(query string) ([]*User, error) {
	return s.repo.Search(strings.ToLower(query))
}

//List users
func (s *manager) List() ([]*User, error) {
	return s.repo.List()
}

//Delete an user
func (s *manager) Delete(id entity.ID) error {
	u, err := s.Get(id)
	if err != nil {
		return err
	}
	if len(u.Books) > 0 {
		return domain.ErrCannotBeDeleted
	}
	return s.repo.Delete(id)
}

//Update an user
func (s *manager) Update(e *User) error {
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
