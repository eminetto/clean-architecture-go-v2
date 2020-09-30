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
