package repository

import (
	"fmt"
	"strings"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//UserInmem in memory repo
type UserInmem struct {
	m map[entity.ID]*entity.User
}

//NewUserInmem create new repository
func NewUserInmem() *UserInmem {
	var m = map[entity.ID]*entity.User{}
	return &UserInmem{
		m: m,
	}
}

//Create an user
func (r *UserInmem) Create(e *entity.User) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *UserInmem) Get(id entity.ID) (*entity.User, error) {
	if r.m[id] == nil {
		// return nil, fmt.Errorf("not found")
		return nil, domain.ErrNotFound
	}
	return r.m[id], nil
}

//Update an user
func (r *UserInmem) Update(e *entity.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *UserInmem) Search(query string) ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.FirstName), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		// return nil, fmt.Errorf("not found")
		return nil, domain.ErrNotFound
	}

	return d, nil
}

//List users
func (r *UserInmem) List() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an user
func (r *UserInmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return fmt.Errorf("not found")
	}
	r.m[id] = nil
	return nil
}
