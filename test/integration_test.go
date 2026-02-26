package test

import (
	"testing"

	"github.com/ziyu-ola/rabbit-test/services"
)

// TestGreeterIntegration is an integration-style test for the greeter service.
func TestGreeterIntegration(t *testing.T) {
	g := services.NewGreeter("rabbit-test")
	got := g.Greet()
	want := "Hello, rabbit-test!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
