package book

import (
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//Book data
type Book struct {
	ID        entity.ID
	Title     string
	Author    string
	Pages     int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
