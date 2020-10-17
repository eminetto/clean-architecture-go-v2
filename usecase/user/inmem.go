package user

import (
	"fmt"
	"strings"

	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.User
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.User{}
	return &inmem{
		m: m,
	}
}

//Create an user
func (r *inmem) Create(e *entity.User) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *inmem) Get(id entity.ID) (*entity.User, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update an user
func (r *inmem) Update(e *entity.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *inmem) Search(query string) ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.FirstName), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}

	return d, nil
}

//List users
func (r *inmem) List() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an user
func (r *inmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return fmt.Errorf("not found")
	}
	r.m[id] = nil
	return nil
}
