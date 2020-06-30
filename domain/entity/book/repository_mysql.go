package book

import (
	"database/sql"
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//mySQLRepo mysql repo
type mySQLRepo struct {
	db *sql.DB
}

//NewMySQLRepository create new repository
func NewMySQLRepository(db *sql.DB) *mySQLRepo {
	return &mySQLRepo{
		db: db,
	}
}

//Create a book
func (r *mySQLRepo) Create(e *Book) (entity.ID, error) {
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

//Get a book
func (r *mySQLRepo) Get(id entity.ID) (*Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where id = ?`)
	if err != nil {
		return nil, err
	}
	var b Book
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
	}
	return &b, nil
}

//Update a book
func (r *mySQLRepo) Update(e *Book) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update book set title = ?, author = ?, pages = ?, quantity = ?, updated_at = ? where id = ?", e.Title, e.Author, e.Pages, e.Quantity, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search books
func (r *mySQLRepo) Search(query string) ([]*Book, error) {
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

//List books
func (r *mySQLRepo) List() ([]*Book, error) {
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

//Delete a book
func (r *mySQLRepo) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from book where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
