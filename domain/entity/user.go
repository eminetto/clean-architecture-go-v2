package entity

import (
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain"
	"golang.org/x/crypto/bcrypt"
)

//User data
type User struct {
	ID        ID
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []ID
}

//NewUser create a new user
func NewUser(email, password, firstName, lastName string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	err = u.Validate()
	if err != nil {
		return nil, domain.ErrInvalidEntity
	}
	return u, nil
}

//AddBook add a book
func (u *User) AddBook(id ID) error {
	_, err := u.GetBook(id)
	if err == nil {
		return domain.ErrBookAlreadyBorrowed
	}
	u.Books = append(u.Books, id)
	return nil
}

//RemoveBook remove a book
func (u *User) RemoveBook(id ID) error {
	for i, j := range u.Books {
		if j == id {
			u.Books = append(u.Books[:i], u.Books[i+1:]...)
			return nil
		}
	}
	return domain.ErrNotFound
}

//GetBook get a book
func (u *User) GetBook(id ID) (ID, error) {
	for _, v := range u.Books {
		if v == id {
			return id, nil
		}
	}
	return id, domain.ErrNotFound
}

//Validate validate data
func (u *User) Validate() error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Password == "" {
		return domain.ErrInvalidEntity
	}

	return nil
}

//ValidatePassword
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
