package password

import "errors"

//FakePassword password
type FakePassword struct{}

//NewFakeService create a new fake password
func NewFakeService() *FakePassword {
	return &FakePassword{}
}

//Generate a new password
func (p *FakePassword) Generate(raw string) (string, error) {
	return raw, nil
}

//Compare compare two passwords
func (p *FakePassword) Compare(p1, p2 string) error {
	if p1 == p2 {
		return nil
	}
	return errors.New("Invalid password")
}
