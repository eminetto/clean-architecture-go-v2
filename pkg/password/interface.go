package password

//UseCase interface
type UseCase interface {
	Generate(raw string) (string, error)
	Compare(p1, p2 string) error
}
