package entity

import (
	"time"
)

func NewFixtureUser() *User {
	return &User{
		ID:        NewID(),
		Email:     "ozzy@metalgods.net",
		Password:  "123456",
		FirstName: "Ozzy",
		LastName:  "Osbourne",
		CreatedAt: time.Now(),
	}
}
