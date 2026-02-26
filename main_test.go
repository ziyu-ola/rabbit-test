package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ziyu-ola/rabbit-test/db"
)

func TestMain(m *testing.M) {
	// Initialize DB once for all tests
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestLookupUsers(t *testing.T) {

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture stderr
	oldErr := os.Stderr
	rErr, wErr, _ := os.Pipe()
	os.Stderr = wErr

	lookupUsers()

	// Restore stdout and stderr
	w.Close()
	os.Stdout = old
	wErr.Close()
	os.Stderr = oldErr

	var buf bytes.Buffer
	io.Copy(&buf, r)
	var bufErr bytes.Buffer
	io.Copy(&bufErr, rErr)

	output := buf.String()
	errOutput := bufErr.String()

	// Verify no errors
	if errOutput != "" {
		t.Errorf("unexpected stderr output: %s", errOutput)
	}

	// Verify all UIDs are present in output
	expectedUIDs := []int{1000, 1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010, 1011, 1012, 1013, 1014, 1015}
	expectedNames := []string{"Alice", "Bob", "Charlie", "Dave", "Eve", "Frank", "Grace", "Hank", "Ivy", "Jack", "Karen", "Leo", "Mia", "Nick", "Olivia", "Paul"}

	for i, uid := range expectedUIDs {
		expectedLine := fmt.Sprintf("uid %d: %s", uid, expectedNames[i])
		if !strings.Contains(output, expectedLine) {
			t.Errorf("output missing expected line: %q", expectedLine)
		}
	}
}

func TestLookupUsers_OutputFormat(t *testing.T) {

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	lookupUsers()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output contains expected format "uid XXXX: Name"
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 16 {
		t.Errorf("expected 16 lines of output, got %d", len(lines))
	}

	for _, line := range lines {
		if !strings.Contains(line, "uid ") || !strings.Contains(line, ": ") {
			t.Errorf("line does not match expected format 'uid X: Name': %q", line)
		}
	}
}

func TestLookupUsers_NoErrors(t *testing.T) {
	// This test verifies that lookupUsers runs without errors
	// DB should already be initialized by TestMain

	// Capture stderr to check for error messages
	oldErr := os.Stderr
	rErr, wErr, _ := os.Pipe()
	os.Stderr = wErr

	// Call lookupUsers (should succeed since DB is initialized)
	lookupUsers()

	wErr.Close()
	os.Stderr = oldErr

	var bufErr bytes.Buffer
	io.Copy(&bufErr, rErr)
	errOutput := bufErr.String()

	// In normal case, there should be no errors
	if errOutput != "" {
		t.Errorf("unexpected error output: %s", errOutput)
	}
}

func TestMain_WithoutArgs(t *testing.T) {
	// Save original args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set args to just program name (no birthday argument)
	os.Args = []string{"test"}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture stderr
	oldErr := os.Stderr
	rErr, wErr, _ := os.Pipe()
	os.Stderr = wErr


	main()

	w.Close()
	os.Stdout = old
	wErr.Close()
	os.Stderr = oldErr

	var buf bytes.Buffer
	io.Copy(&buf, r)
	var bufErr bytes.Buffer
	io.Copy(&bufErr, rErr)

	output := buf.String()

	// Verify greeting is present
	if !strings.Contains(output, "Hello, World!") {
		t.Errorf("expected greeting 'Hello, World!' in output, got: %s", output)
	}

	// Verify user lookups are present
	if !strings.Contains(output, "uid 1000: Alice") {
		t.Errorf("expected user lookup output, got: %s", output)
	}
}

func TestMain_WithValidBirthday(t *testing.T) {
	// Save original args and exit function
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Set args with a valid birthday
	os.Args = []string{"test", "2000-01-01"}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w


	main()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()

	// Verify greeting is present
	if !strings.Contains(output, "Hello, World!") {
		t.Errorf("expected greeting in output")
	}

	// Verify age output is present
	if !strings.Contains(output, "Age:") {
		t.Errorf("expected age output, got: %s", output)
	}

	// Verify user lookups are present
	if !strings.Contains(output, "uid 1000: Alice") {
		t.Errorf("expected user lookup output")
	}
}

func TestMain_WithInvalidBirthday(t *testing.T) {
	// Note: This test cannot fully test os.Exit(1) without refactoring main.go
	// However, we can verify that stderr gets the error message before exit
	// We just can't capture after os.Exit is called

	// This is a limitation of testing functions that call os.Exit
	// In a real scenario, we would refactor main() to return an error
	// instead of calling os.Exit directly
	t.Skip("Skipping test that requires os.Exit mocking - would need code refactoring")
}

func TestLookupUsers_Comprehensive(t *testing.T) {
	// Additional comprehensive test to verify complete functionality
	// Capture both stdout and stderr
	oldOut := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	oldErr := os.Stderr
	rErr, wErr, _ := os.Pipe()
	os.Stderr = wErr

	lookupUsers()

	wOut.Close()
	os.Stdout = oldOut
	wErr.Close()
	os.Stderr = oldErr

	var bufOut bytes.Buffer
	io.Copy(&bufOut, rOut)
	var bufErr bytes.Buffer
	io.Copy(&bufErr, rErr)

	output := bufOut.String()
	errOutput := bufErr.String()

	// Verify no errors
	if errOutput != "" {
		t.Errorf("unexpected stderr: %s", errOutput)
	}

	// Verify complete output
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 16 {
		t.Errorf("expected 16 lines, got %d", len(lines))
	}

	// Verify first and last entries
	if !strings.HasPrefix(lines[0], "uid 1000:") {
		t.Errorf("first line should start with 'uid 1000:', got: %s", lines[0])
	}
	if !strings.HasPrefix(lines[15], "uid 1015:") {
		t.Errorf("last line should start with 'uid 1015:', got: %s", lines[15])
	}
}

func TestLookupUsers_AllExpectedUsers(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	lookupUsers()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Test that all 16 users from 1000-1015 are present
	expectedUsers := map[int]string{
		1000: "Alice",
		1001: "Bob",
		1002: "Charlie",
		1003: "Dave",
		1004: "Eve",
		1005: "Frank",
		1006: "Grace",
		1007: "Hank",
		1008: "Ivy",
		1009: "Jack",
		1010: "Karen",
		1011: "Leo",
		1012: "Mia",
		1013: "Nick",
		1014: "Olivia",
		1015: "Paul",
	}

	for uid, name := range expectedUsers {
		expected := fmt.Sprintf("uid %d: %s", uid, name)
		if !strings.Contains(output, expected) {
			t.Errorf("missing expected output: %q", expected)
		}
	}
}