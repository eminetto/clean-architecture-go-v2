package repository

import (
	"strings"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//BookInmem in memory repo
type BookInmem struct {
	m map[entity.ID]*entity.Book
}

//NewBookInmem create new repository
func NewBookInmem() *BookInmem {
	var m = map[entity.ID]*entity.Book{}
	return &BookInmem{
		m: m,
	}
}

//Create a book
func (r *BookInmem) Create(e *entity.Book) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a book
func (r *BookInmem) Get(id entity.ID) (*entity.Book, error) {
	if r.m[id] == nil {
		// return nil, fmt.Errorf("not found")
		return nil, domain.ErrNotFound
	}
	return r.m[id], nil
}

//Update a book
func (r *BookInmem) Update(e *entity.Book) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search books
func (r *BookInmem) Search(query string) ([]*entity.Book, error) {
	var d []*entity.Book
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.Title), query) {
			d = append(d, j)
		}
	}
	return d, nil
}

//List books
func (r *BookInmem) List() ([]*entity.Book, error) {
	var d []*entity.Book
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete a book
func (r *BookInmem) Delete(id entity.ID) error {
	if r.m[id] == nil {
		// return fmt.Errorf("not found")
		return domain.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
