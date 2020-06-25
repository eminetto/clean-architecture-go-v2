package user

import (
	"database/sql"
	"time"

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
func (r *MySQLRepo) Create(e *User) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into user (id, email, password, first_name, last_name, created_at) 
		values(?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Email,
		e.Password,
		e.FirstName,
		e.LastName,
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
func (r *MySQLRepo) Get(id entity.ID) (*User, error) {
	return nil, nil
}

//Update an user
func (r *MySQLRepo) Update(e *User) error {
	return nil
}

//Search users
func (r *MySQLRepo) Search(query string) ([]*User, error) {

	return nil, nil
}

//List users
func (r *MySQLRepo) List() ([]*User, error) {
	return nil, nil
}

//Delete an user
func (r *MySQLRepo) Delete(id entity.ID) error {
	return nil
}
