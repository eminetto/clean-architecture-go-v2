package presenter

import (
	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//Book data
type Book struct {
	ID       entity.ID `json:"id"`
	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Pages    int       `json:"pages"`
	Quantity int       `json:"quantity"`
}
