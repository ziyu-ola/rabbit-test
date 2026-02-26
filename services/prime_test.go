package services_test

import (
	"fmt"
	"testing"

	"github.com/ziyu-ola/rabbit-test/services"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{n: 1, want: false},
		{n: 2, want: true},
		{n: 3, want: true},
		{n: 4, want: false},
		{n: 5, want: true},
		{n: 6, want: false},
		{n: 7, want: true},
		{n: 8, want: false},
		{n: 9, want: false},
		{n: 10, want: false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			if got := services.IsPrime(tt.n); got != tt.want {
				t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestLoopAndCheckPrimes(t *testing.T) {
	results := services.LoopAndCheckPrimes()

	if len(results) != 10 {
		t.Fatalf("expected 10 results, got %d", len(results))
	}

	expected := map[int]bool{
		1: false, 2: true, 3: true, 4: false, 5: true,
		6: false, 7: true, 8: false, 9: false, 10: false,
	}

	for n, want := range expected {
		if got := results[n]; got != want {
			t.Errorf("LoopAndCheckPrimes()[%d] = %v, want %v", n, got, want)
		}
	}
}
