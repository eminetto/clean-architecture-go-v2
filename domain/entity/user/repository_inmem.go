package user

import (
	"strings"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[entity.ID]*User
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[entity.ID]*User{}
	return &IRepo{
		m: m,
	}
}

//Create an user
func (r *IRepo) Create(e *User) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *IRepo) Get(id entity.ID) (*User, error) {
	if r.m[id] == nil {
		return nil, domain.ErrNotFound
	}
	return r.m[id], nil
}

//Update an user
func (r *IRepo) Update(e *User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *IRepo) Search(query string) ([]*User, error) {
	var d []*User
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.FirstName), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, domain.ErrNotFound
	}

	return d, nil
}

//List users
func (r *IRepo) List() ([]*User, error) {
	var d []*User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an user
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return domain.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
