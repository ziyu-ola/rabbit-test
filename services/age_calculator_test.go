package services_test

import (
	"testing"
	"time"

	"github.com/ziyu-ola/rabbit-test/services"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		birthday string
		want     int
	}{
		{name: "30 years old", birthday: "1996-02-26", want: 30},
		{name: "0 years old (born today)", birthday: "2026-02-26", want: 0},
		{name: "birthday not yet this year", birthday: "1990-12-31", want: 35},
	}

	// Use a fixed reference date so the tests are deterministic.
	now := time.Date(2026, 2, 26, 12, 0, 0, 0, time.UTC)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			birthday, err := time.Parse("2006-01-02", tt.birthday)
			if err != nil {
				t.Fatalf("failed to parse birthday: %v", err)
			}
			got := services.CalculateAgeAt(birthday, now)
			if got != tt.want {
				t.Errorf("CalculateAgeAt() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestAgeFromBirthdayString_InvalidFormat(t *testing.T) {
	_, err := services.AgeFromBirthdayString("not-a-date")
	if err == nil {
		t.Error("expected an error for invalid birthday format, got nil")
	}
}
