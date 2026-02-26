package db_test

import (
	"database/sql"
	"sync"
	"testing"

	"github.com/ziyu-ola/rabbit-test/db"
)

func TestMain(m *testing.M) {
	// Initialize DB once for all tests
	if err := db.InitDB(); err != nil {
		panic(err)
	}
	m.Run()
}

func TestInitDB(t *testing.T) {
	// InitDB should be idempotent due to sync.Once
	err := db.InitDB()
	if err != nil {
		t.Fatalf("InitDB() failed: %v", err)
	}

	if db.DB == nil {
		t.Fatal("InitDB() did not set DB")
	}

	// Verify table exists by querying it
	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		t.Fatalf("failed to query users table: %v", err)
	}

	if count != 16 {
		t.Errorf("expected 16 users, got %d", count)
	}
}

func TestInitDB_Concurrent(t *testing.T) {
	// Test concurrent initialization (should be safe due to sync.Once)
	var wg sync.WaitGroup
	errChan := make(chan error, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := db.InitDB(); err != nil {
				errChan <- err
			}
		}()
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		t.Errorf("concurrent InitDB() failed: %v", err)
	}

	if db.DB == nil {
		t.Fatal("InitDB() did not set DB")
	}
}

func TestGetNameByUid(t *testing.T) {
	// DB should already be initialized by TestMain

	tests := []struct {
		name    string
		uid     int
		want    string
		wantErr bool
	}{
		{name: "Alice_1000", uid: 1000, want: "Alice", wantErr: false},
		{name: "Bob_1001", uid: 1001, want: "Bob", wantErr: false},
		{name: "Charlie_1002", uid: 1002, want: "Charlie", wantErr: false},
		{name: "Dave_1003", uid: 1003, want: "Dave", wantErr: false},
		{name: "Eve_1004", uid: 1004, want: "Eve", wantErr: false},
		{name: "Frank_1005", uid: 1005, want: "Frank", wantErr: false},
		{name: "Grace_1006", uid: 1006, want: "Grace", wantErr: false},
		{name: "Hank_1007", uid: 1007, want: "Hank", wantErr: false},
		{name: "Ivy_1008", uid: 1008, want: "Ivy", wantErr: false},
		{name: "Jack_1009", uid: 1009, want: "Jack", wantErr: false},
		{name: "Karen_1010", uid: 1010, want: "Karen", wantErr: false},
		{name: "Leo_1011", uid: 1011, want: "Leo", wantErr: false},
		{name: "Mia_1012", uid: 1012, want: "Mia", wantErr: false},
		{name: "Nick_1013", uid: 1013, want: "Nick", wantErr: false},
		{name: "Olivia_1014", uid: 1014, want: "Olivia", wantErr: false},
		{name: "Paul_1015", uid: 1015, want: "Paul", wantErr: false},
		{name: "NonExistent_999", uid: 999, want: "", wantErr: true},
		{name: "NonExistent_1016", uid: 1016, want: "", wantErr: true},
		{name: "NonExistent_0", uid: 0, want: "", wantErr: true},
		{name: "NonExistent_Negative", uid: -1, want: "", wantErr: true},
		{name: "NonExistent_LargeValue", uid: 9999, want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.GetNameByUid(tt.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNameByUid(%d) error = %v, wantErr %v", tt.uid, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNameByUid(%d) = %q, want %q", tt.uid, got, tt.want)
			}
		})
	}
}

func TestGetNameByUid_ErrorHandling(t *testing.T) {
	// Test with UID that doesn't exist
	name, err := db.GetNameByUid(9999)
	if err == nil {
		t.Error("expected error for non-existent UID, got nil")
	}
	if name != "" {
		t.Errorf("expected empty name for non-existent UID, got %q", name)
	}

	// Verify the error is sql.ErrNoRows wrapped
	if err != nil && err != sql.ErrNoRows {
		// Check if it's a wrapped error
		if !contains(err.Error(), "query uid") {
			t.Errorf("expected error to contain 'query uid', got: %v", err)
		}
	}
}

func TestGetNameByUid_BoundaryValues(t *testing.T) {
	// Test boundary values
	tests := []struct {
		name    string
		uid     int
		wantErr bool
	}{
		{name: "FirstUID", uid: 1000, wantErr: false},
		{name: "LastUID", uid: 1015, wantErr: false},
		{name: "BeforeFirst", uid: 999, wantErr: true},
		{name: "AfterLast", uid: 1016, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.GetNameByUid(tt.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNameByUid(%d) error = %v, wantErr %v", tt.uid, err, tt.wantErr)
			}
		})
	}
}

func TestInitDB_DataIntegrity(t *testing.T) {
	// Verify all expected names are present in order
	expectedNames := []string{
		"Alice", "Bob", "Charlie", "Dave", "Eve",
		"Frank", "Grace", "Hank", "Ivy", "Jack",
		"Karen", "Leo", "Mia", "Nick", "Olivia", "Paul",
	}

	for i, expectedName := range expectedNames {
		uid := 1000 + i
		name, err := db.GetNameByUid(uid)
		if err != nil {
			t.Errorf("GetNameByUid(%d) failed: %v", uid, err)
			continue
		}
		if name != expectedName {
			t.Errorf("GetNameByUid(%d) = %q, want %q", uid, name, expectedName)
		}
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && hasSubstring(s, substr))
}

func hasSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}