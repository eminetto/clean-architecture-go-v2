package password

//Service interface
type Service interface {
	Generate(raw string) (string, error)
	Compare(p1, p2 string) error
}
