package user

import (
	"fmt"
	"strings"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//IRepo in memory repo
type IRepo struct {
	m map[entity.ID]*entity.User
}

//NewInmemRepository create new repository
func NewInmemRepository() *IRepo {
	var m = map[entity.ID]*entity.User{}
	return &IRepo{
		m: m,
	}
}

//Create an user
func (r *IRepo) Create(e *entity.User) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *IRepo) Get(id entity.ID) (*entity.User, error) {
	if r.m[id] == nil {
		// return nil, fmt.Errorf("not found")
		return nil, domain.ErrNotFound
	}
	return r.m[id], nil
}

//Update an user
func (r *IRepo) Update(e *entity.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *IRepo) Search(query string) ([]*entity.User, error) {
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
func (r *IRepo) List() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an user
func (r *IRepo) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return fmt.Errorf("not found")
	}
	r.m[id] = nil
	return nil
}
