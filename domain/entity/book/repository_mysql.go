package book

import (
	"database/sql"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//MySQLRepo mysql repo
type MySQLRepo struct {
	db *sql.DB
}

//NewMySQLRepoRepository create new repository
func NewMySQLRepoRepository(db *sql.DB) *MySQLRepo {
	return &MySQLRepo{
		db: db,
	}
}

//Create an user
func (r *MySQLRepo) Create(e *Book) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into book (id, title, author, pages, quantity, created_at) 
		values(?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Title,
		e.Author,
		e.Pages,
		e.Quantity,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

//Get an user
func (r *MySQLRepo) Get(id entity.ID) (*Book, error) {
	return nil, nil
}

//Update an user
func (r *MySQLRepo) Update(e *Book) error {
	return nil
}

//Search users
func (r *MySQLRepo) Search(query string) ([]*Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where title like ?`)
	if err != nil {
		return nil, err
	}
	var books []*Book
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	if len(books) == 0 {
		return nil, domain.ErrNotFound
	}
	return books, nil
}

//List users
func (r *MySQLRepo) List() ([]*Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book`)
	if err != nil {
		return nil, err
	}
	var books []*Book
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	if len(books) == 0 {
		return nil, domain.ErrNotFound
	}
	return books, nil
}

//Delete an user
func (r *MySQLRepo) Delete(id entity.ID) error {
	return nil
}
