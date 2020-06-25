package user

import (
	"github.com/eminetto/clean-architecture-go-v2/domain/entity"
	"time"
)

func NewFixtureUser() *User {
	return  &User{
		ID: entity.NewID(),
		Email:"ozzy@metalgods.net",
		FirstName:"Ozzy",
		LastName: "Osbourne",
		CreatedAt: time.Now(),
	}
}