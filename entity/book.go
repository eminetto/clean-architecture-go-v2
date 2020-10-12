package entity

import (
	"time"
)

//Book data
type Book struct {
	ID        ID
	Title     string
	Author    string
	Pages     int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewBook create a new book
func NewBook(title string, author string, pages int, quantity int) (*Book, error) {
	b := &Book{
		ID:        NewID(),
		Title:     title,
		Author:    author,
		Pages:     pages,
		Quantity:  quantity,
		CreatedAt: time.Now(),
	}
	err := b.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return b, nil
}

//Validate validate book
func (b *Book) Validate() error {
	if b.Title == "" || b.Author == "" || b.Pages <= 0 || b.Quantity <= 0 {
		return ErrInvalidEntity
	}
	return nil
}
