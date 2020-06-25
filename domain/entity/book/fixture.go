package book

import (
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

func NewFixtureBook() *Book {
	return &Book{
		ID:        entity.NewID(),
		Title:     "I Am Ozzy",
		Author:    "Ozzy Osbourne",
		Pages:     294,
		Quantity:  1,
		CreatedAt: time.Now(),
	}
}
