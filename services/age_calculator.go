package services

import (
	"fmt"
	"time"
)

// CalculateAge returns the age in years for a given birthday as of today.
func CalculateAge(birthday time.Time) int {
	return CalculateAgeAt(birthday, time.Now())
}

// CalculateAgeAt returns the age in years for a given birthday as of the provided reference time.
func CalculateAgeAt(birthday, now time.Time) int {
	years := now.Year() - birthday.Year()
	// If the birthday hasn't occurred yet this year, subtract one year.
	birthdayThisYear := time.Date(now.Year(), birthday.Month(), birthday.Day(), 0, 0, 0, 0, now.Location())
	if now.Before(birthdayThisYear) {
		years--
	}
	return years
}

// AgeFromBirthdayString parses a birthday string in "YYYY-MM-DD" format and returns the age.
func AgeFromBirthdayString(birthday string) (int, error) {
	t, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, fmt.Errorf("invalid birthday format (expected YYYY-MM-DD): %w", err)
	}
	return CalculateAge(t), nil
}
