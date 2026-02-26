package services_test

import (
	"testing"

	"github.com/ziyu-ola/rabbit-test/services"
)

func TestGreeter_Greet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "World", want: "Hello, World!"},
		{name: "Go", want: "Hello, Go!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := services.NewGreeter(tt.name)
			if got := g.Greet(); got != tt.want {
				t.Errorf("Greet() = %q, want %q", got, tt.want)
			}
		})
	}
}
