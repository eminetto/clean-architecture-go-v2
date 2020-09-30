package entity

import (
	"time"
)

func NewFixtureBook() *Book {
	return &Book{
		ID:        NewID(),
		Title:     "I Am Ozzy",
		Author:    "Ozzy Osbourne",
		Pages:     294,
		Quantity:  1,
		CreatedAt: time.Now(),
	}
}
