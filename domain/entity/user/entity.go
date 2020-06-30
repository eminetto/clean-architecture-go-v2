package user

import (
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

//User data
type User struct {
	ID        entity.ID
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []entity.ID
}
