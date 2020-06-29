package password

import (
	"golang.org/x/crypto/bcrypt"
)

//Password password
type Password struct{}

//NewService create a new fake password
func NewService() *Password {
	return &Password{}
}

//Generate a new password
func (p *Password) Generate(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//Compare compare two passwords
func (p *Password) Compare(p1, p2 string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2))
	if err != nil {
		return err
	}
	return nil
}
