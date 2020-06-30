package book

import (
	"strings"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//iRepo in memory repo
type iRepo struct {
	m map[entity.ID]*Book
}

//NewInmemRepository create new repository
func NewInmemRepository() *iRepo {
	var m = map[entity.ID]*Book{}
	return &iRepo{
		m: m,
	}
}

//Create a book
func (r *iRepo) Create(e *Book) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a book
func (r *iRepo) Get(id entity.ID) (*Book, error) {
	if r.m[id] == nil {
		return nil, domain.ErrNotFound
	}
	return r.m[id], nil
}

//Update a book
func (r *iRepo) Update(e *Book) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search books
func (r *iRepo) Search(query string) ([]*Book, error) {
	var d []*Book
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Title), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, domain.ErrNotFound
	}

	return d, nil
}

//List books
func (r *iRepo) List() ([]*Book, error) {
	var d []*Book
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a book
func (r *iRepo) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return domain.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
