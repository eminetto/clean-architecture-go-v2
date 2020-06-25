package user

import (
	"time"

	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
)

func NewFixtureUser() *User {
	return &User{
		ID:        entity.NewID(),
		Email:     "ozzy@metalgods.net",
		Password:  "123456",
		FirstName: "Ozzy",
		LastName:  "Osbourne",
		CreatedAt: time.Now(),
	}
}
