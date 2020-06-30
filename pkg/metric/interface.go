package metric

import "time"

//CLI define a CLI app
type CLI struct {
	Name       string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

// NewCLI create a new CLI app
func NewCLI(name string) *CLI {
	return &CLI{
		Name: name,
	}
}

//Started start monitoring the app
func (c *CLI) Started() {
	c.StartedAt = time.Now()
}

// Finished app finished
func (c *CLI) Finished() {
	c.FinishedAt = time.Now()
	c.Duration = time.Since(c.StartedAt).Seconds()
}

//HTTP application
type HTTP struct {
	Handler    string
	Method     string
	StatusCode string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

//NewHTTP create a new HTTP app
func NewHTTP(handler string, method string) *HTTP {
	return &HTTP{
		Handler: handler,
		Method:  method,
	}
}

//Started start monitoring the app
func (h *HTTP) Started() {
	h.StartedAt = time.Now()
}

// Finished app finished
func (h *HTTP) Finished() {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt).Seconds()
}

//Service definition
type Service interface {
	SaveCLI(c *CLI) error
	SaveHTTP(h *HTTP)
}
