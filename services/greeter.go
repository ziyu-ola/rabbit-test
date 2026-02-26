package services

import "fmt"

// Greeter is a simple greeting service.
type Greeter struct {
	name string
}

// NewGreeter creates a new Greeter for the given name.
func NewGreeter(name string) *Greeter {
	return &Greeter{name: name}
}

// Greet returns a greeting message.
func (g *Greeter) Greet() string {
	return fmt.Sprintf("Hello, %s!", g.name)
}
